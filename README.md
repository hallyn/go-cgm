# go-cgm

This is a go library for interacting with the cgmanager.  It
supports the following calls:

* * *
Function         | Description
:------          | :----------
Cat              | Get the contents of a cgroup file
Create           | Create a cgroup
GetChildren      | Get list of child cgroups
GetTasks         | List tasks in a cgroup
ListControllers  | List mounted controllers
Ls               | List files offered by a controller
MovePid          | Move a process into a cgroup
Ping             | Ping the cgmanager
Remove           | Remove a cgroup
Set              | Set the value of a cgroup file
* * *

**Arguemnts**

    Cat controller cgroup file
    Create controller cgroup
    GetChildren controller cgroup
    GetTasks controller cgroup
    ListControllers
    Ls controller cgroup
    MovePid controller new-cgroup pid
    Ping
    Remove controller cgroup
    Set controller cgroup file new-value
