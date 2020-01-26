## 数据的导入和导出
### 导出 EXPORT
1. export 可以使用sql select 语句将数据从数据库表提取到文件中
2. 导出格式为DEL IXF WSF 
```
#db2 "Export to /home/db2inst4/test.txt of del select * from mytable"
导出数据到指定路径为del格式
```
### 导入import
```
#db2 "import from /home/db2inst4/test.txt of del insert into mytable"
```
### export导出
```

```