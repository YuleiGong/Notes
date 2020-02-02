# 图
* __图__ 的术语及定义
    * __顶点__ : 顶点又称 __节点__ ，是图的基础部分。它可以有自己的名字，我们称作 __键__
    * __边__ : 边是图的另一个基础部分。两个顶点通过一条边相连，表示它们之间存在关系。边既可以是 __单向__ 的，也可以是 __双向__ 的。如果图中的所有边都是单向的，我们称之为 __有向图__ 。
    * __权重__ : 边可以带权重，用来表示从一个顶点到另一个顶点的成本。
    * __路径__ : 路径是由边连接的顶点组成的序列。无权重路径的长度是路径上的边数，有权重路径的长度是路径上 的边的权重之和。
    * __环__ : 环是有向图中的一条起点和终点为同一个顶点的路径。没有环的图被称为 __无环图__ ，没有环的有向图被称为 __有向无环图__ ，简称为 __DAG__。
    <a href="https://sm.ms/image/3PWNMy4VF9vHAit" target="_blank"><img src="https://i.loli.net/2020/02/01/3PWNMy4VF9vHAit.png" ></a>

## 图的抽象数据类型
### 邻接矩阵表示
* 使用二维矩阵来表示图。每一行和每一列都代表一个顶点。交叉处代表边的权重。邻接矩阵的优点是简单,但是很多单元格是空的，存储并不高效。
<a href="https://sm.ms/image/lAwT3BEa8RMzNWV" target="_blank"><img src="https://i.loli.net/2020/02/01/lAwT3BEa8RMzNWV.png" ></a>

### 邻接表
* 在邻接表的实现中,我们为图对象的所有顶点保存一个主列表。同时为每个顶点维护一个列表，记录了与它相连接的顶点。
* 邻接表能够紧凑的表示稀疏图，此外，邻接表也有助于方便的找到某一个顶点相连的其他所有顶点
<a href="https://sm.ms/image/4w8O52dhBtW76Jg" target="_blank"><img src="https://i.loli.net/2020/02/01/4w8O52dhBtW76Jg.png" ></a>

* python 实现

```python
class Vertex:
    """
    存储图中的每个顶点
    Attributes:
        id:顶点
        connectedTo:与顶点相连的其他顶点
    """

    def __init__(self, key):
        self.id = key
        self.connectedTo = {}

    def addNeighbor(self, nbr, weight=0):
        self.connectedTo[nbr] = weight

    def __str__(self):
        return str(self.id) + ' connectedTo: '\
                + str([x.id for x in self.connectedTo])

    def getConnections(self):
        return self.connectedTo.keys()

    def getId(self):
        return self.id

    def getWeight(self, nbr):
        return self.connectedTo[nbr]
class Graph:
    def __init__(self):
        self.vertList = {}
        self.numVertices = 0

    def addVertex(self,key):
        self.numVertices = self.numVertices + 1
        newVertex = Vertex(key)
        self.vertList[key] = newVertex

        return newVertex

    def getVertex(self, n):
        if n in self.vertList:
            return self.vertList[n]
        else:
            return None

    def __contains__(self, n):
        return n in self.vertList

    def addEdge(self, f, t, const=0):
        if f not in self.vertList:
            nv = self.addVertex(f)
        if t not in self.vertList:
            nv = self.addVertex(t)
        self.vertList[f].addNeighbor(self.vertList[t], const)

    def getVertices(self):
        return self.vertList.keys()

    def __iter__(self):
        return iter(self.vertList.values())


if __name__ == '__main__':
    g = Graph()
    for i in range(6):
        g.addVertex(i)
    print (g.vertList)
    g.addEdge(0,1,5)
    g.addEdge(0,5,2)
    g.addEdge(1,2,4)

    for v in g:
        for w in v.getConnections():
            print ("( %s , %s )" % (v.getId(), w.getId()))
```
## 宽度优先搜索
### 词梯问题
* 游戏会给出一个起始词与终止词,玩家需要更改起始词中的一个字母，获得一个新词，计作一步。然后玩家需要更改所得的新词中的某个字母,再获得一个新词,最终获得终止。以下是从fool到sage的例子

```
fool → pool → poll → pole → pale -> sale -> sage 
```
* 使用图构建一个单词集合,如果两个单词的区别仅在于有一个不同的字母，就用一条边将它们相连。如果能创建这样一个图，那么其中的任意一条连接两个单词的路径就是词梯问题的一个解

<a href="https://sm.ms/image/4fGuRXlxYHyDprE" target="_blank"><img src="https://i.loli.net/2020/02/02/4fGuRXlxYHyDprE.png" ></a>

* 构建单词关系图:当处理列表中的每一个单词时,将它与桶上的标签进行比较。使用下划线作为通配符,我们将 POPE 和 POPS 放入同一个桶中。一旦将所有单词都放入对应的桶中之后，我们就知道，同一个桶中的单词一定是相连的。
<a href="https://sm.ms/image/WGCSkgM7iLpQ8wO" target="_blank"><img src="https://i.loli.net/2020/02/02/WGCSkgM7iLpQ8wO.png" ></a>

```python
from graph import Graph

def buildGraph(wordFile):
    """
    构建单词关系图
    Args:
        wordFile:单词文本(图顶点)
    """
    d = {}
    g = Graph()
    wfile  = open(wordFile,'r')
    #创建词桶
    for line in wfile:
        word = line[:-1]
        for i in range(len(word)):
            bucket = word[:i] + '_' + word[i+1:]
            if bucket in d:
                d[bucket].append(word)
            else:
                d[bucket] = [word]

    #为桶中的单词添加顶和边
    for bucket in d.keys():
        for word1 in d[bucket]:
            for word2 in d[bucket]:
                if word1 != word2:
                    g.addEdge(word1, word2)
```
### 实现宽度优先搜索(BFS)
