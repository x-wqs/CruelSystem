package raft

//
// this is an outline of the API that raft must expose to
// the service (or tester). see comments below for
// each of these functions for more details.
//
// rf = Make(...)
//   create a new Raft server.
// rf.Start(command interface{}) (index, term, isleader)
//   start agreement on a new log entry
// rf.GetState() (term, isLeader)
//   ask a Raft for its current term, and whether it thinks it is leader
// ApplyMsg
//   each time a new entry is committed to the log, each Raft peer
//   should send an ApplyMsg to the service (or tester)
//   in the same server.
//

import (
	//	"bytes"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	//	"6.824/labgob"
	"6.824/labrpc"
)

//
// as each Raft peer becomes aware that successive log entries are
// committed, the peer should send an ApplyMsg to the service (or
// tester) on the same server, via the applyCh passed to Make(). set
// CommandValid to true to indicate that the ApplyMsg contains a newly
// committed log entry.
//
// in part 2D you'll want to send other kinds of messages (e.g.,
// snapshots) on the applyCh, but set CommandValid to false for these
// other uses.
//
type ApplyMsg struct {
	CommandValid bool
	Command      interface{}
	CommandIndex int

	// For 2D:
	SnapshotValid bool
	Snapshot      []byte
	SnapshotTerm  int
	SnapshotIndex int
}

type Role int

const (
	Leader Role = iota
	Follower
	Candidate
)

// raft log
type RFLog struct {
	Term    int
	Index   int
	Command interface{}
}

//
// A Go object implementing a single Raft peer.
//
type Raft struct {
	mu        sync.Mutex          // Lock to protect shared access to this peer's state
	peers     []*labrpc.ClientEnd // RPC end points of all peers
	persister *Persister          // Object to hold this peer's persisted state
	me        int                 // this peer's index into peers[]
	dead      int32               // set by Kill()

	// Your data here (2A, 2B, 2C).
	// Look at the paper's Figure 2 for a description of what
	// state a Raft server must maintain.
	role Role

	term         int
	lastLogIndex int
	lastLogTerm  int
	leaderCommit int

	// timeout channel
	toCh chan bool
	// each term only vote once
	voted bool

	// log apply channel
	applyCh chan ApplyMsg
	// nextIndex
	nextIndexes []int
	// the index that has been confirmed committed
	committedIndex int
	logs           []RFLog
	// the counter to track each commit's res
	committedCounter map[int]int
}

// ### time

const HEARTBEAT_INTERVAL = 100

// 150ms - 300ms
const TIMEOUT_BASE = 150

const TIMEOUT_RANGE = 150

// return currentTerm and whether this server
// believes it is the leader.
func (rf *Raft) GetState() (int, bool) {

	// var term int
	// var isleader bool
	// Your code here (2A).
	return rf.term, rf.role == Leader
}

//
// save Raft's persistent state to stable storage,
// where it can later be retrieved after a crash and restart.
// see paper's Figure 2 for a description of what should be persistent.
//
func (rf *Raft) persist() {
	// Your code here (2C).
	// Example:
	// w := new(bytes.Buffer)
	// e := labgob.NewEncoder(w)
	// e.Encode(rf.xxx)
	// e.Encode(rf.yyy)
	// data := w.Bytes()
	// rf.persister.SaveRaftState(data)
}

//
// restore previously persisted state.
//
func (rf *Raft) readPersist(data []byte) {
	if data == nil || len(data) < 1 { // bootstrap without any state?
		return
	}
	// Your code here (2C).
	// Example:
	// r := bytes.NewBuffer(data)
	// d := labgob.NewDecoder(r)
	// var xxx
	// var yyy
	// if d.Decode(&xxx) != nil ||
	//    d.Decode(&yyy) != nil {
	//   error...
	// } else {
	//   rf.xxx = xxx
	//   rf.yyy = yyy
	// }
}

//
// A service wants to switch to snapshot.  Only do so if Raft hasn't
// have more recent info since it communicate the snapshot on applyCh.
//
func (rf *Raft) CondInstallSnapshot(lastIncludedTerm int, lastIncludedIndex int, snapshot []byte) bool {

	// Your code here (2D).

	return true
}

// the service says it has created a snapshot that has
// all info up to and including index. this means the
// service no longer needs the log through (and including)
// that index. Raft should now trim its log as much as possible.
func (rf *Raft) Snapshot(index int, snapshot []byte) {
	// Your code here (2D).

}

//
// example RequestVote RPC arguments structure.
// field names must start with capital letters!
//
// type RequestVoteArgs struct {
// 	// Your data here (2A, 2B).
// }

