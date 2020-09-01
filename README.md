Solution-
1. Created a list of input intervals. intervalList
2. Sort the intervalList based on increasing starInterval time.
3. Create a new stack and push the first interval from intervalList to stack (intervalStack).
4. Iterate over the intervalList and do the following:
   * If the current interval does not overlap with the stack top, push it.
   * If the current interval overlaps with stack top and ending time of current interval is more than that of stack top, update stack top with the ending time of current interval.
5. We will have the stack contains the merged intervals.


Answers:

1. I liked the logic and time complexity. There could be one more implementation that falls more under the mathematical way to solve it.
We can talk about the mathematical approach over call ifyou like.

2. Might have skipped the code ie. throwing exceptions.

3. O(nlogn) + O(3n) 

4. I will change the implementation and will do it in time complexity of O(nlogn) + o(1).