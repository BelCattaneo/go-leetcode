package main

import (
	"fmt"
)


type Customer struct {
	Id int
}

type TripName struct {
	StartStation string
	EndStation string
}

type TripData struct {
	Customer Customer
	Duration int
}

type OngoingTripData struct {
	StartStation string
	StartTime int
}


type UndergroundSystem struct {
	Trips map[TripName][]TripData
	OngoingTrips map[Customer]OngoingTripData
}


func Constructor() UndergroundSystem {
	return UndergroundSystem{
		Trips:        make(map[TripName][]TripData),
		OngoingTrips: make(map[Customer]OngoingTripData),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
	//A customer with a card ID equal to id, checks in at the station stationName at time t.
	//A customer can only be checked into one place at a time.
	this.OngoingTrips[Customer{Id: id}] = OngoingTripData{
		StartStation: stationName,
		StartTime:    t,
	}
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
	//A customer with a card ID equal to id, checks out from the station stationName at time t.
	customer := Customer{Id: id}
	ongoingTrip := this.OngoingTrips[customer]
	tripName := TripName{
		StartStation: ongoingTrip.StartStation,
		EndStation:   stationName,
	}
	this.Trips[tripName] = append(this.Trips[tripName], TripData{
		Customer: Customer{id},
		Duration: t - ongoingTrip.StartTime,
	})

	delete(this.OngoingTrips, customer)

}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	//Returns the average time it takes to travel from startStation to endStation.
	//The average time is computed from all the previous traveling times from startStation to
	//endStation that happened directly, meaning a check in at startStation followed by a check
	//out from endStation.
	//The time it takes to travel from startStation to endStation may be different from the
	//time it takes to travel from endStation to startStation.
	//There will be at least one customer that has traveled from startStation to endStation
	//before getAverageTime is called.
	tripName := TripName{
		StartStation: startStation,
		EndStation:   endStation,
	}
	var totalTime int
	for _, tripData := range this.Trips[tripName] {
		totalTime += tripData.Duration
	}
	return float64(totalTime / len(this.Trips[tripName]))

}


/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
func main() {
	obj := Constructor()
	obj.CheckIn(45, "Leyton", 3)
	obj.CheckIn(45, "Leyton", 3)
	obj.CheckIn(32, "Paradise", 8)
	obj.CheckIn(27, "Leyton", 10)
	obj.CheckOut(45, "Waterloo", 15)  // Customer 45 "Leyton" -> "Waterloo" in 15-3 = 12
	obj.CheckOut(27, "Waterloo", 20)  // Customer 27 "Leyton" -> "Waterloo" in 20-10 = 10
	obj.CheckOut(32, "Cambridge", 22) // Customer 32 "Paradise" -> "Cambridge" in 22-8 = 14
	param1 := obj.GetAverageTime("Paradise", "Cambridge") // return 14.00000. One trip "Paradise" -> "Cambridge", (14) / 1 = 14
	param2 := obj.GetAverageTime("Leyton", "Waterloo")    // return 11.00000. Two trips "Leyton" -> "Waterloo", (10 + 12) / 2 = 11
	obj.CheckIn(10, "Leyton", 24)
	param3 := obj.GetAverageTime("Leyton", "Waterloo")    // return 11.00000
	obj.CheckOut(10, "Waterloo", 38)  // Customer 10 "Leyton" -> "Waterloo" in 38-24 = 14
	param4 := obj.GetAverageTime("Leyton", "Waterloo")    // return 12.00000. Three trips "Leyton" -> "Waterloo", (10 + 12 + 14) / 3 = 12
	fmt.Println(param1)
	fmt.Println(param2)
	fmt.Println(param3)
	fmt.Println(param4)
}
