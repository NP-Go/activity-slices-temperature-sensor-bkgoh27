package main

import "fmt"

func AverageTemperature(room []float64) float64 {
	var sum, average float64
	for i := 0; i < len(room); i++ {
		sum += room[i]
	}
	average = sum / float64(len(room))
	fmt.Println("The average room temperature is", average)
	return average
}

func main() {
	//slices template
	room1 := []float64{20, 21, 23, 25, 22}
	room2 := []float64{27, 23, 25, 20, 30, 24}
	room3 := []float64{22, 23, 24, 22}

	AverageTemperature(room1)
	AverageTemperature(room2)
	AverageTemperature(room3)
}
