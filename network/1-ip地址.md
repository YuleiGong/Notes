# ifconfig/ip addr

## ifconfig 

* __inet__ 代表了ip地址，通过.分隔成4部分，每个部分8bit。 所以 IP 地址总共是 32 位。
* __inet6__ 有 128 位， 这个就是IPv6，弥补了 32 位Ip地址不够的缺陷。

  ```
  ➜  Notes git:(master) ifconfig lo0
  lo0: flags=8049<UP,LOOPBACK,RUNNING,MULTICAST> mtu 16384
          options=1203<RXCSUM,TXCSUM,TXSTATUS,SW_TIMESTAMP>
          inet 127.0.0.1 netmask 0xff000000
          inet6 ::1 prefixlen 128
          inet6 fe80::1%lo0 prefixlen 64 scopeid 0x1
          nd6 options=201<PERFORMNUD,DAD>
  ➜  Notes git:(master) ifconfig en0
  en0: flags=8863<UP,BROADCAST,SMART,RUNNING,SIMPLEX,MULTICAST> mtu 1500
          options=400<CHANNEL_IO>
          ether 4c:32:75:89:b1:b9
          inet6 fe80::8a8:8171:713f:527a%en0 prefixlen 64 secured scopeid 0x4
          inet 192.168.101.8 netmask 0xffffff00 broadcast 192.168.101.255
          nd6 options=201<PERFORMNUD,DAD>
          media: autoselect
          status: active
  ```

## ip 地址分类

* ![ip地址分类.png](https://i.loli.net/2021/03/15/OVPd7LiWRIz2JeZ.png)
* A B C 三类，主要包含了 __网络号和主机号__。

## 无类型域间选路(CIDR)

* 由于C 类ip地址分类所能代表的主机数量实在太少，于是有了一个折中的方式叫作无类型域间选路，简称CIDR。
* CIDR 打破了原来设计的几类地址的做法，将32位的IP地址一分为二，前面是网络号，后面是主机号。
* inet 10.0.254.25/24，IP地址中有一个斜杠，斜杠后面有个数字24。这种地址表示形式，就是CIDR。后面24的意思是32 位中前24 位是网络号，后8位是主机号。
* 伴随着CIDR存在的，一个是 __广播地址__，10.0.254.255。如果发送这个地址，所有10.0.254 网络里面的机器都可以收到。另一个是 __子网掩码__，255.255.255.0。
* 将 __子网掩码__ 和 IP 地址进行 AND 计算。前面三个255，转成二进制都是1。1和任何数值取AND，都是原来数值，因而前三个数不变，为 10.0.254。后面一个0，转换成二进制是0，0和任何数值取 AND，都是0，因而最后一个数变为0，合起来就是 10.100.122.0。这就是网络号。将子网掩码和IP地址按位计算AND，就可得到 __网络号__。
* 在 IP 地址的后面有个 __scope__，对于eth0这张网卡来讲，是global，说明这张网卡是可以对外的，可以接收来自各个地方的包。对于lo来讲，是host，说明这张网卡仅仅可以供本机相互通信。
* MAC 地址在 IP 地址的上一行是这个被称为MAC 地址，是一个网卡的物理地址，用十六进制，6 个byte 表示。
* MAC 地址的通信范围比较小，局限在一个子网里面。例如，从 192.168.0.2/24 访问 192.168.0.3/24 是可以用 MAC 地址的。一旦跨子网，即从 192.168.0.2/24 到 192.168.1.2/24，MAC 地址就不行了，需要 IP 地址起作用了。

  ```
  ➜  ~ ip addr
  1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1
      link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
      inet 127.0.0.1/8 scope host lo
         valid_lft forever preferred_lft forever
      inet6 ::1/128 scope host
         valid_lft forever preferred_lft forever
  2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
      link/ether 52:54:00:06:69:d9 brd ff:ff:ff:ff:ff:ff
      inet 10.0.254.25/24 brd 10.0.254.255 scope global eth0
         valid_lft forever preferred_lft forever
      inet6 fe80::5054:ff:fe06:69d9/64 scope link
         valid_lft forever preferred_lft forever
  ➜  ~ ifconfig
  eth0      Link encap:Ethernet  HWaddr 52:54:00:06:69:d9
            inet addr:10.0.254.25  Bcast:10.0.254.255  Mask:255.255.255.0
            inet6 addr: fe80::5054:ff:fe06:69d9/64 Scope:Link
            UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
            RX packets:116100001 errors:0 dropped:4854952 overruns:0 frame:0
            TX packets:119060168 errors:0 dropped:0 overruns:0 carrier:0
            collisions:0 txqueuelen:1000
            RX bytes:34064629331 (34.0 GB)  TX bytes:17907072973 (17.9 GB)
  ```


