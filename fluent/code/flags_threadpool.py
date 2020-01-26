#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-07-12 09:05:22
from __future__ import unicode_literals
from __future__ import absolute_import

from concurrent import futures
from flags import save_flag, get_flag,show,main

MAX_WORKERS = 20#设定最多的线程数

def download_one(cc):
    image = get_flag(cc)
    show(cc)
    save_flag(image, cc.lower() + '.gif')
    return cc


def download_many(cc_list):
    workers = min(MAX_WORKERS, len(cc_list))#设定工作线程数
    with futures.ThreadPoolExecutor(workers) as executor:
        res = executor.map(download_one, sorted(cc_list))#map方法会返回一个生成器,因此可以迭代获取各个函数的返回值

    return len(list(res))#获取结果数量,如果有异常,会在此抛出。


if __name__ == '__main__':
    main(download_many)
