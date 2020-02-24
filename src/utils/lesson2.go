package utils

import (
	"math/rand"
)

type VehicleRef struct {
	Mark               string
	YearRelease        int
	TrunkVolume        int64 // litres
	CurrentFilledTrunk int   // 0-100
	IsEngineRunning    bool
	IsWindowsOpen      bool
	IsTruck            bool
}

type Fifo struct {
	nodes []*VehicleRef
	count int
	size  int
	head  int // point to head
	tail  int // point to tail
}

func NewFifo() *Fifo {
	return &Fifo{
		nodes: make([]*VehicleRef, 1),
		size:  1,
	}
}

func (f *Fifo) Push(ref *VehicleRef) {
	if f.head == f.tail && f.count > 0 { // we have one element
		nodes := make([]*VehicleRef, len(f.nodes)+f.size) // expand
		copy(nodes, f.nodes[f.head:])
		copy(nodes[len(f.nodes)-f.head:], f.nodes[:f.head])
		f.head = 0
		f.tail = len(f.nodes)
		f.nodes = nodes
	}
	f.nodes[f.tail] = ref
	f.tail = (f.tail + 1) % len(f.nodes)
	f.count++
}

func (f *Fifo) Pop() *VehicleRef {
	if f.count == 0 {
		return nil
	}
	node := f.nodes[f.head]
	f.head = (f.head + 1) % len(f.nodes)
	f.count--
	return node
}

func Initialize() []VehicleRef {
	vehicles := make([]VehicleRef, 10)
	var marks = []string{"Volkswagen", "Mercedes", "Porche", "Volvo", "Geely",
		"Renault", "LADA", "Ford", "Jeep", "Lexus"}

	rand.Seed(500)
	for l := 0; l < len(marks); l++ {
		vehicles[l].Mark = marks[l]
		vehicles[l].YearRelease = 2000 + rand.Intn(20)
		vehicles[l].TrunkVolume = int64(rand.Intn(500))
		vehicles[l].CurrentFilledTrunk = rand.Intn(100)
		if rand.Int63n(1000)%2 == 0 {
			vehicles[l].IsEngineRunning = true
			vehicles[l].IsWindowsOpen = true
			vehicles[l].IsTruck = true
		} else {
			vehicles[l].IsEngineRunning = false
			vehicles[l].IsWindowsOpen = false
			vehicles[l].IsTruck = false
		}
	}
	return vehicles
}
