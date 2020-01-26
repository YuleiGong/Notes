#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-02-18 21:31:23
from __future__ import unicode_literals
from __future__ import absolute_import

class HauntedBus:
    def __init__(self, passwngers=[]):
        self.passwngers = passwngers

    def pick(self, name):
        self.passwngers.append(name)

    def drop(self, name):
        self.passwngers.remove(name)

if __name__ == '__main__':
    pass
