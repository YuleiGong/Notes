#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-17 22:10:10
from __future__ import unicode_literals
from __future__ import absolute_import

from threading import Thread
from time import sleep

class CookBook(Thread):
    def __init__(self):
        Thread.__init__(self)
        self.message = "Hello Parallel Python CookBook!!"

    def print_message(self):
        print (self.message)

    def run(self):
        print ("Thread Starting")
        x = 0
        while (x < 10):
            print ("Thread Starting")
            sleep(2)
            x += 1
        print ("Thread Ended")

print ("Process Started")
hells_pyhton = CookBook()
hells_pyhton.start()
print ("Precess Ended") 

