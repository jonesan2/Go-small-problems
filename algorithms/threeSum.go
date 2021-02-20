// Brute force: three nested loops (O(N^3))
//
// Algorithm:
// --create a hash of the count of each value (N)
// --check for 3 zeros
// --check each other value for multiple occurrances
//   --if there are 2 or more, look for the value that would create zero sum
// --from the hash, create an array of only unique values, sorted by absolute value
// 	--grab all of the hash keys into a slice
//   --use Quicksort to sort the slice by absolute value
// --check for solutions involving three different numbers
//   --start with anchor on the largest absolute value
//   --move runner, only stopping on values of opposite sign
//   --search for the third number using binary search
//   --stop when the runner reaches a number equal to half of the anchor value
//   --move the anchor in by one
//
// Example
//
// Input: [-15, -9, 8, 7, -6, 4, 4, -2, -2, -2, -2, -2, 1, 1, 0]
// Hash: {
//				 -15: 1,
//				 -9: 1,
//				 8: 1,
//				 7: 1,
//				 -6: 1,
//				 4: 2,
//				 -2: 5,
//				 1: 2,
//				 0: 1
//			 }
//
// --there are not 3 zeros
//

package main

func main() {
}
