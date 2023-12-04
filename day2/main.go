// mbowolick advent-of-code-2023/day2 attempt
//
// https://adventofcode.com/2023/day/2
//
// --- Day 2: Cube Conundrum ---
//
// The Elf shows you a small bag and some cubes which are either
// red, green, or blue. Each time you play this game, he will hide a
// secret number of cubes of each color in the bag, and your goal is
// to figure out information about the number of cubes.

// You play several games and record the information from each game
// (your puzzle input). Each game is listed with its ID number (like the
// 11 in Game 11: ...) followed by a semicolon-separated list of subsets
// of cubes that were revealed from the bag (like 3 red, 5 green, 4 blue).

// For example, the record of a few games might look like this:
// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

// In game 1, three sets of cubes are revealed from the bag (and then put
// back again):
// [1] The first set: 	is 3 blue cubes and 4 red cubes;
// [2] The second set: 	is 1 red cube, 2 green cubes, and 6 blue cubes;
// [3] The third set:		is only 2 green cubes.

// The Elf would first like to know which games would have been possible if
// the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?

// In the example above, games 1, 2, and 5 would have been possible if the bag
// had been loaded with that configuration. However, game 3 would have been
// impossible because at one point the Elf showed you 20 red cubes at once;
// similarly, game 4 would also have been impossible because the Elf showed you
// 15 blue cubes at once. If you add up the IDs of the games that would have been
// possible, you get 8.

// TASK (according to an input):
// Determine which games would have been possible if the bag had been loaded with
// only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs
// of those games?

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) ([]string, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	var lines []string

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	return lines, scanner.Err()
}

func parseGameData(event string) (int, map[string]int, error) {
	// According to an event input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	// Parse the data and return the game index and a map of the colours and their values
	// E.g. returns { 1: { "red": 5, "green": 4, "blue": 9 } } 
	
	event_split := strings.Split(event, ":")

	// Get the game index number
	game_split := strings.Split(event_split[0], " ") // ["Game", "15"]
	game_index, err := strconv.Atoi(game_split[1])
	if err != nil {
		return 0, nil, err
	} 
	
	// For each pull during the game build the gameData map
	gameData := map[string]int{
    "red": 0,
    "green": 0,
    "blue": 0,
	}

	// data_split_raw := strings.TrimSpace(event_split[1])
	data_split := strings.Split(event_split[1], ";")
	fmt.Printf("\nGame %d: \"%s\"\n", game_index, data_split)

	// data_split at this point is: [" 1 red, 2 green, 6 blue", "2 green"]
	// Loop over the split and assess each pull_set
	for _, pull_set := range data_split {
		// example pull_set at this point: " 1 red, 2 green, 6 blue"
		pull_set_split := strings.Split(pull_set, ",") // [" 1 red" , " 2 green", " 6 blue"]
		
		for _, pull_keyvalue := range pull_set_split {
			pull_keyvalue_trim := strings.TrimSpace(pull_keyvalue)
			pull := strings.Split(pull_keyvalue_trim, " ")
			
			// Extract colour
			pull_colour := pull[1] // "red"

			// Extract count
			pull_count_int, _ := strconv.Atoi(pull[0]) // 1
			fmt.Printf("|%s:%d|\n",pull_colour, pull_count_int)
			
			// Update gameData if a value is seen higher than previous
			if pull_count_int > gameData[pull_colour] {
				gameData[pull_colour] = pull_count_int
			}
		}
	}
	return game_index, gameData, nil
}

func main() {
	// According to a set of N pull events
	// Check if the number of the same colour across N is <= POSSIBLE NUMBER
	// If all checked colours are within the POSSIBLE NUMBER the GAME ID WINS.
	// Add all winning IDS. 
	
	fmt.Print("Starting day 3\n\n")
	inputDocLines, err := readInput("input.txt")
	
	if err != nil {
		fmt.Println("Issue reading input:", err.Error())
	}

	only_red := 12
	only_green := 13
	only_blue := 14

	var possible_games []int
	possible_calculation := 0
	for _, game_input := range inputDocLines {
		game_index, gameData, err  := parseGameData(game_input)

		if err != nil {
			fmt.Printf("Something went wrong parsing the game data: %s", err)
		}

		if ( 
			gameData["red"] <= only_red && 
			gameData["green"] <= only_green && 
			gameData["blue"] <= only_blue ) {
			possible_games = append(possible_games, game_index)
			possible_calculation += game_index
			fmt.Printf("{{ POSSIBLE }}\n")
		}
	}
	
	
	fmt.Print("\n\nPossible games include: \n")
	fmt.Printf("%d\n\n", possible_games)
	
	fmt.Print("With a calculation of: \n")
	fmt.Printf("%d\n\n", possible_calculation)
}

