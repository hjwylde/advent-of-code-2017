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

	score := calculateScore(group, 1)

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
	var characters []rune
	var groups []group

	s.next()
	for c := s.peek(); c != '}'; c = s.peek() {
		switch c {
		case '{':
			group := parseGroup(s)

			groups = append(groups, group)
		case '<':
			parseRubbish(s)
		case ',':
			s.next()
		default:
			characters = append(characters, c)

			s.next()
		}
	}
	s.next()

	return newGroup(characters, groups)
}

func parseRubbish(s *scanner) {
	for c := s.next(); c != '>'; c = s.next() {
		if c == '!' {
			s.next()
		}
	}
}

func calculateScore(g group, level int) int {
	totalScore := level
	for _, subgroup := range g.groups {
		totalScore += calculateScore(subgroup, level+1)
	}

	return totalScore
}
