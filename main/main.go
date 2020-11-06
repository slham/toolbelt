package main

import (
	"fmt"
	"github.com/slham/toolbelt/partition"
	"time"
)

const smallSize = 100
const mediumSize = 1000
const largeSize = 10000
const smallPartitions = 25
const mediumPartitions = 250
const largePartitions = 2500

func intArrOfN (num int) []int {
	out := make([]int, num)
	for i := 0; i < num; i++ {
		out[i] = i
	}
	return out
}

func elapsedTime(ip partition.IntPartition, arr []int) time.Duration {
	start := time.Now()
	ip.Do(arr)
	return time.Since(start)
}

func main() {
	small := intArrOfN(smallSize)
	medium := intArrOfN(mediumSize)
	large := intArrOfN(largeSize)

	p := partition.PercentPartition{
		Left:       0,
		Right:      0,
		Segment:    1,
	}
	s := partition.SimplePartition{Partitions: 4}

	p.Partitions, p.Length = smallPartitions, smallSize
	pSmallTime, sSmallTime := elapsedTime(&p, small), elapsedTime(&s, small)
	p.Partitions, p.Length = mediumPartitions, mediumSize
	pMediumTime, sMediumTime := elapsedTime(&p, medium), elapsedTime(&s, medium)
	p.Partitions, p.Length = largePartitions, largeSize
	pLargeTime, sLargeTime := elapsedTime(&p, large), elapsedTime(&s, large)

	fmt.Printf("percent small time: %v. simple small time: %v\n", pSmallTime, sSmallTime)
	fmt.Printf("percent medium time: %v. simple medium time: %v\n", pMediumTime, sMediumTime)
	fmt.Printf("percent large time: %v. simple large time: %v\n", pLargeTime, sLargeTime)
}
