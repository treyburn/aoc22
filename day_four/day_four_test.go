package day_four

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildSector(t *testing.T) {
	type testCase struct {
		name    string
		input   string
		wantErr bool
		want    Sector
	}

	var tests = []testCase{
		{"valid", "1-2", false, Sector{1, 2}},
		{"invalid1", "12", true, Sector{}},
		{"invalid2", "1-2-3", true, Sector{}},
		{"invalid3", "1-2,2-3", true, Sector{}},
		{"invalid4", "1-2,", true, Sector{}},
		{"invalid5", "1- 2", true, Sector{}},
		{"invalid6", "foo-bar", true, Sector{}},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got, gotErr := buildSector(tc.input)

			if tc.wantErr {
				assert.Error(t, gotErr)
			} else {
				assert.NoError(t, gotErr)
			}

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBuildSectorPair(t *testing.T) {
	type testCase struct {
		name    string
		input   string
		wantErr bool
		want    SectorPair
	}

	var tests = []testCase{
		{"valid", "1-2,2-3", false,
			SectorPair{
				first:  Sector{1, 2},
				second: Sector{2, 3},
			},
		},
		{"invalid1", "1", true, SectorPair{}},
		{"invalid2", "1-2", true, SectorPair{}},
		{"invalid3", "1-2,", true, SectorPair{}},
		{"invalid4", "1-2 ,1-2", true, SectorPair{}},
		{"invalid5", "1-2,1-2,5-6", true, SectorPair{}},
		{"invalid6", "1-2,1-foo", true, SectorPair{}},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got, gotErr := buildSectorPair(tc.input)

			if tc.wantErr {
				assert.Error(t, gotErr)
			} else {
				assert.NoError(t, gotErr)
			}

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBuildSectorPairs(t *testing.T) {
	type testCase struct {
		name    string
		input   string
		wantErr bool
		want    []SectorPair
	}

	var tests = []testCase{
		{"valid",
			"2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8",
			false,
			[]SectorPair{
				{Sector{2, 4}, Sector{6, 8}},
				{Sector{2, 3}, Sector{4, 5}},
				{Sector{5, 7}, Sector{7, 9}},
				{Sector{2, 8}, Sector{3, 7}},
				{Sector{6, 6}, Sector{4, 6}},
				{Sector{2, 6}, Sector{4, 8}},
			},
		},
		{"invalid1", "2-4,6-82-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8", true, []SectorPair{}},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got, gotErr := BuildSectorPairs(tc.input)

			if tc.wantErr {
				assert.Error(t, gotErr)
			} else {
				assert.NoError(t, gotErr)
			}

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSectorPair_HasFullOverlap(t *testing.T) {
	type testCase struct {
		name  string
		input SectorPair
		want  bool
	}

	var tests = []testCase{
		{"has full overlap", SectorPair{Sector{1, 4}, Sector{2, 3}}, true},
		{"has full overlap - one value total", SectorPair{Sector{1, 1}, Sector{1, 1}}, true},
		{"has full overlap - one value second", SectorPair{Sector{1, 3}, Sector{2, 2}}, true},
		{"has full overlap - one value first", SectorPair{Sector{1, 1}, Sector{1, 2}}, true},
		{"has partial overlap", SectorPair{Sector{1, 3}, Sector{2, 4}}, false},
		{"has no overlap", SectorPair{Sector{1, 3}, Sector{4, 6}}, false},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got := tc.input.HasFullOverlap()

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCountFullOverlaps(t *testing.T) {
	input := []SectorPair{
		{Sector{2, 4}, Sector{6, 8}},
		{Sector{2, 3}, Sector{4, 5}},
		{Sector{5, 7}, Sector{7, 9}},
		{Sector{2, 8}, Sector{3, 7}},
		{Sector{6, 6}, Sector{4, 6}},
		{Sector{2, 6}, Sector{4, 8}},
	}

	want := 2

	got := CountFullOverlaps(input)

	assert.Equal(t, want, got)
}

func TestIsBetween(t *testing.T) {
	type testCase struct {
		name string
		x    int
		y    int
		z    int
		want bool
	}

	var tests = []testCase{
		{"true 1", 1, 3, 2, true},
		{"true 2", 1, 2, 2, true},
		{"true 3", 2, 2, 2, true},
		{"true 4", 3, 1, 2, true},
		{"false 1", 4, 8, 2, false},
		{"false 2", 4, 8, 12, false},
		{"false 3", 8, 4, 2, false},
		{"false 4", 8, 4, 12, false},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got := isBetween(tc.x, tc.y, tc.z)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSectorPair_HasOverlap(t *testing.T) {
	type testCase struct {
		name  string
		input SectorPair
		want  bool
	}

	var tests = []testCase{
		{"first", SectorPair{Sector{2, 4}, Sector{6, 8}}, false},
		{"second", SectorPair{Sector{2, 3}, Sector{4, 5}}, false},
		{"third", SectorPair{Sector{5, 7}, Sector{7, 9}}, true},
		{"fourth", SectorPair{Sector{2, 8}, Sector{3, 7}}, true},
		{"fifth", SectorPair{Sector{6, 6}, Sector{4, 6}}, true},
		{"sixth", SectorPair{Sector{2, 6}, Sector{4, 8}}, true},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got := tc.input.HasOverlap()

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCountPartialOverlaps(t *testing.T) {
	input := []SectorPair{
		{Sector{2, 4}, Sector{6, 8}},
		{Sector{2, 3}, Sector{4, 5}},
		{Sector{5, 7}, Sector{7, 9}},
		{Sector{2, 8}, Sector{3, 7}},
		{Sector{6, 6}, Sector{4, 6}},
		{Sector{2, 6}, Sector{4, 8}},
	}

	want := 4

	got := CountPartialOverlaps(input)

	assert.Equal(t, want, got)
}