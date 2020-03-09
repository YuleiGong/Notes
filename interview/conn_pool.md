# 连接池设计
## 基本模型
* class Connection 连接类 负责维护单个连接
* class ConnPool 连接池类 负责维护连接池
* 使用实例:

```
pool = ConnPool()
conn = pool.get_conn()
conn.release()
#进程结束
pool.discount()
```

## 模型设计
* class ConnPool 
    * 维护正在使用的连接列表(in_use_connection-正在使用, available_connection-可用)
    * 维护max_connection 数量
    * get_conn 的逻辑：优先从 available_connection-可用取,否者make_connection
    * conn.release:将连接丢回 available_connection in_use_connection中删除该连接
    * conn.discount:当前进程代码结束,连接池丢弃。不要带到下一个进程环境中
    * check_pid(): 在获取连接，release 时，都要检查。主要检查当前的进程环境pid和之前的进程pid是否是一个,如果不是,获取锁,对连接进行reset重置。这种情况针对fork之后在子进程中使用同一个连接池的操作。获取锁的时候,添加超时，避免死锁。原则上不允许，在子进程中使用同一个连接池。

## 改进
* 1.添加with as 语句。 对连接自动release 工作
* 2.对于常用的连接，各个进程，在代码片段中，不在显示的 now =connpool。而是在异步函数的装饰器中，根据url来自动拿到连接池,自动分配连接。避免连接的频繁创建
* 3.check_pid 添加超时，避免死锁。
* 4.释放连接的时候，做检查，添加可以重入锁。确保获取锁和释放锁的引用是同一个。
    

