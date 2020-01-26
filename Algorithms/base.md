# 大O记法

## 算法步骤

* 算法步骤T(n):如果将算法的每一步看成基本单位,那么可以将算法的执行时间描述成解决问题的__步骤数__。

* test1:
    * 上述T(n) = 1+n,当n无线大时,T(n) = n,即算法的步骤是线性增长的。
    * 该算法步骤,后面的n对于算法起了决定性的作用。数量级函数的描述就是,当n增长,T(n)增长最快的部分。
    ```python
    #T(n) = n+1
    def sumOfN(n):

        theSum = 0 #赋值语句可以看做步骤1
        for i in range(1,n+1):
            theSum = theSum + i #累加操作在循环里完成了 n次

        return theSum
    ```

## 大O记法
* 算法步骤的__数量级__ 常被称为大O记法。上述例子执行时间是O(n)。T(n)=5n^2 + 27n + 1005 执行时间可以看做的O(n^2)
* 常见的大O函数
    <a href="https://sm.ms/image/U7ACX3Do2NWFZJI" target="_blank"><img src="https://i.loli.net/2020/01/01/U7ACX3Do2NWFZJI.png" ></a>

* test2
    * 算法步骤T(n)=3n^2 + 2n + 4,很容易看出来n起主导作用,时间复杂度=O(n^2)

    ```python
    a = 5 #1
    b = 6 #1
    #3n^2
    for i in range(n):
        for j in range(n):
            x = i * j
            y = j * j
            z = i * j
    #2n
    for k in range(n):
        w =  a * k + 45
        v = b *b

    d = 33 #1
    ```
