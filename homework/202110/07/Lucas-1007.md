## CH5 设计一致性哈希，1005
- 由 MIT Karger 提出
- 虚拟节点，解决某些过热节点的问题
- 一个物理节点分成几个虚拟节点
- 原来的一致性哈希应用到虚拟节点
- 如果删除一个物理节点，数据不会只移动到顺时针方向的下一个物理节点
- 有了虚拟节点，会移到多个虚拟节点顺时针方向的下一个虚拟节点
- 这样缓解了原来物理节点顺时针方向下一个物理节点过热的问题

```Java
package com.lint.A0520ConsistentHashII;

import java.util.*;

public class Solution {

    private int k;
    private int size = 0;
    private int[] nums;
    private TreeMap<Integer, Integer> tree = new TreeMap<>();

    public Solution(int n, int k) {
        this.k = k;
        nums = new int[n];
        for(int i = 0; i < n; i++) {
            nums[i] = i;
        }

        Random rand = new Random();
        for(int i = 0; i < n; i++) {
            int r = rand.nextInt(i + 1);
            nums[i] ^= nums[r] ^ (nums[r] = nums[i]);
        }
        System.out.println(Arrays.toString(nums));
    }

    public static Solution create(int n, int k) {
        return new Solution(n, k);
    }

    // @return a list of shard ids
    public List<Integer> addMachine(int machindId) {
        List<Integer> ids = new ArrayList<>();
        for(int i = 0; i < k; i++) {
            int id = nums[size++ % nums.length];
            ids.add(id);
            tree.put(id, machindId);
        }
        return ids;
    }

    public int getMachineIdByHashCode(int hashcode) {
        if (tree.isEmpty()) {
            return 0;
        }
        Integer ceiling = tree.ceilingKey(hashcode);
        if (ceiling != null) {
            return tree.get(ceiling);
        }
        return tree.get(tree.firstKey());
    }

    public static void main(String[] args) {
        Solution s = Solution.create(10, 3);
        System.out.println(s.addMachine(1));
        System.out.println(s.getMachineIdByHashCode(4));
        System.out.println(s.addMachine(2));
        System.out.println(s.getMachineIdByHashCode(61));
        System.out.println(s.getMachineIdByHashCode(91));
    }
}

```




















