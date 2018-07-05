//
// service - define the routes and start the service
//
// @author darryl.west@ebay.com
// @created 2017-07-20 12:57:59
//

package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Client interface for generic sidecar clients
type Client interface {
	ProcessLoop(docker *DockerClient)
}

// Service - the service struct
type Service struct {
	cfg      *Config
	handlers *Handlers
	client   Client
}

// ContainerSpec holds the current container specification
type ContainerSpec struct {
	Name      string
	Version   int
	Instances int
}

// NewService create a new service by passing in config
func NewService(config *Config) (*Service, error) {

	handlers := NewHandlers(config)
	svc := Service{config, handlers, nil}

	return &svc, nil
}

// Start start the admin listener and event loop
func (svc Service) Start() error {
	log.Info("start the sidecar service for service-hub...")

	// start the listener in background
	go svc.startServer()

	stop := make(chan bool)
	err := svc.runEventLoop(stop)
	if err != nil {
		log.Error("error running event loop: %v", err)
	}

	return err
}

func (svc Service) runEventLoop(stop chan bool) error {

	docker, err := NewDockerClient(svc.cfg)
	if err != nil {
		log.Warn("docker is not active...")
		return err
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	tick := ticker.C

	count := 0
	log.Info("run the event loop...")

	for {
		select {
		case <-tick:
			if count < 1 {
				// reset the count first
				count = svc.cfg.LoopSeconds

				// invoke the loop processor
				svc.client.ProcessLoop(docker)
			}

			count--

		case <-stop:
			log.Info("received stop signal...")
			return nil
		}
	}
}

func (svc Service) startServer() {
	cfg := svc.cfg
	hnd := svc.handlers

	log.Info("configure the router/handler...")
	router := httprouter.New()
	router.GET("/", hnd.homeHandler)
	router.GET("/status", hnd.statusHandler)
	router.GET("/logger", hnd.getLogLevel)
	router.PUT("/logger/:level", hnd.setLogLevel)

	host := fmt.Sprintf(":%d", cfg.Port)
	log.Info("start listening on port %s", host)

	err := http.ListenAndServe(host, router)
	if err != nil {
		log.Error("http error: %s", err)
	}
}
