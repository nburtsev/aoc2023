package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type CubeHand struct {
	red   int
	green int
	blue  int
}

type GameResult struct {
	id       int
	possible bool
	minPower int
}

func main() {
	var idSum = 0
	var powerSum = 0

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gameResult := processLine(scanner.Text(), CubeHand{red: 12, green: 13, blue: 14})
		if gameResult.possible {
			idSum += gameResult.id
		}
		powerSum += gameResult.minPower
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	println(idSum)
	println(powerSum)
}

func parseGame(game string) CubeHand {
	var result = CubeHand{0, 0, 0}
	colors := strings.Split(game, ",")
	for _, color := range colors {
		var split = strings.Split(strings.TrimSpace(color), " ")
		var count, _ = strconv.Atoi(split[0])
		var color = split[1]
		if color == "red" {
			result.red = count
		} else if color == "green" {
			result.green = count
		} else if color == "blue" {
			result.blue = count
		}
	}
	return result
}

func processLine(line string, bag CubeHand) GameResult {
	split_game := strings.Split(line, ":")
	id, _ := strconv.Atoi(strings.Split(split_game[0], " ")[1])
	games := strings.Split(split_game[1], ";")

	var minGame = CubeHand{0, 0, 0}
	var possible = true

	for _, game := range games {
		game := parseGame(game)
		if game.red > bag.red || game.green > bag.green || game.blue > bag.blue {
			possible = false
		}
		if game.red > minGame.red {
			minGame.red = game.red
		}
		if game.blue > minGame.blue {
			minGame.blue = game.blue
		}
		if game.green > minGame.green {
			minGame.green = game.green
		}
	}
	return GameResult{id, possible, minGame.red * minGame.green * minGame.blue}
}
