package simpos

import "testing"

func TestRunQueue(t *testing.T) {
	testcases := []struct {
		name string
		q    string
	}{
		{
			"reversal queue",
			"reversal",
		},
		{
			"adjustment queue",
			"adjustment",
		},
		{
			"both queues",
			"both",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := RunQueue(tc.q)
			if err != nil {
				t.Errorf("Expected error == nil but got error != nil, error == %v.", err)
			}
		})
	}

	t.Run("Queue error", func(t *testing.T) {
		dummy := "dummy"
		err := RunQueue(dummy)

		if err == nil {
			t.Errorf("Expected error != nil but got error == nil.")
		}

		if err != ErrNoQueueSpecified {
			t.Errorf("Expected error == %v but got error == nil.", ErrNoQueueSpecified)
		}
	})

	// TO-DO: Refactor RunTask in order to test ErrQueueRun
}

func TestRunTask(t *testing.T) {
	// Happy path has been covered in TestRunQueue
	t.Run("Happy path with adjustment", func(t *testing.T) {
		err := runTask(companionTaskUrl, adjustment)
		if err != nil {
			t.Errorf("Expected error == nil but got error != nil, error == %v.", err)
		}
	})

	t.Run("Queue run error", func(t *testing.T) {
		badTaskUrl := "https://xxx.yyy/badProcess"
		task := "task"
		err := runTask(badTaskUrl, task)
		if err == nil {
			t.Errorf("Expected error != nil but got error == nil.")
		}
		if err != ErrQueueRun {
			t.Errorf("Expected error == %v but got error == %v.", ErrQueueRun, err)
		}
	})
}
