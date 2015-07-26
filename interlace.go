// Package interlace is used to compute a geographic block address given latitude and longitude coordinates.
// Currently only available for 64 bit architectures.
// A useful addressing format for latitude and longitude range searching.
// Invented by yours truly, Nicholas (Nick) Roberts [https://github.com/nick11roberts].
// Created on Friday July 10 2015.
package interlace

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const debug = false

// TwoDimensionalInterlace64 takes two float64 arguments (latitude, longitude) and returns an int64.
// The return value is the address to the corresponding geographic block according to this algorithm.
func TwoDimensionalInterlace64(latitude float64, longitude float64) (int64, error) {
	primerValues := interlacePrimer{}
	var failure error //Use for possible errors

	//Make sure input is valid
	if math.Abs(latitude) > 90.0 || math.Abs(longitude) > 180.0 {
		log.Fatalln("Input out of range. Aborting. ")
	}

	//Set negativities, then take the absolute values of latitude and longitude, finally reset latitude and longitude
	if latitude >= 0 {
		primerValues.interlaceString += strconv.Itoa(2)
	} else {
		primerValues.interlaceString += strconv.Itoa(1)
	}
	if longitude >= 0 {
		primerValues.interlaceString += strconv.Itoa(2)
	} else {
		primerValues.interlaceString += strconv.Itoa(1)
	}
	latitude = math.Abs(latitude)
	longitude = math.Abs(longitude)

	//Set decimal positions, remove decimals, update primerValues
	if latitude < 1.0 {
		primerValues.interlaceString += strconv.Itoa(0)
	} else if latitude < 10.0 {
		primerValues.interlaceString += strconv.Itoa(1)
	} else if latitude < 100.0 {
		primerValues.interlaceString += strconv.Itoa(2)
	} else if latitude < 1000.0 {
		primerValues.interlaceString += strconv.Itoa(3)
	} else {
		log.Fatalln("This should never happen (failure at latitude decimal). Aborting. ")
	}
	if longitude < 1.0 {
		primerValues.interlaceString += strconv.Itoa(0)
	} else if longitude < 10.0 {
		primerValues.interlaceString += strconv.Itoa(1)
	} else if longitude < 100.0 {
		primerValues.interlaceString += strconv.Itoa(2)
	} else if longitude < 1000.0 {
		primerValues.interlaceString += strconv.Itoa(3)
	} else {
		log.Fatalln("This should never happen (failure at longitude decimal). Aborting. ")
	}
	latitudeString := fmt.Sprintf("%f", latitude)
	latitudeString = strings.Replace(latitudeString, ".", "", 1)
	latitudeString = strings.TrimPrefix(latitudeString, "0")
	for len(latitudeString) < 7 {
		latitudeString += "0"
	}
	if debug {
		fmt.Println("Latitude integer: " + latitudeString)
	}
	longitudeString := fmt.Sprintf("%f", longitude)
	longitudeString = strings.Replace(longitudeString, ".", "", 1)
	longitudeString = strings.TrimPrefix(longitudeString, "0")
	for len(longitudeString) < 8 {
		longitudeString += "0"
	}
	if debug {
		fmt.Println("Longitude integer: " + longitudeString)
	}

	//Perform the body of the interlace algorithm
	i := 0 //iterator
	j := 0 //index
	for len(primerValues.interlaceString) < 19 {
		if debug {
			fmt.Println("Interlace at i=", i, ": "+primerValues.interlaceString)
		}
		if len(primerValues.interlaceString) > 19 {
			log.Fatalln("Interlace intermediate string max length exceeded. Aborting. ")
		}
		if i%2 != 0 {
			primerValues.interlaceString += string(latitudeString[j])
			if debug {
				fmt.Println("j at even (lat):", j)
			}
			j++
		} else if i%2 == 0 {
			primerValues.interlaceString += string(longitudeString[j])
			if debug {
				fmt.Println("j at odd (lon):", j)
			}
		} else {
			log.Fatalln("This should never happen (failure at interlace body). Aborting. ")
		}
		i++
	}

	//Set interlaceFinal
	primerValues.interlaceFinal, failure = strconv.ParseInt(primerValues.interlaceString, 10, 64)

	if debug {
		fmt.Println("RESULT:", primerValues)
	}
	return primerValues.interlaceFinal, failure
}

type interlacePrimer struct {
	interlaceString string
	interlaceFinal  int64 //Array containing the digits from interlace header and interlace body
}
