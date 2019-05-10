package mergeIntervals;

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
