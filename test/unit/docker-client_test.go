//
// docker client -
//
// @author darryl.west@ebay.com
// @created 2017-09-01 12:57:59
//

package unit

import (
	"app"
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func TestDockerClient(t *testing.T) {
	g := Goblin(t)

	g.Describe("DockerClient", func() {
		cfg := new(app.Config)

		g.It("should create a docker client struct", func() {
			client, err := app.NewDockerClient(cfg)
			g.Assert(err).Equal(nil)
			g.Assert(fmt.Sprintf("%T", client)).Equal("*app.DockerClient")
		})

		g.It("should initialize and start the service", func() {
			client, err := app.NewDockerClient(cfg)
			g.Assert(err).Equal(nil)
			g.Assert(fmt.Sprintf("%T", client)).Equal("*app.DockerClient")
		})

		g.It("should create a container configuration", func() {
			client, _ := app.NewDockerClient(cfg)
			conf := client.CreateContainerConfig("ebay-local/alpine-envoy", "router-1")
			g.Assert(conf.Name).Equal("router-1")
			g.Assert(conf.HostConfig == nil).IsTrue()
			g.Assert(conf.NetworkConfig == nil).IsTrue()
			// fmt.Println(*conf.ContainerConfig)
		})
	})
}
