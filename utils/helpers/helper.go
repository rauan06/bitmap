package helpers

import "fmt"

func BytesToInt(data []byte) int {
	result := int(0)
	for i := 0; i < len(data); i++ {
		result |= int(data[i]) << (8 * i)
	}
	fmt.Println(result)
	return result
}
