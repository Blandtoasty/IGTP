package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	fmt.Println("Welcome to the Intergalactic trading program")
	fmt.Println("Use 'help' to view commands")
	fmt.Println("Your starting money is 200,000")
	fmt.Println("Start the game with 'change'")

	var (
		refinedUranium    int
		naturalResources  int
		stableOganesson   int
		command           string
		buyRequest        string
		inventory         []string
		wallet            int
		sellRequest       string
		pRefinedUranium   int
		pNaturalResources int
		pStableOganesson  int
	)

	wallet = 200000

	for {

		//Commands
		fmt.Scan(&command)

		if command == "buy" {

			fmt.Println("What would you like to buy?")
			fmt.Scan(&buyRequest)

			if buyRequest == "refined" {

				if wallet < refinedUranium {
					fmt.Println("Sorry, you don't have enough money to buy that")
				} else {
					wallet = wallet - refinedUranium
					inventory = append(inventory, "Refined Uranium,")
				}

				continue

			} else if buyRequest == "natural" {

				if wallet < naturalResources {
					fmt.Println("Sorry, you don't have enough money to buy that")
				} else {
					wallet = wallet - naturalResources
					inventory = append(inventory, "Natural Resources,")
				}

				continue

			} else if buyRequest == "stable" {

				if wallet < stableOganesson {
					fmt.Println("Sorry, you don't have enough money to buy that")
				} else {
					wallet = wallet - stableOganesson
					inventory = append(inventory, "Stable Oganesson,")
				}
				continue

			}
			continue

		} else if command == "help" {

			fmt.Println("'buy' Buy something\n'inventory' See your inventory\n'exit' Exit the program\n'prices' Show the prices\n'change' Change the prices\n'wallet' Show how much money you have\n'sell' Sell something\n'past' Show the past prices")

		} else if command == "inventory" {

			fmt.Println(inventory)

		} else if command == "exit" {

			os.Exit(0)

		} else if command == "prices" {

			fmt.Printf("Refined Uranium: %v\nNatural Resources: %v\nStable Oganesson: %v\n//////////////////////////\n", refinedUranium, naturalResources, stableOganesson)

		} else if command == "change" {

			pRefinedUranium = refinedUranium
			pNaturalResources = naturalResources
			pStableOganesson = stableOganesson

			refinedUranium = rand.Intn(250000-100000) + 100000
			naturalResources = rand.Intn(150000-1000) + 1000
			stableOganesson = rand.Intn(500000-200000) + 100000

			fmt.Printf("Refined Uranium: %v\nNatural Resources: %v\nStable Oganesson: %v\n//////////////////////////\n", refinedUranium, naturalResources, stableOganesson)
		} else if command == "wallet" {
			fmt.Println(wallet)
		} else if command == "sell" {
			fmt.Println("What would you like to sell?")
			fmt.Scan(&sellRequest)

			if sellRequest == "refined" {

			}

		} else if command == "past" {
			fmt.Printf("The past prices are:\nRefined uranium: %v\nNatural resources: %v\nStable oganesson: %v\n", pRefinedUranium, pNaturalResources, pStableOganesson)
		}
	}

}
