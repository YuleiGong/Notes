## 安装前准备
###  安装所需依赖
* yum 安装依赖包
```sh
yum install make gcc binutils gcc-c++ compat-libstdc++ elfutils-libelf-devel elfutils-libelf-devel-static ksh libaio libaio-devel numactl-devel sysstat unixODBC unixODBC-devel pcre-devel glibc.i686
```
* 建立用户和用户组
```sh
groupadd oinstall
groupadd dba
useradd -g oinstall -G dba -d /home/oracle oracle
passwd oracle          //设置oracle密码 passwd:oracle
```
* 目录准备及权限调整
```sh
mkdir -p /export/servers/oracle/11.2.0  //数据库系统安装目录
mkdir -p /export/data/oradata    //数据库数据安装目录
mkdir /export/data/oradata_back  //数据备份目录
mkdir /home/oracle/inventory //清单目录
chown -R oracle:oinstall /export/servers/oracle
chown -R oracle:oinstall /home/oracle/inventory
chown -R oracle:oinstall /export/data
chmod -R 775 /export/servers/oracle
chmod-R 775 /export/data
``` 
* 内核参数调整

```sh
vim /etc/sysctl.conf 在文件最后增加
fs.aio-max-nr = 1048576
fs.file-max = 6553600
kernel.shmall = 2097152
kernel.shmmax = 2147483648
kernel.shmmni = 4096
kernel.sem = 250 32000 100 128
net.ipv4.ip_local_port_range = 1024 65000
net.core.rmem_default = 262144
net.core.rmem_max = 4194304
net.core.wmem_default = 262144
net.core.wmem_max = 1048586
保存文件。
/sbin/sysctl -p          //让参数生效
```

* 用户的限制文件修改

```sh
#vim /etc/security/limits.conf 在文件后增加
oracle           soft    nproc           2047
oracle           hard    nproc           16384
oracle           soft    nofile          1024
oracle           hard    nofile          65536
oracle           soft    stack           10240
保存文件。

修改/etc/pam.d/login文件，增加如下：
session  required   /lib64/security/pam_limits.so
session     required      pam_limits.so
修改/etc/profile,增加：
if [ $USER = "oracle" ]; then
 if [ $SHELL = "/bin/ksh" ]; then
  ulimit -p 16384
  ulimit -n 65536
 else
  ulimit -u 16384 -n 65536
 fi
fi
```

## 开始安装
* 解压安装文件到/home/oracle/database目录中
* 复制一份安装应答文件

```sh
cp -R /home/oracle/database/response/db_install.rsp  /home/oracle/database/response/my_db_install.rsp
```

* 修改应答文件

```sh
oracle.install.option=INSTALL_DB_SWONLY #指定安装选项
ORACLE_HOSTNAME=oracle11g.jd.com 
UNIX_GROUP_NAME=oinstall
INVENTORY_LOCATION=/home/oracle/inventory/ #指定清单目录
ORACLE_HOME=/export/servers/oracle/11.2.0
ORACLE_BASE=/export/servers/oracle
oracle.install.db.InstallEdition=EE #指定安装版本为企业版
oracle.install.db.isCustomInstall=false
oracle.install.db.DBA_GROUP=dba
oracle.install.db.OPER_GROUP=dba
DECLINE_SECURITY_UPDATES=true
```

* 开始安装

```sh
./runInstaller -silent -responseFile /home/oracel/database/response/my_db_install.rsp
```

* 成功后，切换root执行下列脚本

```sh
sh /home/oracle/inventory/orainstRoot.sh
sh /export/servers/oracle/11.2.0/root.sh
```

* 更改oracle 的bash_profile

```sh
export ORACLE_SID=orcl
export ORACLE_BASE=/export/servers/oracle
export ORACLE_HOME=$ORACLE_BASE/11.2.0
export LD_LIBRARY_PATH=$ORACLE_HOME/lib:/lib:/usr/lib
PATH=$PATH:$ORACLE_HOME/bin:$HOME/bin
export PATH
```

## 建立数据库
* 新增自己的建库文件

```sh
cd ~oracle/database/response/
cp dbca.rsp my_dbca.rsp
```
* 修改建库文件

```sh
OPERATION_TYPE = "createDatabase"
GDBNAME = "orcl11g"
SID = "orcl"
SYSPASSWORD = "oracle"
SYSTEMPASSWORD = "oracle"
DATAFILEDESTINATION = /export/data/oradata
RECOVERYAREADESTINATION = /export/data/oradata_back
SYSDBAUSERNAME = "system"
SYSDBAPASSWORD = "oracle"
INSTANCENAME = "orcl11g"
CHARACTERSET = "UTF8"
NATIONALCHARACTERSET= "UTF8" 
```
* 创建数据库

```
dbca -silent -responseFile /home/oracel/response/my_dbca.rsp
```
## 配置监听

* 启动监听
```sh
netca /silent /responsefile /home/oracle/database/response/netca.rsp
```

* 修改监听文件
```sh
cd /export/servers/oracle/11.2.0/network/admin/samples
cp listener.ora ../
#修改如下内容
LISTENER =
  (DESCRIPTION_LIST =
    (DESCRIPTION =
      (ADDRESS = (PROTOCOL = IPC)(KEY = EXTPROC1521))
      (ADDRESS = (PROTOCOL = TCP)(HOST = 0.0.0.0)(PORT = 1521))
    )
  )
SID_LIST_LISTENER =
(SID_LIST =
  (SID_DESC =
  (GLOBAL_DBNAME = orcl)
  (SID_NAME = orcl)
  )
)
```
```bash
cd /export/servers/oracle/11.2.0/network/admin/samples
cp tnsnames.ora ../
#注释文件 添加以下内容

```
* 重启监听

```sh
lsnrctl reload
```

* 编辑 /etc/oratab 把 orcl:/export/servers/oracle/11.2.0:N的‘N’，改为‘Y’，这样就可以通过dbstart启动此实例，也可以通过dbshut关闭此实例了。
* sqlplus / as sysdba 连接测试

## python 远程连接
* 需要安装instantclient_11_2 并设置
```sh
export ORACLE_HOME=${PROJ_DIR}/instantclient_11_2
export LD_LIBRARY_PATH=${PROJ_DIR}/instantclient_11_2:$LD_LIBRARY_PATH
export DYLD_LIBRARY_PATH=${LD_LIBRARY_PATH}
```
```python
import cx_Oracle
if __name__ == '__main__':
    db = cx_Oracle.connect('system', 'oracle', '172.16.156.130:1521/orcl')
    cr = db.cursor()
```




