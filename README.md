
# Leetcode_go
Go solutions for Leetcode. (😁Updating) 


## Array

| Title | Solution | Difficulty | Time | Space | Like | Mark |
| ----- | :--------: | :----------: | :----: | :-----: | :---: | :----------: |
|[1. Two Sum](https://leetcode.com/problems/two-sum/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/array/1.Twosum.go)| Easy | O(n)| O(n)|||
|[11. Contaniner with most water](https://leetcode.com/problems/container-with-most-water/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/array/11.Container_With_Most_Water.go)| Medium| O(n)| O(1)||Two pointers (move small)|
|[15. 3 Sum](https://leetcode.com/problems/3sum)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/array/15.3sum.go)| Medium | O(n^2)| O(n)|||
|[16. 3 Sum Closest](https://leetcode.com/problems/3sum-closest)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/array/16.3Sum_Closest.go)| Medium | O(n^2)| O(1)|||
|[33. Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/array/33.Search_in_Rotated_Sorted_Array.go)| Medium | O(lg(n))| O(1)||Binary compare lo|
|[39. Combination Sum](https://leetcode.com/problems/combination-sum)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/array/39.Combination_sum.go)| Medium | O(2^n)| O(n)|||
|[40. Combination Sum II](https://leetcode.com/problems/combination-sum-ii)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/array/40.Combination_sum2.go)| Medium | O(2^n)| O(n)|||
|[53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/array/53.Maximum_Subarray.go)| Medium | O(n)| O(1)||DP or CC|
|[84. Largest Rectangle in Histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/array/84.Largest_Rectangle_in_Histogram.go)| Hard | O(n)| O(n)|🌹|Monotonous stack|
|[122. Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/array/122.Best_Time_to_Buy_and_Sell_II.go)| Easy | O(n)| O(1)|||
|[136. Single number](https://leetcode.com/problems/single-number/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/array/122.Best_Time_to_Buy_and_Sell_II.go)| Easy | O(n)| O(1)|🌹|Bit manipulation|
|[283. Move zeros](https://leetcode.com/problems/move-zeros/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/array/84.Largest_Rectangle_in_Histogram.go)| Easy | O(n)| O(1)|||
|[525. Contiguous Array](https://leetcode.com/problems/contiguous-array/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/array/525.Contiguous_Array.go)| Medium | O(n)| O(n)||diff = 0|

## Dynamic Programing

| Title | Solution | Difficulty | Time | Space | Like | Mark |
| ----- | :--------: | :----------: | :----: | :-----: | :-----: | :----------: |
|[64. Minimum Path Sum](https://leetcode.com/problems/minimum-path-sum/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/dp/64.Minimum_Path_Sum.go)| Medium | O(mn)| O(1)|| see 174 |
|[174. Dungeon Game](https://leetcode.com/problems/dungeon-game/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/dp/174.Dungeon_Game.go)| Hard | O(mn)| O(1)|🌹| bottom-right to top-left|
|[887. Super Egg Drop](https://leetcode.com/problems/super-egg-drop/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/dp/887.Super_Egg_Drop.go)| Hard | O(n^1/2)| O(km)|🌹||

## String

| Title | Solution | Difficulty | Time | Space | Like | Mark |
| ----- | :--------: | :----------: | :----: | :-----: | :-----: | :----------: |
|[49. Group Anagrams](https://leetcode.com/problems/group-anagrams/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/string/49._Group_Anagrams.go)| Medium | O(nk) | O(nk)|| sort/count|
|[678. Valid Parenthesis String](https://leetcode.com/problems/valid-parenthesis-string/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/string/678.Valid_Parenthesis_String.go)| Medium | O(n) | O(1)|🌹||
|[844. Backspace String Compare](https://leetcode.com/problems/backspace-string-compare/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/string/844.Backspace_String_Compare.go)| Medium | O(n) | O(1)|| pointer(end -> start) |

## Tree

| Title | Solution | Difficulty | Time | Space | Like | Mark |
| ----- | :--------: | :----------: | :----: | :-----: | :-----: | :----------: |
|[95. Unique Binary Search Trees II](https://leetcode.com/problems/unique-binary-search-trees-ii/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/tree/95.Unique_Binary_Search_TreesII.go)| Medium | O(n * Gn) | O(n*Gn)||(lo, hi), DFS, Memo|
|[96. Unique Binary Search Trees](https://leetcode.com/problems/unique-binary-search-trees/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/tree/96.Unique_Binary_Search_Trees.go)| Medium | O(n^2) | O(N)||Catalan number|
|[543. Diameter of Binary Tree](https://leetcode.com/problems/backspace-string-compare/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/tree/543.Diameter_of_Binary_Tree.go)| Easy | O(n) | O(n)|||
|[1008. Construct Binary Search Tree from Preorder Traversal](https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/tree/1008.Construct_Binary_Search_Tree_from_Preorde.go)| Medium | O(n) | O(n)|||



## Linked List

| Title | Solution | Difficulty | Time | Space | Like | Mark |
| ----- | :--------: | :----------: | :----: | :-----: | :-----: | :----------: |
|[202. Happy Number](https://leetcode.com/problems/happy-number/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/array/202.Happy_Number.go)| Easy | O(log(n)) | O(1)|🌹|Hare and tortoise|
|[206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)| [Go solution](https://github.com/calmbryan/Leetcode_Go/blob/master/linkedlist/206.Reverse_Linked_List.go)| Easy | O(n) | O(1)|||
