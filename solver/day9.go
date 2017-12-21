package solver

import (
	"fmt"
)

var Day9 = Solver{
	Run: runDay9,
}

type scanner struct {
	input []rune
	index int
}

type group struct {
	characters []rune
	groups     []group
}

func runDay9(input []string) (string, error) {
	s := newScanner(input[0])
	group := parseGroup(&s)

	score := calculateRubbish(group)

	output := fmt.Sprintf("%d", score)

	return output, nil
}

func newScanner(input string) scanner {
	return scanner{
		input: []rune(input),
		index: 0,
	}
}

func (s *scanner) next() rune {
	c := s.peek()
	s.index++

	return c
}

func (s scanner) peek() rune {
	return s.input[s.index]
}

func (s scanner) hasNext() bool {
	return s.index < len(s.input)
}

func newGroup(characters []rune, groups []group) group {
	return group{
		characters: characters,
		groups:     groups,
	}
}

func parseGroup(s *scanner) group {
	var groups []group

	s.next()
	for c := s.peek(); c != '}'; c = s.peek() {
		switch c {
		case '{':
			group := parseGroup(s)

			groups = append(groups, group)
		case '<':
			rubbish := parseRubbish(s)

			groups = append(groups, rubbish)
		case ',':
			s.next()
		default:
			panic("Uh oh")
		}
	}
	s.next()

	return newGroup([]rune{}, groups)
}

func parseRubbish(s *scanner) group {
	var characters []rune

	s.next()
	for c := s.next(); c != '>'; c = s.next() {
		if c == '!' {
			s.next()
		} else {
			characters = append(characters, c)
		}
	}

	return newGroup(characters, []group{})
}

func calculateRubbish(g group) int {
	rubbish := len(g.characters)
	for _, subgroup := range g.groups {
		rubbish += calculateRubbish(subgroup)
	}

	return rubbish
}
