```
https://leetcode.cn/problems/distinct-subsequences/
```

```
115. 不同的子序列
困难

给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数。

测试用例保证结果在 32 位有符号整数范围内。

示例 1：

输入：s = "rabbbit", t = "rabbit"
输出：3
解释：
如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
**rabb**b**it**
**ra**b**bbit**
**rab**b**bit**

示例 2：

输入：s = "babgbag", t = "bag"
输出：5
解释：
如下所示, 有 5 种可以从 s 中得到 "bag" 的方案。 
**ba**b**g**bag
**ba**bgba**g**
**b**abgb**ag**
ba**b**gb**ag**
babg**bag**

提示：

    1 <= s.length, t.length <= 1000
    s 和 t 由英文字母组成
```
