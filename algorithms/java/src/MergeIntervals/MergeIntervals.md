56. Merge Intervals


```
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
```

##### 算法思路：
1. 把int[][] intervals  区间数组按左区间数值大小递增排序成List<Point> list区间列表；
2. List<Point> list区间列表遍历合并到新List<Point> resList，合并规则：
- Point ePoint = resList.get(resList.size()-1); Point p = list.get(i);
- 如果p.start左区间在ePoint<start,end>区间中，则把ePoint.end替换成p.end和ePoint.end中的最大值；
- 如果p.start左区间在ePoint<start,end>区间外，则为新区间，加到结果区间中resList.add(p);


代码如下：

```
package MergeIntervals;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class MergeIntervals {

    public class Point implements Comparable<Point> {
        public int start;
        public int end;

        public Point(int start, int end) {
            this.start = start;
            this.end = end;
        }

        @Override
        public int compareTo(Point o) {
            return this.start - o.start;
        }
    }

    public int[][] listToArray(List<Point> resList) {
        int res[][] = new int[resList.size()][2];
        for (int i = 0; i < resList.size(); i++) {
            res[i][0] = resList.get(i).start;
            res[i][1] = resList.get(i).end;
        }
        return res;
    }

    public List<Point> arrayToList(int[][] intervals) {
        List<Point> list = new ArrayList<Point>(intervals.length);
        for (int i = 0; i < intervals.length; i++) {
            list.add(new Point(intervals[i][0], intervals[i][1]));
        }
        return list;
    }

    public int[][] merge(int[][] intervals) {
        if (intervals.length <= 0) {
            return new int[0][0];
        }
        List<Point> list = arrayToList(intervals);
        Collections.sort(list);
        List<Point> resList = new ArrayList<Point>();
        for (int i = 0; i < list.size(); i++) {
            Point p = list.get(i);
            if (resList.size() <= 0) {
                resList.add(p);
            } else {
                Point ePoint = resList.get(resList.size() - 1);
                if (p.start <= ePoint.end) {
                    ePoint.end = Math.max(ePoint.end, p.end);
                } else {
                    resList.add(p);
                }
            }
        }
        return listToArray(resList);
    }
}

```
