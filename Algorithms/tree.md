# 树
* python 实现树有两种方式 __列表之列表__ , __节点之引用__

## 列表之列表
* 在 __列表之列表__ 的树中,我们将根节点的值作为列表的第一个元素;第二个元素是代表左子树的列表;第三个元素是代表右子树的列表。

<a href="https://sm.ms/image/cz5hxdSNufanmUQ" target="_blank"><img src="https://i.loli.net/2020/01/30/cz5hxdSNufanmUQ.png" ></a>

* 列表之列表的简单表示

```python
myTree = [
    'a',#根节点
    ['b',['d',[],[]],['e',[],[]]],#左字树
    ['c',['f',[],[]]],#右子树
]
out:
(py37) ➜  search-sort git:(develop) ✗ python a.py
根节点:a
左子树:['b', ['d', [], []], ['e', [], []]]
右子树:['c', ['f', [], []]]
```
* 列表之列表的pyton实现

```python
def BinaryTree(r):
    """
    添加二叉树的新节点
    """
    return [r, [], []]

def insertLeft(root, newBranch):
    """
    插入左子树
    """
    t = root.pop(1)
    #如果左子树有内容,需要将旧的左子树作为新节点的左子树
    if len(t) > 1:
        root.insert(1, [newBranch,t,[]])
    else:
        root.insert(1,[newBranch,[],[]])

    return root

def insertRight(root, newBranch):
    """
    插入右子树
    """
    t = root.pop(2)
    if len(t) > 1:
        root.insert(2,[newBranch,[],t])
    else:
        root.insert(2, [newBranch,[],[]])

    return root

def getRootVal(root):
    return root[0]

def setRootVal(root, newVal):
    root[0] = newVal

def getLeftChild(root):
    return root[1]

def getRightChild(root):
    return root[2]

if __name__ == '__main__':
    r = BinaryTree(3)
    insertLeft(r,4)
    insertLeft(r,5)
    insertRight(r,6)
    insertRight(r,7)
    l = getLeftChild(r)
    print (l)
    setRootVal(l,9)
    print (r)
```
## 节点与引用
* 定义一个类,其中有根节点和左右子树的属性。这种表示法遵循面向对象范式。
<a href="https://sm.ms/image/d9AC2Rrmvh3Bu67" target="_blank"><img src="https://i.loli.net/2020/01/30/d9AC2Rrmvh3Bu67.png" ></a>

```python
class BinaryTree:
    """
    二叉树的抽象数据模型
    """

    def __init__(self, rootObj):
        self.key = rootObj
        self.leftChild = None
        self.rightChild = None

    def insertLeft(self, newNode):
        if self.leftChild == None:
            self.leftChild = BinaryTree(newNode)
        else:
            t = BinaryTree(newNode)
            t.left = self.leftChild
            self.leftChild = t


    def insertRight(self,newNode):
        if self.rightChild == None:
            self.rightChild = BinaryTree(newNode)
        else:
            t = BinaryTree(newNode)
            t.right = self.rightChild
            self.rightChild = t

    def getRightChild(self):
        return self.rightChild

    def getLeftChild(self):
        return self.leftChild

    def setRootVal(self, obj):
        self.key = obj

    def getRootVal(self):
        return self.key
out:
In [1]: from binary_tree1 import BinaryTree
In [2]: r = BinaryTree('a')
In [3]: r.getRootVal()
Out[3]: 'a'
In [4]: r.insertLeft('b')
In [5]: r.insertRight('c')
In [6]: print (r.getRightChild().getRootVal())
c
In [7]: r.getRightChild().setRootVal('hello')
In [8]: print (r.getRightChild().getRootVal())
hello
```

## 二叉树的应用
* 解析树
* 树的遍历

### 解析树
* 解析树可以用来表示现实世界中像 __句子__ 和 __数学表达式__ 这样的构造
<a href="https://sm.ms/image/MEm5GNQFXRnVPty" target="_blank"><img src="https://i.loli.net/2020/01/30/MEm5GNQFXRnVPty.png" ></a>
<a href="https://sm.ms/image/VShWEpBf3TMX2li" target="_blank"><img src="https://i.loli.net/2020/01/30/VShWEpBf3TMX2li.png" ></a>

* 根据表达式 __构造解析树__
    * 如果当前标记是(，就为当前节点添加一个左子节点，并下沉至该子节点;
    * 如果当前标记在列表['+', '-', '/', '\*']中,就将当前节点的值设为当前标记对应的运算符;为当前节点添加一个右子节点，并下沉至该子节点;
    * 如果当前标记是数字，就将当前节点的值设为这个数并返回至父节点;
    * 如果当前标记是)，就跳到当前节点的父节点。
