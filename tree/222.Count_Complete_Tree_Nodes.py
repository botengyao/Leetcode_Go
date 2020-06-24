# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def countNodes(self, root: TreeNode) -> int:
        if not root:
            return 0
        temp = root
        depth = 0
        while root:
            depth += 1
            root = root.left
        n = 2 ** (depth - 1)
        return self.findlast(temp, n)
        
    def findlast(self, root, n):
        left, right = 0, n - 1
        while left <= right:
            mid = left + (right - left) // 2
            if self.exists(mid, n, root):
                left = mid + 1
            else: 
                right = mid - 1
        return n + right 
    
    def exists(self, idx, n, node):
        #0 1 2 3 4 x x x 
        #充分利用完全二叉树的性质，根据mid来决定node是向left还是right移动
        left, right = 0, n - 1
        while left < right:
            mid = left + (right - left) // 2
            if mid < idx:
                node = node.right 
                left = mid + 1
            else:
                node = node.left
                right = mid
        return node != None