package main 

//Blockchain : an ordered back-linked list
type Blockchain struct { 
	blocks []*Block
}

// AddBlock : add a new block to blockchain
func (bc *Blockchain) AddBlock(data string) { 
	prevBlock := bc.blocks[len(bc.blocks) - 1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)	
}

// NewGenesisBlock : creates a root node
func NewGenesisBlock() *Block { 
	return NewBlock("Genesis block", []byte{})
}

// NewBlockChain : creates a new blockchain
func NewBlockChain() *Blockchain{ 
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}