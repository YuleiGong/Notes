#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-03-27 22:26:05
from __future__ import unicode_literals
from __future__ import absolute_import

import re
import reprlib

RE_WORD = re.compile('\w+')


class Sentence:

    def __init__(self, text):
        self.text = text
        self.words = RE_WORD.findall(text)

    def __len__(self):
        return len(self.words)


    def __iter__(self):
        for word in self.words:
            yield word
        return 

    def __repr__(self):
        return 'Sentence(%s)' % reprlib.repr(self.text)

class SentenceIterator:
    def __init__(self, words):
        self.words = words
        self.index = 0

    def __next__(self):
        try:
            word = self.words(self.index)
        except IndexError:
            raise StopIteration()
        self.index += 1
        return word

    def __iter__(self):
        return self




if __name__ == '__main__':
    s = iter(Sentence('"The time has come," the Walrus said,'))
    for word in s:
        print (word)




