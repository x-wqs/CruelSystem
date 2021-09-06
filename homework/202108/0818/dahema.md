# MIT 6.824 L13 - Transactions
- https://youtu.be/B6btpukqHpM
- https://pdos.csail.mit.edu/6.824/notes/l-2pc.txt

## slides 1
- 2 phase locking (2 PL)
- 2 phase commit (2 PC)

## slides 2
- problem: cross machine atomic ops
- goal: atomicity with  failures + concurrency 

## slides 3
- Transactions
```
   T1                         T2
begin_x                    begin_x
  add(x, -1)                 t1 <- get(x)
  add(y, 1)                  t2 <- get(y)
Commit                     Commit/Abort
```
- our focus: distributed transaction
- ACID: atomicity, consistency, isolation, durability

## slides 4
- Isolation: 未提交读（Read uncommitted）、提交读（read committed）、可重复读（repeatable read）和串行化（Serializable）
- Serializable: T1 || T2 = T1; T2 or T2; T1, 
- example as slides 3, x = 10, y = 10, 
  - Execution order: T1; T2 => T1: 9, 11, T2: 9, 11.
  - Execution order: T2; T1 => T2: 10, 10, T1: 9, 11. 
- serializablity VS lineraizability

## slides 5 - problem cases
```
    T1                T2
                  t1 <- get(x)
 transfer(x,y)
                  t2 <- get(y)
   9, 11            10, 11
not serializable, becuase is not one of the result in slides 4
```

## slides 6 - concurrency control
- pessimistic (locks)
- otimistic (no locks), abort if not serializable

## slides 7 - tow phase locking
- transactions accquire lock before using
- transactions holds until commit

## slides 8
```
    T1                T2
  Lx
    put(x)
  Rx
                  Lx
                  t1 <- get(x)
                  Ly
                  t2 <- get(y)
                  print x,y
                  Rx, Ry
not serializable,
```

## slides 9 - Dead lock
```
    T1                T2
  Lx
    put(x)
                  Ly
                  t2 <- get(y)
  Ly -> wait
    put(y)
                  Lx  -> wait
                  t1 <- get(x)
```
solution dead lock
- abort with timeout 
- graph to detect cycle
