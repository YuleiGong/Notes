### 报错db2事务日志满
#### SQL0964C  The transaction log for the database is full
***在每次重建数据库后都要设置***
```
#执行如下指令修改日志文件大小
#db2 update db cfg for database_name using LOGFILSIZ 7900 
#db2 update db cfg for detabase_name using LOGPRIMARY 30 
#db2 update db cfg for database_name using LOGSECOND 20
```