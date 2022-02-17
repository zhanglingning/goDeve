package main

func main() {
	bc := NewBlockChain()
	bc.AddBlock("知链")
	bc.AddBlock("区块链")
	bc.AddBlock("人才")
	bc.AddBlock("培养")
	bc.AddBlock("摇篮")

	bc.BlockChainList()
}
