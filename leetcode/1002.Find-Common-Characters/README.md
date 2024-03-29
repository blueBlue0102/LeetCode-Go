# 1002. Find Common Characters

<https://leetcode.com/problems/find-common-characters/>

找出在所有給定字串中都出現過的單字，單字可以重覆  
回傳的陣列不要求順序

每巡過一個字串就建立一次 hash table，排列出各個字元的出現頻率  
取出在所有 hash table 中出現過的單字的最小出現頻率即可

時間複雜度為 $O(n)$，$n$ 是所有字串的字元長度總和  
空間複雜度是 $O(1)$，雖然單個 hash table 的大小會受到字元種類的數量影響，但如果從所有字元種類的數量是已知有限的角度下，hash table 大小的「增長」最終並不會受到字串影響，所以是 $O(1)$

## Takeaway

無