// //
// // example RequestVote RPC reply structure.
// // field names must start with capital letters!
// //
// type RequestVoteReply struct {
// 	// Your data here (2A).
// }

//
// example RequestVote RPC handler.
//
// func (rf *Raft) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {
// 	// Your code here (2A, 2B).
// }

//
// example code to send a RequestVote RPC to a server.
// server is the index of the target server in rf.peers[].
// expects RPC arguments in args.
// fills in *reply with RPC reply, so caller should
// pass &reply.
// the types of the args and reply passed to Call() must be
// the same as the types of the arguments declared in the
// handler function (including whether they are pointers).
//
// The labrpc package simulates a lossy network, in which servers
// may be unreachable, and in which requests and replies may be lost.
// Call() sends a request and waits for a reply. If a reply arrives
// within a timeout interval, Call() returns true; otherwise
// Call() returns false. Thus Call() may not return for a while.
// A false return can be caused by a dead server, a live server that
// can't be reached, a lost request, or a lost reply.
//
// Call() is guaranteed to return (perhaps after a delay) *except* if the
// handler function on the server side does not return.  Thus there
// is no need to implement your own timeouts around Call().
//
// look at the comments in ../labrpc/labrpc.go for more details.
//
// if you're having trouble getting RPC to work, check that you've
// capitalized all field names in structs passed over RPC, and
// that the caller passes the address of the reply struct with &, not
// the struct itself.
//
// func (rf *Raft) sendRequestVote(server int, args *RequestVoteArgs, reply *RequestVoteReply) bool {
// 	ok := rf.peers[server].Call("Raft.RequestVote", args, reply)
// 	return ok
// }

//
// the service using Raft (e.g. a k/v server) wants to start
// agreement on the next command to be appended to Raft's log. if this
// server isn't the leader, returns false. otherwise start the
// agreement and return immediately. there is no guarantee that this
// command will ever be committed to the Raft log, since the leader
// may fail or lose an election. even if the Raft instance has been killed,
// this function should return gracefully.
//
// the first return value is the index that the command will appear at
// if it's ever committed. the second return value is the current
// term. the third return value is true if this server believes it is
// the leader.
//
func (rf *Raft) Start(command interface{}) (int, int, bool) {
	index := rf.lastLogIndex
	term := rf.lastLogTerm
	isLeader := rf.role == Leader
	if !isLeader {
		return index, term, isLeader
	}

	rf.lastLogIndex++
	rf.lastLogTerm = rf.term
	rf.logs = append(rf.logs, RFLog{rf.lastLogIndex, rf.lastLogTerm, command})
	// self commit confirmed
	rf.committedCounter[rf.lastLogIndex]++

	for i := 0; i < len(rf.peers); i++ {
		rf.nextIndexes[i]++
	}

	return rf.lastLogIndex, rf.lastLogTerm, isLeader
}

func (rf *Raft) startOneElection() {
	rf.role = Candidate
	rf.term++
	rf.voted = true
	// DPrintf("%d starts election in term %d", rf.me, rf.term)

	votedChn := make(chan bool)
	granted := 1

	for i := 0; i < len(rf.peers); i++ {
		if i != rf.me {
			go rf.CallRequestVote(i, votedChn)
		}
	}

	for i := 0; i < len(rf.peers)-1; i++ {
		if <-votedChn {
			granted++
		}
	}

	// DPrintf("<<<<<<<<<< %d in term %d  received %d", rf.me, rf.term, granted)

	if granted > len(rf.peers)/2 && rf.role == Candidate {
		rf.role = Leader
		rf.toCh <- true
		DPrintf("%d win in term %d", rf.me, rf.term)

		// start heartbeat
		go rf.startLeaderHeartbeat()

		// start to sync follower's log
		go rf.startSyncAllLogs()
		return
	}
}

// caller
func (rf *Raft) CallRequestVote(serverId int, votedChn chan bool) {
	defer func() {
		// DPrintf("########## %d in term %d timeout from %d", rf.me, rf.term, serverId)
		votedChn <- false
	}()

	args := RequestVoteArgs{}
	reply := RequestVoteReply{}

	args.Term = rf.term
	args.CandidateId = rf.me
	args.LastLogIndex = rf.lastLogIndex
	args.LastLogTerm = rf.lastLogTerm

	rf.peers[serverId].Call("Raft.RequestVote", &args, &reply)

	// DPrintf("for %d, %d can connect", rf.me, serverId)

	if reply.Term > rf.term {
		rf.role = Follower
		// reset timeout timer
		rf.toCh <- true
		// wait to vote the candidate
		rf.voted = false
	}

	votedChn <- reply.VoteGranted

	return
}

