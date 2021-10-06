## CH5 设计一致性哈希，1005
- 原始哈希，数据ID % 节点数
- 节点数变化，引起大量数据移动
- 一致性哈希，节点分布式 360度的圆环上，哈希环
- 数据ID取模，落入圆环范围，然后顺时针，找到对应的服务器节点
- 如果减少节点，则继续按照顺时针，寻找可用节点
- 如果增加节点，则某些节点的数据分一半给新节点，其他节点的数据分布保持不变

### 代码实现
```Java
// 0519-consistent-hash
// Q: https://www.lintcode.com/problem/consistent-hashing

public class Solution {
    /*
     * @param n: a positive integer
     * @return: n x 3 matrix
     */

        // S1: PriorityQueue
        // Time: O(N * LogN)
    // Rank: 4002ms
    // n = 5, [[0, 44, 1], [45, 89, 5], [90, 179, 3], [180, 269, 2], [270, 359, 4]]
    // n = 6, [[0,44,1],[45,89,5],[90,179,3],[180,224,2],[225,269,6],[270,359,4]]
    public List<List<Integer>> consistentHashing(int n) {
        PriorityQueue<List<Integer>> pq = new PriorityQueue<>((a, b) -> firstLargest(a, b));
        pq.add(Arrays.asList(0, 359, 1));
        for (int i = 2; i <= n; i ++) {
                        List<Integer> top = pq.poll();
                        int start = top.get(0), end = top.get(1), old = top.get(2);
                        int mid = (start + end) / 2;
                        pq.add(Arrays.asList(start, mid, old));
                        pq.add(Arrays.asList(mid + 1, end, i));
            System.out.println(pq);
                }
                return new ArrayList<>(pq);
    }

    int firstLargest(List<Integer> a, List<Integer> b) {
        int c = a.get(1) - a.get(0), d = b.get(1) - b.get(0);
        return c == d ? a.get(2) - b.get(2) : d - c;
    }

    // Rank: 2436ms
    public List<List<Integer>> consistentHashingN2(int n) {
        // write your code here
        List<List<Integer>> machineList = new ArrayList<>();
        List<Integer> machine = new ArrayList<>();
        machine.add(0);
        machine.add(359);
        machine.add(1);
        machineList.add(machine);

        for (int mID = 1; mID < n; mID++) {
            List<Integer> newMachine = new ArrayList<>();
            int idxMax = 0;
            int zoneMax = machineList.get(idxMax).get(1) - machineList.get(idxMax).get(0) + 1;
            for (int idx = 1; idx < mID; idx++) {
                int zoneNew = machineList.get(idx).get(1) - machineList.get(idx).get(0) + 1;
                if (zoneNew > zoneMax) {
                    idxMax = idx;
                    zoneMax = zoneNew;
                    break;
                }
            }

            int zoneMaxStart = machineList.get(idxMax).get(0);
            int zoneMaxEnd = machineList.get(idxMax).get(1);
            machineList.get(idxMax).set(1, (zoneMaxEnd + zoneMaxStart) / 2);

            newMachine.add((zoneMaxEnd + zoneMaxStart) / 2 + 1);
            newMachine.add(zoneMaxEnd);
            newMachine.add(mID + 1);

            machineList.add(newMachine);
        }
        return machineList;
    }
}
```
