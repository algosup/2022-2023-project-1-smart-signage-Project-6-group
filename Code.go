package main

import (
	"machine"
	"time"
)

func main() {
	var period uint64 = 1e9 / 1000
	led := machine.LED
	light := machine.PWA0
	tension := machine.PC15
	// antenna := machine.UART0
	// antenna.Configure(machine.UARTConfig{TX: machine.Pin(2), RX: machine.Pin(3)})
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	light.Configure(machine.PinConfig{Mode: machine.PinOutput})
	tension.Configure()

	// var test float64
	// test = tension.get
	light.Configure(machine.PWMConfig{
		Period: period,
	})

	ch, err := light.Channel(light)
	if err != nil {
		println(err.Error())
		return
	}

	for i := 1; i < 255; i++ {
		println("low")
		led.Low()
		time.Sleep(time.Millisecond * 500)
		light.Set(ch, light.Top()/uint32(i))

		println("high")
		led.High()
		time.Sleep(time.Millisecond * 500)

		println("tension is equal to", tension.get)

	}

}
