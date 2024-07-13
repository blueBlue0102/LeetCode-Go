# 198. House Robber

<https://leetcode.com/problems/house-robber/>

在給定的 `nums` 數字陣列中，找出不相鄰節點最大的總和

假設 `dp(i)` 可回傳 `nums[0~i]` 的最大值總和  
當 `len(nums) == 1`: `dp(0) = nums[0]`  
當 `len(nums) == 2`: `dp(1) = max(nums[0], nums[1])`  
當 `len(nums) == 3`: `dp(2) = max(dp(1), nums[2]+dp(0))`  
當 `len(nums) == 4`: `dp(3) = max(dp(2), nums[3]+dp(1))`  

從 base case 推導出規則後，即可解題

## Takeaway

- Dynamic Programming
