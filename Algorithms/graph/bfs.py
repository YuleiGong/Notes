#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-02 17:08:22

from queue import Queue
from vertex import Vertex
from buildGraph import buildGraph


def bfs(g, start):
    """
    宽度优先遍历
    Args:
        g:图数据
        start:起始顶点
    """
    start.setDistance(0)
    start.setPred(None)
    vertQueue = Queue()
    vertQueue.enqueue(start)
    while (vertQueue.size() > 0):
        currentVert = vertQueue.dequeue()
        for nbr in currentVert.getConnections():#相关联的顶点
            if (nbr.getColor() == 'white'): #白色代表还未访问过
                nbr.setColor('gray') #第一次访问,标记为灰色
                nbr.setDistance(currentVert.getDistance() + 1)
                nbr.setPred(currentVert)
                vertQueue.enqueue(nbr)
        currentVert.setColor('black') #完成对白色的顶点访问后,标记为黑色

def traverse(y):
    x = y
    while (x.getPred()):
        print (x.getId())
        x = x.getPred()
    print (x.getId())

if __name__ == '__main__':
    g = buildGraph('wordFile')
    start = g.getVertex('fool')
    bfs(g,start)
    traverse(g.getVertex('sage'))
