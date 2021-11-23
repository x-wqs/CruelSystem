system design tips and intro by scott shi

# Intro
- Readings
  - DDIA: 入门教材
  - paper: spanner, bigtable, etc.
  - hacker news
- Characteristics
  - No standard answers
    - __ALL ABOUT TRADEOFFS__
    - __Goal and Non-Goal,with ASSUMPTIONS__
    - 系统设计体现的是设计, 而不是做题
  - Shared-Nothing architecture (not mainframe)
  - Require experience + practices + discussions
  - business senses: __推荐DynamoDB的paper中关于bussinss senses的讨论, 解释__
  - Articulate & Communicate: 交流和设计能力同样重要
  - OK to not know sth, not OK to argue with interview
  - 题外话: 随着 yoe 增加, articulation 和 influence 重要性++, purely technical skills 重要性--
  - __<3 yoe, 应该更加关注 fundamental knowledge, such as SQL/_NoSQL_, ACID, concurrency models, B+ trees, etc.__
  - __urgent interview requirement__: 整理past projects的深入细节

# Details (DDIA Ch8)
## Unreliable Network
Idempotent for request retries

## Unreliable Clock
