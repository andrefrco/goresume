package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/andrefrco/resume/scripts/pdf"
	"github.com/andrefrco/resume/scripts/serve"
)

func main() {
	mode := flag.String("mode", "", "Mode to run: 'serve' or 'pdf'")
	flag.Parse()

	switch *mode {
	case "serve":
		serve.StartServer()
	case "pdf":
		pdf.GeneratePDF()
	default:
		fmt.Println("Usage: go run scripts/main.go -mode=[serve|pdf]")
		os.Exit(1)
	}
}
