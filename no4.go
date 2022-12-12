package main

import (
	"strings"
)

type OutputFormat struct {
	FormatSatu  string
	FormatDua   string
	FormatTiga  string
	FormatEmpat string
}

func generateOutputFormat(input string) *OutputFormat {
	formatSatu := strings.ToUpper(input)
	formatDua := strings.ToLower(input)
	formatTiga := strings.Title(input)
	formatEmpat := strings.Replace(strings.ToUpper(input), " ", "", -1)

	return &OutputFormat{
		FormatSatu:  formatSatu,
		FormatDua:   formatDua,
		FormatTiga:  formatTiga,
		FormatEmpat: formatEmpat,
	}
}
