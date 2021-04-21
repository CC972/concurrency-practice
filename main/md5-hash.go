package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func main() {
	words := []string{"the", "gecko", "lives", "in", "the", "woods"}
	complexity := "00000"

	superHash := superHash(words, complexity)
	fmt.Println(superHash)
}

func superHash(salts []string, complexity string) int {
	total := 0
	var wg sync.WaitGroup

	for _, salt := range salts {
		wg.Add(1)
		go accumulatingHash(&wg, &total, salt, complexity)
	}

	wg.Wait()

	return total
}

func accumulatingHash(wg *sync.WaitGroup, acc *int, salt string, complexity string) {
	*acc += hash(salt, complexity)
	wg.Done()
}

func hash(salt string, complexity string) int {
	i := 0

	for {
		hashed := md5.Sum([]byte(salt + strconv.Itoa(i)))
		hashed_str := hex.EncodeToString(hashed[:])

		if strings.HasPrefix(hashed_str, complexity) {
			return i
		}

		i++
	}
}
