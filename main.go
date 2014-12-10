package cgm

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// go-cgm ping
// go-cgm gettasks freezer lxc/t1  # show pids
// go-cgm cat freezer lxc/t1 freezer.state
// go-cgm set freezer lxc/t1 freezer.state FROZEn
// go-cgm move freezer lxc/t1 pid1

func usage() {
	fmt.Println("Usage:")
	fmt.Println("ping")
	fmt.Println("gettasks <controller> <cgroup>")
	fmt.Println("cat <controller> <cgroup> <file>")
	fmt.Println("set <controller> <cgroup> <file> <new-value>")
	fmt.Println("move <controller> <cgroup> <pid>")
	os.Exit(0)
}

func run() error {
	if len(os.Args) < 2 || os.Args[1] == "help" {
		usage()
	}

	if os.Args[1] == "ping" {
		err := cgm_ping()
		if err != nil {
			fmt.Println("Error calling ping: ", err)
			os.Exit(1)
		}
		fmt.Println("Ping succeeded")
		os.Exit(0)
	}

	if os.Args[1] == "gettasks" {
		if len(os.Args) < 4 {
			usage()
		}
		l, err := cgm_gettasks(os.Args[2], os.Args[3])
		if err != nil {
			fmt.Println("Error calling gettasks: ", err)
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
