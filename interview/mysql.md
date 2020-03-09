# sql规范 
* select子句中尽量避免使用\*
* where 左侧不要使用函数
* in / not in 少使用会全表扫描  between代替 
* or 少使用。 会全表扫描 使用union代替

* 后端分页:
    * 前端传给后端 页码和每页条数 pagesize:偏移量
    * select * from table limit (page-1)\*pagesize offset pagesize;
    * 数据量过大，可以使用唯一索引或主键 select * from table where demo_id > (pageNo-1)\*pageSize limit pageSize 分页
* 分库分表:
    * 垂直切分:按照字段进行分库分表,不建议使用
    * 水平切分:按照数量进行切分,比如按照时间或id分表，热数据放在一张表中。
    * id取摸分表:数据会比较分散 tb_0 tb_1 tb_2 tb_3 取数的时候，取id%4来确定需要查询的表

