package main

import (
	"fmt"
)

// Define a struct to represent a travel route
type TravelRoute struct {
	Origin      string
	Destination string
	Distance    float64
	TravelTime  string
}

// Function to calculate the distance and travel time between two cities
func calculateTravelInfo(origin, destination string) (float64, string) {
	// Placeholder values for demonstration purposes
	var distance float64
	var travelTime string

	// Placeholder calculations based on origin and destination
	if origin == "Rome" && destination == "Barcelona" {
		distance = 1000.0 // Assigning a placeholder distance value
		travelTime = "10 hours" // Assigning a placeholder travel time
	} else if origin == "Marrakech" && destination == "Lisbon" {
		distance = 500.0 // Assigning a placeholder distance value
		travelTime = "5 hours" // Assigning a placeholder travel time
	} else if origin == "Florence" && destination == "Casablanca" {
		distance = 800.0 // Assigning a placeholder distance value
		travelTime = "8 hours" // Assigning a placeholder travel time
	}

	// Returning the calculated distance and travel time
	return distance, travelTime
}

func main() {
	// Define a slice of TravelRoute structs to store multiple routes
	routes := []TravelRoute{
		{Origin: "Rome", Destination: "Barcelona"},
		{Origin: "Marrakech", Destination: "Lisbon"},
		{Origin: "Florence", Destination: "Casablanca"},
	}

	// Loop through each route and calculate travel info
	for _, route := range routes {
		// Calling the calculateTravelInfo function to get the distance and travel time
		distance, travelTime := calculateTravelInfo(route.Origin, route.Destination)

		// Printing the origin, destination, distance, and travel time for each route
		fmt.Printf("From %s to %s:\n", route.Origin, route.Destination) // Printing the origin and destination cities
		fmt.Printf("Distance: %.2f kilometers\n", distance) // Printing the distance with 2 decimal places
		fmt.Printf("Estimated travel time: %s\n", travelTime) // Printing the estimated travel time
		fmt.Println() 
	}
}
