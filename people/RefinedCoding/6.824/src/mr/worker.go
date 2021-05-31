package mr

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"
)
import "log"
import "net/rpc"
import "hash/fnv"

// KeyValue
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
	success := call("Coordinator.FinishTask", &request, &response)
	return response.Done, success
}

func doMap(mapId int, file string, nReduce int, mapf func(string, string) []KeyValue) {
	f, e := os.Open(file)
	if e != nil {
		fmt.Printf("Failed to open file %v\n", file)
		return
	}

	bytes, e := ioutil.ReadAll(f)
	if e != nil {
		fmt.Printf("Failed to read file %v\n", file)
		return
	}

	kvs := mapf(file, string(bytes))
	files := make([]*os.File, 0, nReduce)
	buffers := make([]*bufio.Writer, 0, nReduce)
	encoders := make([]*json.Encoder, 0, nReduce)

	// write intermediate results
	for i := 0; i < nReduce; i ++ {
		// TODO: temp path
		path := fmt.Sprintf("map-%v-%v-%v", mapId, i, os.Getpid())
		f, e := os.Create(path)
		if e != nil {
			fmt.Printf("Failed to create file %v\n", path)
		}
		buffer := bufio.NewWriter(f)
		files = append(files, f)
		buffers = append(buffers, buffer)
		encoders = append(encoders, json.NewEncoder(buffer))
	}

	// split map result to nReduce intermediate files
	for _, kv := range kvs {
		index := ihash(kv.Key) % nReduce
		e := encoders[index].Encode(&kv)
		if e != nil {
			fmt.Printf("Failed to encode %v\n", &kv)
		}
	}

	for i, buffer := range buffers {
		e := buffer.Flush()
		if e != nil {
			fmt.Printf("Failed to flush buffer for file %v\n", files[i].Name())
		}
	}

	for i, file := range files {
		file.Close()
		e := os.Rename(file.Name(), fmt.Sprintf("map-%v-%v", mapId, i))
		if e != nil {
			fmt.Printf("Failed to rename file %v\n", file.Name())
		}
	}
}

func doReduce(reduceId int, reducef func(string, []string) string) {
	files, e := filepath.Glob(fmt.Sprintf("map-*-%v", reduceId))
	if e != nil {
		fmt.Println("Failed to list map results")
	}

	kvs := make(map[string][]string)
	for _, path := range files {
		f, e := os.Open(path)
		if e != nil {
			fmt.Printf("Failed to open file %v\n", path)
		}

		decode := json.NewDecoder(f)
		for decode.More() {
			var kv KeyValue
			e = decode.Decode(&kv)
			kvs[kv.Key] = append(kvs[kv.Key], kv.Value)
		}
	}

	// write reduce result
	keys := make([]string, 0, len(kvs))
	for key := range kvs {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	path := fmt.Sprintf("reduce-temp-%v-%v", reduceId, os.Getpid())
	file, e := os.Create(path)
	if e != nil {
		fmt.Printf("Failed to create file %v\n", path)
	}

	for _, key := range keys {
		count := reducef(key, kvs[key])
		_, e := fmt.Fprintf(file, "%v %v\n", key, count)
		if e != nil {
			fmt.Printf("Failed to write MapReduce result (%v, %v) to file %v\n", key, count, path)
		}
	}

	file.Close()
	e = os.Rename(path, fmt.Sprintf("mr-out-%v", reduceId))
	if e != nil {
		fmt.Printf("Failed to rename file %v\n", path)
	}
}

// Worker
// main/mrworker.go calls this function.
//
func Worker(mapf func(string, string) []KeyValue, reducef func(string, []string) string) {

	// Your worker implementation here.
	nReduce, success := getReduceCount()
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

		done, success := false, true
		if response.TaskType == MapTask {
			doMap(response.TaskId, response.TaskFile, nReduce, mapf)
			done, success = finishTask(response.TaskId, MapTask)
		} else if response.TaskType == ReduceTask {
			doReduce(response.TaskId, reducef)
			done, success = finishTask(response.TaskId, ReduceTask)
		}

		if done || !success {
			fmt.Println("All tasks done or Coordinator unavailable, Worker quiting ...")
			return
		}

		time.Sleep(time.Second)
	}
}

// CallExample
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
