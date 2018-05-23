/* 5.15 Write variadic functions max and min, analogous to sum. What should these functions do when called with no arguments? Write variants that require at least one argument. */
package main

func max(vals ...int) int {
	if len(vals) == 0 {
		panic("no parameters, at least one parameter required for max() function")
	}
	max := vals[0]
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}
