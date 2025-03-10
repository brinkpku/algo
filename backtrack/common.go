// Package backtrack related
package backtrack

/*
本质上是对多叉树的遍历，常见的问题有子集（组合）、排列等。
核心框架：

result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return

    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择
*/
