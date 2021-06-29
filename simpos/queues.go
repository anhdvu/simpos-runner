package simpos

import (
	"fmt"
	"net/http"
	"strings"
)

func RunQueue(q string) error {
	switch q {
	case reversal, adjustment:
		fmt.Printf("Running %v queue...\n", q)
		err := runTask(companionTaskUrl, q)
		if err != nil {
			return err
		}
	case both:
		fmt.Println("Running reversal and adjustment queue...")
		err := runTask(companionTaskUrl, reversal)
		if err != nil {
			return err
		}
		err = runTask(companionTaskUrl, adjustment)
		if err != nil {
			return err
		}
	default:
		return ErrNoQueueSpecified
	}
	fmt.Println("DONE!")
	return nil
}

func runTask(baseTaskUrl string, task string) error {
	u := baseTaskUrl + strings.Title(task)
	_, err := http.Get(u)
	if err != nil {
		return ErrQueueRun
	}
	return nil
}
