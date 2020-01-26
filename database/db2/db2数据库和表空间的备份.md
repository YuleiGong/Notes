## db2数据库和表空间的备份
### 离线备份
***#db2 backup database work*** 不能在线备份
```
#备份前最好进入实例用户下，在其它目录可能没有权限写入数据
#db2 backup database work
#查看当前连接的数据库
#db2 list application
 No data was returned by Database System Monitor
#ls -ltr
#可以看到,产生了一个备份数据
drwxrwxr-x  3 db2inst4 db2adm4      4096 Jul 19 09:00 db2inst4
drwxrwsr-t 19 db2inst4 db2adm4      4096 Jul 19 09:24 sqllib
-rw-------  1 db2inst4 db2adm4 318865408 Jul 21 01:03 WORK.0.db2inst4.NODE0000.CATN0000.20160721010331.001

```
```
#db2 connect to work 连接数据库
#db2 list appliaction 可以看到 连接了数据库
Auth Id  Application    Appl.      Application Id                                                 DB       # of
         Name           Handle                                                                    Name    Agents
-------- -------------- ---------- -------------------------------------------------------------- -------- -----
DB2INST4 db2bp          164        *LOCAL.db2inst4.160721051208                                   WORK     1

```
```
#再次备份 
#db2 backup database work
SQL1035N  The database is currently in use.  SQLSTATE=57019 数据库正在使用，不能在线备份
#db2 terminate 断开和数据库的连接
#再次备份 
#db2 backup database work 备份成功

```
### 在线备份
***db2 backup database work online***在线备份
```
#db2 connect to work
#db2 backup database work online
#没有打开数据库归档日志，备份失败
SQL2413N  Online backup is not allowed because the database is not recoverable
or a backup pending condition is in effect.
```
#### 打开数据库的归档案日志
```
#创建归档日志文件夹 archive
/home/db2inst4/work/archive
#db2 update db cfg for work using LOGARCHMETH1 "disk:/home/db2inst4/work/archive" 创建归档日志
#重新连接数据库
#db2 connect to work
SQL1116N  A connection to or activation of database "WORK" cannot be made
because of BACKUP PENDING.  SQLSTATE=57019 需要离线备份一次
#db2 backup database work 
#db2 connect to work
#再次测试在线备份
#db2 connect databse work online
Backup successful. The timestamp for this backup image is : 20160721020156 在线备份成功
```
### 在线备份+备份日志(只能针对在线备份)
```
#db2 backup database work online include logs
```
### 指定备份文件的存放路径
```
#db2 backup database work online to /home/db2inst4/db1, /home/db2inst4/db2
```
### 只备份数据库的表空间
```
#db2 "backup database work tablespace (syscatspace) online to /home/db2inst4/db"
```
### 增量备份
```
#db2 backup database work online incremental to /home/db2inst4/db
表示没有开启增量备份
SQL2426N  The database has not been configured to allow the incremental backup
operation. Reason code = "1".
#开启增量备份
#db2 UPDATE DATABASE cfg FOR databse_name USING TRACKMOD YES
#再次备份
#db2 backup database work online incremental to /home/db2inst4/db
表示在开启增量备份后没有对至少一个表空间非增量备份
SQL2426N  The database has not been configured to allow the incremental backup
operation. Reason code = "2".
#执行非增量备份
#db2 backup database work online to /home/db2inst4/db
#再次执行增量备份
#db2 backup database work online incremental to /home/db2inst4/db 成功

```
### delta 备份
```
db2 backup database work online incremental delta to /home/db2inst4/db

```
### 检查备份的完整性
**db2ckbkp**
1. 检查备份文件的一致性和备份文件的可用性
2. 用于显示备份文件的元数据
```
显示备份信息
#db2ckbkp -h WORK.0.db2inst4.NODE0000.CATN0000.20160721024045.001
```