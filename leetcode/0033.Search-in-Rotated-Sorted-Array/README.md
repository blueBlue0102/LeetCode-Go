# 33. Search in Rotated Sorted Array

<https://leetcode.com/problems/search-in-rotated-sorted-array/>

以 `[1 2 3 4 5 6 7]` 的各種 case 為例

`[1 2 3 4 5 6 7]`  
`[7 1 2 3 4 5 6]`  
`[6 7 1 2 3 4 5]`  
`[5 6 7 1 2 3 4]`  
if `nums[mid] < nums[right]`，表示 `mid` 的右邊一定都是大於 `nums[mid]` 的數字  
而 `mid` 左邊的數字則是混亂的不一定  
所以當 `target < nums[mid]` 時，可以確定要往左邊找  
而當 `target > nums[mid]` 時，可以先判斷 `target` 和 `nums[right]` 的關係  
若 `target > nums[right]` 則也要往左找，反之則往右找  

`[4 5 6 7 1 2 3]`  
`[3 4 5 6 7 1 2]`  
`[2 3 4 5 6 7 1]`  
if `nums[mid] > nums[right]`，表示 `mid` 的左邊都是小於 `nums[mid]` 的數字  
（因為轉折點在右邊所以才有機會 `nums[mid] > nums[right]`）  
而 `mid` 右邊的數字則是混亂的不一定  
所以當 `target > nums[mid]` 時，可以確定要往右找  
而當 `target < nums[mid]` 時，可以先判斷 `target` 和 `nums[left]` 的關係  
因為想找 `< nums[mid]` 的值，若 `nums[left] > target`，則表示左邊的值一定都大於 `target`，所以要往右找

## Takeaway

無
