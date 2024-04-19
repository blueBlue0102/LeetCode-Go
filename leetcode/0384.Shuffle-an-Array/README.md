# 384. Shuffle an Array

<https://leetcode.com/problems/shuffle-an-array/>

此題要求將給定的數字陣列進行隨機洗牌  
關於答案的正確與否，整合討論區裡的意見，似乎會判斷每次洗牌結果是否足夠分布  
也就是每種洗牌結果的出現機率都是相等的

須注意 `rand.IntN(n)` 的結果是 `[0,n)` 也就是 `>=0 和 <n`  
而 Fisher-Yates Algorithm 在針對每個 index 洗牌時，也需要包含不換位置的可能性，所以隨機值需要包含 `n`

## Takeaway

- Fisher-Yates Algorithm  
  <https://youtu.be/tLxBwSL3lPQ?si=mIPydZUYrs0BTBIe>  
  看完影片後就能快速理解 Fisher-Yates Algorithm 的作法