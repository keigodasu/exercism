package collatzconjecture

import (
	"fmt"
)

func CollatzConjecture(num int) (int, error) {
	givenNum := num
	count := 0
	for givenNum > 0 {
		if givenNum == 1 {
			return count, nil
		}
		if givenNum%2 == 0 {
			givenNum = givenNum / 2
			count++
		} else {
			givenNum = 3*givenNum + 1
			count++
		}
	}

	return 0, fmt.Errorf("%d coudnl't reach to 1", num)
}
