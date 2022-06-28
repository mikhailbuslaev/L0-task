package parser

import (
	"os"
	"strings"
)

type Parser struct {
	Params Params
}

type Params struct {
	IgnoreRegister bool
}

func (p *Parser) Parse(fileName, pattern string) ([]string, string, error) {
	buf, err := os.ReadFile(fileName)
	if err != nil {
		return []string{""}, pattern, err
	}
	row := string(buf)
	if p.Params.IgnoreRegister {
		row = strings.ToLower(row)
		pattern = strings.ToLower(pattern)
	}
	// split buffer by newline
	rows := strings.Split(strings.ReplaceAll(row, "\r\n", "\n"), "\n")
	return rows, pattern, nil
}
