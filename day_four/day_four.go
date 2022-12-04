package day_four

import (
	"fmt"
	"strconv"
	"strings"
)

type Sector struct {
	lowerBound int
	upperBound int
}

type SectorPair struct {
	first  Sector
	second Sector
}

// HasOverlap returns true if the lower/upper bounds of either sector overlap
func (sp SectorPair) HasOverlap() bool {
	return isBetween(sp.first.lowerBound, sp.first.upperBound, sp.second.lowerBound) ||
		isBetween(sp.first.lowerBound, sp.first.upperBound, sp.second.upperBound) ||
		isBetween(sp.second.lowerBound, sp.second.upperBound, sp.first.lowerBound) ||
		isBetween(sp.second.lowerBound, sp.second.upperBound, sp.first.upperBound)
}

// HasFullOverlap returns true if the lower/upper bounds of any sector is fully overlapped by the other
func (sp SectorPair) HasFullOverlap() bool {
	return (isBetween(sp.first.lowerBound, sp.first.upperBound, sp.second.lowerBound) &&
		isBetween(sp.first.lowerBound, sp.first.upperBound, sp.second.upperBound)) ||
		(isBetween(sp.second.lowerBound, sp.second.upperBound, sp.first.lowerBound) &&
			isBetween(sp.second.lowerBound, sp.second.upperBound, sp.first.upperBound))
}

// buildSector takes a string in the format of `X-Y` and outputs a Sector struct
func buildSector(input string) (Sector, error) {
	split := strings.Split(input, "-")
	if len(split) != 2 {
		return Sector{}, fmt.Errorf("invalid input string for sector: %v", input)
	}

	lower, err := strconv.Atoi(split[0])
	if err != nil {
		return Sector{}, err
	}

	upper, err := strconv.Atoi(split[1])
	if err != nil {
		return Sector{}, err
	}

	return Sector{lowerBound: lower, upperBound: upper}, nil
}

// buildSectorPair takes a string in the format of `X-Y,A-B` and produces a SectorPair struct
func buildSectorPair(input string) (SectorPair, error) {
	split := strings.Split(input, ",")
	if len(split) != 2 {
		return SectorPair{}, fmt.Errorf("invalid input string for sector pair: %v", input)
	}

	s1, err := buildSector(split[0])
	if err != nil {
		return SectorPair{}, err
	}

	s2, err := buildSector(split[1])
	if err != nil {
		return SectorPair{}, err
	}

	return SectorPair{first: s1, second: s2}, nil
}

// BuildSectorPairs takes newline separated strings that are valid for building a SectorPair and returns them as a slice
func BuildSectorPairs(input string) ([]SectorPair, error) {
	pairs := make([]SectorPair, 0)
	split := strings.Split(input, "\n")
	for _, pair := range split {
		if pair != "" {
			sp, err := buildSectorPair(pair)
			if err != nil {
				return pairs, err
			}
			pairs = append(pairs, sp)
		}
	}

	return pairs, nil
}

func CountFullOverlaps(input []SectorPair) int {
	var total int
	for _, pair := range input {
		if pair.HasFullOverlap() {
			total++
		}
	}

	return total
}

func CountPartialOverlaps(input []SectorPair) int {
	var total int
	for _, pair := range input {
		if pair.HasOverlap() {
			total++
		}
	}

	return total
}

// isBetween returns true if the value z is between x and y
func isBetween(x, y, z int) bool {
	return (x <= z && z <= y) || (x >= z && z >= y)
}
