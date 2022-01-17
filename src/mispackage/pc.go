package mispackage

import "fmt"

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPc Pc) Ping() {
	fmt.Println(myPc.Brand, "Pong")
}

func (myPc *Pc) DuplicateRam() {
	myPc.Ram = myPc.Ram * 2
}

func (myPC Pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.Ram, myPC.Disk, myPC.Brand)
}
