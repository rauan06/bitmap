package utils

import (
	"fmt"
	"os"
	"strings"

	"bitmap/utils/helpers"
)

type Bit struct {
	Red   int
	Green int
	Blue  int
}

type BMP struct {
	FileType         string
	FileSizeInBytes  int
	Offset           int
	DibHeaderSize    int
	Width            int
	Height           int
	PixelSize        int
	ImageSizeInBytes int
	Compression      int
	Data             []byte
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

	header, errHeader := createBMP(data)
	if errHeader != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", errHeader)
		os.Exit(1)
	}

	fmt.Println(header)
}

func createBMP(data []byte) (*BMP, error) {
	header := &BMP{
		FileType:         string(data[:2]),
		FileSizeInBytes:  helpers.BytesToInt(data[2:6]),
		Offset:           helpers.BytesToInt(data[10:14]),
		DibHeaderSize:    helpers.BytesToInt(data[14:18]),
		Width:            helpers.BytesToInt(data[18:22]),
		Height:           helpers.BytesToInt(data[22:26]),
		PixelSize:        helpers.BytesToInt(data[28:30]),
		ImageSizeInBytes: helpers.BytesToInt(data[34:38]),
		Compression:      helpers.BytesToInt(data[30:34]),
		// Data:             data[:],
	}
	err := checkHeader(header)
	if err != nil {
		return nil, err
	}

	fmt.Println(header)
	return header, nil
}

func checkHeader(header *BMP) error {
	if header.FileType != "BM" {
		return fmt.Errorf("unsupported file type: %v", header.FileType)
	}
	if header.PixelSize != 24 {
		return fmt.Errorf("unsupported pixel size: %v", header.PixelSize)
	}
	if header.Compression != 0 {
		return fmt.Errorf("only uncompressed files allowed")
	}
	return nil
}
