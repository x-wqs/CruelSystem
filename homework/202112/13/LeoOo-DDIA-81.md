
## B-Tree

B-tree are the ubiquitous index used by RDBMS and NoSQL DB.

K-V pairs sorted by key allows lookup and range queries.

B-Tree vs Log-structured Index
- LI: break the db down into variable-size sements, MB size. Always write a segment sequentially.
- BT: break the db down into fixed-size blocks or pages, 4KB size (can be determined by filesystem block size), read or write one page at a time.
  - disks are also arranged in fixed-size blocks.
- Each page can be identified using an address or location, like a pointer (but on disk), aka a `page reference`.
- One page is designated as the root of the B-tree, whenever you want to look up a key in the index, you start from the root.
  - contains keys and references to child pages.
  - key: e.g. user_id
  - `ref,100;ref,200;ref,300;ref,400;ref,500;ref`
  - each child is responsible for a continuous range of keys.
  - we keep follow the references until we get to a page that contains individual keys (a leaf page).
- Modify value of a key
  - follow the ref, find the leaf page and change the value then write back.
- Add a new key
  - find the leaf page whose range encompasses the new key and add it to that page.
  - if there's enough space, insert the k,v and done.
  - if there isn't enough space, break the leaf page into two pages, insert a new node to parent and update the parent to contain the ref to the two new pages.
- Tree is balanced, with a depth of logn, four-level tree of 4KB pages with a branching factor of 500 can store up to 256TB.
- Deleting while keeping the tree balanced is more challenging.


Reliable B-Tree
- B-tree update a page inplace so that the references are not changed. (LSM tree only append but never modify file inplace)
  - Writing a page is a hardware operation and sometimes e.g. insert, we need to write multiple pages (new leaf pages, parent pages for updated refs).
  - Hardware failure could lead to incomplete writes, corrupted pages, orphan pages.
- Write-ahead log, WAL, aka redo log.
-
