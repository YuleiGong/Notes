## python连接db2数据库
```
#修改端口号
#db2sart
#db2 update dbm cfg using SVCENAME 50001
#设置实例连接方式为tcpip
#db2set DB2COMM=TCPIP -i
#db2stop
```
## 修改连接字符串配置文件
```
DBURL='db2+ibm_db://db2inst4:123456@192.168.5.125:50001/hszzs?charset=utf8'
```
