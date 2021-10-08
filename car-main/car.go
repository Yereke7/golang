package main

import (
	"fmt"
	"car-main/headlights"
	"car-main/stereo"
	"car-main/wheels"
)

func OpenDoor() {
	fmt.Println("Opening door")
}
func main()  {
	OpenDoor()
	headlights.TurnOn()
	stereo.TurnOn()
	stereo.BoostBass()
	wheels.Steer()
	wheels.Accelerate()
}
