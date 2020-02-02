#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-02 13:54:18

from graph import Graph

def buildGraph(wordFile):
    """
    构建单词关系图
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
    return g

if __name__ == '__main__':
    buildGraph('wordFile')

