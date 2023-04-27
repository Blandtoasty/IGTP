package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var inventory []string

	fmt.Println("Welcome to the Intergalactic trading program")
	fmt.Println("Once the prices appear, you can start playing, use 'help' to view commands")

	var (
		refinedUranium   int
		naturalResources int
		stableOganesson  int
		command          string
	)

	//Price loop
	for range time.Tick(time.Second * 20) {

		refinedUranium = rand.Intn(250000-100000) + 100000
		naturalResources = rand.Intn(150000-1000) + 1000
		stableOganesson = rand.Intn(500000-200000) + 100000

		fmt.Printf("Refined Uranium: %v\nNatural Resources: %v\nStable Oganesson: %v\n//////////////////////////\n", refinedUranium, naturalResources, stableOganesson)

		var (
			buyRequest string
		)

		for {
			//Commands
			fmt.Scan(&command)
			if command == "buy" {
				fmt.Println("What would you like to buy?")
				fmt.Scan(&buyRequest)
				if buyRequest == "refined uranium" {
					inventory = append(inventory, "Refined Uranium")
					fmt.Println(inventory)
				}

			}
		}

	}
}
