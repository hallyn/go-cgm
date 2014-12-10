package main

import (
	"fmt"
	"os"
	"github.com/hallyn/go-cgm"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// go-cgm ping
// go-cgm gettasks freezer lxc/t1  # show pids
// go-cgm controllers # list controllers
// go-cgm ls freezer lxc/t1  # list files in such a cgroup
// go-cgm cat freezer lxc/t1 freezer.state
// go-cgm set freezer lxc/t1 freezer.state FROZEn
// go-cgm move freezer lxc/t1 pid1

func usage(cmd *string) {
	if cmd == nil {
		fmt.Println("cat <controller> <cgroup> <file>")
		fmt.Println("controllers")
		fmt.Println("gettasks <controller> <cgroup>")
		fmt.Println("ls <controller> <cgroup>")
		fmt.Println("move <controller> <cgroup> <pid>")
		fmt.Println("ping")
		fmt.Println("set <controller> <cgroup> <file> <new-value>")
		os.Exit(0)
	}
	switch (*cmd) {
	case "gettasks":
		fmt.Println("gettasks <controller> <cgroup>")
		os.Exit(1)
	case "ls":
		fmt.Println("ls <controller> <cgroup> <directory>")
		os.Exit(1)
	case "controller":
		fmt.Println("controllers")
		os.Exit(1)
	}
	fmt.Println("Unknown command: ", *cmd)
	os.Exit(1)
}

func run() error {
	if len(os.Args) < 2 || os.Args[1] == "help" {
		usage(nil)
	}

	switch os.Args[1] {
	case "ping":
		err := cgm.Ping()
		if err != nil {
			fmt.Println("Error calling ping: ", err)
			os.Exit(1)
		}
		fmt.Println("Ping succeeded")
		os.Exit(0)

	case "gettasks":
		if len(os.Args) < 4 {
			usage(&os.Args[1])
			os.Exit(1)
		}
		l, err := cgm.Gettasks(os.Args[2], os.Args[3])
		if err != nil {
			fmt.Println("Error calling gettasks: ", err)
			os.Exit(1)
		}
		for _, v := range *l {
			fmt.Println(v)
		}
		os.Exit(0)

	case "ls":
		if len(os.Args) < 4 {
			usage(&os.Args[1])
			os.Exit(1)
		}
		l, err := cgm.Ls(os.Args[2], os.Args[3])
		if err != nil {
			fmt.Println("Error calling ls: ", err)
			os.Exit(1)
		}
		var v cgm.Cgmfile
		for _, v = range *l {
			fmt.Println(v.Name)
		}
		os.Exit(0)

	case "controllers":
		l, err := cgm.ListControllers()
		if err != nil {
			fmt.Println("Error calling ListControllers: ", err)
			os.Exit(1)
		}
		for _, c := range *l {
			fmt.Println(c)
		}
		os.Exit(0)
	}

	fmt.Println("Not yet implemented")
	usage(nil)
	return nil
}
