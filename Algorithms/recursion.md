# 递归
* __递归__ 是解决问题的一种方法,他将问题不断分解成更小的子问题,直到子问题可以用普通的方法解决。
* 通常情况下,递归会使用一个不停调用自己的函数,在某些问题上,递归可以帮助我们写出非常优雅的解决方案。

## 递归示例
* 求列表之```[1,3,5,7,9]```和,可以使用循环遍历,也可以使用递归
* 步骤分解: 
    * 总和 = (1 + (3 + (5 + (7 + 9)))) 
    * 总和 = (1 + (3 + (5 + 16))) 
    * 总和 = (1 + (3 + 21))
    * 总和 = (1 + 24)
    * 总和 = 25

```python
listSum(numList) = first(numList) + listSum(rest(numList))
"""
递归的列表求和
"""
def list_sum(num_list):
    if len(num_list) == 1:
        return num_list[0]
    else:
        print (num_list)
        return num_list[0] + list_sum(num_list[1:])
```

* 整数装换为任意进制的字符串(可以取余数用栈的后进先出实现),这里使用递归实现
* 步骤分解(十进制->二进制):
    * 10 / 2 = 5 ...0
    * 5 / 2 = 2 ...1
    * 2 / 2 = 1 ... 0
    * 1 < 2 = 0 ... 1

```python
"""
1.不断的做除法运算取余数,直到小于基数
2.通过查表法将余数对应的字符串取出来

"""
def to_str(n,base):
    converString = '0123456789ABCDEF'
    if n<base:
        return converString[n]
    else:
        #可以不做字符串拼接,而是将余数结果压入栈中,最后统一出栈获取结果
        return to_str(n/base,base) + converString[n%base]
```

## 递归三原则
* 递归算法必须有__基本情况__(算法停止递归的条件)
* 递归算法必须改变其状态并向__基本情况__靠近
* 递归算法必须递归地调用自己。


## 复杂的递归问题
### 汉诺塔

```python
"""
汉诺塔
ABC 三个圆盘。A圆盘上按照顶部小,底部小堆积
1:每次只能移动一个盘子
2:大盘子不能放到小盘子上
3:将A上的圆盘移动到B上
"""

def move_tower(n,A,B,C):
    """
    圆盘移动
    Args:
        n:圆盘数量
        A:第一根圆柱
        B:第二根圆柱
        C:第三根圆柱
    """
    if n == 1:
        print ("moving disk:{}->{}".format(A,C))
    else:
        move_tower(n-1,A,C,B) #移动到中间柱子
        print ("moving disk:{}->{}".format(A,C))
        move_tower(n-1,B,A,C) #中间注意依次移动到最后一根柱子
```


