package main

import (
	"github.com/docker/docker/client"
	"log"
	"golang.org/x/net/context"
	"github.com/docker/docker/api/types"
)

var cli *client.Client

func NewClient() {
	var err error

	if cli, err = client.NewEnvClient(); err != nil {
		log.Println("Docker client error", err)
	}
}

func controllers() ([]types.Container) {
	var err error
	var containers []types.Container
	if containers, err = cli.ContainerList(context.Background(), types.ContainerListOptions{}); err != nil {
		log.Println("Docker client error", err)
	}

	return containers
}