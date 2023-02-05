package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	msg := `Usage: %s [options...] arg1 arg2
  arg1: time (s)
  arg2: wattage
  
  e.g. $ mw 210 500
`
	cmd := strings.Split(os.Args[0], "/")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, msg, cmd[len(cmd)-1])
		flag.PrintDefaults()
	}

	if len(os.Args) != 3 {
		showHelp()
		return
	}
	var baseTime, baseWattage int

	//for i, v := range os.Args[1:] {
	//	fmt.Printf("args[%d] -> %s\n", i, v)
	//}
	if v, err := strconv.Atoi(os.Args[1]); err != nil {
		panic(err)
	} else {
		baseTime = v
	}
	if v, err := strconv.Atoi(os.Args[2]); err != nil {
		panic(err)
	} else {
		baseWattage = v
	}

	list := [7]int{300, 500, 600, 700, 800, 1000, 1500}
	for _, wat := range list {
		ratio := float64(wat) / float64(baseWattage)
		tm := float64(baseTime) / ratio
		fmt.Printf(" wattage: %d (w), time: %d (s)\n", wat, int(math.Floor(tm)))
	}
}

func showHelp() {
	flag.Usage()
}
