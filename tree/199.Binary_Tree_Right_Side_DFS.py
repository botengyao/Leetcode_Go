# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

"""
dfs exploring/ traversing the graph going left and as deep as possible, then right when the (whole) left subtree has been visited
at the moment the dfs is "being called", the current node is the right most node at the its depth
"""
class Solution:
    def rightSideView(self, root: TreeNode) -> List[int]:
        h = dict()
        self.travel(root, 0, h)
        return [k[1] for k in sorted(h.items(), key=lambda item:item[0])]
        
    def travel(self, root, level, h):
        if not root:
            return 
        h[level] = root.val
        self.travel(root.left, level + 1, h)
        self.travel(root.right, level + 1, h)