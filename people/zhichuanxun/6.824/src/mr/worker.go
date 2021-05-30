package mr

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)
import "log"
import "net/rpc"
import "hash/fnv"

nReduce :=  0

//
// Map functions return a slice of KeyValue.
//
type KeyValue struct {
	Key   string
	Value string
}

//
// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
//
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

func getReduceCount() (int, bool) {
	response := ReduceCountResponse{}
	success := call("Coordinator.GetReduceCount", ReduceCountRequest{}, &response)
	return response.ReduceCount, success
}

func getTask() (*TaskResponse, bool) {
	// TODO: create worker ID
	request := TaskRequest{ os.Getpid() }
	response := TaskResponse{}
	success := call("Coordinator.GetTask", &request, &response)
	return &response, success
}

func finishTask(taskId int, taskType TaskType) (bool, bool) {
	request := TaskStatusRequest{taskId, taskType, os.Getpid() }
	response := TaskStatusResponse{}
	success = call("Coordinator.FinishTask", &request, &response)
	return response.Done, success
}

func doMap(taskId int, file string, mapf func(string, string) []KeyValue) {
	f, error := os.Open(file)
	if error != nil {
		fmt.Println("Failed to open file %v", file)
		return
	}

	bytes, error := ioutil.ReadAll(f)
	if error != nil {
		fmt.Println("Failed to read file %v", file)
		return
	}

	kvs := mapf(file, string(bytes))

	// TODO: write intermediate results

}

//
// main/mrworker.go calls this function.
//
func Worker(mapf func(string, string) []KeyValue, reducef func(string, []string) string) {

	// Your worker implementation here.
	_, success := getReduceCount()
	if !success {
		fmt.Println("Failed to get the count of reduce tasks, quiting ...")
		return
	}

	// uncomment to send the Example RPC to the coordinator.
	// CallExample()
	for {
		response, success := getTask()
		if !success {
			fmt.Println("Failed to get task, quiting ...")
			return
		}
		if response.TaskType == MapTask {
			doMap(response.TaskId, response.TaskFile, mapf)
			status, success = finishTask(response.TaskId, MapTask)
		} else if response.TaskType == ReduceTask {
			doReduce(response.TaskId, reducef)
			status, success = finishTask(response.TaskId, ReduceTask)
		}

		if status != Done || !success {
			fmt.Println("Faild to update task status, quiting")
			return
		}

		time.Sleep(time.Second)
	}
}

//
// example function to show how to make an RPC call to the coordinator.
//
// the RPC argument and reply types are defined in rpc.go.
//
func CallExample() {

	// declare an argument structure.
	args := ExampleArgs{}

	// fill in the argument(s).
	args.X = 99

	// declare a reply structure.
	reply := ExampleReply{}

	// send the RPC request, wait for the reply.
	call("Coordinator.Example", &args, &reply)

	// reply.Y should be 100.
	fmt.Printf("reply.Y %v\n", reply.Y)
}

//
// send an RPC request to the coordinator, wait for the response.
// usually returns true.
// returns false if something goes wrong.
//
func call(rpcname string, args interface{}, reply interface{}) bool {
	// c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	sockname := coordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}
