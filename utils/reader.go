package utils

import (
	"fmt"
	"os"
	"strings"
)

type Header struct{}

type BMP struct {
	header     Header
	width      int
	height     int
	fileType   string
	size       string
	headerSize string
	DIBheader  string
}

func ReadFile(fileName string) {
	if !strings.HasSuffix(fileName, ".bmp") {
		fmt.Fprintf(os.Stderr, "Error: %s is not bitmap file\n", fileName)
		os.Exit(1)
	}

	fileName = "img/" + fileName

	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
	}

	fmt.Println(string(data))
}
