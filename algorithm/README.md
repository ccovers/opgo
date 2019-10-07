# 面试题目
- 面试遇到的相关问题，总结

## 编程
- [json的Marshal类似接口](json_convert.go)
- [Atoi](atoi.go)
- [Itoa](itoa.go)
- [二分查找](binary_search.go)
- [归并排序](binary_search.go)


## SQL
- 求每个科目的最高分，每个科目分数最高的人
- 索引：=、!=、>、<、LIKE能使用索引吗？ 多列索引什么情况下能使用索引，什么情况下不能？
- 表连接，inner join, outer join, full join 区别
- 事务有什么作用？
- 数据库有哪些锁类型，分别与谁互斥？
- 隔离级别有哪些，默认的隔离级别是什么？


## Linux、Git
- 列出当前目录下所有的.go文件，如果包含子目录呢
ls *.go
find ./ -name *.go

- 查找所有.go文件中包含的chan字符串
grep chan ./ -r

- 将当前目录下的.sql文件都加上.backup后缀
rename .sql .sql.backup ./*

- Git常用命令及工作流程
git add .
git commit -m '注释'
git push
git pull
git merge
git reset --hard '版本号'


http协议，https+加密方式，websocket协议，tcp，udp，hash碰撞，ginrouter结构，redis底层，mysql索引结构，redis集群key问题，切片，map底层，队列，无锁安全队列，数据库锁，b+树


对了，还有笔试题哦，最后一个，让你设计一个日进50g，秒100W+的日志搜集系统，手写一个队列

快速排序算法
堆排序算法
归并排序
二分查找算法
BFPRT(线性查找算法)
DFS（深度优先搜索）
BFS(广度优先搜索)
Dijkstra算法
动态规划算法
朴素贝叶斯分类算法

