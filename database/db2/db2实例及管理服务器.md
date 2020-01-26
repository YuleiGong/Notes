## 实例
### 实例为数据库提供了运行环境
### 实例由一组共享内存和后台进程组成
### 数据库是物理的，而实例是逻辑的
1. 一个db2产品可以创建多个实例
2. 每个实例都有自己的配置文件
3. 每个实例都有自己的共享内存和进程结构
4. 实际生成环境中，实例个数由情况而定
### 创建实例
1. linux中，实例和用户名一样，数据库中的文件存放在/home/username目录下
2. 用户状态必须正常
```
#创建实例的指令 -u后为受防护的用户+实例用户 username(实例)
sudo ./db2icrt -u db2fenc4 db2inst4
```
### 连接和断开实例
####  操作数据库要先连接实例，才能connect数据库
```
#断开实例
#db2 detach
#连接实例
#db2 attach to db2inst4
#停止实例
#db2stop
#如果有用户连接进来不能停止。强制停止
#db2stop force
```
### 查看实例参数(配置文件)
```
#查看实例的配置
#db2 get dbm cfg
#修改实例配置(某些参数需要重启实例才生效)
#db2 update dbm cfg using 参数名字+参数配置
#修改参数为默认值
#db2 reset dbm cfg
```
### 删除实例
1. sudo 或者root用户
2. 删除前，需要db2stop实例
3. instance 目录
```
#删除实例
#./db2idrop db2inst4
```
### db2其它实例命令
#### 查看db2创建了那些实例
```
#db2ilist
#db2开机自己启动实例
#db2iauto
```
### 实例目录
1. 实例目录在用户创建实例后自动创建
2. 位于/home/db2inst4/sqllib 即用户目录下面
3. /home/db2inst4/sqllib/db2dump db2diag.log 记录的该实例下db2的错误日志
4. /home/db2inst4/sqllib/sqldbdir db2系统数据库目录
5. db2systm 数据库管理配置文件
6. db2nodes.cfg db2节点配置文件


