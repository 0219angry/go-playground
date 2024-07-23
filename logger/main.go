package main

import (
	"github.com/0219Angry/go-playground/logger/first"
	"github.com/0219Angry/go-playground/logger/second"
)

func main() {
	first := first.InitFirst("Counter1")
	second := second.InitSecond("Counter2")

	first.CountUp()
	first.CountUp()
	first.CountUp()
	second.CountUp()
	second.CountUp()

}
