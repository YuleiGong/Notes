#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-02-17 21:41:36
from __future__ import unicode_literals
from __future__ import absolute_import

class Bus:
    def __init__(self, passengers=None):
        if passengers is None:
            self.passengers = []
        else:
            self.passengers = list(passengers)

    def pick(self,name):
        self.passengers.append(name)

    def drop(self,name):
        self.passengers.remove(name)

if __name__ == '__main__':
   pass 

