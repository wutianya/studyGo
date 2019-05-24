/*
package main

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/daemon/cluster/executor/container"
	"cmd/go/internal/str"
	"github.com/docker/docker/api/server/router/container"
	"context"
	"fmt"



	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli ,err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	for _, container := range containers {
		fmt.Printf("%s %s %s\n", container.ID[:12], container.Image, container.State)
	}
}
*/

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	dockerHosts    string = "tcp://192.168.0.253:2376"
	dockerAPI      string = "v1.27"
	defaultHeadres        = map[string]string{"User-Agent": "engine-cli-1.0"}
)

func main() {
	container, err := GetContainersInfo()
	if err != nil {
		panic(err)
	}
	cli, _ := client.NewClient(dockerHosts, dockerAPI, nil, defaultHeadres)
	body, _ := cli.ContainerStats(context.Background(),container.ID,false)
	defer body.Body.Close()
	content ,err := ioutil.ReadAll(body.Body)
	// fmt.Println(string(context))
	stats := types.StatsJSON{}
	err = json.Unmarshal([]byte(content), &stats) 
	fmt.Println(container.ID[:12],float32(stats.MemoryStats.Usage) / 1024 /1024 )
}

func GetContainersInfo() (types.Container, error) {
	var container types.Container
	cli, err := client.NewClient(dockerHosts, dockerAPI, nil, defaultHeadres)
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	for _, v := range containers {
		container = v
	}
	return container, err
}
