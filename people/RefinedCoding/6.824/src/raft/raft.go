package raft

//
// this is an outline of the API that raft must expose to
// the service (or tester). see comments below for
// each of these functions for more details.
//
// rf = Make(...)
//   create a new Raft server.
// rf.Start(command interface{}) (index, term, isLeader)
//   start agreement on a new log entry
// rf.GetState() (term, isLeader)
//   ask a Raft for its current term, and whether it thinks it is leader
// ApplyMsg
//   each time a new entry is committed to the log, each Raft peer
//   should send an ApplyMsg to the service (or tester)
//   in the same server.
//

import (
	"fmt"
	"math/rand"
	//	"bytes"
	"sync"
	"sync/atomic"
	"time"

	//"time"

	//	"6.824/labgob"
	"6.824/labrpc"
)

type State int

const (
	FOLLOWER = iota
	CANDIDATE
	LEADER

	HEARTBEAT_INTERVAL    = 100
	ELECTION_INTERVAL_MIN = 400
	ELECTION_INTERVAL_MAX = 500
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
	votee int
	voted int
	state State
	term  int

	voteCh   chan struct{}
	appendCh chan struct{}

	electionTimer *time.Timer
}

// return currentTerm and whether this server
// believes it is the leader.
func (rf *Raft) GetState() (int, bool) {

	var term int
	var isLeader bool
	// Your code here (2A).
	rf.mu.Lock()
	defer rf.mu.Unlock()

	term = rf.term
	isLeader = rf.state == LEADER
	return term, isLeader
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
// Fix: labgob error: lower-case field Term of VoteRequest in RPC or persist/snapshot will break your Raft
type VoteRequest struct {
	// Your data here (2A, 2B).
	Term      int
	Candidate int
}

//
// example RequestVote RPC response structure.
// field names must start with capital letters!
//
type VoteResponse struct {
	// Your data here (2A).
	Term  int
	Voted bool
}

// TODO-RF: from Leader to follower ?
type AppendEntriesRequest struct {
	Term   int
	Leader int
}

type AppendEntriesResponse struct {
	Term    int
	Success bool
}

// TODO-learn: https://blog.csdn.net/Tommenx/article/details/78313435

//
// example RequestVote RPC handler.
//
func (rf *Raft) RequestVote(request *VoteRequest, response *VoteResponse) {
	// Your code here (2A, 2B).
	rf.mu.Lock()
	defer rf.mu.Unlock()
	//fmt.Printf("handleRequestVote: local: %d, remote: %d, r.term: %d\n", rf.me, request.Candidate, request.Term)

	if request.Term < rf.term {
		response.Voted = false // not vote source ? VoteRequest for source ? apply to be Candidate ?
		response.Term = rf.term
	} else if request.Term > rf.term {
		rf.term = request.Term
		rf.votee = request.Candidate
		rf.state = FOLLOWER
		response.Voted = true
		//rf.updateState(FOLLOWER)
	} else {
		if rf.votee == -1 {
			rf.votee = request.Candidate
			response.Voted = true
		} else {
			response.Voted = false
		}
	}

	if response.Voted {
		go func() {
			rf.voteCh <- struct{}{}
		}()
	}
}

//
// example code to send a RequestVote RPC to a server.
// server is the index of the target server in rf.peers[].
// expects RPC arguments in args.
// fills in *response with RPC response, so caller should
// pass &response.
// the types of the args and response passed to Call() must be
// the same as the types of the arguments declared in the
// handler function (including whether they are pointers).
//
// The labrpc package simulates a lossy network, in which servers
// may be unreachable, and in which requests and replies may be lost.
// Call() sends a request and waits for a response. If a response arrives
// within a timeout interval, Call() returns true; otherwise
// Call() returns false. Thus Call() may not return for a while.
// A false return can be caused by a dead server, a live server that
// can't be reached, a lost request, or a lost response.
//
// Call() is guaranteed to return (perhaps after a delay) *except* if the
// handler function on the server side does not return.  Thus there
// is no need to implement your own timeouts around Call().
//
// look at the comments in ../labrpc/labrpc.go for more details.
//
// if you're having trouble getting RPC to work, check that you've
// capitalized all field names in structs passed over RPC, and
// that the caller passes the address of the response struct with &, not
// the struct itself.
//
func (rf *Raft) sendRequestVote(server int, request *VoteRequest, response *VoteResponse) bool {
	success := rf.peers[server].Call("Raft.RequestVote", request, response)
	return success
}

func (rf *Raft) broadcastVoteRequest() {
	// Fix: WARNING: DATA RACE
	// /mnt/d/code/CruelSystem/people/RefinedCoding/6.824/src/raft/raft.go:264 +0x5c
	rf.mu.Lock()
	request := VoteRequest{Term: rf.term, Candidate: rf.me}
	state := rf.state			// Fix: WARNING: DATA RACE
	rf.mu.Unlock()

	for i, _ := range rf.peers {
		if i != rf.me {
			go func(peer int) {
				var response VoteResponse
				if state == CANDIDATE && rf.sendRequestVote(peer, &request, &response) {
					rf.mu.Lock()
					defer rf.mu.Unlock()

					if response.Voted {
						rf.voted++
					} else {
						if response.Term > rf.term {
							rf.term = response.Term
							rf.state = FOLLOWER
							rf.votee = -1
							// rf.updateState(FOLLOWER)
						}
					}
				} else {
					fmt.Println("Failed to send vote request to ", peer)
				}
			}(i)
		}
	}
}

func (rf *Raft) broadcastAppendEntries() {
	request := AppendEntriesRequest{Term: rf.term, Leader: rf.me}
	for i, _ := range rf.peers {
		if i != rf.me {
			go func(peer int) {
				var response AppendEntriesResponse
				if rf.state == LEADER && rf.sendAppendEntries(peer, &request, &response) {
					rf.mu.Lock()
					defer rf.mu.Unlock()
					if !response.Success && response.Term > rf.term {
						rf.term = response.Term
						rf.updateState(FOLLOWER)
					}
				}
			}(i)
		}
	}
}

// TODO-Lab2B
func (rf *Raft) sendAppendEntries(peer int, request *AppendEntriesRequest, response *AppendEntriesResponse) bool {
	return false
}

func (rf *Raft) AppendEntries(request *AppendEntriesRequest, response *AppendEntriesResponse) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	if request.Term < rf.term {
		response.Success = false
		response.Term = rf.term
	} else if request.Term > rf.term {
		rf.term = request.Term
		rf.updateState(FOLLOWER)
		response.Success = true
	} else {
		response.Success = true
	}

	go func() {
		rf.appendCh <- struct{}{}
	}()
}

