package utils

import (

	"context"
	"io"
	"os"
        "fmt"

	"github.com/docker/docker/api/types"
//	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
//	"github.com/docker/docker/pkg/stdcopy"
)

func PullImage(imagename string) {

    ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        panic(err)
    }

    fmt.Println("Pulling latest Jenkins image from docker hub")
    reader, err := cli.ImagePull(ctx, imagename, types.ImagePullOptions{})
    if err != nil {
        panic(err)
    }
    io.Copy(os.Stdout, reader)
}
