# 网络与web编程
## 本地安装httpbin

```sh
pip install httpbin gunicorn
gunicorn httpbin:app
```

## 与http服务交互
* 简单的请求使用urllib实现

```python
def main():
    #url = 'http://localhost:8000/get'
    url = 'http://localhost:8000/post'
    parms = {
        'name1':'value1',
        'name2':'value2'
    }
    querystring = parse.urlencode(parms)
    #get
    #u = request.urlopen(url + '?' + querystring)
    #post
    headers = {
        'User-agent':'none/ofyourbusiness',
        'Spam' : 'Eggs'
    }
    req = request.Request(url, querystring.encode('ascii'), headers=headers)
    u = request.urlopen(req)
    resp = u.read()
    print (resp)

```
* 复杂的使用requests实现
* 可以获取到不同形式的response 二进制,字符串,json

```python
def main1():
    import requests
    url = 'http://localhost:8000/post'
    parms = {
        'name1':'value1',
        'name2':'value2'
    }
    headers = {
        'User-agent':'none/ofyourbusiness',
        'Spam' : 'Eggs'
    }
    resp = requests.post(url, data=parms, headers=headers)
    print (resp.text)
    print (resp.json())
    print (resp.content)
```

## 创建TCP服务器
* 使用socketserver创建简单的TCP单线程服务器

```python
#server
from socketserver import BaseRequestHandler, TCPServer

class EchoHandler(BaseRequestHandler):

    def handle(self):
        print ('Got connection from', self.client_address)
        while  True:
            msg = self.request.recv(8192)
            if not msg:
                break
            self.request.send(msg)

if __name__ == '__main__':
    serv = TCPServer(('', 20000), EchoHandler)
    serv.serve_forever()
```
* client

```python
from socket import socket, AF_INET, SOCK_STREAM
def main():
    s = socket(AF_INET, SOCK_STREAM)
    s.connect(('', 20000))
    print (s.send(b'Hello'))
    print (s.recv(8192))

if __name__ == '__main__':
    main()
out_server:
Got connection from ('127.0.0.1', 62852)
out_client:
5                     
b'Hello' 
```
* 使用线程池多线程来管理TCP服务器

```python
if __name__ == '__main__':
    from threading import Thread
    NWORKERS = 16
    serv = TCPServer(('', 20000), EchoHandler)
    for n in range(NWORKERS):
        t = Thread(target=serv.serve_forever)
        t.daemon = True 
        t.start()

    serv.serve_forever()
```
