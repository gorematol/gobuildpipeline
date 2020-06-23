package utils

import (

	"context"
	"io"
	"os"
        "fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	docker = initClient() 
        ctx = context.Background()
)


func initClient() (cli *client.Client) {

    	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    	if err != nil {
        	panic(err)
    	}
	return cli
}
 

func PullImage(imagename string) {

    fmt.Println("Pulling latest Jenkins image from docker hub")
    reader, err := docker.ImagePull(ctx, imagename, types.ImagePullOptions{})
    if err != nil {
        panic(err)
    }
    io.Copy(os.Stdout, reader)
}

//func RunContainer(contid string) {
