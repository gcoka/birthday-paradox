package main

import (
	"fmt"

	"github.com/gcoka/birthday-paradox/birthday"
)

func main() {
	for n := 2; n <= 50; n++ {

		count := 0

		for i := 0; i < 10000; i++ {

			c := birthday.NewClass(n)

			checkResult := birthday.CheckSameBirthday(c)
			if checkResult {
				count++
			}
		}

		fmt.Println(n, "人クラス ", float64(count)/10000.0*100.0)
	}
}
