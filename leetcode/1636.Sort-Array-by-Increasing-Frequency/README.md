# 1636. Sort Array by Increasing Frequency

<https://leetcode.com/problems/sort-array-by-increasing-frequency/>

影響排序的因素有"數字出現的頻率"以及"數字大小"  
為了得知數字出現的頻率，因此勢必要掃過整個 array 才能做到  
所以建立一個 hash table 來記錄各數字的出現次數（$O(n)$）  
之後對 `nums` 以 hash table 的值來進行 comparsion sort（$O(n \cdot log(n))$）

## Takeaway

- Sort
