//
// handlers tests
//
// @author darryl.west <darwest@ebay.com>
// @created 2017-07-21 08:35:20
//

package unit

import (
	"app"
	"testing"

	. "github.com/franela/goblin"
)

func TestHandlers(t *testing.T) {
	g := Goblin(t)

	g.Describe("Handlers", func() {
		app.CreateLogger()

		g.It("should create a handler struct", func() {
			cfg := new(app.Config)
			hnd := app.NewHandlers(cfg)
			g.Assert(hnd != nil).IsTrue()
		})

		g.It("should return a valid status object", func() {
			hnd := app.NewHandlers(new(app.Config))
			g.Assert(hnd != nil).IsTrue()
			// need a response writer, request and httprouter  mocks
		})
	})
}
