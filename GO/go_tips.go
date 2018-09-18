/*

Go tips
=======

*/

// 1) Ternary operator:
// Go doesn't have ternary operator (as by Sep 2018).
// There is a hack to solve this in 1 line:

var ternary = map[bool]string{ true:"TRUE", false:"FALSE" }[true]

// as a result ternary would store TRUE value.
