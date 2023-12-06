package main

import (
	"aoc2023"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Triplet [3]int

func (t Triplet) Map(key int) int {
	//log.Println("Mapping", key, t)
	if key < t[0] {
		return key
	}
	if key >= t[0]+t[2] {
		return key
	}

	return t[1] + (key - t[0])
}

type Triplets []Triplet

func (t Triplets) Len() int {
	return len(t)
}

func (t Triplets) Less(i, j int) bool {
	return t[i][0] < t[j][0]
}

func (t Triplets) Swap(i, j int) {
	buf := t[i]
	t[i] = t[j]
	t[j] = buf
}

func (t Triplets) Map(key int) int {
	pos := sort.Search(len(t), func(i int) bool {
		return t[i][0] > key
	})

	if pos != 0 {
		pos -= 1
	}
	return t[pos].Map(key)
}

type Almanac struct {
	Seeds                 []int
	SeedToSoil            Triplets
	SoilToFertilizer      Triplets
	FertilizerToWater     Triplets
	WaterToLight          Triplets
	LightToTemperature    Triplets
	TemperatureToHumidity Triplets
	HumidityToLocation    Triplets
}

func (al *Almanac) GetTriplets() []Triplets {
	return []Triplets{
		al.SeedToSoil,
		al.SoilToFertilizer,
		al.FertilizerToWater,
		al.WaterToLight,
		al.LightToTemperature,
		al.TemperatureToHumidity,
		al.HumidityToLocation,
	}
}

func (al *Almanac) GetSeedLocation(seed int) int {
	result := seed

	for _, triplet := range al.GetTriplets() {
		//prev := result
		result = triplet.Map(result)
		//log.Println(seed, prev, result)
	}

	return result
}

func parseInts(s string) []int {
	result := make([]int, 0, 0)

	for _, val := range strings.Split(s, " ") {
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, num)
	}

	return result
}

func parseTriplets(s string) Triplets {
	result := Triplets{}
	for _, line := range strings.Split(s, "\n")[1:] {
		ints := parseInts(line)
		result = append(result, Triplet{ints[1], ints[0], ints[2]})
	}
	sort.Sort(result)
	return result
}

func parseInput(input []byte) *Almanac {
	groups := strings.Split(string(input), "\n\n")

	seeds := parseInts(strings.TrimPrefix(groups[0], "seeds: "))

	almanac := &Almanac{Seeds: seeds}

	for i, group := range groups[1:] {
		trip := parseTriplets(group)
		switch i {
		case 0:
			almanac.SeedToSoil = trip
		case 1:
			almanac.SoilToFertilizer = trip
		case 2:
			almanac.FertilizerToWater = trip
		case 3:
			almanac.WaterToLight = trip
		case 4:
			almanac.LightToTemperature = trip
		case 5:
			almanac.TemperatureToHumidity = trip
		case 6:
			almanac.HumidityToLocation = trip
		default:
			panic(i)
		}
	}
	//log.Println(almanac)
	return almanac
}

func PartOne(al *Almanac) int {
	result := math.MaxInt
	for _, seed := range al.Seeds {
		location := al.GetSeedLocation(seed)
		//log.Println(seed, location)
		result = aoc2023.Min(result, location)
	}

	return result
}

func PartTwo(al *Almanac) int {
	result := math.MaxInt
	for i := 0; i < len(al.Seeds); i += 2 {
		left := al.Seeds[i]
		right := left + al.Seeds[i+1]

		for j := left; j < right; j += 1 {
			location := al.GetSeedLocation(j)
			result = aoc2023.Min(result, location)
		}
		log.Println("Seed done", i)
	}
	return result
}

func main() {
	input := aoc2023.ReadBytes("day5/input")
	//log.Println(parseInput(input).GetSeedLocation(14))
	//log.Println(PartOne(parseInput(input)))
	log.Println(PartTwo(parseInput(input)))
}
