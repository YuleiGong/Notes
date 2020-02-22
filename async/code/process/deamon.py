#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-21 11:27:25
from __future__ import unicode_literals
from __future__ import absolute_import

import multiprocessing
import time

def foo():
    name = multiprocessing.current_process().name
    print ("Starting {}".format(name))
    while True:
        pass
    print ("Exiting {}".format(name))



if __name__ == '__main__':
    background = multiprocessing.Process(name="background_process", target=foo)
    background.daemon = True
    no_background =  multiprocessing.Process(name="no_background_process", target=foo)
    no_background.daemon = False

    background.start()
    no_background.start()
