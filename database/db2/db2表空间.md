## db2表空间管理
### 表空间概念
1. 表空间是数据库逻辑层和物理层的中间桥梁
2. 表空间是用户逻辑层(包括表和索引)的存储空间
3. 表空间在物理层对应多个容器(包括文件，目录，设备)
4. 数据库可以有多个表空间，表空间可以有多个容器
### 表空间类型
#### 目录表空间
1. 存储db2系统编目，即"数据字典"
2. 默认名称： syscatspace
#### 系统临时表空间
1. 用于存储分组，排序，连接，创建索引等操作的中间结果(example:排序一般在内存中完成,当内存不足，就会存在系统临时表空间)
2. 数据库应至少有一个这样的表空间
3. 默认名称： tempspacel
#### 用户临时表空间
1. 默认不会创建
2. 存放创建表的临时数据
#### 用户表空间
1. 存储用户对象(表 索引等)的空间
2. 默认名称：userpace1
### 表空间存储类型
#### SMS系统管理表空间
1. 表空间物理层的容器是目录或者文件夹
2. 表空间的分配和管理由操作系统的文件系统完成
3. 可优化选项少，且性能不好
#### DMS数据库管理表空间
1. 表空间物理层的容器可以使用文件或者裸设备
2. 存储空间的分配和管理由数据库完成
3. DMS默认定义为大型表空间，还可以定义常规大小的表空间
#### 自动存储管理的表空间
1. 数据库管理器负责创建和扩展容器
2. 将有存储管理的表空间转为数据库统一管理
3. 简化了表空间存储的管理
### 创建表空间
```
#进入实例
#db2start
#连接数据库
db2 connnect to work
#创建一个最简单的表空间
#db2 create tablespace worksimple
#创建一个大型表空
db2create large tablespace worklarge
#创建用户临时表空间
#db2 create user temporary tablespace workusertemp
#创建4k页大小的表空间
#db2 create tablespace work4k pagesize 4k
```
### 查看表空间信息
```
#db2 list tablespaces
#更多内容
#db2 list tablespaces show detail
```
### 查看表空间容器
```
#ID为表空间的ID可以db2 list tablespaces查看
#db2 list tablespace containers for ID
#db2 list tablespace containers for ID showdetail
```
