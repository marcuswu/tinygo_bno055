//go:build (sam && atsamd51) || (sam && atsame5x)

package main

import (
	"device/sam"
	"errors"
	"fmt"
	"machine"
	"time"

	"github.com/marcuswu/tinygo_bno055/imu_data"
	"tinygo.org/x/drivers/bno055"
)

var (
	errI2CWriteTimeout       = errors.New("I2C timeout during write")
	errI2CReadTimeout        = errors.New("I2C timeout during read")
	errI2CBusReadyTimeout    = errors.New("I2C timeout on bus ready")
	errI2CSignalStartTimeout = errors.New("I2C timeout on signal start")
	errI2CSignalReadTimeout  = errors.New("I2C timeout on signal read")
	errI2CSignalStopTimeout  = errors.New("I2C timeout on signal stop")
	errI2CAckExpected        = errors.New("I2C error: expected ACK not NACK")
	errI2CBusError           = errors.New("I2C bus error")
)

const (
	// SERCOM_FREQ_REF is always reference frequency on SAMD51 regardless of CPU speed.
	SERCOM_FREQ_REF       = 48000000
	SERCOM_FREQ_REF_GCLK0 = 120000000

	// Default rise time in nanoseconds, based on 4.7K ohm pull up resistors
	riseTimeNanoseconds = 125

	// wire bus states
	wireUnknownState = 0
	wireIdleState    = 1
	wireOwnerState   = 2
	wireBusyState    = 3

	// wire commands
	wireCmdNoAction    = 0
	wireCmdRepeatStart = 1
	wireCmdRead        = 2
	wireCmdStop        = 3
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
	OPERATION_MODE_NDOF   = 0x0c
)

const i2cTimeout = 2500

type myI2C struct {
	machine.I2C
}

func (i2c *myI2C) sendAddress(address uint16, write bool) error {
	data := (address << 1)
	if !write {
		data |= 1 // set read flag
	}

	// wait until bus ready
	timeout := i2cTimeout
	// println("sendAddress using timeout", timeout)
	for !i2c.Bus.STATUS.HasBits(wireIdleState<<sam.SERCOM_I2CM_STATUS_BUSSTATE_Pos) &&
		!i2c.Bus.STATUS.HasBits(wireOwnerState<<sam.SERCOM_I2CM_STATUS_BUSSTATE_Pos) {
		timeout--
		if timeout == 0 {
			return errI2CBusReadyTimeout
		}
	}
	i2c.Bus.ADDR.Set(uint32(data))

	return nil
}

func (i2c *myI2C) signalStop() error {
	i2c.I2C.Bus.CTRLB.SetBits(wireCmdStop << sam.SERCOM_I2CM_CTRLB_CMD_Pos) // Stop command
	timeout := i2cTimeout
	for i2c.I2C.Bus.SYNCBUSY.HasBits(sam.SERCOM_I2CM_SYNCBUSY_SYSOP) {
		timeout--
		if timeout == 0 {
			return errI2CSignalStopTimeout
		}
	}
	return nil
}

func (i2c *myI2C) signalRead() error {
	i2c.I2C.Bus.CTRLB.SetBits(wireCmdRead << sam.SERCOM_I2CM_CTRLB_CMD_Pos) // Read command
	timeout := i2cTimeout
	for i2c.I2C.Bus.SYNCBUSY.HasBits(sam.SERCOM_I2CM_SYNCBUSY_SYSOP) {
		timeout--
		if timeout == 0 {
			return errI2CSignalReadTimeout
		}
	}
	return nil
}

func (i2c *myI2C) readByte() byte {
	for !i2c.I2C.Bus.INTFLAG.HasBits(sam.SERCOM_I2CM_INTFLAG_SB) {
	}
	return byte(i2c.I2C.Bus.DATA.Get())
}

func (i2c *myI2C) ReadRegister(address uint8, register uint8, data []byte) error {
	return i2c.Tx(uint16(address), []byte{register}, data)
}

