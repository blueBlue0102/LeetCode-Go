# 2225. Find Players With Zero or One Losses

<https://leetcode.com/problems/find-players-with-zero-or-one-losses/>

由於最後輸出的結果需要按照玩家編號由小到大，如果採用 hash map 來記錄各玩家的分數的話，可能就會發生需要進行 sort  
尤其在 Go 中的 map 是無序的

因此使用 array 的方式來記錄分數，index 即為玩家編號，這樣輸出時便不需要進行排序

## Takeaway

無
