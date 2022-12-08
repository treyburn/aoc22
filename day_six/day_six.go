package day_six

type Scanner struct {
	cap     int
	set     map[rune]struct{}
	current []rune
}

func NewScanner(cap int) *Scanner {
	s := make(map[rune]struct{})
	c := make([]rune, 0)

	return &Scanner{
		cap:     cap,
		set:     s,
		current: c,
	}
}

func (s *Scanner) FindMarker(input []rune) int {
	var marker int
	for _, r := range input {
		marker++

		for !s.isUnique(r) {
			s.dropLeft()
		}

		s.add(r)

		if s.len() == s.cap {
			break
		}

	}
	return marker
}

func (s *Scanner) len() int {
	return len(s.current)
}

func (s *Scanner) isUnique(in rune) bool {
	_, found := s.set[in]

	return !found
}

func (s *Scanner) add(in rune) {
	s.set[in] = struct{}{}

	s.current = append(s.current, in)
	if s.len() > s.cap {
		s.dropLeft()
	}
}

func (s *Scanner) dropLeft() {
	if s.len() == 0 {
		return
	}

	left := s.current[0]
	s.current = s.current[1:]

	delete(s.set, left)
}