// Tx does a single I2C transaction at the specified address.
// It clocks out the given address, writes the bytes in w, reads back len(r)
// bytes and stores them in r, and generates a stop condition on the bus.
func (i2c *myI2C) Tx(addr uint16, w, r []byte) error {
	var err error
	if len(w) != 0 {
		// send start/address for write
		i2c.sendAddress(addr, true)

		// wait until transmission complete
		timeout := i2cTimeout
		// println("Tx using timeout", timeout)
		for !i2c.Bus.INTFLAG.HasBits(sam.SERCOM_I2CM_INTFLAG_MB) {
			timeout--
			if timeout == 0 {
				return errI2CWriteTimeout
			}
		}

		// ACK received (0: ACK, 1: NACK)
		if i2c.I2C.Bus.STATUS.HasBits(sam.SERCOM_I2CM_STATUS_RXNACK) {
			return errI2CAckExpected
		}

		// write data
		for _, b := range w {
			err = i2c.WriteByte(b)
			if err != nil {
				return err
			}
		}

		err = i2c.signalStop()
		if err != nil {
			return err
		}
	}
	if len(r) != 0 {
		// send start/address for read
		i2c.sendAddress(addr, false)

		// wait transmission complete
		for !i2c.I2C.Bus.INTFLAG.HasBits(sam.SERCOM_I2CM_INTFLAG_SB) {
			// If the peripheral NACKS the address, the MB bit will be set.
			// In that case, send a stop condition and return error.
			if i2c.I2C.Bus.INTFLAG.HasBits(sam.SERCOM_I2CM_INTFLAG_MB) {
				i2c.I2C.Bus.CTRLB.SetBits(wireCmdStop << sam.SERCOM_I2CM_CTRLB_CMD_Pos) // Stop condition
				return errI2CAckExpected
			}
		}

		// ACK received (0: ACK, 1: NACK)
		if i2c.I2C.Bus.STATUS.HasBits(sam.SERCOM_I2CM_STATUS_RXNACK) {
			return errI2CAckExpected
		}

		// read first byte
		r[0] = i2c.readByte()
		for i := 1; i < len(r); i++ {
			// Send an ACK
			i2c.I2C.Bus.CTRLB.ClearBits(sam.SERCOM_I2CM_CTRLB_ACKACT)

			i2c.signalRead()

			// Read data and send the ACK
			r[i] = i2c.readByte()
		}

		// Send NACK to end transmission
		i2c.I2C.Bus.CTRLB.SetBits(sam.SERCOM_I2CM_CTRLB_ACKACT)

		err = i2c.signalStop()
		if err != nil {
			return err
		}
	}

	return nil
}

// WriteByte writes a single byte to the I2C bus.
func (i2c *myI2C) WriteByte(data byte) error {
	// Send data byte
	i2c.Bus.DATA.Set(data)

	// wait until transmission successful
	timeout := i2cTimeout
	for !i2c.Bus.INTFLAG.HasBits(sam.SERCOM_I2CM_INTFLAG_MB) {
		// check for bus error
		if i2c.Bus.STATUS.HasBits(sam.SERCOM_I2CM_STATUS_BUSERR) {
			return errI2CBusError
		}
		timeout--
		if timeout == 0 {
			return errI2CWriteTimeout
		}
	}

	if i2c.Bus.STATUS.HasBits(sam.SERCOM_I2CM_STATUS_RXNACK) {
		return errI2CAckExpected
	}

	return nil
}

func (i2c *myI2C) WriteRegister(address uint8, register uint8, data []byte) error {
	buf := make([]uint8, len(data)+1)
	buf[0] = register
	copy(buf[1:], data)
	return i2c.Tx(uint16(address), buf, nil)
}

