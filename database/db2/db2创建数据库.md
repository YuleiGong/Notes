## db2创建数据库
### 基本概念
1. 数据库的最小存储单位是数据页
2. 区：由整数倍的数据页组成
3. 表空间容器：有多个区组成
4. 表空间：多个表空间容器组成

### 创建数据库
```
#都使用默认设置
#db2 list database directory
```
### 默认是数据库目录
```
/home/db2inst4/db2inst4/NODE0000/SQL00001
#cd SQL00001
#存储了表空间信息
#SQLSPCS.1
#SQLSPCS.2
#存储器组控制文件
#SQLSGF.1
#SQLSGF.2
#全局配置文件
#SQLDBCON
```

