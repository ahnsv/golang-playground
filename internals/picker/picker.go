package picker

import (
	"log"
	"math/rand"
	"time"
)

func ParseDateString(layout string, dateString string) time.Time {
	t, err := time.Parse(layout, dateString)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return t
}

func GenerateRandomNumber(maxNumber int, srcNum int64) int {
	src := rand.NewSource(srcNum)
	r := rand.New(src)
	num := r.Intn(maxNumber)
	return num
}

func PickOne(s []string, srcNum int64) string {
	return s[GenerateRandomNumber(len(s), srcNum)]
}
