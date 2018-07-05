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
	"path"
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

// ParseArgs parse the command line args
func ParseArgs() *Config {
	dflt := NewDefaultConfig()

	vers := flag.Bool("version", false, "show the version and exit")
	level := flag.Int("loglevel", dflt.LogLevel, "set the server's log level 0..5 for trace..error, default info=2")
	port := flag.Int("port", dflt.Port, "set the server's listening port")
	loopSeconds := flag.Int("loop-seconds", dflt.LoopSeconds, "set the real time loop seconds, default=10")

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

	log.SetLevel(cfg.LogLevel)

	return &cfg
}
