package mr

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"net/rpc"
	"os"
	"sort"
	"strconv"
	"time"
)

//
// Map functions return a slice of KeyValue.
//
const RETRY_TIMEOUT = 10

type KeyValue struct {
	Key   string
	Value string
}

type ByKey []KeyValue

// for sorting by key.
func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

//
// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
//
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}


//
// main/mrworker.go calls this function.
//
func Worker(mapf func(string, string) []KeyValue,
	reducef func(string, []string) string) {

	// Your worker implementation here.
	for {
		reply := CallAssignTask()
		if reply.State == Retry {
			time.Sleep(time.Second*RETRY_TIMEOUT)
		} else if reply.State == Free {
			// break
		} else {
				if reply.Type == MapTask {
				DoMap(mapf, reply.Id, reply.Filepath, reply.NReduce)
			} else {
				DoReduce(reducef, reply.Id, reply.NMap)
			}
		}
	}
	// uncomment to send the Example RPC to the coordinator.
	// CallExample()
}

func DoMap(mapf func(string, string) []KeyValue, Id int, filepath string, nReduce int) {
	// filepath = "./tmp" + filepath
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("cannot open %v", filepath)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("cannot read %v", filepath)
	}
	file.Close()
	// map(k1, v1) -> list(k2, v2)
	kva := mapf(filepath, string(content))
	// split by key
	kvs := make([][]KeyValue, nReduce)
	for _, kv := range kva {
		ridx := ihash(kv.Key) % nReduce
		kvs[ridx] = append(kvs[ridx], kv)
	}

	// write to immediate file
	for reduceId, kvl := range kvs {
		filename := "mr-" + strconv.Itoa(Id) + "-" + strconv.Itoa(reduceId)
		file,_ = os.Create(filename)
		enc := json.NewEncoder(file)
		for _, kv := range kvl {
			err := enc.Encode(&kv)
			if (err != nil) {
				log.Fatal("file encode error")
			}
		}
	}
	CallUpdateTaskState(Id, Done)
}

func DoReduce(reducef func(string, []string) string, Id int, nMap int) {
	// all file pathes
	reduceTaskId := Id - nMap
	filelist := make([]string, 0)
	for i := 0; i < nMap; i++ {
		filelist = append(filelist, "mr-" + strconv.Itoa(i) + "-" + strconv.Itoa(reduceTaskId))
	}

	// read json data
	intermediate := []KeyValue{}
	for _, filepath := range filelist{
		file, err := os.Open(filepath)
		defer file.Close()
		if err != nil {
			log.Fatal("file open error")
		}
		dec := json.NewDecoder(file)
		for {
			var kv KeyValue
			if err := dec.Decode(&kv); err != nil {
				break
			}
			intermediate = append(intermediate, kv)
		}
	}

	sort.Sort(ByKey(intermediate))
	
	oname := "mr-out-" + strconv.Itoa(reduceTaskId)
	ofile, _ := os.Create(oname)
	defer ofile.Close()

	// copy and paste
	i := 0
	for i < len(intermediate) {
		j := i + 1
		for j < len(intermediate) && intermediate[j].Key == intermediate[i].Key {
			j++
		}
		values := []string{}
		for k := i; k < j; k++ {
			values = append(values, intermediate[k].Value)
		}
		output := reducef(intermediate[i].Key, values)

		// this is the correct format for each line of Reduce output.
		fmt.Fprintf(ofile, "%v %v\n", intermediate[i].Key, output)

		i = j
	}

	CallUpdateTaskState(Id, Done)
}

func CallAssignTask() AssignTaskReply {
	args := AssignTaskArgs{}
	reply := AssignTaskReply{}
	call("Coordinator.AssignTask", &args, &reply)
	return reply
}

func CallUpdateTaskState(Id int, State TaskState) {
	args := UpdateTaskArgs{}
	reply := UpdateTaskReply{}

	args.Id = Id
	args.State = State
	call("Coordinator.UpdateTaskState", &args, &reply)
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
