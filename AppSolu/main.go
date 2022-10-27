package main

import (
	"time"
)

func main() {
	debug := InitDebug()

	lora := InitLoRa()

	println("start")

	ATWrite(lora, "AT+JOIN")

	time.Sleep(time.Second * 10)

	//ATWrite(lora, `AT+MSG="ALGOSUP"`)

	time.Sleep(time.Second * 10)

	joinOut := ReadBuffer(lora, debug)
	println(joinOut)

	println("done")
}
