/*
 
Fibonacci number implementation without recursion.
Expected complexity: O(n)
Expected memory:     O(1)
Input: position number in Fibonacci sequence, ex 1, 3 or 10th.
Output: Fib value at position 1,3 or 10 equals to 0, 1 or 34.

*/

#include <iostream>

int main() {
  int num; 
  int res; 
  int n1 = 0;
  int n2 = 1;
  
  std::cout << "Input Fib sequence number to get:" << std::endl;
  std::cin >> num;
  
  switch (num) {
    case 1: 
      res = n1;
      break;
    case 2: 
      res = n2;
      break;
    default:
      for (int i = 0; i < (num - 2); i++) {
        res = n1 + n2;
        n1 = n2;
        n2 = res;
      }
  }
  std::cout << "Fib member number " << num << " equals: " << res << std::endl;
  return 0;
}
