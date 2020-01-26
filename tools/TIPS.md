# TIPS
## svn
* 回滚

```
svn merge -r 最新版本号:要恢复的版本号 文件名

```
* 查看改动

```
svn cat -r 783 aml.py|less
```

## brew
* 下载最新的__brew__包 ```brew update```
* 查看那些包可以更新 ```brew outdated```
* 锁定不想更新的包 ```brew pin mysql```
* 解锁```brew unpin mysql```
* 更新 ```brew upgrade ```
* 包清理 ```brew cleanup ```

## oracle
* 改变日期的显示格式

```sql
alter session set nls_date_format='yyyy-mm-dd hh24:mi:ss';
```

* 添加字段

```sql
alter table test1
add (name varchar2(30) default ‘无名氏’ not null);
```

* 修改字段属性

```sql
alter table test modify(name varchar(255));
```

* 登陆sqlplus

```
sqlplus / as sysdba
conn name/passwd@db_name
```

* 修改字段名字

```
alter table [表名]  rename column 旧的字段名 to 新的字段名;
```


* ORA-00054

```sql
sqlplus / as sysdba
select session_id from v$locked_object;
SELECT sid, serial#, username, osuser FROM v$session where sid = session_id;
ALTER SYSTEM KILL SESSION 'sid,serial#';
commit;
```
* ORA-28001 密码过期

```sql
sqlplus / as sysdba
alter user risk identified by risk account unlock;

```
* 重启

```
sqlplus / as sysdba
startup
lsnrctl start #启动监听
```
* oracle 建立用户

```
create user risk identified by risk;
grant connect,resource to risk;
create tablespace risk datafile '/export/data/oradata/orcl11g/risk01.dbf' size 2000m autoextend on
alter user risk default tablespace risk;
```

* 全库导入导出

```bash
#导出dmp文件 为空的表不会被导出
exp kreport3_user/kreport3_pwd@cbs
#导入  确保数据库为空
imp kreport3_user/kreport3_pwd@cbs file=kreport.dmp full=y ignore=y
```



## linux 

* 设置swapfile

```sh
su root
mkdir /home/ssd/swap
touch swapfile
dd if=/dev/zero of=/home/ssd/swap/swapfile bs=1024M count=128
mkswap swapfile
swapon /home/ssd/swap/swapfile
cat /etc/fstab
#添加启动
/home/ssd/swap/swapfile swap swap defaults 0 0
```

* grep

```sh
grep 'from risk' -r * --exclude-dir='.svn' --include='*.py'
```

## redis

```
redis-cli
//删除当前数据库中的所有Key
flushdb
//删除所有数据库中的key
flushall
```

## celery

```
#purge 指定队列
celery -A WORKER amqp queue.purge QUEUE_NAME
```


## svn ignore
* 忽略__文件夹__

```sh
svn propset svn:ignore .env ./
svn up
svn ci
```
* 忽略 __文件__

```sh
export SVN_EDITOR=vim
svn propedit svn:ignore /product
#在忽略的文件夹中加入忽略文件 前提是/product已经在版本库
svn ci
```
* 批量忽略 __文件__

```sh
svn propedit svn:ignore modules/*
#依次填写文件夹
svn ci 

```
* 客户端忽略

```
vim ~/.subversion/config
修改global-ignores 属性
```

## svn 切换url

```sh
svn relocate svn://node43
```

## hexo

```
#重新部署
hexo d -g
hexo new page "tags"
```
## ssh 免密登陆需要注意设置ssh 和 .ssh/authorized_keys权限

```
sudo chmod 700 .ssh/
sudo chmod 644 .ssh/authorized_keys
```
