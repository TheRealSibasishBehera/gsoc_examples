package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/fs"

	"os/exec"
	"path/filepath"
	"regexp"

	"strconv"

	"golang.org/x/sys/unix"
)

var (
	re             = regexp.MustCompile(`libpod-(.*?)\.scope`)
	cgroupPath     = "/sys/fs/cgroup"
	cGroupIDToPath = map[uint64]string{}
	byteOrder      binary.ByteOrder
	unknownPath    string = "unknown"
)

func getPathFromcGroupID(cgroupId uint64) (string, error) {
	// if p, ok := cGroupIDToPath[cgroupId]; ok {
	// 	// fmt.Println(p)
	// 	return p, nil
	// }

	err := filepath.WalkDir(cgroupPath, func(path string, dentry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !dentry.IsDir() {
			return nil
		}
		handle, _, err := unix.NameToHandleAt(unix.AT_FDCWD, path, 0)
		fmt.Println(handle)
		if err != nil {
			return fmt.Errorf("error resolving handle: %v", err)
		}
		// cGroupIDToPath[byteOrder.Uint64(handle.Bytes())] = path
		// if handle == nil {
		// 	return nil
		// }
		fmt.Println(handle.Type())
		fmt.Println(handle.Bytes())
		fmt.Println(path)
		fmt.Println(byteOrder.Uint64(handle.Bytes()))
		return nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to find cgroup id: %v", err)
	}
	if p, ok := cGroupIDToPath[cgroupId]; ok {
		// fmt.Println(p)
		return p, nil
	}

	cGroupIDToPath[cgroupId] = unknownPath
	// fmt.Println(cGroupIDToPath[cgroupId])
	return cGroupIDToPath[cgroupId], nil
}

func getPathFromcGroupID2(cgroupId uint64) (string, error) {
	// find /sys/fs/cgroup -inum $cgroupID

	cmd := exec.Command("/usr/bin/find", "/sys/fs/cgroup",
		"-inum", strconv.Itoa(int(cgroupId)),
	)

	var out bytes.Buffer
	// var stderr bytes.Buffer
	cmd.Stdout = &out
	// cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + "Path not found")
		return unknownPath, err
	}
	fmt.Println("Directory contents : ", out.String())
	return out.String(), nil

}
func main() {

	getPathFromcGroupID(29066)

}
