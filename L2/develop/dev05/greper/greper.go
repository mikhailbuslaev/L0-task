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
// 3 different grep algorithms
type FixedGreper struct{} // fixed greper compare all row with pattern

type DefaultGreper struct { // default greper returns pieces of row, were we find pattern
	Params Params
}

type InvertGreper struct{} // invert greper returns not matched rows

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
		rowArr := strings.Split(rows[i], "")// here we begin comparing by loop over letters
		rowLength := len(rowArr)// length of row
		patternArr := (strings.Split(pattern, ""))// split pattern too
		patternLength := len(patternArr)// length of pattern
		for j := 0; j < rowLength-patternLength+1; j++ {// here we go loop over row and trying to catch pattern
			word := strings.Join(rowArr[j:j+patternLength], "")
			if word == pattern {

				if d.Params.LineNum {
					result = append(result, fmt.Sprintf("%d", i))
					continue
				}

				if d.Params.BeforeLength > i || d.Params.AfterLength > len(rows[i])-i {// this is 2 params from app flags
					return []string{""}, fmt.Errorf("afterLength or beforeLength too far")
					// before length is length before start of matched word
					// after length is length after end of matched word
				}

				leftBorder := (i - d.Params.BeforeLength) % (i + 1)// left border is start, from we return row
				rightBorder := (i + d.Params.AfterLength) % (rowLength - i)// right border is end
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
