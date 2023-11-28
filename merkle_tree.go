package main

import (
	"crypto/sha256"
)

type MerkleTree struct {
	RootNode *MerkleNode
}
type MerkleNode struct {
	LeftNode  *MerkleNode
	RightNode *MerkleNode
	Data      []byte
}

func NewMerkleNode(leftNodePointer *MerkleNode, rightNodePointer *MerkleNode, data []byte) *MerkleNode {
	merkleNode := MerkleNode{}
	//Neu la node la
	if leftNodePointer == nil && rightNodePointer == nil {
		hash := sha256.Sum256(data)
		merkleNode.Data = hash[:]
	} else {
		var prevHash []byte
		//Sap xep
		if string(leftNodePointer.Data) < string(rightNodePointer.Data) {
			prevHash = append(leftNodePointer.Data, rightNodePointer.Data...)
		} else {
			prevHash = append(rightNodePointer.Data, leftNodePointer.Data...)
		}
		hash := sha256.Sum256(prevHash)
		merkleNode.Data = hash[:]
	}
	merkleNode.LeftNode = leftNodePointer
	merkleNode.RightNode = rightNodePointer
	return &merkleNode
}

func NewMerkleTree(datas [][]byte) *MerkleTree {
	merkleTree := &MerkleTree{}
	var nodes []MerkleNode

	n := len(datas)
	//Sao chep tu node cuoi cung
	tempNode := datas[len(datas)-1]
	//Neu transaction khong la boi cua 2 thi them mot node la ban sao cua tracsaction cuoi cung
	for (n & (n - 1)) != 0 {
		datas = append(datas, tempNode)
		n = len(datas)
	}

	//Chuyen tat ca du lieu thanh node la va them vao slice nodes
	for _, data := range datas {
		node := NewMerkleNode(nil, nil, data)
		nodes = append(nodes, *node)
	}

	//Xay dung cac node con lai (khong phai node la)
	for len(nodes) > 1 {
		var newLevel []MerkleNode

		for j := 0; j < len(nodes); j += 2 {
			newNode := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *newNode)
		}

		nodes = newLevel
	}

	//Node con lai duy nhat nay chinh la node goc
	merkleTree.RootNode = &nodes[0]

	return merkleTree
}

// Xac thuc giao dich
func (m *MerkleTree) Verify(data []byte) bool {
	var path [][]byte
	//Lay Merkle Path
	bool := m.MerklePath(data, &path)
	if !bool {
		return false
	}
	hashSlice := sha256.Sum256(data)
	hash := hashSlice[:]
	for _, value := range path {
		var prevHash []byte
		//Sap xep
		if string(hash) < string(value) {
			prevHash = append(hash, value...)
		} else {
			prevHash = append(value, hash...)
		}
		hashArray := sha256.Sum256(prevHash)
		hash = hashArray[:]
	}
	//So sanh voi Merkle Root
	return string(hash) == string(m.RootNode.Data)
}

func (m *MerkleTree) MerklePath(data []byte, path *[][]byte) bool {
	return m.RootNode.merklePath(data, path)
}

func (n *MerkleNode) merklePath(data []byte, path *[][]byte) bool {
	//Check node la
	if n.LeftNode == nil && n.RightNode == nil {
		hash := sha256.Sum256(data)
		if string(n.Data) == string(hash[:]) {
			return true
		}
		return false
	}

	if n.LeftNode != nil {
		var isPathL bool
		isPathL = n.LeftNode.merklePath(data, path)
		if isPathL {
			*path = append(*path, n.RightNode.Data)
			return true
		}
	}

	if n.RightNode != nil {
		var isPathR bool
		isPathR = n.RightNode.merklePath(data, path)
		if isPathR {
			*path = append(*path, n.LeftNode.Data)
			return true
		}
	}

	//Node khong ton tai trong Merkle Tree
	return false
}
