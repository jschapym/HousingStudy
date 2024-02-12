package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat"
)

// record represents the structure of our data
type record struct {
	neighborhood string
	crim         float64
	zn           float64
	indus        float64
	chas         float64
	nox          float64
	rooms        float64
	age          float64
	dis          float64
	rad          float64
	tax          float64
	ptratio      float64
	lstat        float64
	mv           float64
}

type Result struct {
	XName    string
	A        float64
	B        float64
	RSquared float64
}

type IterationResult struct {
	Iteration   int
	CrimResult  Result
	RoomsResult Result
}

// checkErr is a simple error check helper function
func checkErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// readData reads the CSV file and returns three slices: crim, rooms, and mv
func readData(csvPath string) ([]float64, []float64, []float64, error) {
	// Open the file
	data, err := os.ReadFile(csvPath)
	if err != nil {
		return nil, nil, nil, err
	}

	// Convert \r line breaks to \n
	dataStr := strings.ReplaceAll(string(data), "\r", "\n")

	// Create a new CSV reader reading from the string
	reader := csv.NewReader(strings.NewReader(dataStr))
	reader.Comma = ','
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true

	// Read the header line
	_, err = reader.Read()
	if err != nil {
		return nil, nil, nil, err
	}

	// Prepare slices for crim, rooms, and mv
	var crim, rooms, mv []float64

	// Read the rest of the data
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, nil, err
		}

		// Parse the fields of the line into the record structure
		rec := record{}
		rec.neighborhood = line[0]
		rec.crim, err = strconv.ParseFloat(line[1], 64)
		checkErr(err)
		rec.zn, err = strconv.ParseFloat(line[2], 64)
		checkErr(err)
		rec.indus, err = strconv.ParseFloat(line[3], 64)
		checkErr(err)
		rec.chas, err = strconv.ParseFloat(line[4], 64)
		checkErr(err)
		rec.nox, err = strconv.ParseFloat(line[5], 64)
		checkErr(err)
		rec.rooms, err = strconv.ParseFloat(line[6], 64)
		checkErr(err)
		rec.age, err = strconv.ParseFloat(line[7], 64)
		checkErr(err)
		rec.dis, err = strconv.ParseFloat(line[8], 64)
		checkErr(err)
		rec.rad, err = strconv.ParseFloat(line[9], 64)
		checkErr(err)
		rec.tax, err = strconv.ParseFloat(line[10], 64)
		checkErr(err)
		rec.ptratio, err = strconv.ParseFloat(line[11], 64)
		checkErr(err)
		rec.lstat, err = strconv.ParseFloat(line[12], 64)
		checkErr(err)
		rec.mv, err = strconv.ParseFloat(line[13], 64)
		checkErr(err)

		// Append the data to the corresponding slices
		crim = append(crim, rec.crim)
		rooms = append(rooms, rec.rooms)
		mv = append(mv, rec.mv)
	}

	return crim, rooms, mv, nil
}

func performRegression(n int, crim, rooms, mv []float64, verbose bool) {
	for i := 0; i < n; i++ {
		// Perform linear regression for each iteration
		iteration := i + 1
		fmt.Printf("Iteration %d:\n", iteration)

		// Calculate regression for Crim vs. Median Value
		a1, b1 := stat.LinearRegression(crim, mv, nil, false)
		rsq1 := stat.RSquared(crim, mv, nil, a1, b1)
		fmt.Printf("Crim vs Median Value: %.2f + %.2f * Crim, R-squared: %.2f\n", a1, b1, rsq1)

		// Calculate regression for Rooms vs. Median Value
		a2, b2 := stat.LinearRegression(rooms, mv, nil, false)
		rsq2 := stat.RSquared(rooms, mv, nil, a2, b2)
		fmt.Printf("Rooms vs Median Value: %.2f + %.2f * Rooms, R-squared: %.2f\n", a2, b2, rsq2)
	}
}

func main() {
	verbose := flag.Bool("verbose", false, "Print the iterations if verbose is set")
	flag.Parse()

	// Set the path to the CSV file
	csvPath := "/Users/jaredchapman/HousingStudyWOCon/boston.csv" // Update this with the actual path

	// Read the data
	crim, rooms, mv, err := readData(csvPath)
	checkErr(err)

	// Perform linear regression calculation 100 times
	performRegression(100, crim, rooms, mv, *verbose)
}
