package main

import (
	"fmt"
	"machine"
	"time"
)

const (
	CycleTime   = 100 // milliseconds
	NumEntries  = 1000
	REG_CHIP_ID = 0x00
	Address     = 0x29
	AddressAlt  = 0x28
)

func main() {
	time.Sleep(1 * time.Second)
	println("Configuring I2C0")
	machine.I2C0.Configure(machine.I2CConfig{})
	err := machine.I2C0.Configure(machine.I2CConfig{
		SCL: machine.SCL_PIN,
		SDA: machine.SDA_PIN,
	})
	if err != nil {
		println("Could not configure I2C:", err)
		return
	}
	time.Sleep(100 * time.Millisecond)

	w := []byte{0x00}
	r := make([]byte, 1)
	err = machine.I2C0.Tx(AddressAlt, w, r)
	for err != nil {
		err = machine.I2C0.Tx(AddressAlt, w, r)
		if err != nil {
			println("could not interact with I2C device:", err)
		}
	}
	println("CHIP_ID", fmt.Sprintf("0x%x", r[0]))
}
