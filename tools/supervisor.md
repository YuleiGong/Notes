# supervisor
* 安装

```python
pip install supervisor
```

## supervisor 配置
* 配置分两部分
    __server__端 supervisor
    __client__端 supervisorctl和app
* 运行 ```echo_supervisord_conf``` 输出配置文件,也可以重定向
```
echo_supervisord_conf > supervisord.conf
```
* 配置文件内容

```
[unix_http_server]
file=/home/dev/supervisor.sock   ; socket文件位置,supervisorctl 会使用
;chmod=0700                 ; socket 文件的读写权限 默认0700 
;chown=nobody:nogroup       ; socket 文件所属

;[inet_http_server]         ; 提供了web 管理页面
;port=127.0.0.1:9001        ; ip_address:port specifier, *:port for all iface
;username=user              ; default is no username (open server)
;password=123               ; default is no password (open server)

[supervisord]
logfile=/home/dev/supervisord.log ; 日志文件默认位置
logfile_maxbytes=50MB        ; 日志大小 超过会切分
logfile_backups=10           ; 日志保留数量
loglevel=info                ; 日志级别
pidfile=/home/dev/supervisord.pid ; supervisord pidfile; default supervisord.pid
nodaemon=false               ; 是否前台启动,默认后台启动
minfds=1024                  ; 可以打开的文件描述符最小值
minprocs=200                 ; 可以打开的进程数最小值

[supervisorctl]
serverurl=unix://supervisor.sock ; 通过unix socket连接 supervisord 路径与unix_http_server file 路径一致
;serverurl=http://127.0.0.1:9001 ; 通过http方式连接

;包含其他配置文件
[include]
files = conf/*.conf ;可以是.conf .ini
```
* 启动 supervisord
```
supervisord -c supervisord.conf
(py27_env) [dev@localhost ~]$ ps aux |grep 'sup'
root         16  0.0  0.0      0     0 ?        S    Jun15   0:05 [sync_supers]
dev       68962  0.0  0.3 139364  5868 pts/2    S+   01:51   0:00 vim supervisor.md
dev       69222  0.0  0.7 208716 13956 ?        Ss   02:13   0:00 /home/dev/env/py27_env/bin/python2.7 /home/dev/env/py27_env/bin/supervisord -c supervisord.conf
```
## app 配置
* 添加需要管理的配置文件,这些配置文件都可以写到supervisord.conf 中。然后通过include 引入
* app配置中的日志文件夹需要提前建好,并具有777权限

```
[program:hello]
directory = /home/dev ; 程序的启动目录
command = sh hello.sh  ; 启动命令
autostart = true     ; 在 supervisord 启动的时候也自动启动
startsecs = 5        ; 启动 5 秒后没有异常退出，就当作已经正常启动了
autorestart = true   ; 程序异常退出后自动重启
startretries = 3     ; 启动失败自动重试次数，默认是 3
redirect_stderr = true  ; 把 stderr 重定向到 stdout，默认 false
user = dev
stdout_logfile_maxbytes = 20MB  ; stdout 日志文件大小，默认 50MB
stdout_logfile_backups = 20     ; stdout 日志文件备份数
; stdout 日志文件，需要注意当指定目录不存在时无法正常启动，所以需要手动创建目录（supervisord 会自动创建日志文件）
stdout_logfile = /home/dev/logs/hello.log
stopasgroup = true
killasgroup = true
```
* hello.sh 在shell 中source好python 环境变量

```sh
source ~dev/env/py27_env/bin/activate
python hello.py
```

## 运行

```
supervisord -c supervisord.conf
supervisorctl -c supervisord.conf
supervisorctl status
supervisorctl stop
supervisorctl start
sapervisorctl reread
sapervisorctl update
```
