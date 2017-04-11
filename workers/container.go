package workers

// Docker Engine API
import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	dockerapi "github.com/docker/docker/client"
)

var cli *dockerapi.Client
var err error

func init() {
	cli, err = dockerapi.NewEnvClient()

	if err != nil {
		fmt.Print("error")
	}

}

// ListContainers Display running containers
func ListContainers(running bool) {
	options := types.ContainerListOptions{All: running}
	containers, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		panic(err)
	}

	if len(containers) == 0 {
		fmt.Println("No containers available.")
	}
	for _, c := range containers {
		fmt.Println(c.Names)
	}
}
