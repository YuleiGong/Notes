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

    def __len__(self):
        return len(self.words)

    def __iter__(self):
        for match in RE_WORD.finditer(self.text):
            yield match.group()


if __name__ == '__main__':
    s = iter(Sentence('"The time has come," the Walrus said,'))
    for word in s:
        print (word)




