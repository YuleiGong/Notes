#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-27 21:35:05

def recMc(coinValueList, change):
    """
    找零金额-贪婪算法
    Args:
        coinValueList:零钱面值列表 [1,5,10,25]
        change:待找零钱总额
    Returns:
        minCoins:找零所需硬币数量
    """
    minCoins = change
    if change in coinValueList:
        return 1
    else:
        _coins = [c for c in coinValueList if c <= change]
        for i in _coins:
            numCoins = 1 + recMc(coinValueList,change-i)
            if numCoins < minCoins:
                minCoins = numCoins

    return minCoins

if __name__ == '__main__':
    print (recMc([1,5,10,21,25],63))

