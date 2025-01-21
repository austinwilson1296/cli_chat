package auth

import (
	"crypto/rand"
	"fmt"
)

func GenerateRandomId(numBits int) int {
	num := make([]byte,numBits)

	randNum,err := rand.Read(num)
	if err != nil {
		fmt.Println("error:",err)
		return 0
	}
	return int(randNum)
}