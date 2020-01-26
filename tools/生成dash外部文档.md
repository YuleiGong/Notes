# 生成dash docset 文档
* 1 下载文档代码库的doc 或docs 文件,应该是rst文件
* 2 ```pip install doc2dash & pip install sphinx_rtd_theme ```
* 3 进入相应的docs 目录执行 ```make html ```
* 4 执行 doc2dash -n [name] docs/build/html 生成docset文档 
    * -n :指定目标名字
* 5 导入dash

