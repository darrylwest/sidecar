//
// handlers - methods to handle requests
//
// @author darryl.west <darwest@ebay.com>
// @created 2018-05-09 08:35:20
//

package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-zoo/bone"
)

// Handlers the handlers struct for configuration
type Handlers struct {
	cfg       *Config
}

// NewHandlers create the new handlers object
func NewHandlers(cfg *Config) *Handlers {
	hnd := &Handlers{}
	hnd.cfg = cfg

	return hnd
}

// HomeHandler closure that returns the home page
func (hnd *Handlers) HomeHandler() http.HandlerFunc {
	log.Info("home page handler...")

	return func(w http.ResponseWriter, r *http.Request) {
		blob := GetStatusAsJSON(hnd.cfg)

		log.Info("status: %s", blob)

		log.Info(blob)
		fmt.Fprintf(w, "%s\n\r", blob)
	}
}

// StatusHandler closure that returns the service status
func (hnd *Handlers) StatusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		blob := GetStatusAsJSON(hnd.cfg)

		log.Info("status: %s", blob)

		log.Info(blob)
		fmt.Fprintf(w, "%s\n\r", blob)
	}
}

// GetLogLevel returns the current log level
func (hnd *Handlers) GetLogLevel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"%s\":\"%d\"}\n\r", "loglevel", log.GetLevel())
	}
}

// SetLogLevel returns the current log level
func (hnd *Handlers) SetLogLevel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		value := bone.GetValue(r, "level")
		if value == "" {
			hnd.writeErrorResponse(w, "must supply a level between 0 and 5")
			return
		}

		level, err := strconv.Atoi(value)
		if err != nil {
			log.Warn("attempt to set log level to invalid value: %d, ignored...", level)
			hnd.writeErrorResponse(w, err.Error())
			return
		}

		if level < 0 {
			level = 0
		}

		if level > 5 {
			level = 5
		}

		log.SetLevel(level)

		fmt.Fprintf(w, "{\"%s\":\"%d\"}\n\r", "loglevel", log.GetLevel())
	}
}

// CreateResponseWrapper cxreate a map
func (hnd Handlers) CreateResponseWrapper(status string) map[string]interface{} {
	wrapper := make(map[string]interface{})
	wrapper["status"] = status
	wrapper["version"] = "1.0"
	wrapper["ts"] = time.Now()

	return wrapper
}

func (hnd Handlers) writeJSONBlob(w http.ResponseWriter, wrapper map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	blob, err := json.Marshal(wrapper)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), 501)
		return
	}

	log.Debug("blob: %s", blob)
	fmt.Fprintf(w, "%s\n\r", blob)
}

func (hnd Handlers) writeErrorResponse(w http.ResponseWriter, str string) {
	log.Warn(str)
	http.Error(w, str, 501)
}
