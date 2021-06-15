package simpos

import (
	"fmt"
	"net/http"
)

func RunQueues(q string) {
	switch q {
	case reversal:
		fmt.Println("Running reversal queue...")
		_, err := http.Get(reversalQueue)
		if err != nil {
			fmt.Println("ERROR: Unable to run reversal queue", err)
		}
	case adjustment:
		fmt.Println("Running adjustment queue...")
		_, err := http.Get(adjustmentQueue)
		if err != nil {
			fmt.Println("ERROR: Unable to run adjustment queue", err)
		}
	case both:
		fmt.Println("Running reversal and adjustment queue...")
		_, err := http.Get(reversalQueue)
		if err != nil {
			fmt.Println("ERROR: Unable to run reversal queue", err)
		}
		_, err = http.Get(adjustmentQueue)
		if err != nil {
			fmt.Println("ERROR: Unable to run adjustment queue", err)
		}
	default:
		fmt.Println("No queue specified. Please check typo.")
	}
	fmt.Println("DONE!")
}
