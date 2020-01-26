# celery 生产环境使用

* celery的配置文件统一放到settings目录中
```python
/Users/gongyulei/work/sn_dev/kanas/settings
celeryconfig.py
```

* 使用__redis__ 作为broker 和 backed

```python
#celeryconfig
#!/usr/bin/env python
# -*- coding: utf-8 -*-

from __future__ import unicode_literals
from __future__ import absolute_import


broker_url = 'redis://192.168.100.32:6379/8'

#必须开启result_backend才能获取task执行的状态信息
result_backend = 'redis://192.168.100.32:6379/8'

result_expires = 10

timezone = 'Asia/Shanghai'

imports = (
    'tasks.export_excel',
    'tasks.import_excel',
    'tasks.test_task'
)

#任务的路由，按需修改'*'的路由到指定的query，celery的启动脚本中的Queue的配置做对应的修改
task_routes = {
        'celery.ping': 'default',
        '*': 'weps_celery'
        }
```

* 使用__superisor__统一维护worker,-c 指定了任务数,-n 制定了workname,-A指定celery APP

```
/Users/gongyulei/work/aml/etc/supervisord
# vim: set ft=dosini syntax=dosini :
[program:celery_draco]
;directory = 
command = celery worker --purge -A dracon -l info -c 16 -n 'draco@%%h'
autostart = true
startsecs = 5
autorestart = true
startretries = 3
stopsignal = TERM
stopwaitsecs = 10
stopasgroup = true
killasgroup = true
redirect_stderr = true
stdout_logfile_maxbytes = 100MB
stdout_logfile_backups = 3
stdout_logfile = var/log/celery_draco.log
;environment=

```

* 在PYTHONPATH 所在路径范围编写异步任务处理函数 这将会与celert worker -A 对应,初始化好celery app后,直接使用修饰器即可使用celery

```python
app = Celery('worker')
app.config_from_object('settings.celeryconfig')
from celery import Celery, Task

@app.task(base=True)
def test():
    pass
```

* 使用__rabbitmq__作为broker,使用redis作为backed,通过routing_key设定不同异步函数的优先级,提升负载。

```
#!/usr/bin/env python
# -*- coding: utf-8 -*-

from __future__ import unicode_literals
from __future__ import absolute_import
from kombu import Queue
from kombu import Exchange

import settings

result_serializer = 'json'


process_log_fmt = """
    [%(asctime)s: %(levelname)s/%(processName)s %(filename)s %(lineno)d] %(message)s
""".strip()
task_log_fmt = """[%(asctime)s: %(levelname)s/%(processName)s %(filename)s %(lineno)d] %(task_name)s[%(task_id)s]: %(message)s"""

broker_url = settings.get_corp('default', 'CELERY_BROKER')
backend_url = settings.get_corp('default', 'CELERY_LAKE')

celeryd_prefetch_multiplier=1

worker_max_tasks_per_child=1000

task_queues = (
    Queue('kanas_sync_h',  exchange=Exchange('celery', type='direct'), routing_key='kanas_sync_h'),
    Queue('kanas_sync_m',  exchange=Exchange('celery', type='direct'), routing_key='kanas_sync_m'),
    Queue('kanas_sync_l',  exchange=Exchange('celery', type='direct'), routing_key='kanas_sync_l'),
)

task_routes = ([
    ('klake.worker.sync_database',  {'queue': 'kanas_sync_h'}),
    ('klake.worker.sync_schema',    {'queue': 'kanas_sync_h'}),
    ('klake.worker.sync_table',     {'queue': 'kanas_sync_m'}),
    ('klake.worker.sync_sql',       {'queue': 'kanas_sync_m'}),
    ('klake.worker.sync_snippet',   {'queue': 'kanas_sync_l'}),
    ('klake.worker.drop_database',  {'queue': 'kanas_sync_h'}),
    ('klake.worker.drop_table',     {'queue': 'kanas_sync_l'}),
    ('klake.worker.stats_database', {'queue': 'kanas_sync_h'}),
    ('klake.worker.stats_table',    {'queue': 'kanas_sync_m'}),
    ('*',                           {'queue': 'kanas_sync_l'})
 ],)
```

* 使用__superisor__统一维护worker,-c 指定了任务数,-n 制定了workname,-A指定celery APP,-Q选项,指定了不同类型的worker绑定的任务,需要和celeryconfig.py对照


```python
[program:ksync_h]
;directory = 
command = celery -A klake.worker worker -Q kanas_sync_h --concurrency=1 -l info -f var/log/kanas_sync_h.log -E -n 'kanas_sync_h@node44_kanas' --heartbeat-interval 30 --prefetch-multiplier 1
autostart = true
startsecs = 5
autorestart = true
startretries = 3
stopsignal = TERM
stopwaitsecs = 10
stopasgroup = true
killasgroup = true
redirect_stderr = true
stdout_logfile_maxbytes = 100MB
stdout_logfile_backups = 3
stdout_logfile = var/log/kanas_sync_h.log
```
## other
* 监控,使用celery flower 监控任务执行情况。

* 可以使用celery -A WORKER amqp queue.purge QUEUE_NAME 清空队列数据




