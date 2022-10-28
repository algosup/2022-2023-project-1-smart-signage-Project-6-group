package main

import (
	"machine"
)

func InitDebug() machine.UART {
	debug := machine.UART1
	debug.SetBaudRate(9600)

	return *debug
}

func InitLoRa() machine.UART {
	lora := machine.UART2
	lora.Configure(machine.UARTConfig{
		BaudRate: 9600,
		RX:       machine.A3,
		TX:       machine.A2,
	})

	return *lora
}

func ATWrite(device machine.UART, cmd string) {
	device.Write([]byte("AT+" + cmd + "\r\n"))
}

func ReadBuffer(device machine.UART) string {
	var out string
	for device.Buffered() > 0 {
		b, _ := device.ReadByte()
		out += string(b)
	}

	return out
}
