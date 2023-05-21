package main

import (
	"machine"
	"time"
)

const (
	CycleTime   = 100 // milliseconds
	NumEntries  = 1000
	REG_CHIP_ID = 0x00
	Address     = 0x29
	AddressAlt  = 0x28

	BNO055Id = 0xa0
	OPR_MODE = 0x3d

	OPERATION_MODE_CONFIG = 0x00
	OPERATION_MODE_NDOF   = 0x0C
)

func main() {
	time.Sleep(1 * time.Second)
	println("Configuring I2C0")
	err := machine.I2C0.Configure(machine.I2CConfig{
		SCL: machine.SCL_PIN,
		SDA: machine.SDA_PIN,
	})
	if err != nil {
		println("Could not configure I2C:", err)
		return
	}
	time.Sleep(100 * time.Millisecond)

	bus := machine.I2C0

	r := []byte{0}
	println("Wait for boot")
	timeout := 850
	for timeout = 850; timeout > 0; timeout -= 10 {
		if bus.Tx(uint16(AddressAlt), nil, r) == nil {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	if timeout <= 0 {
		println("Timeout waiting for BNO055 to boot")
	}
	println("Booted")

	// Check chip id
	bus.ReadRegister(AddressAlt, REG_CHIP_ID, r)
	if r[0] != BNO055Id {
		println("Failed initial chip id check")
		time.Sleep(time.Second) // wait further for boot
		bus.ReadRegister(AddressAlt, REG_CHIP_ID, r)
		if r[0] != BNO055Id {
			println("Timeout waiting for BNO055 to identify")
		}
	}
	time.Sleep(30 * time.Millisecond)

	println("Setting operation mode config")
	err = bus.WriteRegister(AddressAlt, OPR_MODE, []byte{OPERATION_MODE_CONFIG})
	if err != nil {
		// Consistent error here err => "I2C timeout during write"
		println("Failed to switch to config operation mode:", err)
	}
	time.Sleep(30 * time.Millisecond)
	println("Finished configuration")
}
