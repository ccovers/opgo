# 面试题目
- 面试遇到的相关问题，总结

## 编程
- [json的Marshal类似接口](json_convert.go)
- [Atoi](atoi.go)
- [Itoa](itoa.go)
- [二分查找](binary_search.go)


## SQL
- 求每个科目的最高分，每个科目分数最高的人
	student
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
| id 	| name 	| obj	| score	|
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
| 1		| 红	红	| 语文	| 90	|
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
| 2 	| 红红 	| 数学 	| 80 	|
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
| 3 	| 东东 	| 语文 	| 80 	|
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
| 4 	| 东东 	| 数学 	| 90 	|
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
| 5 	| 梅梅 	| 语文 	| 90 	|
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
SELECT obj, name, score
FROM student AS t
INNER JOIN (SELECT obj, MAX(score) AS socre FROM student group obj)
	AS tmp ON tmp.obj=t.obj AND tmp.score=t.score
ORDER BY obj, name;

- 索引：=、!=、>、<、LIKE能使用索引吗？ 多列索引什么情况下能使用索引，什么情况下不能？
- 表连接，inner join, outer join, full join 区别

- 事务有什么作用？
同一个事务内的操作，要么全部完成，要么一个都不做

- 数据库有哪些锁类型，分别与谁互斥？
共享锁
排它锁
除了共享锁与共享锁，其它组合全互斥

- 隔离级别有哪些，默认的隔离级别是什么？
未提交读
提交读
重复读（默认）
序列化读


## Linux、Git
- 列出当前目录下所有的.go文件，如果包含子目录呢
ls *.go
find ./ -name *.go

- 查找所有.go文件中包含的chan字符串
grep chan ./ -r

- 将当前目录下的.sql文件都加上.backup后缀
rename .sql .sql.backup ./*

## 知识点
gin
b+树
红黑树
B-树（B树）：多路搜索树，每个结点存储M/2到M个关键字，非叶子结点存储指向关键字范围的子结点；所有关键字在整颗树中出现，且只出现一次，非叶子结点可以命中；
B+树：在B-树基础上，为叶子结点增加链表指针，所有关键字都在叶子结点中出现，非叶子结点作为叶子结点的索引；B+树总是到叶子结点才命中；
B*树：在B+树基础上，为非叶子结点也增加链表指针，将结点的最低利用率从1/2提高到2/3；

链表的各种操作：逆序（部分逆序、按某种条件逆序）、判断是否有环，环的入口节点、删除指定节点等。
二叉树的各种操作：各种非递归的遍历操作（前中后、层）、二叉树的公共祖先、根据前中后的遍历结果来重构二叉树等等。
队列、栈相关操作：最小栈、来队列来实现栈等。


Cloud Native 技术栈
Docker，Kubernetes，Prometheus 这些技术的设计思想和编程范式
Go 语言现在唯一可以说占有绝对统治地位就是 Kubernetes 及其生态

https://mp.weixin.qq.com/s/P5j8Dx3-d7rfNMZg4sQUyw
https://mp.weixin.qq.com/s/ISIy3EO5SyZpBN86ATcqCw
https://mp.weixin.qq.com/s/9r-oqjBuTPutxBdxBoZe6g
https://mp.weixin.qq.com/s/fHct-RlGcjt6v7pVB5qJqg

https://pjmike.github.io/2018/08/05/MySQL%E5%B8%B8%E7%94%A8%E8%AF%AD%E5%8F%A5-%E4%B8%80/


深入理解计算机系统
数据结构与算法分析
统计学习方法
机器学习
深度学习
动手学深度学习
高性能 MySQL
Redis 设计与实现
Hadoop 权威指南
编程之美
编程珠玑

