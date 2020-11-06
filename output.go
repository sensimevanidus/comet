package comet

import (
	"fmt"
	"sync"
)

var outputMx sync.Mutex

func print(message string, verbose, show bool) {
	if verbose && !show {
		return
	}

	outputMx.Lock()
	defer outputMx.Unlock()

	fmt.Print(message)
}

func println(message string, verbose, show bool) {
	if verbose && !show {
		return
	}

	outputMx.Lock()
	defer outputMx.Unlock()

	fmt.Println(message)
}

func fail(i int, step *testStep, err error) {
	outputMx.Lock()
	defer outputMx.Unlock()

	if step != nil {
		fmt.Printf("--- FAIL: %s (0.00s) [#%d]\n", step.String(), i+1)
	} else {
		fmt.Printf("--- FAIL: step (0.00s) [#%d]\n", i+1)
	}

	if err != nil {
		fmt.Printf("    %s\n", err.Error())
	}
}

func success(i int, step testStep, verbose bool) {
	outputMx.Lock()
	defer outputMx.Unlock()

	if verbose {
		fmt.Printf("--- PASS: %s (0.00s) [#%d]\n", step.String(), i+1)
	}
}
