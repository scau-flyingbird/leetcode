##### 171. Excel Sheet Column Number

Given a column title as appear in an Excel sheet, return its corresponding column number.

For example:

```
    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28 
    ... 
```

**Example 1:**

```
Input: "A"
Output: 1
```

**Example 2:**

```
Input: "AB"
Output: 28
```

**Example 3:**

```
Input: "ZY"
Output: 701
```

##### 算法思路：

1. 26进制转10进制。

##### 代码：

```
package excelSheetColumnNumber;

public class ExcelSheetColumnNumber {
    public int titleToNumber(String s) {
        if(s == null || s.length() == 0){
            return 0;
        }
        int sum = 0;
        for(char c: s.toCharArray()){
            sum = sum * 26 + (c - 'A' + 1);
        }
        return sum;
    }
}

```

