## 数据库与表空间的恢复
20160724220525
### 数据库恢复示例
1. 因为没有包含日志文件，属于不完全恢复
```
#db2 "create database testdb"
#db2 connect to testdb
#db2 "ctrate table mytable(id int)"
#db2 "insert into mytable values(1)"
#在db2用户下创建归档日志文件
#mkdir -p /db2inst4/testdb/archive
#打开归档数据文件
#db2 update db cfg for testdb using LOGARCHMETH1 "disk:/home/db2inst4/testdb/archive"
#在线备份时离线备份一次
#db2 backup database testdb
#db2 backup database testdb online
#记住时间戳，也查看历史记录文件
#db2 list history all for testdb
#恢复数据库
#db2 restore database testdb taken at 20160724220525
#db2 connect to testdb
SQL1117N  A connection to or activation of database "TESTDB" cannot be made 
because of ROLL-FORWARD PENDING.  SQLSTATE=57019
表示需要回滚操作
#db2 rollforward database testdb to end of backup and complete
#db2 connect to testdb
#db2 "select * from mytable"
ID         
-----------
          1

  1 record(s) selected.

恢复成功

```
### 数据库恢示例
1. 包含日志文件。完全恢复
2. 第一个示例中已经插入数据2。在备份的时候包含日志文件。应该能够恢复
```
#db2 backup database testdb online include logs
#恢复数据库
#db2 restore database testdb taken at 20160724220529
#回滚
db2 rollforward database testdb to end of logs and complete
#db2 connect to testdb
#db2 "select * from mytable"
```
### 表空间恢复
```
#查看表空间
#db2 list tablespaces
 Op Obj Timestamp+Sequence Type Dev Earliest Log Current Log  Backup ID
 -- --- ------------------ ---- --- ------------ ------------ --------------
  B  D  20160724220533001   N    D  S0000004.LOG S0000004.LOG  
 ----------------------------------------------------------------------------
  Contains 2 tablespace(s):

  00001 SYSCATSPACE                                                           
  00002 USERSPACE1
可以看到备份里2张表空间
#恢复表空间
#db2 "restore database testdb tablespace (USERSPACE1) online taken at 20160724220533"
#查看表空间
#db2 list tablespaces
#可以看到表空间属于回滚状态
Tablespace ID                        = 2
 Name                                 = USERSPACE1
 Type                                 = Database managed space
 Contents                             = All permanent data. Large table space.
 State                                = 0x0080
   Detailed explanation:
     Roll forward pending
#前滚表空间
#db2 "rollforward database testdb to end of logs and complete tablespace (USERSPACE1)"
#db2 list tablespaces 正常
```
### 数据库的增量恢复
```
#全量在线备份
#db2 backup database testdb online to /home/db2inst4/db include logs
#增量备份
db2 backup database testdb online incremental to /home/db2inst4/db include logs
#恢复
#db2 restore database testdb incremental automatinc taken at 20160724220545
此语句回自动寻找备份文件，时间戳为最后一次增量备份低时间戳
#回滚
db2 rollforward database testdb to end of logs and complete 
```