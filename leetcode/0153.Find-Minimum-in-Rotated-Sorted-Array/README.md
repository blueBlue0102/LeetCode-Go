# 153. Find Minimum in Rotated Sorted Array

<https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/>

已知給定的 `nums` 中的數字都不會重覆，然後會經過隨機 `n` 次的翻轉  
題目要求以 $O(logn)$ 找出 `nums` 中的最小值

以 `[1 2 3 4 5 6 7]` 為例，來探討各種不同次翻轉後的情況下，比較 nums[left], nums[mid] and nums[right] 的值的關係  
翻轉 `7` 次，`[1 2 3 4 5 6 7]`，nums[left] < nums[mid] && nums[right] > nums[mid]，nums[left] 就是最小值  

翻轉 `1` 次，`[7 1 2 3 4 5 6]`，nums[left] > nums[mid] && nums[right] > nums[mid]，最小值在中位數左側  
翻轉 `2` 次，`[6 7 1 2 3 4 5]`，nums[left] > nums[mid] && nums[right] > nums[mid]，最小值在中位數左側  
翻轉 `3` 次，`[5 6 7 1 2 3 4]`，nums[left] > nums[mid] && nums[right] > nums[mid]，最小值在中位數左側  

翻轉 `4` 次，`[4 5 6 7 1 2 3]`，nums[left] < nums[mid] && nums[right] < nums[mid]，最小值在中位數右側  
翻轉 `5` 次，`[3 4 5 6 7 1 2]`，nums[left] < nums[mid] && nums[right] < nums[mid]，最小值在中位數右側  
翻轉 `6` 次，`[2 3 4 5 6 7 1]`，nums[left] < nums[mid] && nums[right] < nums[mid]，最小值在中位數右側  

觀察以上規律  
當【nums[left] < nums[mid] && nums[right] > nums[mid]】時，其 nums[left] 就會是最小值  
當【nums[left] > nums[mid] && nums[right] > nums[mid]】時，最小值會在中位數（包含）左側，所以設定 right = mid  
當【nums[left] < nums[mid] && nums[right] < nums[mid]】時，最小值會在中位數（不包含）右側，所以設定 left = mid + 1  

再考慮 Edge Case，思考 left, right 和 mid 的定義，是否會重疊  
`[1 2]`: left = 0, nums[left] = 1, mid = 0, nums[mid] = 1, right = 1, nums[right] = 2  
為了得以判定 left 是最小值，也就是想符合原先【nums[left] < nums[mid] && nums[right] > nums[mid]】的條件  
所以修改條件，left 和 mid 是有重疊的可能性，變成【nums[left] <= nums[mid] && nums[right] > nums[mid]】  
`[2 1]`: left = 0, nums[left] = 2, mid = 0, nums[mid] = 2, right = 1, nums[right] = 1  
想要將`[2 1]`拆成`[1]`，想要符合原先【nums[left] < nums[mid] && nums[right] < nums[mid]】，來觸發`left = mid + 1`  
所以觀察得出，left 和 mid 有重疊的可能，所以條件要改成【nums[left] <= nums[mid] && nums[right] < nums[mid]】

結論：  
當【nums[left] <= nums[mid] && nums[right] > nums[mid]】時，其 nums[left] 就會是最小值  
當【nums[left] >= nums[mid] && nums[right] > nums[mid]】時，最小值會在中位數（包含）左側，所以設定 right = mid  
當【nums[left] <= nums[mid] && nums[right] < nums[mid]】時，最小值會在中位數（不包含）右側，所以設定 left = mid + 1  

---

事後看解答發現，還可以再精簡判斷條件  
`[1 2 3 4 5 6 7]`  
`[7 1 2 3 4 5 6]`  
`[6 7 1 2 3 4 5]`  
`[5 6 7 1 2 3 4]`  
`[4 5 6 7 1 2 3]`  
`[3 4 5 6 7 1 2]`  
`[2 3 4 5 6 7 1]`  

可以觀察出其中，只要 nums[right] > nums[mid]，就可以推測出最小數在左側（含中位數），更新 right = mid  
只要 nums[right] < nums[mid] 就可以推測出最小數在右側（不含中位數），更新 left = mid + 1  
`[2 1]` 的情況下，希望觸發 left = mid + 1，所以 nums[right] < nums[mid] 改成 nums[right] <= nums[mid]  

如此迭代，最後可以得出 left 會是最小值

## Takeaway

無
