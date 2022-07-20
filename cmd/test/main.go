package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/podman/v3/pkg/bindings"
	"github.com/containers/podman/v3/pkg/bindings/containers"
)

func main() {
	conn, err := bindings.NewConnection(context.Background(), "unix://run/USERID#/podman/podman.sock")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	containerList, err := containers.List(conn, new(containers.ListOptions))
	for _, container := range containerList {
		println(container.ID)
	}
	//inspectData, err := containers.Inspect(conn, "a054b2438e2948812aba97533603bc57cd7b4bd9c03498c0a2abdde603993698", new(containers.InspectOptions).WithSize(true))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Print the container ID
	//fmt.Println(inspectData.ID)
}
