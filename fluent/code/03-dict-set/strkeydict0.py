#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-04-14 22:02:29
from __future__ import unicode_literals
from __future__ import absolute_import

class StrKeyDict0(dict):

    def __missing__(self, key):
        if isinstance(key, str):
            raise KeyError(key)
        return self[str(key)]

    def get(self, key, default=None):
        try:
            return self[key]
        except KeyError:
            return default

    def __contains__(self,key):
        return key in self.keys() or str(key) in self.keys()


if __name__ == '__main__':
    d = StrKeyDict0([('2', 'tow'),('4','four')])
    print (d['2'])
    print (d[4])
    print (d.get('2'))
    print (d.get(4))
    print (d.get(1,'N/A'))




