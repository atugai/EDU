/*
 
Fibonacci number implementation without recursion.
Expected complexity: O(n)
Expected memory:     O(1)
Input: number N of first members Fib sequence. 
Output: First N members of Fib.

*/

#include <iostream>

// Prints first N number of Fibonacci sequence.
void fib(int n) {
  unsigned long int res;
  unsigned long int n1 = 0;
  unsigned long int n2 = 1;

  std::cout << "First " << n << " numbers of Fib: " << std::endl;

  switch (n) {
    case 1:
      res = n1;
      std::cout << res << std::endl;
      break;
    case 2:
      res = n2;
      std::cout << res << std::endl;
      break;
    default:
      std::cout << n1 << std::endl;
      std::cout << n2 << std::endl;
      for (int i = 0; i < (n - 2); i++) {
        res = n1 + n2;
        n1 = n2;
        n2 = res;
        std::cout << res << std::endl;
      }
  }
}


int main() {
  int num; 
  
  std::cout << "Number of first Fib members:" << std::endl;
  std::cin >> num;
  
  fib(num);
  return 0;
}
