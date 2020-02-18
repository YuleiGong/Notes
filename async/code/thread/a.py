#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-18 13:48:59
from __future__ import unicode_literals
from __future__ import absolute_import

import threading
import time

def function(i):
    print ("function called by thread {}".format(i))
    time.sleep(i)
    return

threads = []

for i in range(5):
    t = threading.Thread(target=function,args=(i,))
    threads.append(t)
    t.start()
[t.join() for t in threads]

