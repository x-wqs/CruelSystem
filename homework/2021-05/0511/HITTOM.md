3.5 master mantains a table of O(M * R) map.(status of locations of immediates)  
    usually map split size = 64M, M = 2e5, R = 5000, machines = 2000(R = k * machines, k is small constant)

3.6 machines with bad facility cause straggler, so when remaining few tasks, backup them. this promotes 44% performance.

4 refinements:

	 self-designed partitioning functions

	 guaranteen keys increasing ordered

	 combain less frequence keys in one map machine, process local reduce before real reduce, cut down network cost

	 self-designed input or output type, by add self implementation in Reader or Writer interface

	 atomically produce auxiliary files

	 store bad records, skip next time

	 local debug by small-scale testing

	 show status information

	 input/output counters and other self-defined counters

5.2 distributed grep: size = 1T, M = 15000, R = 1, cost time = 150s

5.3 distributed sort: size = 1T, M = 15000, R = 4000, cost time = 891s. when backup disabled, cost time increase 44% due to straggler

5.5 mock machine failure by killing 200 workers(machine still functioning properly), just increase 5% time

6 first version of MapReduce library done in 2003 Feb, instances grow to 900 2004 Sep in Google.

