package cgm

import (
	"github.com/guelfey/go.dbus"
)

func Ping() error {
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

func GetChildren(controller string, cgroup string) (*[]string, error) {
	c, err := dbus.Dial("unix:path=/sys/fs/cgroup/cgmanager/sock")
	if err != nil {
		return nil, err
	}
	err = c.Auth(nil)
	if err != nil {
		return nil, err
	}
	obj := c.Object("org.linuxcontainers.cgmanager0_0", "/org/linuxcontainers/cgmanager")
	call := obj.Call("org.linuxcontainers.cgmanager0_0.ListChildren", 0, controller, cgroup)
	if call.Err != nil {
		return nil, call.Err
	}
	var l []string
	err = call.Store(&l)
	if err != nil {
		return nil, err
	}
	return &l, nil
}

func Gettasks(controller string, cgroup string) (*[]int32, error) {
	c, err := dbus.Dial("unix:path=/sys/fs/cgroup/cgmanager/sock")
	if err != nil {
		return nil, err
	}
	err = c.Auth(nil)
	if err != nil {
		return nil, err
	}
	obj := c.Object("org.linuxcontainers.cgmanager0_0", "/org/linuxcontainers/cgmanager")
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

type Cgmfile struct {
	Name string
	Owner uint32
	Group uint32
	Perms uint32
}

func Ls(controller, cgroup string) (*[]Cgmfile, error) {
	c, err := dbus.Dial("unix:path=/sys/fs/cgroup/cgmanager/sock")
	if err != nil {
		return nil, err
	}
	err = c.Auth(nil)
	if err != nil {
		return nil, err
	}
	obj := c.Object("org.linuxcontainers.cgmanager0_0", "/org/linuxcontainers/cgmanager")
	call := obj.Call("org.linuxcontainers.cgmanager0_0.ListKeys", 0, controller, cgroup)
	if call.Err != nil {
		return nil, call.Err
	}
	var l []Cgmfile
	err = call.Store(&l)
	if err != nil {
		return nil, err
	}
	return &l, nil
}

func ListControllers() (*[]string, error) {
	c, err := dbus.Dial("unix:path=/sys/fs/cgroup/cgmanager/sock")
	if err != nil {
		return nil, err
	}
	err = c.Auth(nil)
	if err != nil {
		return nil, err
	}
	obj := c.Object("org.linuxcontainers.cgmanager0_0", "/org/linuxcontainers/cgmanager")
	call := obj.Call("org.linuxcontainers.cgmanager0_0.ListControllers", 0)
	if call.Err != nil {
		return nil, call.Err
	}
	var l []string
	err = call.Store(&l)
	if err != nil {
		return nil, err
	}
	return &l, nil
}

func Cat(controller, cgroup, file string) (*string, error) {
	c, err := dbus.Dial("unix:path=/sys/fs/cgroup/cgmanager/sock")
	if err != nil {
		return nil, err
	}
	err = c.Auth(nil)
	if err != nil {
		return nil, err
	}
	obj := c.Object("org.linuxcontainers.cgmanager0_0", "/org/linuxcontainers/cgmanager")
	call := obj.Call("org.linuxcontainers.cgmanager0_0.GetValue", 0, controller, cgroup, file)
	if call.Err != nil {
		return nil, call.Err
	}
	var l string
	err = call.Store(&l)
	if err != nil {
		return nil, err
	}
	return &l, nil
}

func Set(controller, cgroup, file, value string) error {
	c, err := dbus.Dial("unix:path=/sys/fs/cgroup/cgmanager/sock")
	if err != nil {
		return err
	}
	err = c.Auth(nil)
	if err != nil {
		return err
	}
	obj := c.Object("org.linuxcontainers.cgmanager0_0", "/org/linuxcontainers/cgmanager")
	call := obj.Call("org.linuxcontainers.cgmanager0_0.SetValue", 0, controller, cgroup, file, value)
	if call.Err != nil {
		return call.Err
	}
	return nil
}
