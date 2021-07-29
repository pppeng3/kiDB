# kiDB
尝试学着写一个nosql数据库

## 已完成
+ skiplist已基本实现, benchmark结果非常满意！
## 筹备中
+ 实现LSM持久化
  + 每个文件包含一个bloom filter
  + 存储的是指令而不是值，这样才能以append-only的形式实现删除操作
    + 暂定格式{key, value, operation}, operation可以是insert或delete。
    + 删除key时插入一条{key, value, delete}
