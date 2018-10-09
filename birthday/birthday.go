package birthday

import (
	"math/rand"
	"sort"
	"time"
)

var seed = time.Now().UnixNano()
var rng = rand.New(rand.NewSource(seed))

func NewBirthday() int {
	return rng.Intn(364) + 1
}

func NewClass(n int) []int {
	c := make([]int, n)
	for i := 0; i < n; i++ {
		c[i] = NewBirthday()
	}
	return c
}

func CheckSameBirthday(class []int) bool {
	for i := 0; i < len(class); i++ {
		for j := 0; j < len(class); j++ {
			if i == j {
				continue
			}
			if class[i] == class[j] {
				return true
			}
		}
	}
	return false
}

func CheckSameBirthday2(class []int) bool {
	sort.Ints(class)
	for i := 0; i < len(class)-1; i++ {
		if class[i] == class[i+1] {
			return true
		}
	}
	return false
}

func CheckSameBirthday3(class []int) bool {
	mem := make(map[int]int)
	for i := 0; i < len(class)-1; i++ {
		_, ok := mem[class[i]]
		if ok {
			return true
		}
		mem[class[i]]++
	}
	return false
}
