package main

import "fmt"

type Computer interface {
	InsertDataThroughUSB()
}


////////////////////////////////////////////
type USB struct{
	data string
}

func (u *USB) InsertDataThroughUSB() {
	fmt.Println(u.data)
}

////////////////////////////////////////////
type MicroUSB struct {
	data string
}

func (m *MicroUSB) InsertDataThroughMicroUSB() {
	fmt.Println(m.data)
}
////////////////////////////////////////////

type ConnectorAdapter struct {
	microUSB *MicroUSB
}

func (c *ConnectorAdapter) InsertDataThroughUSB() {
	c.microUSB.InsertDataThroughMicroUSB()
}
////////////////////////////////////////////

func main() {
	USB:= USB{data: "Some data through USB"}
	microUSB:= MicroUSB{data: "Some data through microUSB"}
	addapter:= ConnectorAdapter{ microUSB: &microUSB}

	GetData(&USB)
	GetData(&addapter)

}

func GetData (com Computer) {
	com.InsertDataThroughUSB()
}

