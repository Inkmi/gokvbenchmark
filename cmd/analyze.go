package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Define the struct
type Benchmark struct {
	Name  string
	NsOp  float64 // Changed to float64
	BOp   int
	Alloc int
}

// AveragedBenchmark struct to hold the averaged data
type AveragedBenchmark struct {
	Name     string
	AvgNsOp  float64
	AvgBOp   float64
	AvgAlloc float64
}

func main() {
	// Open the file
	file, err := os.Open("results.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var benchmarks []Benchmark

	// Use a scanner to read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Benchmark") {
			parts := strings.Fields(line)

			// Extract data and convert to appropriate types
			nsOp, _ := strconv.ParseFloat(strings.TrimSuffix(parts[2], "ns/op"), 64) // Parse as float64
			bOp, _ := strconv.Atoi(parts[4])
			alloc, _ := strconv.Atoi(parts[6])

			// Create a struct and append to the array
			benchmarks = append(benchmarks, Benchmark{
				Name:  parts[0],
				NsOp:  nsOp,
				BOp:   bOp,
				Alloc: alloc,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Calculate the averages
	averagedBenchmarks := averageBenchmarks(benchmarks)

	sort.Slice(averagedBenchmarks, func(i, j int) bool {
		return averagedBenchmarks[i].Name < averagedBenchmarks[j].Name
	})

	// Print the averaged benchmarks (for demonstration)
	fmt.Println("Name;GET AvgNsOp;GET AvgBOp;GET AvgAlloc;PUT AvgNsOp;PUT AvgBOp;PUT AvgAlloc")

	for i := 0; i < len(averagedBenchmarks); i += 2 {
		get := averagedBenchmarks[i]
		put := averagedBenchmarks[i+1]

		stripped := strings.TrimSuffix(get.Name, "_Get") // Extract base name

		benchrun := strings.Split(stripped, "/")[1]
		data := strings.Split(benchrun, "_")
		baseName := data[0]
		benchmarkName := strings.TrimPrefix(benchrun, baseName+"_")

		fmt.Printf("%s;%s;%s;%s;%s;%s;%s;%s\n",
			baseName,
			benchmarkName,
			strings.Replace(fmt.Sprintf("%0.2f", get.AvgNsOp), ".", ",", -1),
			strings.Replace(fmt.Sprintf("%0.2f", get.AvgBOp), ".", ",", -1),
			strings.Replace(fmt.Sprintf("%0.2f", get.AvgAlloc), ".", ",", -1),
			strings.Replace(fmt.Sprintf("%0.2f", put.AvgNsOp), ".", ",", -1),
			strings.Replace(fmt.Sprintf("%0.2f", put.AvgBOp), ".", ",", -1),
			strings.Replace(fmt.Sprintf("%0.2f", put.AvgAlloc), ".", ",", -1),
		)
	}
}

// Function to average benchmarks
func averageBenchmarks(benchmarks []Benchmark) []AveragedBenchmark {
	sumMap := make(map[string]Benchmark)
	countMap := make(map[string]int)

	// Summing up the values
	for _, b := range benchmarks {
		if _, exists := sumMap[b.Name]; exists {
			sumMap[b.Name] = Benchmark{
				Name:  b.Name,
				NsOp:  sumMap[b.Name].NsOp + b.NsOp,
				BOp:   sumMap[b.Name].BOp + b.BOp,
				Alloc: sumMap[b.Name].Alloc + b.Alloc,
			}
			countMap[b.Name]++
		} else {
			sumMap[b.Name] = b
			countMap[b.Name] = 1
		}
	}

	var averagedBenchmarks []AveragedBenchmark

	// Calculating the averages
	for name, sum := range sumMap {
		count := float64(countMap[name])
		averagedBenchmarks = append(averagedBenchmarks, AveragedBenchmark{
			Name:     name,
			AvgNsOp:  sum.NsOp / count,
			AvgBOp:   float64(sum.BOp) / count,
			AvgAlloc: float64(sum.Alloc) / count,
		})
	}

	return averagedBenchmarks
}