// handler
func (rf *Raft) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {
	rf.mu.Lock()
	rf.mu.Unlock()
	reply.Term = rf.term

	// find a newer candidate
	if args.Term > rf.term {
		reply.VoteGranted = true
		if args.LastLogTerm <= rf.lastLogTerm && args.LastLogIndex < rf.lastLogIndex {
			// reject out of date
			// DPrintf("%d in term %d reject %d (term: %d) out-of-date", rf.me, rf.term, args.CandidateId, args.Term)
			reply.VoteGranted = false
		}
		// DPrintf("%d in term %d vote %d (term: %d)", rf.me, rf.term, args.CandidateId, args.Term)
		rf.role = Follower
		rf.toCh <- true
		rf.term = args.Term
		// voted for the new term
		// won't affect if the leader is out of date?
		rf.voted = true
		return
	}

	if args.Term == rf.term {
		if rf.voted {
			// DPrintf("%d in term %d reject %d (term: %d)", rf.me, rf.term, args.CandidateId, args.Term)
			reply.VoteGranted = false
			return
		}
		// candidate always voted when have the same term
		reply.VoteGranted = true
		return
	}

	if args.Term < rf.term {
		reply.VoteGranted = false
		return
	}
	return
}

func (rf *Raft) startLeaderHeartbeat() {
	// DPrintf("%d start leader heartbeat", rf.me)
	for {
		// out-of-date leader
		if rf.role != Leader {
			return
		}
		// in case of leader's suicide
		rf.toCh <- true
		// send empty AppendEntries
		for i := 0; i < len(rf.peers); i++ {
			if i != rf.me {
				go rf.CallAppendEntries(i, true)
			}
		}
		time.Sleep(time.Duration(HEARTBEAT_INTERVAL) * time.Millisecond)
	}
}

func (rf *Raft) startSyncAllLogs() {
	// reset follower trackers
	rf.mu.Lock()
	nextIdx := rf.lastLogIndex + 1
	for i := 0; i < len(rf.peers); i++ {
		rf.nextIndexes[i] = nextIdx
	}
	rf.committedCounter = make(map[int]int)
	rf.mu.Unlock()

	for i := 0; i < len(rf.peers); i++ {
		if i != rf.me {
			go rf.syncFollowerLogs(i)
		}
	}
}

// sync one follower
func (rf *Raft) syncFollowerLogs(serverId int) {
	for {
		if rf.role != Leader {
			// re-election
			return
		}
		if rf.nextIndexes[serverId] <= rf.lastLogIndex {
			rf.CallAppendEntries(serverId, false)
		}
		time.Sleep(time.Duration(HEARTBEAT_INTERVAL) * time.Millisecond)
	}
}

func (rf *Raft) CallAppendEntries(serverId int, isHeartbeat bool) {
	args := AppendEntriesArgs{}
	reply := AppendEntriesReply{}

	rf.mu.Lock()
	args.Term = rf.term
	args.LeaderId = rf.me
	args.LeaderCommit = rf.leaderCommit

	// the log idx try to append
	tryIndex := rf.nextIndexes[serverId]
	args.PrevLogIndex = rf.logs[tryIndex-1].Index
	args.PrevLogTerm = rf.logs[tryIndex-1].Term
	args.Entries = []RFLog{}

	if !isHeartbeat {
		// batch append log entries
		for i := tryIndex; i < rf.lastLogIndex; i++ {
			args.Entries = append(args.Entries, rf.logs[i])

		}
	}
	rf.mu.Unlock()

	rf.peers[serverId].Call("Raft.AppendEntries", &args, &reply)

	if reply.Term > rf.term {
		// re-election
		rf.role = Follower
		return
	}

	rf.mu.Lock()
	if !reply.Success {
		// got the reply, not timeout
		// disagree with the idx
		rf.nextIndexes[serverId]--
	} else {
		if !isHeartbeat {
			// if not a heartbeat, means an entry applied
			// update next idx
			rf.nextIndexes[serverId] += len(args.Entries)
			// update commit counter
			appendedIndex := tryIndex + len(args.Entries) - 1
			if rf.committedIndex < appendedIndex {
				// TODO: can we merge these two loop?
				for i := tryIndex; i <= appendedIndex; i++ {
					rf.committedCounter[i]++
				}
				for i := rf.committedIndex + 1; i <= appendedIndex; i++ {
					if rf.committedCounter[tryIndex] > len(rf.peers) / 2 {
						rf.leaderCommit = tryIndex
					}
				}
			}
		}
	}
	rf.mu.Unlock()
}

