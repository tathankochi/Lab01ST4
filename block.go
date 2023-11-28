package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	//Thoi gian
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	var transactions []byte
	for i := 0; i < len(b.Transactions); i++ {
		transactions = append(transactions, b.Transactions[i].Data...)
	}
	//Noi lai thanh 1 mang duy nhat
	headers := bytes.Join([][]byte{b.PrevBlockHash, transactions, timestamp}, []byte{})
	//Bam headers
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// Lay cac transaction o dang strin
func (b Block) GetTransactionsString() []string {
	var transactionStrings []string
	for i := 0; i < len(b.Transactions); i++ {
		transactionString := string(b.Transactions[i].Data)
		transactionStrings = append(transactionStrings, transactionString)
	}
	return transactionStrings
}

func NewBlock(datas []string, prevBlockHash []byte) *Block {
	var transactions []*Transaction
	for _, data := range datas {
		transaction := &Transaction{Data: []byte(data)}
		transactions = append(transactions, transaction)
	}

	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock([]string{"Genesis Block"}, []byte{})
}
