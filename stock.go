package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/piquette/finance-go/quote"
	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		logrus.Fatalf("No argument given. Please input one stock symbol. Example: %v IBM", os.Args[0])
	}

	smbl := flag.Args()[0]
	q, _ := quote.Get(smbl)

	fmt.Printf("------- %v -------\n", q.ShortName)
	fmt.Printf("Current Price: $%v\n", q.Ask)
}