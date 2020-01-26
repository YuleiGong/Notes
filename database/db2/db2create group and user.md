## create group
### 创建组和用户，并为用户指定组
```
#sudo groupadd db2adm4
#受防护的组
#sudo groupadd db2fen4
#实例用户
#sudo useradd -d /home/db2inst4 -m db2inst4 -g db2adm4
#受防护的用户 属于受防护的用户组
#sudo useradd -d /home/db2fenc4 -m db2fenc4 -g db2fen4
```
### 设置密码 123456
```
#sudo passwd db2inst4
#sudo passwd db2fenc4
```
### 创建实例 进入安装目录/opt/db2/instance
```
#sudo ./db2icrt -u db2fenc4 db2inst4
```
### 检验 出现successfully 表示成功
```
#su db2inst4
#启用实例
#db2start
#db2 create database work
#查看当前实例存在的数据库
#db2 list database directory
#连接
#db2 connect  to work
#查看表空间
#db2 list tablespaces
```
