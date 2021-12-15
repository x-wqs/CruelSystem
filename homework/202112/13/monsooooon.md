designing data intensive application p78-

# 维护sstable
- 内存中保存一个有序的数据结构(BTree/AVLTree), 写入数据记录到内存中
- 如果crash了以后内存数据丢失? Write-Ahead-Log, 每次写入内存, 先用append的方式记录到WAL里才算成功

# 构建LSM-Tree
- 什么是LSM-Tree? keeping a cascade of SSTables that are merged in the background
- 可以学习的数据库: LevelDB/RocksDB, 都使用了SSTable和LSM-Tree
- LSM_Tree查询的问题: 如果Value不存在, 则需要遍历所有层级的Segment和Mem才能知道
- 基于Bloom Filter做优化: 如果判定不存在, 一定不存在. 如果判定存在, 可能不存在
- LSM_Tree里面如何做 segment的 Compaction/Merge? 何时/以怎样的频率做? 主要有两种策略:
  - Leveled
    - 示例: LevelDB/RocksDB
    - the key range is split up into smaller SSTables and older data is moved into separate “levels,” which allows the compaction to proceed more incrementally and use less disk space.
  - Size-Tiered
    - 示例: HBase
    - newer and smaller SSTables are successively merged into older and larger SSTables.

# B Tree
- B-trees break the database down into fixed-size blocks or pages, traditionally 4 KB in size (sometimes bigger), and read or write one page at a time.
- 和disk的读写pattern比较匹配(每次读固定大小的page)
- disk上的page也有地址, 因此page和page之间可以通过指针链接起来, 形成Tree的结构
- 每一个page是一个多叉树的节点, page里面的key按照递增顺序排列
- Branching Factor: 一个page里放多少val, 取决于page的大小, 通常是几百
- BTree的增删改查: 
  - Update: 读取page到内存, 修改, 写入disk
  - Create: 找到range内的page, 插入, 如果超过了page的限制, 需要进行split, 写入两个page, 并且更新parent pages
- BTree通常的深度不超过4层

# Make B Tree Reliable
- 由于BTree的写入涉及到disk的(possibly, 多次)更新, 例如insert需要更新3个page, 如果中途disk crash, 则会出现inconsistency
- 解决:WAL/Redo Log, 每一个BTree的操作之前都需要记录下操作内容. 如果crash, 基于WAL可以回溯之前的步骤
- 多线程写B Tree: protecting the tree’s data structures with latches

# B Tree和LSM Tree的比较
- LSM-trees are typically faster for writes, whereas B-trees are thought to be faster for reads
- To Be Cont.
