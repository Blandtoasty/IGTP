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
	fmt.Println("The prices also show if it is an increase compared to the last prices and if it's affordable (A) or too expensive (TE)")

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
		aOrNoR            string
		aOrNoN            string
		aOrNoO            string
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

			} else {
				fmt.Println("What you typed is not correct")
			}
			continue

		} else if command == "help" {

			fmt.Println("'buy' Buy something\n'inventory' See your inventory\n'exit' Exit the program\n'prices' Show the prices\n'change' Change the prices\n'wallet' Show how much money you have\n'sell' Sell something\n'past' Show the past prices")

		} else if command == "inventory" {

			fmt.Println(inventory)

		} else if command == "exit" {

			os.Exit(0)

		} else if command == "prices" {

			fmt.Printf("Refined Uranium: %v(%v, %v)\nNatural Resources: %v(%v, %v)\nStable Oganesson: %v(%v, %v)\n//////////////////////////\n", refinedUranium, plusOrMinR, aOrNoR, naturalResources, plusOrMinN, aOrNoN, stableOganesson, plusOrMinO, aOrNoO)

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

			if wallet >= refinedUranium {
				aOrNoR = "A"
			} else if wallet <= refinedUranium {
				aOrNoR = "TE"
			}

			if wallet >= naturalResources {
				aOrNoN = "A"
			} else if wallet <= naturalResources {
				aOrNoN = "TE"
			}

			if wallet >= stableOganesson {
				aOrNoO = "A"
			} else if wallet <= stableOganesson {
				aOrNoO = "TE"
			}

			fmt.Printf("//////////////////////////\nRefined Uranium: %v(%v, %v)\nNatural Resources: %v(%v, %v)\nStable Oganesson: %v(%v, %v)\n//////////////////////////\n", refinedUranium, plusOrMinR, aOrNoR, naturalResources, plusOrMinN, aOrNoN, stableOganesson, plusOrMinO, aOrNoO)

		} else if command == "wallet" {
			fmt.Println(wallet)
		} else if command == "sell" {
			fmt.Println("What would you like to sell?")
			fmt.Scan(&sellRequest)

			if sellRequest == "refined" {

				for i, v := range inventory {
					if v == "refined uranium," {
						inventory = append(inventory[:i], inventory[i+1:]...)
						wallet = wallet + refinedUranium
						break
					}
				}

			} else if sellRequest == "natural" {

				for i, v := range inventory {
					if v == "natural resources," {
						inventory = append(inventory[:i], inventory[i+1:]...)
						wallet = wallet + naturalResources
						break
					}
				}

			} else if sellRequest == "stable" {

				for i, v := range inventory {
					if v == "stable oganesson," {
						inventory = append(inventory[:i], inventory[i+1:]...)
						wallet = wallet + stableOganesson
						break
					}
				}

			} else {
				fmt.Println("What you typed is not correct")
			}
		} else if command == "past" {
			fmt.Printf("The past prices are:\nRefined uranium: %v\nNatural resources: %v\nStable oganesson: %v\n", pRefinedUranium, pNaturalResources, pStableOganesson)
		} else {
			fmt.Println("What you typed is not correct")
		}
	}

}
