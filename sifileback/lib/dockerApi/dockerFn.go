package dockerApi

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type Kod struct {
	Image string // "xaljer/kodexplorer"
	DockerPort int //80
	HostPort int
	DockerDir string //"/var/www/html"
	HostDir string

}


func StopContainer(id string)  error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	if err := cli.ContainerStop(ctx, id, nil); err != nil {
		return err
	}
	return nil
}

func RunContainer(k Kod) (id string, err error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}
	_, err = cli.ImagePull(ctx, k.Image, types.ImagePullOptions{})
	if err != nil {
		return "", err
	}
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port(fmt.Sprintf("%d/tcp",k.DockerPort)): []nat.PortBinding{
				{
					HostIP: "0.0.0.0",
					HostPort: fmt.Sprintf("%d",k.HostPort),
				},
			},
		},
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: k.HostDir,
				Target: k.DockerDir,
			},
		},
	}
	resp, err :=
		cli.ContainerCreate(ctx, &container.Config{
			Image:  k.Image,
			Tty:     false,
			ExposedPorts: nat.PortSet{
				nat.Port(fmt.Sprintf("%d/tcp",k.DockerPort)): struct{}{},
			},
		}, hostConfig, nil, "")
	if err != nil {
		return "", err
	}
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}
	return resp.ID,nil
}