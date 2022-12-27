package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWorld = "*"

var transforms = []string{
	otherWorld + "app",
	otherWorld + "site",
	otherWorld + "time",
	"get" + otherWorld,
	"go" + otherWorld,
	"lets" + otherWorld,
	otherWorld + "hq",
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.ReplaceAll(t, otherWorld, s.Text()))
	}
}
