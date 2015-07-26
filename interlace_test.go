package interlace

import (
	"testing"
)

var interlaceTests = []struct {
	latitude  float64
	longitude float64
	interlace int64
}{
	// Both positive, both in furthest decimal places
	{90.000000, 180.000000, 2223198000000000000},
	{89.999999, 179.999999, 2223187999999999999},
	{89.111111, 179.111111, 2223187991111111111},

	// Both positive, latitude decimal in middle case, longitude in furthest case
	{9.000000, 180.000000, 2213198000000000000},
	{8.999999, 179.999999, 2213187999999999999},
	{8.911111, 179.111111, 2213187991111111111},

	// Both positive, latitude decimal in closest case, longitude in furthest case
	{0.900000, 180.000000, 2203198000000000000},
	{0.899999, 179.999999, 2203187999999999909},
	{0.891111, 179.111111, 2203187991111111101},

	// Both positive, latitude in furthest case, longitude in far middle case
	{90.000000, 18.000000, 2222198000000000000},
	{89.999999, 17.999999, 2222187999999999999},
	{89.111111, 17.911111, 2222187991111111111},

	// Both positive, latitude in furthest case, longitude in far middle case
	{90.000000, 1.800000, 2221198000000000000},
	{89.999999, 1.799999, 2221187999999999990},
	{89.111111, 1.791111, 2221187991111111110},

	// Both positive, latitude in furthest case, longitude in far middle case
	{90.000000, 0.180000, 2220198000000000000},
	{89.999999, 0.179999, 2220187999999999090},
	{89.111111, 0.179111, 2220187991111111010},

	// Both negative, both in furthest decimal places
	{-90.000000, -180.000000, 1123198000000000000},
	{-89.999999, -179.999999, 1123187999999999999},
	{-89.111111, -179.111111, 1123187991111111111},

	// Both negative, latitude decimal in middle case, longitude in furthest case
	{-9.000000, -180.000000, 1113198000000000000},
	{-8.999999, -179.999999, 1113187999999999999},
	{-8.911111, -179.111111, 1113187991111111111},

	// Both negative, latitude decimal in closest case, longitude in furthest case
	{-0.900000, -180.000000, 1103198000000000000},
	{-0.899999, -179.999999, 1103187999999999909},
	{-0.891111, -179.111111, 1103187991111111101},

	// Both negative, latitude in furthest case, longitude in far middle case
	{-90.000000, -18.000000, 1122198000000000000},
	{-89.999999, -17.999999, 1122187999999999999},
	{-89.111111, -17.911111, 1122187991111111111},

	// Both negative, latitude in furthest case, longitude in far middle case
	{-90.000000, -1.800000, 1121198000000000000},
	{-89.999999, -1.799999, 1121187999999999990},
	{-89.111111, -1.791111, 1121187991111111110},

	// Both negative, latitude in furthest case, longitude in far middle case
	{-90.000000, -0.180000, 1120198000000000000},
	{-89.999999, -0.179999, 1120187999999999090},
	{-89.111111, -0.179111, 1120187991111111010},

	// Negative latitude, positive longitude, both in furthest decimal places
	{-90.000000, 180.000000, 1223198000000000000},
	{-89.999999, 179.999999, 1223187999999999999},
	{-89.111111, 179.111111, 1223187991111111111},

	// Negative latitude, positive longitude, latitude decimal in middle case, longitude in furthest case
	{-9.000000, 180.000000, 1213198000000000000},
	{-8.999999, 179.999999, 1213187999999999999},
	{-8.911111, 179.111111, 1213187991111111111},

	// Negative latitude, positive longitude, latitude decimal in closest case, longitude in furthest case
	{-0.900000, 180.000000, 1203198000000000000},
	{-0.899999, 179.999999, 1203187999999999909},
	{-0.891111, 179.111111, 1203187991111111101},

	// Negative latitude, positive longitude, latitude in furthest case, longitude in far middle case
	{-90.000000, 18.000000, 1222198000000000000},
	{-89.999999, 17.999999, 1222187999999999999},
	{-89.111111, 17.911111, 1222187991111111111},

	// Negative latitude, positive longitude, latitude in furthest case, longitude in far middle case
	{-90.000000, 1.800000, 1221198000000000000},
	{-89.999999, 1.799999, 1221187999999999990},
	{-89.111111, 1.791111, 1221187991111111110},

	// Negative latitude, positive longitude, latitude in furthest case, longitude in far middle case
	{-90.000000, 0.180000, 1220198000000000000},
	{-89.999999, 0.179999, 1220187999999999090},
	{-89.111111, 0.179111, 1220187991111111010},

	// Positive latitude, negative longitude, both in furthest decimal places
	{90.000000, -180.000000, 2123198000000000000},
	{89.999999, -179.999999, 2123187999999999999},
	{89.111111, -179.111111, 2123187991111111111},

	// Positive latitude, negative longitude, latitude decimal in middle case, longitude in furthest case
	{9.000000, -180.000000, 2113198000000000000},
	{8.999999, -179.999999, 2113187999999999999},
	{8.911111, -179.111111, 2113187991111111111},

	// Positive latitude, negative longitude, latitude decimal in closest case, longitude in furthest case
	{0.900000, -180.000000, 2103198000000000000},
	{0.899999, -179.999999, 2103187999999999909},
	{0.891111, -179.111111, 2103187991111111101},

	// Positive latitude, negative longitude, latitude in furthest case, longitude in far middle case
	{90.000000, -18.000000, 2122198000000000000},
	{89.999999, -17.999999, 2122187999999999999},
	{89.111111, -17.911111, 2122187991111111111},

	// Positive latitude, negative longitude, latitude in furthest case, longitude in far middle case
	{90.000000, -1.800000, 2121198000000000000},
	{89.999999, -1.799999, 2121187999999999990},
	{89.111111, -1.791111, 2121187991111111110},

	// Positive latitude, negative longitude, latitude in furthest case, longitude in far middle case
	{90.000000, -0.180000, 2120198000000000000},
	{89.999999, -0.179999, 2120187999999999090},
	{89.111111, -0.179111, 2120187991111111010},
}

func TestInterlace(t *testing.T) {

	for _, tt := range interlaceTests {
		s, err := TwoDimensionalInterlace64(tt.latitude, tt.longitude)
		if err != nil {
			t.Errorf("TwoDimensionalInterlace64 returned an error. ")
		}
		if s != tt.interlace {
			t.Errorf("TwoDimensionalInterlace64(%v, %v) => %v, want %v", tt.latitude, tt.longitude, s, tt.interlace)
		}
	}

}

func BenchmarkInterlace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoDimensionalInterlace64(-9.9999, 123.5555456)
	}
}
