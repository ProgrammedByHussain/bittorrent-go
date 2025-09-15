package main

import (
	"log"
	"os"

	"bittorrent-go/torrentfile"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run . <torrent-file> <output-file>")
	}
	
	inPath := os.Args[1]
	outPath := os.Args[2]

	file, err := os.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tf, err := torrentfile.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}