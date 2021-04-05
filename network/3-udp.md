# UDP

* __区别__ : TCP 是面向连接的，UDP 是面向无连接的。

## UDP 头

* UDP 头比较简单。只需要源端口号和目标端口号。
![udp头.jpg](https://i.loli.net/2021/03/17/pMPFTlQ8iOVmyoK.jpg)

## UDP 使用场景
* 需要资源少，在网络情况比较好的内网，或者对于丢包不敏感的应用。
* 不需要一对一沟通，建立连接，而是可以广播的应用。 (DHCP)
* 需要处理速度快，时延低，可以容忍少数丢包，但是要求即便网络拥塞，也毫不退缩，一往无前的时候。(视频直播)
