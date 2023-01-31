package main

import (
	"bytes"
	"flag"
	"os"

	"github.com/OliverosA/indexer/client"
	"github.com/OliverosA/indexer/files"
)

func main() {
	// get the flag from cmd
	inboxFlag := flag.Bool("inbox", false, "Untar inbox directory only")
	flag.Parse()

	var tarPath string

	// getting the file from cmd for untar
	if !*inboxFlag {
		tarPath = os.Args[1]
	} else {

		tarPath = os.Args[2]
	}

	ExecuteApp(tarPath, *inboxFlag)

}

// create func for profiling app
func ExecuteApp(tarPath string, inboxFlag bool) {
	//processing the file
	processedFileData := files.ProcessFiles(tarPath, inboxFlag)
	indexingProcessResponse := client.DocumentsBulkCreation(processedFileData)

	println(bytes.NewBuffer(indexingProcessResponse).String())
}
