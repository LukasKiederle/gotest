package philosophers


type Table struct {
	// channel for take requests - answer is sent over the reservedCh.
	takeCh chan int
	// channel to put fork requests.
	putCh chan int
	// direct communication between philosopher and table for getting fork response.
	reservedCh []chan bool
	forkInUse  []bool
	nbrOfSeats int
}

// NewLexer constructor
func NewTable(nbrOfSeats int) *Table {
	table := new(Table)
	table.putCh = make(chan int)
	table.takeCh = make(chan int)

	table.reservedCh = make([]chan bool, nbrOfSeats)
	for i := 0; i < nbrOfSeats; i++ {
		table.reservedCh[i] = make(chan bool)
	}
	table.forkInUse = make([]bool, nbrOfSeats)

	table.nbrOfSeats = nbrOfSeats
	return table
}

// Function run() contains the main loop for assigning forks and starting philosophers.
func (t *Table) run() {

	for {
		select {
		case requestedFork := <-t.takeCh:
			{
				if !t.forkInUse[requestedFork] && !t.forkInUse[(requestedFork+1)%t.nbrOfSeats] {
					// both forks are not in use -> reserve
					t.forkInUse[requestedFork] = true
					t.forkInUse[(requestedFork+1)%t.nbrOfSeats] = true
					t.reservedCh[requestedFork] <- true
				} else {
					t.reservedCh[requestedFork] <- false // not valid try again --> see loop in takeForks.
				}
			}
		case putFork := <-t.putCh:
			{
				// put forks
				t.forkInUse[putFork] = false
				t.forkInUse[(putFork+1)%t.nbrOfSeats] = false
			}
		}
	}
}


func (t *Table) getTakeChannel() chan int {
	return t.takeCh
}

func (t *Table) getPutChannel() chan int {
	return t.putCh
}

func (t *Table) getReservedChannel(id int) chan bool {
	return t.reservedCh[id]
}