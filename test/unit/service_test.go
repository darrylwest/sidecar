//
// service
//
// @author darryl.west@ebay.com
// @created 2017-07-20 12:57:59
//

package unit

import (
	"fmt"
	"app"
	"testing"

	. "github.com/franela/goblin"
)

func TestService(t *testing.T) {
	g := Goblin(t)

	g.Describe("Service", func() {

		g.It("should create a service struct", func() {
			cfg := new(app.Config)
			service, err := app.NewService(cfg)
			g.Assert(err).Equal(nil)
			g.Assert(fmt.Sprintf("%T", service)).Equal("*app.Service")
		})
	})
}
