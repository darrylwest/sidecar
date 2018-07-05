//
// handlers - methods to handle requests
//
// @author darryl.west <darwest@ebay.com>
// @created 2017-07-21 08:35:20
//

package app

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	// "strings"
)

// Handlers the handlers struct for configuration
type Handlers struct {
	cfg *Config
}

// NewHandlers create the new handlers object
func NewHandlers(config *Config) *Handlers {
	hnd := Handlers{config}

	return &hnd
}

func (hnd *Handlers) homeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "{\"%s\":\"%s\"}\n\r", "sidecar", "ok")
}

func (hnd *Handlers) statusHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	json := GetStatusAsJSON(hnd.cfg)
	log.Info(json)
	fmt.Fprintf(w, "%s\n\r", json)
}

// log level handlers Get: /logger -> show level; Put: /logger/n set level to n
func (hnd *Handlers) getLogLevel(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "{\"%s\":\"%d\"}\n\r", "loglevel", log.GetLevel())
}

// set the log level
func (hnd *Handlers) setLogLevel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	level, err := strconv.Atoi(ps.ByName("level"))
	if err == nil && level > 0 && level < 5 {
		log.SetLevel(level)
	} else {
		log.Warn("attempt to set log level to invalid value: %s, ignored...", level)
	}

	fmt.Fprintf(w, "{\"%s\":\"%d\"}\n\r", "loglevel", log.GetLevel())
}
