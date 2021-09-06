package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import (
	"os"
	"strconv"
)

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
type AssignTaskState int
const (
	// no tasks assigned but some of tasks are in progress
	Retry AssignTaskState = iota
	// do task
	Work
	// all done no more tasks
	Free
)
type AssignTaskArgs struct {
}

type AssignTaskReply struct {
	State AssignTaskState
	Id       int
	Type     TaskType
	Filepath string
	NReduce int
	NMap int
}

type UpdateTaskArgs struct {
	Id    int
	State TaskState
}

type UpdateTaskReply struct {
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
