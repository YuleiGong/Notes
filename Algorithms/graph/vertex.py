#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-01 12:56:46

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