func (rf *Raft) updateState(state State) {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	if rf.state != state {
		fmt.Printf("server: %d, Term: %d, currentState: %d, newState: %d\n", rf.me, rf.term, rf.state, state)
		rf.state = state
		if state == FOLLOWER {
			rf.votee = -1
			//} else if state == CANDIDATE {
			//	rf.startElection()
		}
	}
}

//func (rf *Raft) startElection() {
//	rf.Term++
//	rf.votee = rf.me
//	rf.Voted = 1
//	rf.electionTimer.Reset(rf.randElectionTimeout())
//	rf.broadcastVoteRequest()
//}

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
	index := -1
	term := -1
	isLeader := true

	// Your code here (2B).

	return index, term, isLeader
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

		// Your code here to check if a Leader election should
		// be started and to randomize sleeping time using
		// time.Sleep().
		rf.mu.Lock()
		state := rf.state
		rf.mu.Unlock()

		switch state {
		case FOLLOWER:
			select {
			case <-rf.appendCh:
			case <-rf.voteCh:
			case <-time.After(rf.randElectionTimeout()):
				rf.updateState(CANDIDATE)	// Fix: function with lock
			}

		case CANDIDATE:
			rf.mu.Lock()
			rf.term++
			rf.votee = rf.me
			rf.voted = 1
			rf.mu.Unlock()

			go rf.broadcastVoteRequest()

			select {
			case <-time.After(rf.randElectionTimeout()):
			case <-rf.appendCh:
				rf.updateState(FOLLOWER)
			case <-rf.voteCh: // TODO: wonElectionChain, grantVoteChain
				rf.updateState(LEADER)
			}

		case LEADER:
			// TODO-Lab-2B
			// rf.broadcastAppendEntries()
			time.Sleep(HEARTBEAT_INTERVAL)
		}
	}
}

func (rf *Raft) randElectionTimeout() time.Duration {
	rand.Seed(time.Now().UnixNano())
	return time.Duration(ELECTION_INTERVAL_MIN + rand.Intn(ELECTION_INTERVAL_MAX-ELECTION_INTERVAL_MIN))
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
func Make(peers []*labrpc.ClientEnd, me int, persister *Persister, applyCh chan ApplyMsg) *Raft {
	rf := &Raft{}
	rf.peers = peers
	rf.persister = persister
	rf.me = me

	// Your initialization code here (2A, 2B, 2C).
	rf.state = FOLLOWER
	rf.votee = -1
	rf.voteCh = make(chan struct{})
	rf.appendCh = make(chan struct{})

	// initialize from state persisted before a crash
	rf.readPersist(persister.ReadRaftState())

	// start ticker goroutine to start elections
	go rf.ticker()

	return rf
}

// election/leader ticker
// https://github.com/JellyZhang/mit-6.824-2021/blob/master/src/raft/raft.go

// leader loop
// https://github.com/viktorxhzj/6.824/blob/master/src/raft/raft.go

// start election
// https://github.com/h1deOnBush/6.824-2021/blob/master/src/raft/raft.go

// applier / ticker
// https://github.com/crimson-gao/MIT-6.824-spring2021/blob/master/src/raft/raft.go

// timeout watcher / log applier
// https://github.com/MrGuin/6.824-labs-2021/blob/master/src/raft/raft.go

// ticker / NotifyApplyCh
// https://github.com/vnguyen12/6.824-Distributed-Systems/blob/main/src/raft/raft.go

// Image Update
// https://github.com/lzlaa/6.824/blob/master/raft/raft.go
