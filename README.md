# kiDB
尝试学着写一个nosql数据库

## 已完成
+ skiplist已基本实现, benchmark结果非常满意！
## 筹备中
+ 通过protobuf序列化所有的写记录
+ 实现LSM持久化
  + 按不同的数据结构分类，存储到文件(每种数据结构写一份pb用来序列化)
  + 每个文件包含一个bloom filter
  + 存储的是指令而不是值，这样才能支持删除操作的持久化
    + 暂定格式{key, value, operation}, operation可以是insert或delete。
    + 删除key时插入一条{key, value, delete}