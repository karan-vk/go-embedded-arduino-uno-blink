package main

import (
	"machine"
	"time"
)

func main() {

	led := machine.D13
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledSwitch := true
	counter := 0

	for {
		led.Set(ledSwitch)
		ledSwitch = !ledSwitch
		delay(1000)
		counter++
		println("The count is ", counter)

	}

}

func delay(t uint32) {
	time.Sleep(time.Duration(1000000 * t))
}
