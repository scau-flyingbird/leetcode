##### 121. Best Time to Buy and Sell Stock

Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

**Example 1:**

```
Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.  
```

**Example 2:**

```
Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
```

##### 算法思路：

1. 遍历i从0到prices.length-1,中途记录0到i-1中的最小值min，最大利润max；
2. 如果prices[i]<min，则min=prices[i];
3. 如果prices[i]-min>max，则max=prices[i]-min;
4. 最后返回max，则为最大利润。

**复杂度：**O(n)

##### 代码：

```
package bestTimetoBuyandSellStock;

public class BestTimetoBuyandSellStock {

    /**
     * leetcode-cn.com #121. Best Time to Buy and Sell Stock
     * [link]: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
     * @param prices
     * @return
     */
    public int maxProfit(int[] prices) {
        int max = 0 , min = Integer.MAX_VALUE;
        for(int price:prices){
            if(price < min) {min = price;}
            int res = price - min;
            if(res > max) {max = res;}
        }
        return max;
    }
}

```
