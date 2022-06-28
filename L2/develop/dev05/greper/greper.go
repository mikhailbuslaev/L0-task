package greper

import (
	"fmt"
	"strings"
)

type Greper interface {
	Grep([]string, string) ([]string, error)
}

type Params struct {
	BeforeLength int
	AfterLength  int
	LineNum      bool // return only indexes
}

type FixedGreper struct{}

type DefaultGreper struct {
	Params Params
}

type InvertGreper struct{}

func (f *FixedGreper) Grep(rows []string, pattern string) ([]string, error) {
	result := make([]string, 0, len(rows))
	for i := range rows {
		if rows[i] == pattern {
			result = append(result, rows[i])
		}
	}
	return result, nil
}

func (d *DefaultGreper) Grep(rows []string, pattern string) ([]string, error) {
	result := make([]string, 0, len(rows))
	for i := range rows {
		rowArr := strings.Split(rows[i], "")
		rowLength := len(rowArr)
		patternArr := (strings.Split(pattern, ""))
		patternLength := len(patternArr)
		for j := 0; j < rowLength-patternLength+1; j++ {
			word := strings.Join(rowArr[j:j+patternLength], "")
			if word == pattern {

				if d.Params.LineNum {
					result = append(result, fmt.Sprintf("%d", i))
					continue
				}

				if d.Params.BeforeLength > i || d.Params.AfterLength > len(rows[i])-i {
					return []string{""}, fmt.Errorf("afterLength or beforeLength too far")
				}

				leftBorder := (i - d.Params.BeforeLength) % (i + 1)
				rightBorder := (i + d.Params.AfterLength) % (rowLength - i)
				result = append(result, rows[leftBorder:i]...)
				result = append(result, rows[i])
				result = append(result, rows[i:rightBorder]...)
			}
		}
	}
	return result, nil
}

func (g InvertGreper) Grep(rows []string, pattern string) ([]string, error) {
	var n int
	for i := range rows {
		if rows[i] == pattern {
			n++
			switch i {
			case 0:
				copy(rows, rows[i+1:])
			case len(rows) - 1:
				copy(rows, rows[:i-1])
			default:
				rows = append(rows[i-1:], rows[i+1:]...)
				i--
			}
		}
	}
	updatedRows := make([]string, len(rows)-n)
	copy(updatedRows, rows[:len(rows)-n])

	return updatedRows, nil
}
