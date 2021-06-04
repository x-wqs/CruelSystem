package mr

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"
	"time"
)

const TIMEOUT = 10

type TaskType int

const (
	MapTask TaskType = iota
	ReduceTask
)

type TaskState int

const (
	UnAssigned TaskState = iota
	InProgress
	Done
	Error
)

type Task struct {
	// simplify id by int rather than uuid
	Id int
	Type TaskType
	State TaskState
	Filepath string
}

type Coordinator struct {
	// Your definitions here.
	
	// 
	tasks []Task
	mu sync.Mutex
	done int
	nMap int
	nReduce int
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

func (c *Coordinator) CheckTimeoutAndError(Id int) {
	time.AfterFunc(TIMEOUT*time.Second, func() {
		c.mu.Lock()
		if c.tasks[Id].State == InProgress {
			c.tasks[Id].State = UnAssigned
		} else if c.tasks[Id].State == Error {
			// e.g. open file
			errors.New("error")
		}
		c.mu.Unlock()
	})
}

func (c *Coordinator) AssignTask(args *AssignTaskArgs, reply *AssignTaskReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	st, ed := -1, -1
	// reduce tasks are blocked by map tasks
	// TODO: non-block implementation
	if c.done < c.nMap {
		st = 0
		ed = c.nMap
	} else {
		st = c.nMap
		ed = c.nMap + c.nReduce
	}
	for Id := st; Id < ed; Id++ {
		if c.tasks[Id].State == UnAssigned {
			// t := ""
			// if c.tasks[Id].Type == MapTask {
			// 	t = "Map"
			// } else {
			// 	t = "Reduce"
			// }
			// fmt.Printf("%vtask%v assigned", t, Id)
			reply.State = Work
			reply.Id = Id
			reply.Type = c.tasks[Id].Type
			reply.Filepath = c.tasks[Id].Filepath
			// TODO: merge nreduce and nmap into one rpc
			reply.NReduce = c.nReduce
			reply.NMap = c.nMap

			c.tasks[Id].State = InProgress
			go c.CheckTimeoutAndError(Id)
			return nil
		}
	}
	if c.done < len(c.tasks) {
		// case 1:
		// map tasks have not been finished yet
		// ask reduce worker to retry
		// case 2:
		// when some tasks still in progress 
		// which might timeout after a while, 
		// ask worker to retry later
		reply.State = Retry
	} else {
		reply.State = Free
	}
	return nil
}

func (c *Coordinator) UpdateTaskState(args *UpdateTaskArgs, reply *UpdateTaskReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.tasks[args.Id].State = args.State
	if args.State == Done {
		c.done++
	}
	return nil
}


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
	c.mu.Lock()
	defer c.mu.Unlock()
	ret = (c.done == len(c.tasks))
	// fmt.Println("done")
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
	nMap := len(files)
	c.tasks = make([]Task, nMap + nReduce)
	c.done = 0
	c.nMap = nMap
	c.nReduce = nReduce

	for i := 0; i < nMap; i++ {
		c.tasks[i] = Task {
			Id: i,
			Type: MapTask,
			State: UnAssigned,
			Filepath: files[i],
		}
	}

	for i := 0; i < nReduce; i++ {
		c.tasks[nMap + i] = Task {
			Id: nMap + i,
			Type: ReduceTask,
			State: UnAssigned,
			Filepath: "",
		}
	}

	c.server()
	return &c
}
