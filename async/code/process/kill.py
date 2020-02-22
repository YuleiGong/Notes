#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-21 11:35:41
from __future__ import unicode_literals
from __future__ import absolute_import

import multiprocessing
import time

def foo():
    print ("Starting function")
    time.sleep(0.1)
    print ("Finished function")


if __name__ == '__main__':
   p = multiprocessing.Process(target=foo)
   print ("Process berore excution: {}".format(p.is_alive()))
   p.start()
   print ("Process runing: {}".format(p.is_alive()))
   p.terminate()
   time.sleep(0.1)
   print ("Process terminate: {}".format(p.is_alive()))
   p.join()
   print ("Process joined: {}".format(p.is_alive()))
   print ("Process exit code: {}".format(p.exitcode))


