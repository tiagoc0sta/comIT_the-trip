package main // Package declaration, main is the entry point for the executable program

import (
	"fmt" // Importing the fmt package for formatted I/O
)

type TravelRoute struct {

}


// Function to calculate the distance and travel time between two cities
func calculateTravelInfo(origin, destination string) (float64, string) {

	// Placeholder values for demonstration purposes
	distance := 0.0 // Initializing the distance variable
	travelTime := "" // Initializing the travelTime variable

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
origin := "Rome" // Setting the origin city
destination := "Barcelona" // Setting the destination city

// Calling the calculateTravelInfo function to get the distance and travel time
distance, travelTime := calculateTravelInfo(origin, destination)

// Printing the origin, destination, distance, and travel time
fmt.Printf("From %s to %s:\n", origin, destination) // Printing the origin and destination cities
fmt.Printf("Distance: %.2f kilometers\n", distance) // Printing the distance with 2decimal places
fmt.Printf("Estimated travel time: %s\n", travelTime) // Printing the estimated travel time

}