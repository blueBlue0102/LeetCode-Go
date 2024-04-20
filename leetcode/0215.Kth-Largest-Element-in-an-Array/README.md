# 215. Kth Largest Element in an Array

<https://leetcode.com/problems/kth-largest-element-in-an-array/>

由於 Leetcode 上 test case 的更新  
若使用純粹的 quick select 會在最後一個 test case 陷入 worst cast，時間複雜度 $O(n^2)$，導致 TLE

一種解法是 DNF + quick select，先隨機選擇一個 pivot 值並透過 DNF 將 `nums` 切成三個區間  
若 k 落在 left 和 right 之間就直接回傳 pivot 值，反之根據 k 的值去左或右區間內繼續搜尋  
這樣做法的好處在於，若當所選的 pivot 大量重複時，且當欲尋找的目標不等於 pivot，在下次的搜尋中就可以略過這些重複的大量 pivot

## Takeaway

- Quick Select  
  <https://youtu.be/AqMiMkPOutQ?si=_3EjglpC8L4AhkrX>
