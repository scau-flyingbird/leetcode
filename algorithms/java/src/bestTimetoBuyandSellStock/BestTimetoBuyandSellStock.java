package bestTimetoBuyandSellStock;

public class BestTimetoBuyandSellStock {

    /**
     * leetcode-cn.com #121. Best Time to Buy and Sell Stock
     * [link]: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
     *
     * 算法思路：
     * 1.遍历i从0到prices.length-1,中途记录0到i-1中的最小值min，最大利润max；
     * 2.如果prices[i]<min，则min=prices[i];
     * 3.如果prices[i]-min>max，则max=prices[i]-min;
     * 4.最后返回max，则为最大利润。
     *
     * 复杂度：O(n)
     *
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
