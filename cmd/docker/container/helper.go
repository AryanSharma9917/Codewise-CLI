package container

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func dockerClient() (context.Context, *client.Client) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	checkErr(err)
	defer cli.Close()

	return ctx, cli
}

// runningContainerList returns a list of rumming containers
func runnnigContainerList(cli *client.Client, ctx context.Context) []types.Container {
	containerList, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	checkErr(err)
	return containerList
}
