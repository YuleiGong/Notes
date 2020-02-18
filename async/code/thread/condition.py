#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-18 16:52:14
from __future__ import unicode_literals
from __future__ import absolute_import

from threading import Thread, Condition
import time

items = []
condition = Condition()

class consumer(Thread):

    def __init__(self):
        super(consumer,self).__init__()

    def consume(self):
        global condition
        global items
        condition.acquire()
        if len(items) == 0:
            condition.wait()
            print ("Consumer notify: no item to consume")
        items.pop()
        print ("Consumer notify : consumed 1 item")
        print ("Consumer notify : items to consumer are {}".format(len(items)))
        condition.notify()
        condition.release()

    def run(self):
        for i in range(0,20):
            time.sleep(2)
            self.consume()

class producer(Thread):
    def __init__(self):
        super(producer,self).__init__()

    def produce(self):
        global condition
        global items
        condition.acquire()
        if len(items) == 10:
            condition.wait()
            print ("Produce notify: items producted are {}".format(len(items)))
            print ("Produce notify: stop the production !!")
        items.append(1)
        print ("Produce notify : total items producted {}".format(len(items)))
        condition.notify()
        condition.release()
    def run(self):
        for i in range(0,20):
            time.sleep(1)
            self.produce()

if __name__ == '__main__':
    producer = producer()
    consumer = consumer()
    producer.start()
    consumer.start()
    producer.join()
    consumer.join()
