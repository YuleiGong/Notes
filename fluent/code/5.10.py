#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-02-17 15:22:41
from __future__ import unicode_literals
from __future__ import absolute_import


def tag(name, *content, cls=None, **attrs):
    if cls is not None:
        attrs['class'] = cls
    if attrs:
        attr_str = ''.join(' %s="%s"' % (attr, value)
                           for attr, value in sorted(attrs.items()))
    else:
        attr_str = ''

    if content:
        return '\n'.join('<%s%s>%s</%s>' % (name, attr_str, c, name) for c in content)
    else:
        return '<%s%s />' % (name, attr_str)

if __name__ == '__main__':
    #print (tag('br'))
    #print (tag('p', 'hello'))
    print (tag('xx', 'xxx', cls=1, content="testing",s1='2'))


