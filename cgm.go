package cgm

import (
	"fmt"

	"github.com/guelfey/go.dbus"
)

func cgm_ping() error {
	c, err := dbus.Dial("unix:path=/sys/fs/cgroup/cgmanager/sock")
	if err != nil {
		return err
	}
	err = c.Auth(nil)
	if err != nil {
		return err
	}
	obj := c.Object("org.linuxcontainers.cgmanager0_0", "/org/linuxcontainers/cgmanager")
	call := obj.Call("org.linuxcontainers.cgmanager0_0.Ping", 0, int32(1))
	if call.Err != nil {
		return call.Err
	}
	return nil
}

func cgm_gettasks(controller string, cgroup string) (*[]int32, error) {
	c, err := dbus.Dial("unix:path=/sys/fs/cgroup/cgmanager/sock")
	if err != nil {
		return nil, err
	}
	err = c.Auth(nil)
	if err != nil {
		return nil, err
	}
	obj := c.Object("org.linuxcontainers.cgmanager0_0", "/org/linuxcontainers/cgmanager")
	fmt.Println("2, obj is ", obj)
	call := obj.Call("org.linuxcontainers.cgmanager0_0.GetTasks", 0, controller, cgroup)
	if call.Err != nil {
		return nil, call.Err
	}
	var l []int32
	err = call.Store(&l)
	if err != nil {
		return nil, err
	}
	return &l, nil
}

