'''
The set [1,2,3,...,n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:

"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.

Note:

Given n will be between 1 and 9 inclusive.
Given k will be between 1 and n! inclusive.
Example 1:

Input: n = 3, k = 3
Output: "213"
'''


import math
class Solution:
    def getPermutation(self, n: int, k: int) -> str:
        # 9 前面8个
        # 3! = 6
        # 计算前面有多少元素，从右向左看，计算A[i]右面比A[i]小的元素数量，乘以右面元素总数的阶乘
        # 2314 = 0 * 1! + 1 * 2! + 1 * 3! = 8 前面8个
        # 213 = 0 * 1! + 1 * 2! = 2 前面2个
        # 算第一个数是多少 k 
        candidate = [i for i in range(1, n + 1)] #[1,2,3,4,5,6]
        res = ""
        k = k - 1 #eg. k = 9 前面需要8个：k - 1
        for i in range(n, 0, -1):
            temp, remain = self.getfirst(k, i, candidate)
            res += str(temp)
            #print(k, candidate, res, " ", k // math.factorial(i - 1), " ",  math.factorial(i - 1))
            k = remain
        return res
        
    def getfirst(self, k, n, candidate):
        f = math.factorial(n - 1)
        number = k // f #计算前面需要有几个比选择的数小的
        k = k % f #更新新的k
        
        cnt = 0
        #在剩下的元素中选择有number个小的数，用过的被置为0会跳过
        for i in range(len(candidate)):
            if candidate[i] == 0:
                continue
            if cnt == number:
                temp = candidate[i]
                candidate[i] = 0
                break
            cnt += 1
        
        return temp, k
"""
5
40

39 [1, 0, 3, 4, 5] 2   1   24
15 [1, 0, 3, 0, 5] 24   2   6
3 [1, 0, 0, 0, 5] 243   1   2
1 [1, 0, 0, 0, 0] 2435   1   1
0 [0, 0, 0, 0, 0] 24351   0   1
"""