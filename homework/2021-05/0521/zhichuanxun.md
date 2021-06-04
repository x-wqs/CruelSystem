# gfs

Building fault tolerance storage

app = stateless
storage holds persistent state

Why hard

high performance => many servers => constant faults => fault tolerance => replication => potential inconsistences => strong consistence

protocol

gfs: fs for map-reduce

large data set
automatic sharding
all apps see same fs
ft

increase the version number before lease