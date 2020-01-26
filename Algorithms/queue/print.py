#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-15 20:31:44
import random

from queue import Queue

"""
打印任务队列
"""

class Printer:
    """
    打印机
    """
    def __init__(self, ppm):
        """
        初始化打印速度
        Args:
            ppm:每分钟打印多少页
        """
        self.pagerate = ppm
        self.currentTask = None
        self.timeRemaining = 0

    def tick(self):
        if self.currentTask != None:
            self.timeRemaining = self.timeRemaining - 1
            if self.timeRemaining <= 0:
               self.currentTask = None

    def busy(self):
        """
        任务繁忙检测
        Returns:
            打印机繁忙,返回True
            打印机空闲,返回False
        """
        if self.currentTask != None:
            return True
        else:
            return False

    def startNext(self, newtask):
        self.currentTask = newtask
        #打印花费的时间
        self.timeRemaining = newtask.getPages() * 60 /self.pagerate

class Task:
    def __init__(self, time):
        self.timestamp = time
        self.pages = random.randrange(1,21)

    def getStamp(self):
        return self.timestamp

    def getPages(self):
        return self.pages

    def waitTime(self,currenttime):
        """
        任务等待的时间
        Args:
            currenttime:当前时间
        """
        return currenttime - self.timestamp

def simulation(numSeconds,pagesPerMinute):
    """
    打印任务模拟
    Args:
        numSeconds:numSeconds 时间内模拟打印
        pagesPerMinute:打印速度 pages/min
    """
    labprinter = Printer(pagesPerMinute)
    printQueue = Queue()
    waitingtimes = []

    for currentSecond in range(numSeconds):
        if newPrintTask():
            task = Task(currentSecond)
            printQueue.enqueue(task)
        #打印机空闲且队列不为空
        if (not labprinter.busy()) and (not printQueue.isEmpty()):
            #下一个任务
            nexttask = printQueue.dequeue()
            waitingtimes.append(nexttask.waitTime(currentSecond))
            labprinter.startNext(nexttask)
        labprinter.tick()

    averageWait = sum(waitingtimes)/len(waitingtimes)
    print ("Average wait %6.2f secs %3d tasks remaining."%(averageWait,printQueue.size()))




def newPrintTask():
    """
    平均没180秒一个随机任务
    """
    num = random.randrange(1,181)
    if num == 180:
        return True
    else:
        return False

if __name__ == '__main__':
    for i in range(10):
        simulation(3600,5)
