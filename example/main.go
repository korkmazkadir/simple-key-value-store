package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"sync"
	"time"

	"../../simplekeyvalue"
)

type pair struct {
	key   string
	value []byte
}

func main() {

	pairs := createRandomPairs(100, 1000000)

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(index int, wg *sync.WaitGroup) {
			defer wg.Done()

			c, err := simplekeyvalue.NewClient()
			if err != nil {
				log.Println(err)
				return
			}

			putAllpairs(pairs, c)
			getAllPairs(pairs, c)

			fmt.Println(index)

		}(i, &wg)

	}

	wg.Wait()

}

func createRandomByteArray(length int) []byte {
	byteArray := make([]byte, length)
	rand.Read(byteArray)
	return byteArray
}

func createRandomPairs(numberOfPairs int, sizeOfValue int) []pair {

	var pairs []pair

	h := sha256.New()
	for i := 0; i < numberOfPairs; i++ {
		h.Reset()
		bytes := createRandomByteArray(sizeOfValue)
		pairs = append(pairs, pair{key: string(h.Sum(bytes)), value: bytes})
	}

	return pairs
}

func putAllpairs(pairs []pair, c *simplekeyvalue.Client) {

	for _, p := range pairs {
		start := time.Now()
		c.Put(p.key, p.value)
		log.Printf("%s\n", time.Since(start))
	}

}

func getAllPairs(pairs []pair, c *simplekeyvalue.Client) {

	h := sha256.New()
	for _, p := range pairs {
		h.Reset()
		start := time.Now()
		value := c.Get(p.key)
		elapsed := time.Since(start)
		result := p.key == string(h.Sum(value))
		log.Printf("%t\t%s\n", result, elapsed)
	}

}
