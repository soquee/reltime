package reltime

import (
	"fmt"
	"time"
)

func ExampleAgo() {
	s := Ago(-24 * time.Hour)
	fmt.Println(s)

	// Output:
	// about a day ago
}
