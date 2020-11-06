package partition

import (
	"fmt"
	"github.com/meirf/gopart"
)

type IntPartition interface {
	Next()
	Before([]int)
	Process([]int)
	After([]int)
	Do([]int)
}

type PercentPartition struct {
	Left       int
	Right      int
	Segment    int
	Partitions int
	Length     int
}

func (p *PercentPartition) Next() {
	p.Right = p.Length * p.Segment / p.Partitions
}

func (p *PercentPartition) Before(them []int) {
}

func (p *PercentPartition) Process(them []int) {
	fmt.Println(them)
}

func (p *PercentPartition) After(them []int) {
}

func (p *PercentPartition) Do(them []int) {
	p.Next()
	for p.Right != p.Length {
		p.Before(them[p.Left:p.Right])
		p.Process(them[p.Left:p.Right])
		p.After(them[p.Left:p.Right])
		p.Left = p.Right
		p.Segment ++
		p.Next()
	}
	p.Before(them[p.Left:p.Right])
	p.Process(them[p.Left:p.Right])
	p.After(them[p.Left:p.Right])
}

type SimplePartition struct {
	Partitions int
}

func (p *SimplePartition) Next() {}

func (p *SimplePartition) Before(them []int) {
}
func (p *SimplePartition) Process(them []int)  {
	fmt.Println(them)
}
func (p *SimplePartition) After(them []int)  {
}

func (p *SimplePartition) Do(them []int) {
	for idxRange := range gopart.Partition(len(them), p.Partitions) {
		segment := them[idxRange.Low:idxRange.High]
		p.Before(segment)
		p.Process(segment)
		p.After(segment)
	}
}