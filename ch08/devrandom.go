package main

import (
	"encoding/binary"
	"log"
	"os"
)

func main() {
	f, err := os.Open("/dev/random")
	defer f.Close()

	if err != nil {
		log.Panicln(err)
	}

	var seed int64
	binary.Read(f, binary.LittleEndian, &seed)
	log.Println("Seed:", seed)
}
