package philosophers

import (
	"fmt"
	"time"
)

type Philosopher struct {
	id int
	table *Table
}

// Run loops forever.
func (p *Philosopher) run() {
	for {
		p.takeForks()
		p.eat()
		p.putForks()
		p.think()
	}
}

func NewPhilosopher(id int, table *Table) *Philosopher {
	philosopher := new(Philosopher)
	philosopher.id = id
	philosopher.table = table
	return philosopher
}

// Take forks by channeling our id to the table and wait until the table returns true on the reserved channel.
func (p *Philosopher) takeForks() {
	// try to get forks from table
	gotForks := false
	for !gotForks {
		p.table.getTakeChannel() <- p.id
		gotForks = <-p.table.getReservedChannel(p.id)
	}
}

func (p *Philosopher) eat() {
	fmt.Printf("Phil %d is eating", p.id)
	time.Sleep(1 * time.Second)
}

func (p *Philosopher) putForks() {
	p.table.getPutChannel() <- p.id
}

func (p *Philosopher) think() {
	fmt.Printf("Phil %d is thinking", p.id)
	time.Sleep(1 * time.Second)
}