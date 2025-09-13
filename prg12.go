// Problem Definition:
// You are building a Player Gold Tracker.
// Define a struct called Player with fields:
// name (string)
// gold (int)
// Create a map where the keys are player names and the values are their corresponding Player structs.
// Add at least 3 players into the map.
// Write functions (methods if you want) to:
// Print all players and their gold.
// Find and print the player with the highest gold.
// Increase gold of a specific player (simulate collecting more gold).
package main

import "fmt"

type Player struct {
	name string
	gold int
}

var playerMap = make(map[string]Player)

func GoldGame() {
	for {
		var input string
		fmt.Println("****************************************************************")
		fmt.Println("enter 1 to Display all the players")
		fmt.Println("enter 2 to add player ")
		fmt.Println("enter 3 to know the player with highest gold")
		fmt.Println("enter 4 to exit the game")
		fmt.Scan(&input)
		if input == "4" {
			break
		} else {
			Options(input)
		}
	}
}

func Options(input string) {
	switch input {
	case "1":
		displayPlayers()
	case "2":
		addPlayers()
	case "3":
		highestGoldPlayer()

	}
}

func addPlayers() {
	var Pname string
	var gold int
	fmt.Println("****************************************************************")
	fmt.Println("enter the name of the player :- ")
	fmt.Scan(&Pname)
	fmt.Println("enter the amount of  gold :- ")
	fmt.Scan(&gold)

	playerMap[Pname] = Player{Pname, gold}
}

func displayPlayers() {
	for names, players := range playerMap {
		fmt.Println("****************************************************************")
		fmt.Println("player's name :- ", names, "amount of gold player have :- ", players.gold)
	}
}

func highestGoldPlayer() {
	var maxGold = 0
	var maxGoldPlayer string

	for playerName, i := range playerMap {
		if i.gold > maxGold {
			maxGold = i.gold
			maxGoldPlayer = playerName
		}
	}
	fmt.Println("****************************************************************")
	fmt.Println("player with maximum gold is :- ", maxGoldPlayer, "and amount of gold is:-", maxGold)
}
