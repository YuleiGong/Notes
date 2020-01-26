# 网络IPC:套接字
* 套接字通信:即可用于计算机间通信,也可用于计算机内通信

## 套接字描述符
* 套接字是通信端点的抽象。套接字描述符在UNIX系统中被当做是一种__文件描述符__。
* 调用__socket__可以创建一个套接字,返回套接字描述符。

```c
int socket(int domain, int type, int protocol)
```
* __domain__(域)确定通信的特性。各个域的参数都以AF开头(address familyi-地址族)。
    <a href="https://sm.ms/image/RZgABUpenDIVfcz" target="_blank"><img src="https://i.loli.net/2019/12/25/RZgABUpenDIVfcz.png" ></a>
* __type__ 可以确定域的类型
    <a href="https://sm.ms/image/RZgABUpenDIVfcz" target="_blank"><img src="https://i.loli.net/2019/12/25/RZgABUpenDIVfcz.png" ></a>
    * SOCK_STREAM 是有保障的(即能保证数据正确传送到对方)面向连接的SOCKET，多用于资料(如文件)传送,字节流要求在交互数据之前,在本地套接字和通信的对等进程的套接字之间建立一个逻辑连接。SOCK_STREAM提供字节流服务,所以应用程序分辨不出报文的界限,要得到所有发送过来的数据,可能需要经过若干次函数调用。
    * SOCK_DGRAM 是无保障的面向消息的socket,主要用于在网络上发广播信息。对于数据报接口,两个对等进程之间不需要逻辑连接,只需要向对等进程所使用的套接字送出一个报文。

## 建立连接
* 在处理一个面向连接的网络服务,在开始交换数据以前,需要在客户端和服务器之间建立一个连接。(使用connect函数来建立连接)
* 在尝试连接服务器时,出于各种原因,可能会连接失败。要想一个连接请求成功,要连接的计算机必须是开启的,并且是正在运行,服务器的等待队列要有足够的空间
* 如果套接字描述符是非阻塞模式,如果连接不能马上建立,connect会放回-1,并设置错误标志,应用程序可以使用poll或者select来判断文件描述符何时可写。如果可写,连接完成。

```c
int connect(int sockfd, const struct sockaddr *addr, socklen_t len)
```

```c
//处理连接错误的示例
/*如果套接字失败,需要关闭套接字,如果想重试,必须打开一个新的套接字*/
#include "apue.h"
#include <sys/socket.h>

#define MAXSLEEP 128

int
connect_retry(int domain, int type, int protocol,
              const struct sockaddr *addr, socklen_t alen)
{
	int numsec, fd;

	for (numsec = 1; numsec <= MAXSLEEP; numsec <<= 1) {
		if ((fd = socket(domain, type, protocol)) < 0)
			return(-1);
		if (connect(fd, addr, alen) == 0) {
			/*
			 * Connection accepted.
			 */
			return(fd);
		}
		close(fd);

		/*
		 * Delay before trying again.
		 */
		if (numsec <= MAXSLEEP/2)
			sleep(numsec);
	}
	return(-1);
}
```
* 服务器调用listen来宣告他愿意接受连接请求

```
//参数backlog指定了系统该进程所要入队的未完成连接请求的数量
int listen(int sockfd, int backlog);
```
* 一旦服务器调用了listen,所用的套接字就能够接收连接请求。使用accept函数获得连接请求,并建立连接。
* 该函数返回连接到调用connect的客户端描述符。传递给accept的原始套接字不会关联到这个连接,而是继续保持可用状态并接收其他连接请求。
* 如果没有连接请求在等待,accept会阻塞,直到一个请求过来。如果sockfd处于非阻塞模式,accept会返回-1,并设置ERROR,可以使用poll或select来等待一个请求到来。

```c
int accept(int sockfd, struct sockaddr *restrict addr,socklen_t *restrict len);
```
* 初始化套接字服务

```c
//初始化套接字服务进程
#include "apue.h"
#include <errno.h>
#include <sys/socket.h>

int
initserver(int type, const struct sockaddr *addr, socklen_t alen,
  int qlen)
{
	int fd;
	int err = 0;

	if ((fd = socket(addr->sa_family, type, 0)) < 0)
		return(-1);
	if (bind(fd, addr, alen) < 0)
		goto errout;
	if (type == SOCK_STREAM || type == SOCK_SEQPACKET) {
		if (listen(fd, qlen) < 0)
			goto errout;
	}
	return(fd);

errout:
	err = errno;
	close(fd);
	errno = err;
	return(-1);
}

```

## 数据传输
