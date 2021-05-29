package mr

import (
	"log"
	"sync"
)
import "net"
import "os"
import "net/rpc"
import "net/http"

type TaskType int
type TaskStatus int

// https://www.educative.io/edpresso/what-is-an-enum-in-golang
const (
	MapTask TaskType = iota
	ReduceTask
	None
	Exit
)

const (
	New TaskStatus = iota
	InProgress
	Done
)

type Task struct {
	Id int
	Type TaskType
	File string
	WorkerId int
	Status TaskStatus
}

type Coordinator struct {
	// Your definitions here.
	mutex sync.Mutex
	mapTasks []Task
	reduceTasks []Task
}

func (c *Coordinator) SelectTask(workerId int) *Task {
	var tasks []Task
	if len(c.mapTasks) > 0 {
		tasks = c.mapTasks
	} else {
		tasks = c.reduceTasks
	}
	for _, task := range tasks {
		if task.Status == New {
			task.Status = InProgress
			return &task
		}
	}
	return &Task { -1, None, "", -1, Done }
}

// Your code here -- RPC handlers for the worker to call.

//
// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
//
func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}

func (c *Coordinator) TaskRequestHandler(request *TaskRequest, response *TaskResponse) error {
	c.mutex.Lock()
	var task Task = c.SelectTask()
	task.WorkerId = request.WorkerId;

	response.TaskId = task.Id
	response.TaskType = task.Type
	response.TaskFile = task.File

	c.mutex.Unlock()
	// TODO: thread
	return nil
}

// cp people/RefinedCoding/6.824/src/mr/rpc.go homework/0528/RefinedCoding-rpc.go
// cp people/RefinedCoding/6.824/src/mr/coordinator.go homework/0529/RefinedCoding-coordinator.go

//
// start a thread that listens for RPCs from worker.go
//
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
//
func (c *Coordinator) Done() bool {
	ret := false

	// Your code here.


	return ret
}

//
// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
//
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{}

	// Your code here.


	c.server()
	return &c
}
