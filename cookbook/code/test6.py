#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-07-18 15:29:50
from __future__ import unicode_literals
from __future__ import absolute_import


import threading

class SharedCounter:

    def __init__(self, inital_value = 0):
        self._value = inital_value
        self._value_lock = threading.Lock()

    def incr(self, delta=1):
        """
        每次只有一个线程可以执行with语句中的代码块
        """
        with self._value_lock:
            self._value += delta

    def decr(self, delta=1):
        with self._value_lock:
            self._value -= delta





