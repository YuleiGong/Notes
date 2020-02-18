#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-18 18:18:57
from __future__ import unicode_literals
from __future__ import absolute_import

import time
from threading import Thread, Event
import random

items = []
event = Event()


class Consumer(Thread):
    def __init__(self, items, event):
        super(Consumer,self).__init__()
        self.items = items
        self.event = event

    def run(self):
        while True:
            time.sleep(2)
            self.event.wait() #接收到通知set之后,取出
            item = self.items.pop()
            print ("Consumer notify : {} poped from list by {}".format(item, self.name))
class Producer(Thread):
    def __init__(self, items, event):
        super(Producer,self).__init__()
        self.items = items
        self.event = event

    def run(self):
        global item
        for i in range(100):
            time.sleep(2)
            item = random.randint(0,256)
            self.items.append(item)
            print ("Prodcue notify: item {} appended to list by {}".format(item, self.name))
            print ("Produce notify : event set by {}".format(self.name))
            self.event.set() #发出事件通知
            print ('Produce notify : event cleared by {}'.format(self.name))
            self.event.clear()

if __name__ == '__main__':
    t1 = Producer(items, event)
    t2 = Consumer(items, event)
    t1.start()
    t2.start()
    t1.join()
    t2.join()


