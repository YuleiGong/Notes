#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-01 14:04:18
from vertex import Vertex

class Graph:
    def __init__(self):
        self.vertices = {}
        self.numVertices = 0
        
    def addVertex(self,key):
        self.numVertices = self.numVertices + 1
        newVertex = Vertex(key)
        self.vertices[key] = newVertex
        return newVertex
    
    def getVertex(self,n):
        if n in self.vertices:
            return self.vertices[n]
        else:
            return None

    def __contains__(self,n):
        return n in self.vertices
    
    def addEdge(self,f,t,cost=0):
            if f not in self.vertices:
                nv = self.addVertex(f)
            if t not in self.vertices:
                nv = self.addVertex(t)
            self.vertices[f].addNeighbor(self.vertices[t],cost)
    
    def getVertices(self):
        return list(self.vertices.keys())
        
    def __iter__(self):
        return iter(self.vertices.values())


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
