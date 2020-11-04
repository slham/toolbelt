package main

import (
	"fmt"
	"github.com/slham/toolbelt/partition"
	"time"
)

func main() {
	small := []int{0,1,2,3,4,5,6,7,8,9,10}
	medium := make([]int, 100)
	for i := 0; i < 100; i++ {
		medium[i] = i
	}
	large := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		large[i] = i
	}
	p := partition.PercentPartition{
		Left:       0,
		Right:      0,
		Segment:    1,
		Partitions: 3,
		Length:     len(small),
	}
	s := partition.SimplePartition{Partitions: 4}

	start := time.Now()
	p.Do(small)
	fmt.Printf("percent small took %v\n", time.Since(start))

	start = time.Now()
	s.Do(small)
	fmt.Printf("simple small took %v\n", time.Since(start))

	start = time.Now()
	p.Partitions = 25
	p.Length = 100
	p.Do(medium)
	fmt.Printf("percent medium took %v\n", time.Since(start))

	start = time.Now()
	s.Do(medium)
	fmt.Printf("simple medium took %v\n", time.Since(start))

	start = time.Now()
	p.Partitions = 250
	p.Length = 1000
	p.Do(large)
	fmt.Printf("percent large took %v\n", time.Since(start))

	start = time.Now()
	s.Do(large)
	fmt.Printf("percent large took %v\n", time.Since(start))
}
