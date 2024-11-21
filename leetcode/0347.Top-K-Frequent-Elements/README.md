# 347. Top K Frequent Elements

<https://leetcode.com/problems/top-k-frequent-elements/>

給定整數陣列 `nums` 和整數 `k`  
回傳 `nums` 中前 `k` 個高頻率出現的數字  
回傳的數字可以用任何順序出現  
答案已保證是獨一無二的，且要求時間複雜度必須優於 $O(n \cdot log(n))$

假設 `nums` 長度是 `n`  
則出現頻率最高的次數也就是 `n`，意思是「頻率」數值的範圍有限  
因此，可以建立一個長度為 `n` 的 bucket，將相應出現頻率的數字放到指定的 index  

一開始先用 $O(n)$ 掃 `nums` 將所有數字及其頻率記錄在 hash table  
接著掃一次這個 hash table，將出現頻率記錄在相應 bucket  
最後從 bucket 的最後取出前 `k` 個 element

這樣子的時間和空間複雜度都只需要 $O(n)$

## Takeaway

- Hash Table
- Counting Sort
