#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-02 18:23:27
from graph import Graph

def knightGraph(bdSize):
    ktGraph = Graph()
    for row in range(bdSize):
        for col in range(bdSize):
            nodeId = posToNodeId(row, col, dbSize)

