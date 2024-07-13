# 417. Pacific Atlantic Water Flow

<https://leetcode.com/problems/pacific-atlantic-water-flow/>

給定的 `heights[r][c]` 中的值所代表的是高度  
下雨時，雨水可以流向較低或等高的位置  
要找出雨水可以同時流向太平洋和大西洋的位置點

太平洋在左上，可以流向太平洋的定義是 `row=0 || col=0`  
大西洋在左上，可以流向大西洋的定義是 `row=r || col=c`  
因此最左下角和最右上角這兩個點，由於同時鄰近兩個海域，所以必定可以同時流向

想像水流的形狀是一棵樹，每一個點都可以長出一棵樹  
太平洋樹和大西洋樹的交集即為答案

樹的找法，如果是從高往低處將會很困難，因為很難得知該節點的子節點是否能夠到達海洋  
所以反向思考，從海邊的那側開始，從低處往高處找  
若鄰居的高度相等或較高，就標記為可流通  
最終即可將整個 map 中太平洋或大西洋可流通的位置都標註出來

## Takeaway

- Graph Traversal