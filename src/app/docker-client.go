//
// docker-client thin wrapper with convenience methods for querying docker info, images, containers, networks, etc
//
// @author darryl.west <darwest@ebay.com>
// @created 2017-09-01 08:24:30
//

package app

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

// DockerClient thin wrapper around docker's client
type DockerClient struct {
	*client.Client
	cfg *Config
}

// ContainerConfig contains container, host, network configs and name
type ContainerConfig struct {
	ContainerConfig *container.Config
	HostConfig      *container.HostConfig
	NetworkConfig   *network.NetworkingConfig
	Name            string
}

// NewDockerClient returns an instance of docker client
func NewDockerClient(config *Config) (*DockerClient, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Error("%s", err)
		return nil, err
	}

	docker := DockerClient{cli, config}

	return &docker, nil
}

// GetContainers - return the list of all containers with their state
func (cli DockerClient) GetContainers() ([]types.Container, error) {
	opts := types.ContainerListOptions{}
	opts.All = true
	containers, err := cli.ContainerList(context.Background(), opts)
	if err != nil {
		log.Error("container error %v", err)
	}

	log.Debug("total container count: %d", len(containers))

	return containers, err
}

// InspectContainer returns the json container
func (cli DockerClient) InspectContainer(id string) types.ContainerJSON {
	json, err := cli.ContainerInspect(context.Background(), id)
	if err != nil {
		panic(err)
	}

	return json
}

// StartContainer start a container
func (cli DockerClient) StartContainer(container types.Container) error {
	id := container.ID
	opts := types.ContainerStartOptions{}

	return cli.ContainerStart(context.Background(), id, opts)
}

// RestartContainer restart the container with ouptional timeout duration
func (cli DockerClient) RestartContainer(container types.Container, timeout *time.Duration) error {
	id := container.ID

	return cli.ContainerRestart(context.Background(), id, timeout)
}

// CreateContainer creates a new container and returns:
// func (cli DockerClient) CreateContainer(conf ContainerConfig) (types.Container, error) {
// }

// CreateContainerConfig returns a standard container configuration
func (cli DockerClient) CreateContainerConfig(image, name string) ContainerConfig {
	conf := ContainerConfig{
		ContainerConfig: &container.Config{Image: image},
		HostConfig:      nil,
		NetworkConfig:   nil,
		Name:            name,
	}

	return conf
}

// GetNetworks return all the networks
func (cli DockerClient) GetNetworks() ([]types.NetworkResource, error) {
	opts := types.NetworkListOptions{}
	networks, err := cli.NetworkList(context.Background(), opts)
	if err != nil {
		log.Error("network %v", err)
	}

	log.Info("networks conut: %d", len(networks))

	return networks, err
}

// InspectNetwork return the json data
func (cli DockerClient) InspectNetwork(id string) types.NetworkResource {
	opts := types.NetworkInspectOptions{}
	network, err := cli.NetworkInspect(context.Background(), id, opts)
	if err != nil {
		log.Error("network %v", err)
	}

	return network
}

// ShowNetworks send network configurations to designated output
func (cli DockerClient) ShowNetworks() {
	networks, err := cli.GetNetworks()
	if err != nil {
		panic(err)
	}

	for _, network := range networks {
		fmt.Printf("%s %s %s\n", network.ID[:10], network.Name, network.Driver)
		fmt.Printf("%v\n\n", network)
		data := cli.InspectNetwork(network.ID)
		fmt.Printf("%v\n", data)
	}
}

// ShowContainers send container defs to specified output
func (cli DockerClient) ShowContainers() {
	containers, err := cli.GetContainers()
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s %s\n", container.ID[:10], container.State, container.Names[0])
		data := cli.InspectContainer(container.ID)
		fmt.Printf("created: %s\n", data.Created)
		fmt.Printf("path   : %s\n", data.Path)
		fmt.Printf("network: %v\n", data.NetworkSettings)
		fmt.Printf("args   : %v\n", data.Args)
		fmt.Printf("state  : %v\n", data.State)
		fmt.Printf("image  : %v\n", data.Image)
		fmt.Printf("name   : %v\n", data.Name)
		fmt.Printf("restart: %v\n", data.RestartCount)

		fmt.Println("")
	}
}

// RunContainer execute the container run command and return the id
func (cli DockerClient) RunContainer(cmd *exec.Cmd) (string, error) {

	out, err := cmd.Output()
	if err != nil {
		log.Error("command error: %v", err)
		return "", err
	}

	id := strings.Split(string(out), "\n")[0]
	log.Info("created container: %s", id)

	return id, nil
}
