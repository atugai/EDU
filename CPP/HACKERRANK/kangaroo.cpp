/*

[hackerrank.com] Kangaroo
https://www.hackerrank.com/challenges/kangaroo

You are given two kangaroos on a number line ready to jump in the positive direction (i.e, toward positive infinity).
* The first kangaroo starts at x1 location and moves at a rate of v1 meters per jump.
* The second kangaroo starts at x2 location and moves at a rate of v2 meters per jump.

Figure out a way to get both kangaroos at the same location at the same time.

*/

#include <cmath>

using namespace std;

// Checks wheter both kangaroos would be at same position after equal number of jumps.
bool kangaroo(int x1, int v1, int x2, int v2) {
  int delta = abs(float(x1 - x2));
  for (;;) {
    x1 += v1;
    x2 += v2;
    if (x1 == x2) {
      return true;
    }
    int newDelta = abs(float(x1 - x2));
    if (newDelta < delta) {
      delta = newDelta;
      continue;
     }
    return false;
  }
}
