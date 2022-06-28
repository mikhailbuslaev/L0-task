package printer

import "fmt"

type Printer struct {
	Params Params
}

type Params struct {
	CountOption bool
}

func (p *Printer) Print(rows []string) {
	if p.Params.CountOption {
		fmt.Println(len(rows))
		return
	}
	for _, v := range rows {
		fmt.Println(v)
	}
}
