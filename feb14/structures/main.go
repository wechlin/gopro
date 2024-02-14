package main

import "fmt"

type Km float64

func (km Km) ToMiles() float64 {
	return float64(km) * 0.628
}

type tank struct {
	model        string
	topSpeedKmHr float64
}

func (t *tank) String() string {
	return fmt.Sprintf("%s (%3.2fkm/h)", t.model, t.topSpeedKmHr)
}

func (t *tank) TurboCharge(extraSpeed float64) {
	t.topSpeedKmHr += extraSpeed
}

func main() {
	var abrams tank = tank{model: "M1 Abrams", topSpeedKmHr: 72.0}
	var tiger tank = tank{model: "Tiger", topSpeedKmHr: 48.3}

	var distance = Km(750)

	fmt.Println(distance.ToMiles())
	fmt.Println(abrams.model)
	fmt.Println(tiger)
	tiger.TurboCharge(10.1)
	fmt.Println(tiger)
}
