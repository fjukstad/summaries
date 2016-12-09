package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	"github.com/fjukstad/kvik/eutils"
	"github.com/fjukstad/kvik/genenames"
)

func main() {
	inputFilename := flag.String("i", "", "Input filename. Text file with one gene name per line.")
	outputFilename := flag.String("o", "", "Output filename.")
	gene := flag.String("symbol", "", "Get summary for a single gene")

	flag.Parse()

	if *gene != "" {
		doc, err := genenames.GetDoc(*gene)
		if err != nil {
			fmt.Println("Could not convert gene symbol to entrez id", err)
			return
		}

		geneSummary, err := eutils.GeneSummary(doc.EntrezId)
		if err != nil {
			fmt.Println(err)
			return
		}
		summary := geneSummary.Summary
		fmt.Println(summary)
		return
	}

	if *inputFilename == "" || *outputFilename == "" {
		fmt.Println("Error: Please specify both input and output filenames.")
		return
	}

	f, err := os.Open(*inputFilename)
	if err != nil {
		fmt.Println("Error: Could not open input file", err)
		return
	}

	var summaries []string
	var genes []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		gene := scanner.Text()
		genes = append(genes, gene)
		doc, err := genenames.GetDoc(gene)
		if err != nil {
			fmt.Println("Warning: Skipping gene"+gene, err)
			summaries = append(summaries, "")
			continue
		}

		geneSummary, err := eutils.GeneSummary(doc.EntrezId)
		if err != nil {
			fmt.Println("Warning: Could not get summary for", gene, "continuing")
			summaries = append(summaries, "")
			continue
		}
		summary := geneSummary.Summary
		summaries = append(summaries, summary)
	}

	f, err = os.OpenFile(*outputFilename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error: Could not open output file", err)
		return
	}

	w := csv.NewWriter(f)

	for i, _ := range genes {
		record := []string{genes[i], summaries[i]}
		err = w.Write(record)
		if err != nil {
			fmt.Println("Error: Could not write record to output file", err)
			return
		}
	}

	w.Flush()
	err = w.Error()
	if err != nil {
		fmt.Println("Error: Error writing output file")
	}

}
