
##### 155. Min Stack
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

-   push(x) -- Push element x onto stack.
-   pop() -- Removes the element on top of the stack.
-   top() -- Get the top element.
-   getMin() -- Retrieve the minimum element in the stack.

**Example:**
```  
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
```  
##### 算法思路：  
使用两个栈stack(保存原栈数据),minStack(保存当原栈在入栈过程中动态的最小值数据，有新最小值入栈到原栈才插入minStack，minStack中栈顶元素即为当前栈最小值。加入当前栈出栈数值为minStack栈顶元素，则minStack也出栈)。
关键函数：
push(int x)：如果入栈x小于等于minStack栈顶元素，则x入栈stack同时也入栈minStack。否则只入栈stack不入栈minStack。
pop()：加入stack.pop()出栈数值等于minStack.peek()栈顶数值，则minStack栈也要同时出栈minStack.pop()。
  
##### 代码：  
  
```  
package minStack;

import java.util.Stack;

class MinStack {

    private Stack<Integer> stack;
    private Stack<Integer> minStack;

    public MinStack() {
        this.stack = new Stack<Integer>();
        this.minStack = new Stack<Integer>();
    }

    public void push(int x) {
        stack.push(x);
        if (minStack.isEmpty()) {
            minStack.push(x);
        } else {
            if (x <= minStack.peek()) {
                minStack.push(x);
            }
        }
    }

    public void pop() {
        int num = stack.pop();
        if (num == minStack.peek()) {
            minStack.pop();
        }
    }

    public int top() {
        return stack.peek();
    }

    public int getMin() {
        return minStack.peek();
    }
}
```
