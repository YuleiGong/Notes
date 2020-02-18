#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-18 19:13:13
from __future__ import unicode_literals
from __future__ import absolute_import

from threading import Thread,Event
from queue import Queue
import time
import random

class Producer(Thread):
    def __init__(self,queue):
        super(Producer,self).__init__()
        self.queue = queue

    def run(self):
        for i in range(10):
            item = random.randint(0,256)
            self.queue.put(item)
            print ("Producer notify: item {} appended to queue by {}".format(item,self.name))
            time.sleep(1)

class Consumer(Thread):
    def __init__(self, queue):
        super(Consumer,self).__init__()
        self.queue = queue

    def run(self):
        while True:
            item = self.queue.get()
            print ("Consumer notify {} popped from queue by {}".format(item, self.name))
            self.queue.task_done()

if __name__ == '__main__':
    queue = Queue()
    t1 = Producer(queue)
    t2 = Consumer(queue)
    t3 = Consumer(queue)
    t4 = Consumer(queue)

    t1.start()
    t2.start()
    t3.start()
    t4.start()

    t1.join()
    t2.join()
    t3.join()
    t4.join()



