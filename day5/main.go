package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	result1 := solution("input.txt")
	println(result1)

	result2 := solution3("input.txt")
	println(result2)
}

func solution(fileName string) int {

	seeds, map_of_maps := readMap(fileName)

	var closest_location = int(^uint(0) >> 1)
	for _, seed := range seeds {

		seed_number, _ := strconv.Atoi(seed)

		location := getLocation(map_of_maps, seed_number)

		if location < closest_location {
			closest_location = location
		}

	}

	return closest_location
}

// 500 seconds single thread
func solution2(fileName string) int {

	seeds, map_of_maps := readMap(fileName)

	var closest_location = int(^uint(0) >> 1)

	for i := 0; i < len(seeds); i += 2 {

		seed_start, _ := strconv.Atoi(seeds[i])
		seed_range, _ := strconv.Atoi(seeds[i+1])

		for seed_number := seed_start; seed_number < seed_start+seed_range; seed_number++ {

			location := getLocation(map_of_maps, seed_number)
			if location < closest_location {
				closest_location = location
			}
		}

	}

	return closest_location
}

// 320 seconds
func solution3(fileName string) int {

	seeds, map_of_maps := readMap(fileName)

	ch := make(chan int, len(seeds)/2)

	wg := sync.WaitGroup{}

	for i := 0; i < len(seeds); i += 2 {

		seed_start, _ := strconv.Atoi(seeds[i])
		seed_range, _ := strconv.Atoi(seeds[i+1])

		wg.Add(1)

		go func(seed_start int, seed_range int, wg *sync.WaitGroup) {

			closest_location := int(^uint(0) >> 1)
			for seed_number := seed_start; seed_number < seed_start+seed_range; seed_number++ {

				location := getLocation(map_of_maps, seed_number)
				if location < closest_location {
					closest_location = location
				}
			}
			ch <- closest_location
			wg.Done()
		}(seed_start, seed_range, &wg)
	}
	wg.Wait()
	close(ch)

	the_closest_location := int(^uint(0) >> 1)
	for loc := range ch {
		if loc < the_closest_location {
			the_closest_location = loc
		}
	}

	return the_closest_location
}

func getLocation(map_of_maps map[string][][]int, seed_number int) int {
	soil := getMapping(map_of_maps, "seed-to-soil", seed_number)
	fertilizer := getMapping(map_of_maps, "soil-to-fertilizer", soil)
	water := getMapping(map_of_maps, "fertilizer-to-water", fertilizer)
	light := getMapping(map_of_maps, "water-to-light", water)
	temperature := getMapping(map_of_maps, "light-to-temperature", light)
	humidity := getMapping(map_of_maps, "temperature-to-humidity", temperature)
	location := getMapping(map_of_maps, "humidity-to-location", humidity)

	return location
}

func readMap(fileName string) ([]string, map[string][][]int) {

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var seeds []string

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	var map_id string

	// kind_of_map -> [dest_start, source_start, range]
	var map_of_maps = map[string][][]int{}
	for _, line := range lines {
		if strings.HasPrefix(line, "seeds:") {
			seeds = strings.Split(strings.Split(line, ": ")[1], " ")
			continue
		}
		if strings.HasSuffix(line, "map:") {
			map_id = strings.Split(line, " ")[0]
			continue
		}
		if line == "" {
			continue
		}

		numbers := []int{}
		for _, value := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(value)
			numbers = append(numbers, n)
		}
		map_of_maps[map_id] = append(map_of_maps[map_id], numbers)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return seeds, map_of_maps
}
func getMapping(map_of_maps map[string][][]int, map_id string, index int) int {

	input_map := map_of_maps[map_id]

	for _, mapping := range input_map {

		source_start, destination_start, length := mapping[1], mapping[0], mapping[2]
		// if index is not within our mapping we try next one
		if index < source_start || index >= source_start+length {
			continue
		}

		return destination_start + (index - source_start)

	}

	// if we tried all of them then its unmapped
	return index
}
