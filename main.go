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
		plusOrMinR        string
		plusOrMinN        string
		plusOrMinO        string
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
					inventory = append(inventory, "refined uranium,")
				}

				continue

			} else if buyRequest == "natural" {

				if wallet < naturalResources {
					fmt.Println("Sorry, you don't have enough money to buy that")
				} else {
					wallet = wallet - naturalResources
					inventory = append(inventory, "natural resources,")
				}

				continue

			} else if buyRequest == "stable" {

				if wallet < stableOganesson {
					fmt.Println("Sorry, you don't have enough money to buy that")
				} else {
					wallet = wallet - stableOganesson
					inventory = append(inventory, "stable oganesson,")
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

			fmt.Printf("Refined Uranium: %v(%v)\nNatural Resources: %v(%v)\nStable Oganesson: %v(%v)\n//////////////////////////\n", refinedUranium, plusOrMinR, naturalResources, plusOrMinN, stableOganesson, plusOrMinO)

		} else if command == "change" {

			pRefinedUranium = refinedUranium
			pNaturalResources = naturalResources
			pStableOganesson = stableOganesson

			refinedUranium = rand.Intn(250000-100000) + 100000
			naturalResources = rand.Intn(150000-1000) + 1000
			stableOganesson = rand.Intn(500000-200000) + 100000

			if refinedUranium > pRefinedUranium {
				plusOrMinR = "+"
			} else if refinedUranium < pRefinedUranium {
				plusOrMinR = "-"
			}
			if naturalResources > pNaturalResources {
				plusOrMinN = "+"
			} else if naturalResources < pNaturalResources {
				plusOrMinN = "-"
			}

			if stableOganesson > pStableOganesson {
				plusOrMinO = "+"
			} else if stableOganesson < pStableOganesson {
				plusOrMinO = "-"
			}

			fmt.Printf("Refined Uranium: %v(%v)\nNatural Resources: %v(%v)\nStable Oganesson: %v(%v)\n//////////////////////////\n", refinedUranium, plusOrMinR, naturalResources, plusOrMinN, stableOganesson, plusOrMinO)
		} else if command == "wallet" {
			fmt.Println(wallet)
		} else if command == "sell" {
			fmt.Println("What would you like to sell?")
			fmt.Scan(&sellRequest)

			if sellRequest == "refined" {

				if wallet >= refinedUranium {
					wallet = wallet + refinedUranium

					for i, v := range inventory {
						if v == "refined uranium," {
							inventory = append(inventory[:i], inventory[i+1:]...)
							break
						}
					}
				} else {
					fmt.Println("You don't have enough money to but that")
				}
			} else if sellRequest == "natural" {

				if wallet >= naturalResources {
					wallet = wallet + naturalResources

					for i, v := range inventory {
						if v == "natural resources," {
							inventory = append(inventory[:i], inventory[i+1:]...)
							break
						}
					}
				}

			} else if sellRequest == "stable" {

				if wallet >= stableOganesson {

					wallet = wallet + stableOganesson

					for i, v := range inventory {
						if v == "stable oganesson," {
							inventory = append(inventory[:i], inventory[i+1:]...)
							break
						}
					}
				}

			}
		} else if command == "past" {
			fmt.Printf("The past prices are:\nRefined uranium: %v\nNatural resources: %v\nStable oganesson: %v\n", pRefinedUranium, pNaturalResources, pStableOganesson)
		}
	}

}
