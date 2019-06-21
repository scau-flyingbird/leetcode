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
