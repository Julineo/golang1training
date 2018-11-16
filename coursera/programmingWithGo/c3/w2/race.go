/*
Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.

Here, we can see that the getNumber function is setting the value of i in a separate goroutines. We are also returning i from the function without any knowledge of which of our goroutines has completed. So now, there are three operations that are taking place:

The value of i is being set to 5
The value of i is being set to 4
The value of i is being returned from the function
Now depending on which of these three operations completes first, our value printed will be either 0 (the default integer value) or 5 or 4.

This is why it’s called a data race : the value returned from getNumber changes depending on which of the operations 1 or 2 finish first.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(getNumber())
}

func getNumber() int {
	var i int
	go func() {
		i = 5
	}()
	go func() {
		i = 4
	}()

	return i
}