func main() {
	time.Sleep(1 * time.Second)
	println("Configuring I2C0")
	err := machine.I2C0.Configure(machine.I2CConfig{
		// Frequency: 400 * machine.KHz,
		SCL: machine.SCL_PIN,
		SDA: machine.SDA_PIN,
	})
	if err != nil {
		println("Could not configure I2C:", err)
		return
	}
	time.Sleep(100 * time.Millisecond)

	bus := &myI2C{*machine.I2C0}

	/*r := []byte{0}
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
		println("Failed initial chip id check", r[0])
		time.Sleep(time.Second) // wait further for boot
		bus.ReadRegister(AddressAlt, REG_CHIP_ID, r)
		if r[0] != BNO055Id {
			println("Timeout waiting for BNO055 to identify")
		}
	}
	time.Sleep(100 * time.Millisecond)

	println("Setting operation mode config")
	err = bus.WriteRegister(AddressAlt, OPR_MODE, []byte{OPERATION_MODE_CONFIG})
	if err != nil {
		// Consistent error here err => "I2C timeout during write"
		println("Failed to switch to config operation mode:", err)
	}
	time.Sleep(30 * time.Millisecond)
	println("Finished configuration")*/
	println("Setting up IMU")
	imu := bno055.New(bus)
	imu.Address = bno055.AddressAlt

	println(fmt.Sprintf("Using BNO055 address %x", imu.Address))
	// time.Sleep(1000 * time.Millisecond)

	time.Sleep(100 * time.Millisecond)
	// for connected := false; !connected; connected = imu.Connected() {
	// 	print(".")
	// 	time.Sleep(30 * time.Millisecond)
	// }
	// println("")

	println("Configuring IMU")
	err = imu.Configure(bno055.DefaultConfig())
	if err != nil {
		println("Failed to configure IMU:", err)
		return
	}
	time.Sleep(100 * time.Millisecond)

	connected := imu.Connected()
	if !connected {
		return
	}

	fmt.Printf("Test writing to sd card...\n")
	fmt.Printf("I hope I see this\n")

	// Main loop
	for {
		data, err := readIMU(imu)
		if err != nil {
			fmt.Printf("Failed to read from IMU: %v\n", err)
			continue
		}
		fmt.Printf("got IMU data: %v\n", data)
		// TODO: time this -- it may take 10ms so we may need to reduce the CycleTime sleep below
		// f, err := fs.OpenFile("imuData.dat", os.O_CREATE|os.O_RDWR|os.O_APPEND)
		// if err != nil {
		// 	fmt.Printf("Failed to open file on sd card: %v\n", err)
		// }
		// err = writeToSD(writer, data, os.Stdout)
		// err = writeToSD(writer, data, f)
		// if err != nil {
		// 	fmt.Printf("Failed to write to file: %v\n", err)
		// }
		// f.Close()
		time.Sleep(CycleTime * time.Millisecond)
	}
}

func readIMU(imu bno055.Device) (imu_data.IMUEntry, error) {
	imuData := imu_data.IMUEntry{}
	if !imu.Connected() {
		return imuData, errors.New("IMU is not connected")
	}
	// Read from sensors
	XVal, YVal, ZVal, err := imu.ReadLinearAcceleration()
	if err != nil {
		return imuData, errors.New(fmt.Sprintf("Error reading linear acceleration: %v", err))
	}
	imuData.Entry.Acceleration.X = float64(XVal) / 1000000.0
	imuData.Entry.Acceleration.Y = float64(YVal) / 1000000.0
	imuData.Entry.Acceleration.Z = float64(ZVal) / 1000000.0

	XVal, YVal, ZVal, err = imu.ReadRotation()
	if err != nil {
		return imuData, errors.New(fmt.Sprintf("Error reading rotation: %v", err))
	}
	imuData.Entry.Rotation.X = float64(XVal) / 6000000.0
	imuData.Entry.Rotation.Y = float64(YVal) / 6000000.0
	imuData.Entry.Rotation.Z = float64(ZVal) / 6000000.0

	XVal, YVal, ZVal, err = imu.ReadMagneticField()
	if err != nil {
		return imuData, errors.New(fmt.Sprintf("Error reading magnetic field: %v", err))
	}
	imuData.Entry.Magnetometer.X = float64(XVal)
	imuData.Entry.Magnetometer.Y = float64(YVal)
	imuData.Entry.Magnetometer.Z = float64(ZVal)
	imuData.Entry.EntryTime = time.Now().UnixMilli()

	return imuData, nil
}
