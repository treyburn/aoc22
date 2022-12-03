package day_three

import (
	"errors"
	"fmt"
	"strings"
)

type Ruck struct {
	raw []rune
	itemsPerComp int
	comp1 map[rune]int
	comp2 map[rune]int
}

// NewRuck creates a ruck struct from a string
func NewRuck(input string) *Ruck {
	c1 := make(map[rune]int)
	c2 := make(map[rune]int)

	raw := []rune(input)
	itemsPerComp := len(raw)/2

	return &Ruck{raw: raw, itemsPerComp: itemsPerComp, comp1: c1, comp2: c2}
}

// DeDupe returns the duplicate value shared between the ruck compartments while lazily loading the compartments
func (r *Ruck) DeDupe() rune {
	var duplicate rune
	for idx, item := range r.raw {
		if idx < r.itemsPerComp {
			_, ok := r.comp1[item]
			if !ok {
				r.comp1[item] = 1
			} else {
				r.comp1[item]++
			}
		} else {
			if _, ok := r.comp1[item]; ok {
				duplicate = item
				continue
			}
			_, ok := r.comp2[item]
			if !ok {
				r.comp2[item] = 1
			} else {
				r.comp2[item]++
			}
		}
	}
	r.raw = make([]rune, 0)
	return duplicate
}

// IsDeDuped informs the caller if the DeDupe func has been called or not
func (r *Ruck) IsDeDuped() bool {
	return len(r.raw) == 0
}

// UniqueItems returns a list of all unique items between the compartments of ruck and a bool indicating if DeDupe has been run yet
func (r *Ruck) UniqueItems() ([]rune, bool) {
	items := make([]rune, 0)

	if !r.IsDeDuped() {
		return items, false
	}

	for item := range r.comp1 {
		items = append(items, item)
	}

	for item := range r.comp2 {
		items = append(items, item)
	}

	return items, true
}

// BuildRucks takes a newline separated string and creates a slice of Rucks
// 	where each line in the string is turned into a Ruck
func BuildRucks(input string) []*Ruck {
	rucks := make([]*Ruck, 0)

	split := strings.Split(input, "\n")
	for _, line := range split {
		if line != "" {
			ruck := NewRuck(line)
			rucks = append(rucks, ruck)
		}
	}

	return rucks
}

// GetPriority returns the priority of the given rune where a = 1, z = 26, A = 27, Z = 52
// 	only handles a-zA-Z - any other character returns -1
func GetPriority(r rune) int {
	if r < 65 || r > 122 { // filter bad values
		return -1
	}
	if r > 90 { // lowercase
		return int(r) - 96
	}
	return int(r) - 38 // uppercase
}

// SumDupePriorities takes a slice of rucks, de-dupes each ruck, gets the priority of the dupe
//	and returns the sum of the valid priorities
func SumDupePriorities(rucks []*Ruck) int {
	var total int

	for _, ruck := range rucks {
		dupe := ruck.DeDupe()
		priority := GetPriority(dupe)
		if priority != -1 {
			total += priority
		}
	}

	return total
}

type Group struct {
	one *Ruck
	two *Ruck
	three *Ruck
}

func NewGroup(one, two, three *Ruck) Group {
	return Group{one: one, two: two, three: three}
}

func BuildGroups(rucks []*Ruck) ([]Group, error) {
	groups := make([]Group, 0)

	numRucks := len(rucks)

	if numRucks % 3 != 0 {
		return groups, fmt.Errorf("invalid count of groups %v - must be evenly divisible by 3", numRucks)
	}

	for i := 0; i < numRucks; i += 3 {
		one := rucks[i]
		two := rucks[i + 1]
		three := rucks[i + 2]
		groups = append(groups, NewGroup(one, two, three))
	}

	return groups, nil
}

// ListBadges returns all the unique badges held between each member
func (g Group) ListBadges() ([]rune, error) {
	badges := make([]rune, 0)

	// nil check
	if g.one == nil || g.two == nil || g.three == nil {
		return badges, errors.New("invalid ruck(s) present in group")
	}

	items, ok := g.one.UniqueItems()
	if !ok {
		return badges, errors.New("dedupe not run on ruck(s)")
	}
	badges = append(badges, items...)

	items, ok = g.two.UniqueItems()
	if !ok {
		return badges, errors.New("dedupe not run on ruck(s)")
	}
	badges = append(badges, items...)

	items, ok = g.three.UniqueItems()
	if !ok {
		return badges, errors.New("dedupe not run on ruck(s)")
	}
	badges = append(badges, items...)

	return badges, nil
}

// GetGroupBadge returns the common badge shared between all members of the group
func GetGroupBadge(group Group) (rune, error) {
	common := make(map[rune]int)

	badges, err := group.ListBadges()
	if err != nil {
		return -1, err
	}

	for _, badge := range badges {
		_, ok := common[badge]
		if !ok {
			common[badge] = 1
		} else {
			common[badge]++
		}
	}

	for badge, count := range common {
		if count == 3 {
			return badge, nil
		}
	}

	return -1, errors.New("no common badge shared in the group")
}

func SumBadgePriorities(groups []Group) (int, error) {
	var sum int
	for _, group := range groups {
		badge, err := GetGroupBadge(group)
		if err != nil {
			return sum, err
		}
		priority := GetPriority(badge)
		if priority != -1 {
			sum += priority
		}
	}
	return sum, nil
}