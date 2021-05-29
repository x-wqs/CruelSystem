package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import "os"
import "strconv"

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

// Add your RPC definitions here.

type TaskRequest struct {
	WorkerId int
}

type TaskResponse struct {
	TaskId int
	TaskType TaskType
	TaskFile string
}

type TaskStatusRequest struct {
	TaskId int
	TaskType TaskType
	WorkerId int
}

type TaskStatusResponse struct {
	Done bool
}

type ReduceCountRequest struct {
}

type ReduceCountResponse struct {
	ReduceCount int
}

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/824-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
