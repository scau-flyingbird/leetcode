/*
 * @lc app=leetcode.cn id=480 lang=java
 *
 * [480] 滑动窗口中位数
 */

// @lc code=start
class Solution {
    public double[] medianSlidingWindow(int[] nums, int k) {
        DualHeap dualHeap = new DualHeap(k);
        for(int i = 0; i < k; i++){
            dualHeap.insert(nums[i]);
        }

        double[] ret = new double[nums.length - k +1];
        ret[0] = dualHeap.getMiddeNum();
        for(int i = k, idx = 1; i < nums.length; i++, idx++){
            dualHeap.insert(nums[i]);
            dualHeap.ensure(nums[i-k]);
            ret[idx] = dualHeap.getMiddeNum();
        }

        return ret;
    }
}

class DualHeap {

    //大顶堆，存较小的
    private PriorityQueue<Integer> small;

    //小顶堆，存较大的
    private PriorityQueue<Integer> large;

    //哈希表，延迟删除
    private HashMap<Integer, Integer> delayed;

    //k个
    private int k;
    
    private int smallSize, largeSize;

    public DualHeap(int  k){
        this.small = new PriorityQueue<Integer>(new Comparator<Integer>() {
            public int compare(Integer num1, Integer num2) {
                return num2.compareTo(num1);
            }
        });
        this.large = new PriorityQueue<Integer>(new Comparator<Integer>() {
            public int compare(Integer num1, Integer num2) {
                return num1.compareTo(num2);
            }
        });
        this.delayed = new HashMap<>();
        this.k = k;
        this.smallSize = 0;
        this.largeSize = 0;
    }

    public double getMiddeNum(){
        if((k & 1) == 1){
            return small.peek() * 1.0;
        }else {
            return ((double)small.peek() + large.peek())/2;
        }
    }

    public void insert(int num){
        if(small.isEmpty() || num <= small.peek()){
            smallSize++;
            small.offer(num);
        }else {
            largeSize++;
            large.offer(num);
        }
        makeBalance();
    }

    public void ensure(int num){
        delayed.put(num, delayed.getOrDefault(num, 0) + 1);
        if(num <= small.peek()){
            smallSize--;
            if(num == small.peek()){
                prune(small);
            }
        }else {
            largeSize--;
            if(num == large.peek()){
                prune(large);
            }
        }
        makeBalance();
    }

    //不断的弹出堆顶元素并且更新哈希表
    public void prune(PriorityQueue<Integer> heap){
        while(!heap.isEmpty()){
            int num = heap.peek();
            if(delayed.containsKey(num)){
                delayed.put(num, delayed.get(num) - 1);
                if(delayed.get(num) == 0){
                    delayed.remove(num);
                }
                heap.poll();
            }else {
                break;
            }
        }
    }

    public void makeBalance(){
        if(smallSize > largeSize + 1){
            large.offer(small.poll());
            smallSize--;
            largeSize++;
            prune(small);
        }else if(largeSize > smallSize) {
            small.offer(large.poll());
            largeSize--;
            smallSize++;
            prune(large);
        }
    }
}
// @lc code=end

