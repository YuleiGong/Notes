# 标准I/O库
## 流和FILE对象
* 所有的__I/O函数__都是围绕__文件描述符__的,打开一个文件,即返回一个__文件描述符__
* 对于 __标准I/O库__,操作是围绕__流(stream)__。当使用标准I/O库打开或者创建一个文件时,我们已使一个流与一个文件相关联
* 对于__ASCII字符集__,一个字符用一个字节表示。对于__国际字符集__ ,一个字符可用多个字节表示
* 标准I/O文件流可用于__单字节或多字节(宽)__字符集,流的__定向__决定了所读写的字符是单字节还是多字节的
* 当打开一个流时,标准I/O函数fopen返回一个指向__FILE对象__的指针(也称__文件指针__),该对象通常是一个结构体,包含了标准I/O库管理该流所需的信息(文件描述符,流缓冲区指针,缓冲区长度,缓冲区字符数 等)
* 对一个进程预定义了三个流,并且这三个流可以自动的被进程使用: __标准输入,标准输出,标准错误__,通过预定义的文件指针stdin,stdout,stderr加以引用。

## 缓冲
* 标准I/O库提供__缓冲__的目的是尽可能减少read write的调用次数,提供了如下3种类型的缓冲:
    * __全缓冲__: 该情况下，在填满标准I/O缓冲区后才执行实际的I/O操作,对于驻留在磁盘上的文件通常是由标准I/O库实施__全缓冲__
    * __行缓冲__: 当在输入和输出中遇到__换行符__时,标准I/O库执行I/O操作。只有在写一行后才进行实际的I/O操作。当流涉及到终端的时候(标准输入和标准输出),通常使用__行缓冲__
    * __不带缓冲__: 标准I/O库不对字符进行缓冲存储,标准错误流stderr通常是不带缓冲的,这就使得出错信息可以尽快显示出来,而不管它们是否含有一个换行符
* __冲洗flush__: 指标准I/O缓冲区的写操作。缓冲区可由标准I/O例程在缓冲区满时自动冲洗,或者显示调用fflush冲洗一个流,在unix中,flush有两种概念:
    * 在标准IO库方面缓冲区的内容写到磁盘
    * 在终端驱动程序方面,flush(刷清)丢弃已存储在缓冲区中的内容
* 普适规律:
    * 标准错误是不带缓冲的
    * 如果是指向终端设备的流,则是行缓冲的,否则是全缓冲的。

## 打开流
* 可以使用```fopen freopen fdopen ```打开一个标准的I/O流。
* 除非引用终端设备,流被默认打开是__全缓冲__的,引用__终端设备__ 流是行缓冲的。
* 当一个进程被__正常终止__,则所有带__未写缓冲数据__的IO流都被__冲洗__,所有打开的标准I/O流都被__关闭__

## 读和写流
* 一旦打开了流,可以进行三种不同类型的__读写操作__
    * 每次一个字符的I/O,一次读写一个字符
    * 每次一行IO。
    * 直接I/O,每次IO操作读或写某种数量的对象,而某个对象具有指定的长度,也称为__二进制IO__
* 在大多数实现中,为每个流在__FILE__指针对象中维护了 __出错标志和文件结束标志__

## 二进制IO
* 如果进行二进制IO操作,我们可以一次读或写一个完整的结构(使用fwrite fread)
    * 读或写一个__二进制数组__ 
    * 读或写一个__结构体__

    *** 结合起来就可以读或写一个结构数组 ***
* 使用二进制IO的问题在于,只能读在同一系统上已写的数据

## 实现细节
* __标准IO库__最终都要调用IO例程,每个标准I/O流都有一个与其相关联的文件描述符。

* macos 中三个标准流与终端连接,一个文件流

```
➜  lesson_5 ./a.out
enter any character
hello world
one line to standard error
stream = stdin, line buffered, buffer size = 4096
stream = stdout, line buffered, buffer size = 4096
stream = stderr, unbuffered, buffer size = 1
stream = /etc/passwd, fully buffered, buffer size = 65536
```
* macos 中标准流分别定向到文件

```
➜  lesson_5 ./a.out < /etc/group > std.out 2>std.err
➜  lesson_5 cat std.err
one line to standard error
➜  lesson_5 cat std.out
enter any character
stream = stdin, fully buffered, buffer size = 65536
stream = stdout, fully buffered, buffer size = 65536
stream = stderr, unbuffered, buffer size = 1
stream = /etc/passwd, fully buffered, buffer size = 65536
```
* 当__标准输入,标准输出__,指向终端,默认是__行缓冲__的。__标准错误__ 是不带缓冲的
* 普通的文件系统是__全缓冲__的

## 临时文件
* 使用__ tmpnam tmpfile __ 创建临时文件
    * __tmpnam__ 只是创建了一个文件名字符串
    * __tmpfile__ 创建一个临时二进制文件,在关闭文件或程序结束后自动删除

```c
#include "apue.h"

int 
main()
{
    char name[L_tmpnam], line[MAXLINE];
    FILE *fp;

    printf("%s\n", tmpnam(NULL));
    tmpnam(name);
    printf("%s\n", name);
    
    if ((fp = tmpfile()) == NULL)
        err_sys("tmpfile error");
    //指定的字符写入文件流
    fputs("one line of output \n", fp);
    //重设流位置为文件开头
    rewind(fp);
    //读取文件流中的数据到line中
    if (fgets(line, sizeof(line), fp) == NULL)
        err_sys("fgets error");
    fputs(line, stdout);
    exit(0);
}
```
* __mkdtmp__ 创建了一个目录,该目录有一个唯一的名字
* __mkstmp__ 创建了一个文件,该文件有一个唯一的名字,该文件和__tmpfile__创建的不同,并不会自动删除



