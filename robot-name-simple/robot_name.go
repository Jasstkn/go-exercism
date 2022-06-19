package robotname

import (
	"fmt"
	"math/rand"
	"time"
)
// Define the Robot type here.

type Robot struct {
	name string
}

func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// randomly generate string
func randSeq(n int) string {
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

func generateNamePool(size int) []int {
	pool := make([]int, size)
	for i := 0; i < size; i++ {
		pool[i] = i
	}
	rand.Shuffle(size, func(i, j int) {
		pool[i], pool[j] = pool[j], pool[i]
	})
	return pool
}


func (r *Robot) Name() (string, error) {
	if len(r.name) > 0 {
		return r.name, nil
	}
	r.name = randSeq(2) + fmt.Sprintf("%.3d", rand.Intn(999))
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
