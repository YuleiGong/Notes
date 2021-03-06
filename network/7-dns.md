# DNS

* DNS 树状结构
![nds_tree.jpg](https://i.loli.net/2021/03/21/9LWsJQGpywSnfxo.jpg)

## DNS 解析流程

![dns解析流程.jpg](https://i.loli.net/2021/03/21/mvlFKo5sh1if3RC.jpg)

## 负载均衡

### 内部负载均衡

* 某个应用要访问另外一个应用，如果配置另外一个应用的 IP 地址，那么这个访问就是一对一的。但是当被访问的应用撑不住的时候，我们其实可以部署多个。只要修改配置IP为域名就可以了。在域名解析的时候，我们只要配置策略，这次返回第一个 IP，下次返回第二个 IP，就可以实现负载均衡了。

### 全局负载均衡

* 为了保证我们的应用高可用，往往会部署在多个机房，每个地方都会有自己的IP地址。当用户访问某个域名的时候，这个 IP 地址可以轮询访问多个数据中心。如果一个数据中心因为某种原因挂了，只要在DNS服务器里面，将这个数据中心对应的IP地址删除，就可以实现一定的高可用。

## 缺陷

* 域名缓存问题：缓存更新不及时。
* 域名转发问题： 会导致跨运营商访问。
* 出口 NAT 问题:  会导致跨运营商访问。
* 域名更新问题
* 解析延迟问题

# HTTP DNS

* HttpDNS 其实就是，不走传统的DNS解析，而是自己搭建基于HTTP协议的 DNS 服务器集群，分布在多个地点和多个运营商。当客户端需要 DNS 解析的时候，直接通过 HTTP 协议进行请求这个服务器集群，得到就近的地址。

## HTTP DNS 工作模式

![http_dns.jpg](https://i.loli.net/2021/03/21/p4PhGsYnEe5HVBb.jpg)

* 使用 HttpDNS 需要绕过默认的 DNS 路径，就不能使用默认的客户端。使用 HttpDNS 的，往往是手机应用，需要在手机端嵌入支持 HttpDNS 的客户端 SDK。
* 在客户端的 SDK 里动态请求服务端，获取 HttpDNS 服务器的 IP 列表，缓存到本地。随着不断地解析域名，SDK 也会在本地缓存 DNS 域名解析的结果。当手机应用要访问一个地址的时候，首先看是否有本地的缓存，如果有就直接返回。这个缓存和本地 DNS 的缓存不一样的是，这个是手机应用自己做的，而非整个运营商统一做的。如何更新、何时更新，手机应用的客户端可以和服务器协调来做这件事情。
* 手机客户端自然知道手机在哪个运营商、哪个地址。由于是直接的 HTTP 通信，HttpDNS 服务器能够准确知道这些信息，因而可以做精准的全局负载均衡。

