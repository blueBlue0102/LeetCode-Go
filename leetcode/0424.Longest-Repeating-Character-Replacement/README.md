# 424. Longest Repeating Character Replacement

<https://leetcode.com/problems/longest-repeating-character-replacement/>

給定字串 `s` 及整數 `k`，`s` 只由大寫字母組成；`0 <= k <= s.length`  
最多允許對 `s` 中的字元進行 `k` 次的替換，找出最長連續相同字母的長度

由於是要找最長的 substring，而 substring 必為一個連續的字串  
所以可以使用 sliding window 的概念，window 的範圍是目前找到的最長距離  
隨著 right 每向右一步，就要根據是否還有 replace 的機會來決定是否要更新 left

## Takeaway

- Sliding Window
- Hash Map
