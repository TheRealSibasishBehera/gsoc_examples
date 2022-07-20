//package main
//
//import (
//	"context"
//	"fmt"
//	"github.com/containers/podman/v3/pkg/bindings"
//	"github.com/containers/podman/v3/pkg/bindings/containers"
//	"os"
//)
//
//func main() {
//	//fmt.Println("Starting")
//	//ctx, err := bindings.NewConnection(context.Background(), "unix:/run/podman/podman.sock")
//	//
//	//if err != nil {
//	//	log.Fatalf("cannot connect to podman :%v", err)
//	//}
//
//	sock_dir := os.Getenv("XDG_RUNTIME_DIR")
//	if sock_dir == "" {
//		sock_dir = "/var/run"
//	}
//	socket := "unix:" + sock_dir + "/podman/podman.sock"
//
//	// Connect to Podman socket
//	ctx, err := bindings.NewConnection(context.Background(), socket)
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	fmt.Println(ctx)
//
//	//podList, err := pods.List(ctx, nil)
//	//
//	//if err != nil {
//	//	log.Fatalf("cannot get pods:%v", err)
//	//}
//	//for _, pod := range podList {
//	//	for _, container := range pod.Containers {
//	//		data, err := containers.Inspect(ctx, container.Id, nil)
//	//
//	//		if err != nil {
//	//			log.Fatalf("cannot get container details:%v", err)
//	//		}
//	//		fmt.Println(data)
//	//		fmt.Printf("Container Cgroup path %s \n", data.State.CgroupPath)
//	//	}
//	//	fmt.Printf("Pod %+v\n", pod)
//	//}
//
//	//var latestContainers = 1
//	//containerLatestList, err := containers.List(ctx, &containers.ListOptions{
//	//	Last: &latestContainers,
//	//})
//	//if err != nil {
//	//	fmt.Println(err)
//	//	os.Exit(1)
//	//}
//	//fmt.Printf("Latest container is %s\n", containerLatestList[0].Names[0])
//	//
//	// Container inspect
//	ctrData, err := containers.Inspect(ctx, "0c302956e51a", &containers.InspectOptions{})
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	fmt.Printf("Container uses image %s\n", ctrData.ImageName)
//	fmt.Printf("Container running status is %s\n", ctrData.State.Status)
//
//}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/containers/podman/v3/pkg/bindings"
	"github.com/containers/podman/v3/pkg/bindings/containers"
)

func main() {
	fmt.Println("Starting")
	ctx, err := bindings.NewConnection(context.Background(), "unix:/run/podman/podman.sock")

	if err != nil {
		log.Fatalf("cannot connect to podman :%v", err)
	}

	fmt.Println(ctx)

	//podList, err := pods.List(ctx, nil)
	containerList, err := containers.List(ctx, nil)

	if err != nil {
		log.Fatalf("cannot get pods:%v", err)
	}
	//for _, pod := range podList {
	//	for _, container := range pod.Containers {
	//		data, err := containers.Inspect(ctx, container.Id, nil)
	//
	//		if err != nil {
	//			log.Fatalf("cannot get container details:%v", err)
	//		}
	//fmt.Printf("Container %+v\n", data)
	//fmt.Println()
	//fmt.Println()
	//fmt.Println()
	//fmt.Println(data.State.CgroupPath)
	//fmt.Println()
	//fmt.Println()
	//fmt.Println()
	//	}
	//	fmt.Printf("Pod %+v\n", pod)
	//}

	for _, container := range containerList {
		data, err := containers.Inspect(ctx, container.ID, nil)
		fmt.Println(data)

		if err != nil {
			log.Fatalf("cannot get container:%v", err)
		}
		fmt.Printf("Container %+v\n", data)
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println(data.State.CgroupPath)
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}

}
