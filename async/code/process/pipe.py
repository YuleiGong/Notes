#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-21 12:16:31
from __future__ import unicode_literals
from __future__ import absolute_import

import multiprocessing

def create_items(pipe):
    output_pipe,_ = pipe #recv send
    for item in range(10):
        output_pipe.send(item)
    output_pipe.close()

def multiply_items(pip_1,pip_2):
    close, input_pipe = pipe_1
    close.close()
    output_pipe,_ = pip_2
    try:
        while True:
            item = input_pipe.recv()
            output_pipe.send(item*item)
    except EOFError as e:
        output_pipe.close()


if __name__ == '__main__':
    #send
    pipe_1 = multiprocessing.Pipe(True)
    process_pipe_1 = multiprocessing.Process(target=create_items, args=(pipe_1,))
    process_pipe_1.start()
    #recv
    pipe_2 = multiprocessing.Pipe(True)
    process_pipe_2 = multiprocessing.Process(target=multiply_items, args=(pipe_1,pipe_2))
    process_pipe_2.start()
    pipe_1[0].close()
    pipe_2[0].close()

    try:
        while True:
            print (pipe_2[1].recv())
    except EOFError:
        print ("End")
