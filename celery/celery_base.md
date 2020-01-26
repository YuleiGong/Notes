# celery base
## celery 包含组件
* Celery __Beat__：任务调度器，Beat进程会读取配置文件的内容，周期性地将配置中到期需要执行的任务发送给任务队列。
* Celery __Worker__：执行任务的消费者，通常会在多台服务器运行多个消费者来提高执行效率。
* __Broker__：消息代理，或者叫作消息中间件，接受任务生产者发送过来的任务消息，存进队列再按序分发给任务消费方（通常使用rabbitmq或redis）。
    * RabbitMQ功能齐全、稳定、耐用并且容易安装。对于生产环境来说是一个很好的选择。
    * Redis也功能齐全，但是在突然中止或者电源故障的情况下更容易出现数据丢失
* __Producer__：调用了Celery提供的API、函数或者装饰器而产生任务并交给任务队列处理的都是任务生产者。
* __Result Backend__：任务处理完后保存状态信息和结果，以供查询。

![](celery_base.png)

## 简单例子
* 目录结构```/Users/gongyulei/code/python_example/proj```
* app.py

```pyhton
from celery import Celery
app = Celery('proj', include=['proj.tasks'])
app.config_from_object('proj.celeryconfig')
if __name__ == '__main__':
    app.start()
```
* celeryconfig.py

```python
BROKEN_URL = "amqp://guest:guest@localhost:5672//"
CELERY_RESULT_BACKEND = 'redis://localhost:6379'
CELERY_TASK_SERIALIZER = 'msgpack'
CELERY_RESULT_SERIALIZER = 'json'
CELERY_TASK_RESULT_EXPIRES = 60 * 60 * 24
CELERY_ACCEPT_CONTENT = ['json', 'msgpack']
```

* tasks.py

```python
from proj.celery import app
@app.task
def add(x, y):
    return x + y
```
* 启动__worker__(消费者)```celery -A proj.app worker -l info```
* 解释器里面执行任务

```python
>> from proj.tasks import add
>>> r = add.delay(1, 3)
>>> r
<AsyncResult: ed80929a-bbd5-4d4e-b837-aa523fdec398>
>>> r.result
4
>>> r.status
u'SUCCESS'
>>> r.backend
<celery.backends.redis.RedisBackend object at 0x10be2e090>
>>> from celery.result import AsyncResult
>>> AsyncResult('ed80929a-bbd5-4d4e-b837-aa523fdec398').get()
4
```
## 指定队列
* 默认的celery 使用了一个叫celery的队列

```python
Listing queues
celeryev.cda8fa57-266a-43cb-8114-0cd68e653836   0
celery@gongyuleideMacBook-Pro.local.celery.pidbox       0
celery  0
```