func min(x, y int) int {
	if x < y {
			return x
	}
	return y
}

// handler
func (rf *Raft) AppendEntries(args *AppendEntriesArgs, reply *AppendEntriesReply) {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	if rf.term > args.Term {
		// a isolated leader may send heartbeat
		// reject if in a new term
		return
	}
	// ########## heartbeat ##########
	rf.role = Follower
	// reset timeout
	rf.toCh <- true
	rf.term = args.Term
	// ########## heartbeat ##########


	// ########## log replication ##########
	rf.mu.Lock()
	defer rf.mu.Unlock()
	reply.Term = rf.term
	if args.PrevLogIndex > rf.lastLogIndex {
		// follower has less log than the leader
		// tell leader to reduce the nextIdx
		reply.Success = false
		return
	}
	if args.PrevLogIndex == rf.lastLogIndex && args.PrevLogTerm == rf.lastLogTerm {
		// entry matches!!
		reply.Success = true
		if len(args.Entries) > 0 {
			// if not a heartbeat, add log
			// lazy drop all the log after and append new log
			for i := 0; i < len(args.Entries); i++ {
				newLog := args.Entries[i]
				if cap(rf.logs) <= newLog.Index{
					rf.logs = append(rf.logs, newLog)
				} else {
					rf.logs[newLog.Index] = newLog
				}
				rf.lastLogIndex = newLog.Index
				rf.lastLogTerm = newLog.Term
			}

		}		
		// update committedIndex if leader committed
		if rf.committedIndex < args.LeaderCommit {
			for i := rf.committedIndex; i < min(args.LeaderCommit, len(rf.logs)); i++ {
				rf.applyLog(i)
			}
		}
		rf.committedIndex = args.LeaderCommit
		return 
	}

	reply.Success = false
	return
}

func(rf *Raft) applyLog(logIdx int) {

	// TODO: snapshoot
	msg := ApplyMsg{}
	msg.CommandValid = true
	msg.Command = rf.logs[logIdx].Command
	msg.CommandIndex = logIdx
	
	rf.applyCh <- msg
	
}

//
// the tester doesn't halt goroutines created by Raft after each test,
// but it does call the Kill() method. your code can use killed() to
// check whether Kill() has been called. the use of atomic avoids the
// need for a lock.
//
// the issue is that long-running goroutines use memory and may chew
// up CPU time, perhaps causing later tests to fail and generating
// confusing debug output. any goroutine with a long-running loop
// should call killed() to check whether it should stop.
//
func (rf *Raft) Kill() {
	atomic.StoreInt32(&rf.dead, 1)
	// Your code here, if desired.
}

func (rf *Raft) killed() bool {
	z := atomic.LoadInt32(&rf.dead)
	return z == 1
}

// The ticker go routine starts a new election if this peer hasn't received
// heartsbeats recently.
func (rf *Raft) ticker() {
	for rf.killed() == false {

		// Your code here to check if a leader election should
		// be started and to randomize sleeping time using
		// time.Sleep().
		select {
		case <-rf.toCh:
			continue
		case <-time.After(
			time.Duration(TIMEOUT_BASE+rand.Int31n(TIMEOUT_RANGE)) *
				time.Millisecond):

			go rf.startOneElection()
		}
	}
}

//
// the service or tester wants to create a Raft server. the ports
// of all the Raft servers (including this one) are in peers[]. this
// server's port is peers[me]. all the servers' peers[] arrays
// have the same order. persister is a place for this server to
// save its persistent state, and also initially holds the most
// recent saved state, if any. applyCh is a channel on which the
// tester or service expects Raft to send ApplyMsg messages.
// Make() must return quickly, so it should start goroutines
// for any long-running work.
//
func Make(peers []*labrpc.ClientEnd, me int,
	persister *Persister, applyCh chan ApplyMsg) *Raft {
	rf := &Raft{}
	rf.peers = peers
	rf.persister = persister
	rf.me = me

	// Your initialization code here (2A, 2B, 2C).
	rf.role = Follower

	rf.term = 0
	rf.lastLogIndex = 0
	rf.lastLogTerm = 0
	rf.leaderCommit = 0

	rf.toCh = make(chan bool)
	rf.voted = false

	rf.applyCh = applyCh
	rf.nextIndexes = make([]int, len(rf.peers))
	rf.committedIndex = 0
	rf.logs = make([]RFLog, 0)
	// sentry entry
	rf.logs = append(rf.logs, RFLog{0, 0, nil})
	// initialize from state persisted before a crash
	rf.readPersist(persister.ReadRaftState())

	// start ticker goroutine to start elections
	go rf.ticker()

	return rf
}
