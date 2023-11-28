package main

type Transaction struct {
	Data []byte
}

// Chuyen tu Kieu []string sang [][]byte
func String2ToByte2(datas []string) [][]byte {
	byte2 := make([][]byte, len(datas))
	for i, value := range datas {
		byte2[i] = []byte(value)
	}
	return byte2
}
