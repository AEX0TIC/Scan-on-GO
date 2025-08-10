package main

import (
	"Scan-on-GO/scanner"
	"flag"
	"fmt"
	"os"
)

func printBanner() {
	fmt.Println(" ##### ###### ###### #     #       ###### #     #        ###### ######")
	fmt.Println("#      #      #    # # #   #       #    # # #   #       #       #    #")
	fmt.Println("###### #      ###### #  #  # ##### #    # #  #  # ##### #    ## #    #")
	fmt.Println("     # #      #    # #   # #       #    # #   # #       #     # #    #")
	fmt.Println("#####  ###### #    # #    ##       ###### #    ##        ###### ######")
	fmt.Println("")

}

func main() {
	printBanner()

	target := flag.String("target", "", "Target IP or domain")
	start := flag.Int("start", 1, "Start port")
	end := flag.Int("end", 1024, "End Port")
	timeout := flag.Int("timeout", 1000, "Timeout in ms")
	save := flag.String("save", "", "Save output to file")

	flag.Parse()

	if *target == "" {
		fmt.Println("Please provide a target using -target")
		os.Exit(1)
	}

	scanner.Scan(*target, *start, *end, *timeout, *save)
}