* 表达式 (3 + (4 \* 5)) 拆分成列表,['(', '3', '+', '(', '4', '\*', '5', ')', ')'],构造解析树的过程如果
<a href="https://sm.ms/image/J5el4hS1zZMXwvb" target="_blank"><img src="https://i.loli.net/2020/01/30/J5el4hS1zZMXwvb.png" ></a>
* 子节点可以使用 getLeftChild getRihtChild 追踪。追踪当前节点的父节点则需要使用栈,每当需要下沉节点时,需要将当前节点压入栈中,当要返回父节点时,把栈中中的数据返回就行。

* 构造表达式解析树 python实现

```python
def buildParseTree(fpexp):
    """
    * 如果当前标记是(，就为当前节点添加一个左子节点，并下沉至该子节点;
    * 如果当前标记在列表['+', '-', '/', '\*']中,就将当前节点的值设为当前标记对应的运算符;为当前节点添加一个右子节点，并下沉至该子节点;
    * 如果当前标记是数字，就将当前节点的值设为这个数并返回至父节点;
    * 如果当前标记是)，就跳到当前节点的父节点。
    Args:
        fpexp:表达式,注意使用空格分隔每个字符 exp: "( 3 + ( 4 * 5 ) )"
    Returns:
        eTree:二叉树对象
    """
    fplist = fpexp.split()
    pStack = Stack()
    eTree = BinaryTree('')
    pStack.push(eTree)
    currentTree = eTree

    for i in fplist:
        if i == '(':
            currentTree.insertLeft('')
            pStack.push(currentTree)
            currentTree = currentTree.getLeftChild()
        elif i not in '+-*/)': #数字
            currentTree.setRootVal(eval(i))
            parent = pStack.pop()
            currentTree = parent
        elif i in '+-*/': 
            currentTree.setRootVal(i)
            currentTree.insertRight('')
            pStack.push(currentTree)
            currentTree = currentTree.getRightChild()
        elif i == ')':
            currentTree = pStack.pop()
        else:
            raise ValueError("Unknow Operator: " + i)

    return eTree
```
* 计算二叉解析树的函数

```python
def evalute(parseTree):
    """
    计算二叉解析树
    Args:
        parseTree:表达式二叉树
    Returns:
        计算结果
    """
    opers = {'+':operator.add,'-':operator.sub,
             '*':operator.mul,'/':operator.truediv
    }
    leftC = parseTree.getLeftChild()
    rightC = parseTree.getRightChild()

    if leftC and rightC:
        fn = opers[parseTree.getRootVal()]
        return fn(evalute(leftC), evalute(rightC))
    else:
        return parseTree.getRootVal()
```

### 树的遍历
* __前序遍历__ :先访问根节点,然后递归的前序遍历左子树,最后递归的前序遍历右子树
    * __前序遍历__ :外部函数

    ```python
    from .ParseTree import buildParseTree
    def preorder(tree):
        """
        前序遍历
            Args:
                tree:二叉树对象
        """
        if tree:
            print (tree.getRootVal())
            preorder(tree.getLeftChild())
            preorder(tree.getRightChild())

    ```
    * __前序遍历__ :类方法

    ```python
    def preorder(self):
        print (self.key)
        if self.leftChild:
            self.leftChild.preorder()
        if self.rightChild:
            self.rightChild.preorder()
    ```

* __中序遍历__ : 先递归地中序遍历左子树，然后访问根节点，最后递归地中序遍历右子树。

    ```python
    from ParseTree import buildParseTree

    def inorder(tree):
        if tree != None:
            inorder(tree.getLeftChild())
            print (tree.getRootVal())
            inorder(tree.getRightChild())


    ```

* __后序遍历__ : 先递归地后序遍历右子树，然后递归地后序遍历左子树，最后访问根节点。

    ```python
    from ParseTree import buildParseTree

    def postorder(tree):
        if tree != None:
            postorder(tree.getLeftChild())
            postorder(tree.getRightChild())
            print (tree.getRootVal())
    ```


## 利用二叉堆实现优先级队列
* 队列有一个重要的变体,叫做 __优先级队列__ 。和队列一样,优先级队列从头部移除元素，不过元素的逻辑顺序由 __优先级__ 决定。优先级最高的元素在最前，优先级最低的元素在最后。
* 二叉堆有俩个最常见的变体: __最小堆__ (最小的元素一直在队首) 与 __最大堆__ (最大的元素一直在队首)。
###  二叉堆的实现
* 二叉堆必须是一颗平衡二叉树: __平衡二叉树__ 是指，其根节点的左右子树含有数量大致相等的节点。在实现二叉堆时，我们通过创建一颗 __完全二叉树__ 来维持树的平衡。在完全二叉树中,除了最底层，其他每一层的节点都是满的。在最底层，我们从左往右填充节点。

