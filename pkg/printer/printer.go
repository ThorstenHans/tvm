package printer

import (
	"fmt"

	"github.com/fatih/color"
)

type Printer interface {
	Error(err error, message string)
	Errorf(err error, format string, a ...any)
	Hint(message string)
	Hintf(format string, a ...any)
	Success(msg string)
	Successf(format string, a ...any)
}

type defaultPrinter struct {
	verbose bool
}

func (p defaultPrinter) Errorf(err error, format string, a ...any) {
	p.Error(err, fmt.Sprintf(format, a...))
}

func (p defaultPrinter) Error(err error, message string) {
	color.Red(message)

	if p.verbose {
		color.Red(fmt.Sprintf("Error: %s", err))
	}
}

func (p defaultPrinter) Hintf(format string, a ...any) {
	p.Hint(fmt.Sprintf(format, a...))
}

func (p defaultPrinter) Hint(msg string) {
	color.Yellow(msg)
}

func (p defaultPrinter) Successf(format string, a ...any) {
	p.Success(fmt.Sprintf(format, a...))
}

func (p defaultPrinter) Success(msg string) {
	color.Green(msg)
}

func NewPrinter(verbose bool) Printer {
	return &defaultPrinter{verbose: verbose}
}
