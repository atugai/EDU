package hackerrank

import (
  "testing"
)

func TestKangaroo(t *testing.T) {
  tests := []struct{
    desc string
    x1, v1, x2, v2 int 
    want bool
  }{  
    {   
      desc: "kangaroo 1 meets 2", 
      x1: 0,
      v1: 3,
      x2: 4,
      v2: 2,
      want: true,
    }, {
      desc: "kangaroo 1 meets 2", 
      x1: 2,
      v1: 1,
      x2: 1,
      v2: 2,
      want: true,
    }, {
      desc: "kangaroo 1 doesn't meet 2", 
      x1: 0,
      v1: 2,
      x2: 5,
      v2: 3,
      want: false,
    }, {
      desc: "kangaroo 1 doesn't meet 2",
      x1: 21,
      v1: 6,
      x2: 47,
      v2: 3,
      want: false,
    }, {
      desc: "kangaroo 1 doesn't meet 2",
      x1: 43,
      v1: 2,
      x2: 70,
      v2: 2,
      want: false,
    },  
  }   

  for i, tc := range tests {
    if got := Kangaroo(tc.x1, tc.v1, tc.x2, tc.v2); got != tc.want {
      t.Errorf("TestKangaroo(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
     }   
   }   
}