<a href="https://sm.ms/image/ADoP3MuUYEQz4yw" target="_blank"><img src="https://i.loli.net/2020/01/31/ADoP3MuUYEQz4yw.png" ></a://sm.ms/image/ADoP3MuUYEQz4yw" target="_blank"><img src="https://i.loli.net/2020/01/31/ADoP3MuUYEQz4yw.png" ></a>
* 我们可以使用列表数据结构来表示 __完全二叉树__ 对于位置为P的节点
    * 左子节点:2p,右子节点: 2p+1 
    * 父节点: p // 2

<a href="https://sm.ms/image/zsuUH5XNoj9eCJI" target="_blank"><img src="https://i.loli.net/2020/01/31/zsuUH5XNoj9eCJI.png" ></a>

* 二叉堆是有序的: __堆的有序性__ : 对于堆中的任意元素x及其父元素p,p都不大于x。完全二叉树具备堆的有序性。

```python
class BinHeap:
    """
    二叉堆的实现
    Attributes:
        heapList:初始化二叉堆列表
        currentSize:二叉堆的大小
    """

    def __init__(self):
        self.heapList = [0]
        self.currentSize = 0

    def buildHeap(self,alist):
        """
        通过列表构建二叉堆
        """
        i = len(alist) // 2
        self.currentSize = len(alist)
        self.heapList = [0] + alist[:]
        while (i>0):
            self.percDown(i)
            i = i - 1


    def insert(self,k):
        """
        二叉堆中插入新元素,列表中插入元素,可以保证完全数的性质,但破坏了堆的结构。
        需要逐次比较与父元素的大小,移动元素。
        Args:
            k:插入的元素
        """
        self.heapList.append(k)
        self.currentSize = self.currentSize + 1
        self.percUp(self.currentSize)

    def percUp(self, i):
        while i // 2 > 0:
            if self.heapList[i] < self.heapList[i // 2]:
                tmp = self.heapList[i // 2]
                self.heapList[i // 2] = self.heapList[i]
                self.heapList[i] = tmp
            i = i // 2

    def delMin(self):
        """
        移除堆的最小元素:
        列表的根节点是最小元素,可以直接移除。移除后,需要保证二叉堆的结构性和有序性
        结构性:取出列表中的最后一个元素,将其移动到根节点的位置。
        有序性:逐次比较子节点,移动位置。保证堆的有序性
        Returns:
            retval:返回最小元素
        """
        retval = self.heapList[1]
        self.heapList[1] = self.heapList[self.currentSize]
        self.currentSize = self.currentSize - 1
        self.heapList.pop()
        self.percDown(1)
        return retval

    def percDown(self,i):
        while (i * 2) <= self.currentSize:
            mc = self.minChild(i)
            if self.heapList[i] > self.heapList[mc]:
                tmp = self.heapList[i]
                self.heapList[i] = self.heapList[mc]
                self.heapList[mc] = tmp
            i = mc

    def minChild(self, i):
        if i * 2 + 1 > self.currentSize:
            return i * 2
        else:
            if self.heapList[i*2] < self.heapList[i*2+1]:
                return i * 2
            else:
                return i * 2 + 1
```

## 二叉搜索树
* __二叉搜索树__ 依赖于这样一个性质:小于父节点的键都在左子树中，大于父节点的键则都在右子树中。我们称这个性质为二叉搜索性，可以用这个性质来实现 __映射__

<a href="https://sm.ms/image/gPUDYeRnNZXVh7A" target="_blank"><img src="https://i.loli.net/2020/01/31/gPUDYeRnNZXVh7A.png" ></a>
* 二叉树的节点数据结构

```python
class TreeNode:

    def __init__(self,key,val,left=None,right=None,parent=None):
        self.key = key
        self.payload = val
        self.leftChild = left
        self.rightChild = right
        self.parent = parent

    def hasLeftChild(self):
        return self.leftChild

    def hasRightChild(self):
        return self.rightChild

    def isLeftChild(self):
        return self.parent and \
                self.parent.leftChild == self

    def isRightChild(self):
        return self.parent and \
                self.parent.rightChild == self

    def isRoot(self):
        return not self.parent

    def isLeaf(self):
        return not (self.rightChild or self.leftChild)

    def hasAnyChildren(self):
        return self.rightChild or self.leftChild

    def hasBothChildren(self):
        return self.rightChild and self.leftChild

    def replaceNodeData(self, key, value, lc, rc):
        self.key = key
        self.payload = value
        self.leftChild = lc
        self.rightChild = rc
        if self.hasLeftChild():
            self.leftChild.parent = self
        if seld.hasRightChild():
            self.rightChild.parent = self
```
## AVL树
