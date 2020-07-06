#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-04-14 21:44:30
from __future__ import unicode_literals
from __future__ import absolute_import

import sys
import re

WORD_RE = re.compile('\w+')

index = {}
with open(sys.argv[1], encoding='utf-8') as fp:
    for line_no, line in enumerate(fp, 1):
        for match in WORD_RE.finditer(line):
            word = match.group()
            column_no = match.start()+1
            location = (line_no, column_no)
            index.setdefault(word,[]).append(location)

# print in alphabetical order
for word in sorted(index, key=str.upper):  # <4>
    print(word, index[word])

