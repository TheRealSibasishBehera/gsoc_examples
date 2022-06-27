package main

import (
	"context"
	"fmt"
	"syscall"

	"github.com/containers/podman/v3/pkg/bindings"
	"github.com/containers/podman/v3/pkg/bindings/containers"
	"github.com/containers/podman/v3/pkg/domain/entities"

	"log"
)

type PodmanContainerLister struct{}

//type ContainerData struct {
//}

func StartingPodmanSocket() *context.Context {
	fmt.Println("Starting")
	ctx, err := bindings.NewConnection(context.Background(), "unix:/run/podman/podman.sock")
	if err != nil {
		log.Fatalf("cannot connect to podman :%v", err)
	}
	return &ctx
}

func (k *PodmanContainerLister) ListContainers(ctx context.Context) []entities.ListContainer {
	containerList, err := containers.List(ctx, nil)

	if err != nil {
		log.Fatalf("cannot get pods:%v", err)
	}
	return containerList
}

func GetCGroupPathFromContainerID(ctx *context.Context, nameOrID string) string {
	data, err := containers.Inspect(*ctx, nameOrID, nil)
	if err != nil {
		log.Fatalf("cannot get container:%v", err)
	}
	CGroupPath := data.State.CgroupPath
	return CGroupPath
}

func getInodeOfAFile(fileName string) uint64 {
	var stat syscall.Stat_t
	if err := syscall.Stat(fileName, &stat); err != nil {
		panic(err)
	}
	return stat.Ino
}

func GetInodefOfCGroup(CGroupPath string) uint64 {
	ino_val := getInodeOfAFile(CGroupPath)
	return ino_val
}
func main() {
	//fmt.Println(getInodeOfAFile("/machine.slice/libpod-79992c0ba595b9fe3de651f549596b67ee846555186d4300ee0f3e65a4844103.scope"))
	ctx := StartingPodmanSocket()
	conlister := new(PodmanContainerLister)
	conList := conlister.ListContainers(*ctx)
	for _, container := range conList {
		cgroupPath := GetCGroupPathFromContainerID(ctx, container.ID)
		inode_num := GetInodefOfCGroup("/sys/fs/cgroup" + cgroupPath)
		fmt.Printf("\nCGroup : %s image:  %s Time: %s : ID : %s\n\n", cgroupPath, container.Image, container.Created, container.Names, container.ID)
		fmt.Println("Inode Number of cgroup path is ", inode_num)
		//fmt.Println(container.Names)

	}
}
