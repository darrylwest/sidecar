//
// handlers tests
//
// @author darryl.west <darwest@ebay.com>
// @created 2017-08-21 08:35:20
//

package unit

import (
	"fmt"
	"app"
	"testing"

	. "github.com/franela/goblin"
)

func TestConfig(t *testing.T) {
	g := Goblin(t)

	g.Describe("Config", func() {
		app.CreateLogger()

		g.It("should create a config struct", func() {
			cfg := new(app.Config)
			g.Assert(fmt.Sprintf("%T", cfg)).Equal("*app.Config")
		})

		g.It("should create a context struct with defaults set", func() {
			cfg := app.NewDefaultConfig()
			g.Assert(cfg != nil).IsTrue()
			g.Assert(cfg.Port).Equal(8001)
			g.Assert(cfg.LogLevel >= 2).IsTrue()
			g.Assert(cfg.LoopSeconds).Equal(10)
			// g.Assert(cfg.Home).Equal(path.Join(os.Getenv("HOME"), "TestAutomation", "bolt"))
		})

		g.It("should parse an empty command line and return default config", func() {
			cfg := app.ParseArgs()
			g.Assert(cfg != nil).IsTrue()
		})

		g.It("should create a standard pid filename")
		g.It("should return true if the pid file exists")
	})
}
