package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/nick11roberts/interlace"
)

func main() {

	if len(os.Args) <= 1 {
		log.Fatalln("\nUsage:\ntestinterlace <latitude> <longitude>")
	}

	latitude, err := strconv.ParseFloat(os.Args[1], 64)
	longitude, err := strconv.ParseFloat(os.Args[2], 64)

	if err != nil {
		log.Fatalln("Error parsing arguments. ")
	}

	result, err := interlace.TwoDimensionalInterlace64(latitude, longitude)

	if err != nil {
		log.Fatalln("Error. Interlace failure")
	}

	fmt.Println(result)

}
