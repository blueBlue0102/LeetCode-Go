# 148. Sort List

<https://leetcode.com/problems/sort-list/>

以下紀錄自己的思考過程

## 一、暴力解

暴力解 Insertion Sort($O(n^2)$)  
雖然邏輯正確，但是 TLE

## 二、放到 array 排序完再更新 list

把 List 的資料放入 array 後進行排序，再一次更新整個 List

思考由於題目並不要求要進行 node 的 swap，而是可以單純將 val 進行 swap 即可  
而且由於 Linked List 的特性，所以必定至少需要進行一次 $O(n)$ 來得知所有的值  
所以推測第一次巡視的同時又要同時進行更新是不可能的  
因此需要進行兩次整個 Linked List 的巡視  
~~一次會將所有的值進行記錄放入 array 並同時排序，另一次則拿排序後的值進行 Linked List 的更新  
這個做法需要考慮的是如何有效率的持續維護 array 的排序~~  

一次巡完最後再排序即可  
因為若使用 Binary Search Tree 來維護 array，則每次 insert 的時間複雜度是 $O(log n)$  
放完整個 List 需要 $O(n \cdot log n)$，和 quick sort 或 merge sort 的時間複雜度一樣  
所以一次巡完後再排序即可

這樣的時間複雜度是 $O(n)+O(n \cdot log n)+O(n)=O(n \cdot log n)$  
空間複雜度是 $O(n)$  

## 三、Merge Sort

看解答後才知道可以用 Merge Sort  
又能再分成 Top Down 和 Bottom Up 的寫法  
Top Down 的寫法採用遞迴，時間複雜度是 $O(n \cdot log n)$，空間複雜度是 $O(n)$

> 「空間複雜度是 $O(n)$ 而不是 $O(n \cdot log n)$」的解釋可參考  
> <https://stackoverflow.com/questions/10342890/merge-sort-time-and-space-complexity/28641693#28641693>

Bottom Up 做法之後有空再研究

## Takeaway

- Sort
- Linked List
