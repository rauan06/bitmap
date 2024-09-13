package utils

import (
	"fmt"
	"os"
	"strings"
)

func FlagHandler() {
	args := os.Args

	if len(args) == 1 {
		printUsageWithExitCode(0)
	}

	command := args[1]

	switch command {
	case "header":
		header(args)
	case "apply":
		// TODO: apply
	default:
		fmt.Fprintf(os.Stderr, "Uknown command: %s\n", command)
	}
}

func header(args []string) {
	if len(args) == 2 {
		printUsageWithExitCode(1)
	}

	fileName := args[2]

	if strings.HasPrefix(fileName, "-h") || strings.HasPrefix(fileName, "--h") {
		fmt.Println("Usage:\n" +
			"  bitmap header <source_file>\n\n" +
			"Description:\n" +
			"  Prints bitmap file header information")
		os.Exit(0)
	}

	ReadFile(fileName)
}

func printUsageWithExitCode(code int) {
	fmt.Println("Usage:\n" +
		"  bitmap <command> [arguments]\n\n" +
		"The commands are:\n" +
		"  header    prints bitmap file header information\n" +
		"  apply     applies processing to the image and saves it to the file")
	os.Exit(code)
}
