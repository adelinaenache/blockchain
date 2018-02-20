package main

func main() {
	bc := NewBlockChain()
	cli := &CLI{bc}

	defer bc.db.Close()

	cli.Run()
}
