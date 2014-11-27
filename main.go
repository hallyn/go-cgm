package main

import (
	"fmt"
	"os"

	"github.com/guelfey/go.dbus"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// go-cgm ping
// go-cgm ls freezer lxc/t1  # show pids
// go-cgm cat freezer lxc/t1 freezer.state
// go-cgm set freezer lxc/t1 freezer.state FROZEn
// go-cgm move freezer lxc/t1 pid1

func usage() {
	fmt.Println("Usage:")
	fmt.Println("ping")
	fmt.Println("ls <controller> <cgroup>")
	fmt.Println("cat <controller> <cgroup> <file>")
	fmt.Println("set <controller> <cgroup> <file> <new-value>")
	fmt.Println("move <controller> <cgroup> <pid>")
	os.Exit(0)
}

func do_ping() error {
	c, err := dbus.Dial("unix:path=/sys/fs/cgroup/cgmanager/sock")
	if err != nil {
		fmt.Println("Error connecting to cgmanager: ", err)
		return err
	}
	obj := c.Object("org.linuxcontainers.cgmanager0_0", "/org/linuxcontainers/cgmanager")
	call := obj.Call("org.linuxcontainers.cgmanager0_0.Ping", 0, int32(1))
	if call.Err != nil {
		return call.Err
	}
	return nil
}

func do_ls(controller string, cgroup string) (*[]int32, error) {
	c, err := dbus.Dial("unix:path=/sys/fs/cgroup/cgmanager/sock")
	if err != nil {
		fmt.Println("Error connecting to cgmanager: ", err)
		return nil, err
	}
	fmt.Println("1")
	obj := c.Object("org.linuxcontainers.cgmanager0_0", "/org/linuxcontainers/cgmanager")
	fmt.Println("2, obj is ", obj)
	call := obj.Call("org.linuxcontainers.cgmanager0_0.GetTasks", 0, controller, cgroup)
	fmt.Println("3")
	if call.Err != nil {
		return nil, call.Err
	}
	var l *[]int32
	err = call.Store(&l)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func run() error {
	if len(os.Args) < 2 || os.Args[1] == "help" {
		usage()
	}

	if os.Args[1] == "ping" {
		err := do_ping()
		if err != nil {
			fmt.Println("Error calling ping: ", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if os.Args[1] == "ls" {
		if len(os.Args) < 4 {
			usage()
		}
		l, err := do_ls(os.Args[2], os.Args[3])
		if err != nil {
			fmt.Println("Error calling ls: ", err)
			os.Exit(1)
		}
		for _, v := range *l {
			fmt.Println(v)
		}
		os.Exit(0)
	}

	fmt.Println("Not yet implemented")
	usage()
	return nil
}