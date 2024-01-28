package main

import (
	"os"
	"strings"
)

func parseFontFile(path string) (map[string]grid, error) {
	fontMap := make(map[string]grid)
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	stringSlice := strings.Split(string(bytes), "\n")
	curChar := ""
	curGrid := make(grid, 0)
	for _, s := range stringSlice {
		if len(s) == 1 {
			curChar = s
		} else if len(s) == 0 {
			fontMap[curChar] = curGrid
			curGrid = make(grid, 0)
		} else {
			curGrid = append(curGrid, s)
		}
	}
	fontMap["CLEAR"] = grid{
		"........", "........", "........", "........",
		"........", "........", "........", "........",
	}
	return fontMap, nil
}
