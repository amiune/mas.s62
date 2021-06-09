package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// This file is for the mining code.
// Note that "targetBits" for this assignment, at least initially, is 33.
// This could change during the assignment duration!  I will post if it does.

// Mine mines a block by varying the nonce until the hash has targetBits 0s in
// the beginning.  Could take forever if targetBits is too high.
// Modifies a block in place by using a pointer receiver.
func (self *Block) Mine(targetBits uint8) {
	// your mining code here
	// also feel free to get rid of this method entirely if you want to
	// organize things a different way; this is just a suggestion

	var newBlock Block
	newBlock.PrevHash = self.Hash()
	newBlock.Name = "hernan"

	start := time.Now()
	for {
		newBlock.Nonce = strconv.FormatInt(rand.Int63(), 10)
		if CheckWork(newBlock, 25) {
			fmt.Printf("New Block Hash: %s\n", newBlock.Hash().ToString())
			SendBlockToServer(newBlock)
			break
		}
		if time.Since(start).Seconds() > 5 {
			break
		}
	}

}

// CheckWork checks if there's enough work
func CheckWork(bl Block, targetBits uint8) bool {
	h := bl.Hash()

	for i := uint8(0); i < targetBits; i++ {
		// for every bit from the MSB down, check if it's a 1.
		// If it is, stop and fail.
		// Could definitely speed this up by checking bytes at a time.
		// Left as excercise for the reader...?
		if (h[i/8]>>(7-(i%8)))&0x01 == 1 {
			return false
		}
	}
	return true
}
