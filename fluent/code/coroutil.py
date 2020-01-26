#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-07-10 09:15:52
from __future__ import unicode_literals
from __future__ import absolute_import

from functools import wraps

def coroutine(func):
    """装饰器:向前执行到第一个yield表达式,预激func"""
    @wraps(func)
    def primer(*args,**kwargs):
        gen = func(*args,**kwargs)
        next(gen)
        return gen
    return primer
