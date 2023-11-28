package main

import "fmt"

func main() {
	Block1Datas := []string{"Loc gui 1 BTC cho Alice", "Alice gui 2 BTC cho Bob", "Bob gui 1 BTC cho Eve", "Eve gui 5 BTC cho Eve", "Jame gui 20 BTC cho Alice"}
	Block2Datas := []string{"Loc gui 1 BTC cho Alice", "Loc gui 2 BTC cho Loc", "Bob gui 4 BTC cho Eve", "Bob gui 5 BTC cho Bob"}
	bc := NewBlockchain()
	bc.AddBlock(Block1Datas)
	bc.AddBlock(Block2Datas)
	for _, block := range bc.blocks {
		fmt.Println("PrevBlockHash: ", block.PrevBlockHash)
		fmt.Println("Data: ", block.Transactions[0].Data)
		fmt.Println("Hash: ", block.Hash)
		fmt.Println("")
	}

	//MerkleTree de xac thuc cac Transactions trong Block1
	mt := NewMerkleTree(String2ToByte2(bc.blocks[1].GetTransactionsString()))
	if mt.Verify([]byte("Eve gui 5 BTC cho Eve")) {
		fmt.Println("Tracsactions da dc xac thuc")
	} else {
		fmt.Println("Khong tim thay Transactions")
	}
	if mt.Verify([]byte("Eve gui 6 BTC cho Eve")) {
		fmt.Println("Tracsactions da dc xac thuc")
	} else {
		fmt.Println("Khong tim thay Transactions")
	}
}
