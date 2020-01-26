---
title: 3-文件IO
date: 2019-12-19 19:35:35
categories: 
- apue/Unix

---

# 文件I/O
* 本章的函数都被称为 __不带缓冲的I/O__
## 文件描述符
* 对于内核，所有打开的文件都通过__文件描述符__引用
* 0,1,2是unix系统shell的标准文件描述符, 分别对应
    * 0:__标准输入__ STDIN_FILENO
    * 1:__标准输出__ STDOUT_FILENO
    * 2:__标准错误__ STRDEE_FILENO

## 文件偏移量(current file offset)
* 每一个打开的文件都有一个相关联的__当前文件偏移量__，文件的读写操作都是从当前偏移量开始的
* lseek可以__显示__的设置打开文件的偏移量，返回新的文件偏移量，如果文件描述符指向的是一个管道，网络套接字，则不能设置偏移量返回-1
* 对于__普通文件__,其文件偏移量一定是一个__非负整数__,所以在测试时,应该使用-1

```c
//测试标准输入能否设置便宜量
#include "apue.h"
int 
main(void)
{
    if (lseek(STDIN_FILENO, 0, SEEK_CUR) == -1) 
        printf("cannot seek\n");
    else
        printf("seek OK \n");
    exit(0);
}
out:
➜  apue ./a.out < /etc/passwd
seek OK
➜  apue cat < /etc/passwd |./a.out
cannot seek
```

* 文件的偏移量可以大于文件的当前长度。如果设置的偏移量大于当前文件的长度，对文件的下一次写会加长文件，没有写的字节置为0, 形成一个 __空洞__,可能会引起磁盘使用的增加

```c
#include "apue.h"
#include <fcntl.h>

char	buf1[] = "abcdefghij";
char	buf2[] = "ABCDEFGHIJ";

int
main(void)
{
	int		fd;

	if ((fd = creat("file.hole", FILE_MODE)) < 0)
		err_sys("creat error");

	if (write(fd, buf1, 10) != 10)
		err_sys("buf1 write error");
	/* offset now = 10 */

	if (lseek(fd, 16384, SEEK_SET) == -1)
		err_sys("lseek error");
	/* offset now = 16384 */

	if (write(fd, buf2, 10) != 10)
		err_sys("buf2 write error");
	/* offset now = 16394 */

	exit(0);
}
out:
#空洞部分被补0
➜  lesson_3 ll file.hole
-rw-r--r--  1 gongyulei  staff    16K Aug 12 22:53 file.hole
➜  lesson_3 od -c file.hole
0000000    a   b   c   d   e   f   g   h   i   j  \0  \0  \0  \0  \0  \0
0000020   \0  \0  \0  \0  \0  \0  \0  \0  \0  \0  \0  \0  \0  \0  \0  \0
*
0040000    A   B   C   D   E   F   G   H   I   J
0040012

```

## IO效率
* 进程关闭时,会关闭所打开的文件描述符
* 大多数文件系统在读取文件会采取某种预读技术,试图读入比缓冲区更多的数据到内存中。所以读入某些大小缓冲区的数据消耗时间是差不多的。

## 文件共享 __相关资料:p61__
* 内核使用3种数据结构表示打开文件，他们之间的关系决定了在文件共享时，一个文件对另外一个文件的影响
    1. 每个进程在进程表中都有一个进程表项,包含了一张进程打开文件的 __文件描述表__ 包含了打开的 __文件描述符标志__ ,指向一个文件表项的 __指针__
    2. 内核为每个打开的文件维护了一张文件表( __文件表项__ ) 包含了文件的 __状态标志__(读，写，等), 当前文件的 __偏移量__ ,指向该文件的 __V节点表项__ 指针
    3. __V节点__ 包含了文件信息和文件的操作函数指针,同时包含了 __i节点__(存储文件长度)

## 原子操作问题
* 进程A B共同写一个文件,有2个文件表项,但是共享一个v节点。A调用lseek设置偏移量为1500,内核切换进程B,也设置1500偏移量,然后B调用wrtie写100字节数据，偏移量变为1600,i节点的文件长度更新为1600。此时A的文件偏移量还是1500，调用write写数据，就会把B的数据覆盖
* 在这过程中，先定位到文件的尾端(lseek),然后写(write),使用的是2个不同的函数，在这2个函数调用之间内核可能会挂起进程，去执行另外一个进程。可以在打开文件的时候设置 __O_APPEND__ 标志，每次写操作前，都将当期偏移量设置到文件的末尾
* __原子操作__ 是指由多步执行的一个操作，则要么执行完所有操作，要么一步也不执行
## dup dup2
* dup 和 dup2可以复制一个文件描述符,文件描述符共享同一个文件表项，包括文件状态标志,当前文件偏移量。
* dup2可以显示的指定文件描述符值
## sync fsync 
* 向文件写入数据,数据会先写入缓冲区(__页高速缓存__)，然后排入队列，依次写入磁盘,这种方式称为延迟写
* __sync__ 会将缓冲区的数据排入队列就结束，不等待写入磁盘，内核会有一个update定时调用sync，定时冲洗缓冲区
* __fsync__ 只对文件描述符fb的文件起作用，写磁盘结束才返回

## fcntl
* fcntl 可以改变或获取文件的描述符和状态标志,这在无法得到shell打开的文件名时很有用,只需要知道描述符,就修改属性

```c
#include "apue.h"
#include <fcntl.h>

int
main(int argc, char *argv[])
{
	int	    val;

	if (argc != 2)
		err_quit("usage: a.out <descriptor#>");

	if ((val = fcntl(atoi(argv[1]), F_GETFL, 0)) < 0)
		err_sys("fcntl error for fd %d", atoi(argv[1]));

	switch (val & O_ACCMODE) {
	case O_RDONLY:
		printf("read only");
		break;

	case O_WRONLY:
		printf("write only");
		break;

	case O_RDWR:
		printf("read write");
		break;

	default:
		err_dump("unknown access mode");
	}

	if (val & O_APPEND)
		printf(", append");
	if (val & O_NONBLOCK)
		printf(", nonblocking");
	if (val & O_SYNC)
		printf(", synchronous writes");

#if !defined(_POSIX_C_SOURCE) && defined(O_FSYNC) && (O_FSYNC != O_SYNC)
	if (val & O_FSYNC)
		printf(", synchronous writes");
#endif

	putchar('\n');
	exit(0);
}
```

* 1:标准输入 2:标准输出 3:错误 5:shell <>:标识用文件描述符5打开数据
* 在使用fctnl修改文件描述符或文件状态值时,需要先获得现在的标志值,然后按照期望修改
* 通过使用fsync 或者改变文件状态标志位O_SYNC时可以获得 __延迟写__

```c
out :
➜  lesson_3 ./a.out 0</dev/tty
read only
➜  lesson_3 ./a.out 1>temp.foo 
➜  lesson_3 cat temp.foo 
write only
➜  lesson_3 ./a.out 2 2>>temp.foo  
write only, append
➜  lesson_3 ./a.out 5 5<>temp.foo 
read write
```
