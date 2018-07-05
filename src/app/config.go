//
// config  - application specification and CLI parsing
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-07-20 17:56:46

package app

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

// Config the config structure
type Config struct {
	Port        int
	LogLevel    int
	LoopSeconds int
	Home        string
}

// NewDefaultConfig default settings
func NewDefaultConfig() *Config {
	cfg := new(Config)

	cfg.Port = 8001
	cfg.LogLevel = 3
	cfg.LoopSeconds = 10
	cfg.Home = "/opt/"

	return cfg
}

// ShowHelp dump out the use/command line options
func ShowHelp() {
	fmt.Printf("\n%s USE:\n\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Printf("\n%s Version %s\n", os.Args[0], Version())
}

func (cfg Config) pidFileExists(pidfile string) bool {
	file, err := os.Open(pidfile)
	if err != nil {
		return false
	}

	defer file.Close()
	pid, err := cfg.readPidFile(file)

	return err == nil && pid > 0
}

func (cfg Config) readPidFile(file *os.File) (int, error) {
	// read the pid and determine if the process is dead
	data := make([]byte, 20)
	n, err := file.Read(data)
	if err != nil {
		log.Warn("data in pid file has errors: %s", err)
		return 0, err
	}

	data = data[:n]

	log.Info("pid file contains: %s", data)

	return strconv.Atoi(string(data))
}

// CreatePidFilename creates a pid filename from the service name
func (cfg Config) CreatePidFilename() string {
	return fmt.Sprintf("%s/test-hub-sidecar.pid", cfg.Home)
}

func (cfg Config) daemonize() {
	// check to see if sidecar is running now...
	fn := cfg.CreatePidFilename()
	if cfg.pidFileExists(fn) {
		log.Warn("pid file %s exists, aborting daemon...", fn)
		return
	}

	log.Info("re-launch %s in background...", os.Args[0])

	cmd := exec.Command(os.Args[0])
	for i := 1; i < len(os.Args); i++ {
		if strings.HasSuffix(os.Args[i], "-start") == false {
			cmd.Args = append(cmd.Args, os.Args[i])
		}
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Start()

	log.Info("save pid %d to file: %s\n", cmd.Process.Pid, fn)
	file, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Error("error opening pid file: %s", err)
		return
	}

	defer file.Close()
	_, err = file.WriteString(strconv.Itoa(cmd.Process.Pid))
	if err != nil {
		log.Error("error writing to pid file: %s", err)
	}
}

// stop a sidecar process
func (cfg Config) stop() {
	fn := cfg.CreatePidFilename()

	file, err := os.Open(fn)
	if err != nil {
		log.Error("error reading pid file: %s %s", fn, err)
		return
	}

	defer func() {
		file.Close()
		os.Remove(fn)
	}()

	pid, err := cfg.readPidFile(file)
	if err != nil {
		log.Error("error getting pid from file: %s %s", fn, err)
		return
	}

	log.Info("kill %d ...", pid)
	cmd := exec.Command("kill", "-2", strconv.Itoa(pid))
	err = cmd.Start()
	if err != nil {
		log.Error("error stopping: %s", err)
		return
	}

	cmd.Wait()
}

// ParseArgs parse the command line args
func ParseArgs() *Config {
	dflt := NewDefaultConfig()

	vers := flag.Bool("version", false, "show the version and exit")
	level := flag.Int("loglevel", dflt.LogLevel, "set the server's log level 0..5 for trace..error, default info=2")
	port := flag.Int("port", dflt.Port, "set the server's listening port")
	loopSeconds := flag.Int("loop-seconds", dflt.LoopSeconds, "set the real time loop seconds, default=10")
	start := flag.Bool("start", false, "start the service and run as a daemon")
	stop := flag.Bool("stop", false, "stop the process and remove the pid file")

	flag.Parse()

	if *vers == true {
		fmt.Printf("Version %s\n", Version())
		return nil
	}

	log.Info("%s Version: %s\n%s\n", path.Base(os.Args[0]), Version(), appLogo())

	cfg := Config{
		Port:        *port,
		LogLevel:    *level,
		LoopSeconds: *loopSeconds,
	}

	if *start {
		cfg.daemonize()
		time.Sleep(time.Millisecond * 500)
		os.Exit(0)
	}

	if *stop {
		cfg.stop()
		time.Sleep(time.Millisecond * 500)
		os.Exit(0)
	}

	log.SetLevel(cfg.LogLevel)

	return &cfg
}
