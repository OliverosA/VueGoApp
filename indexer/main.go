package main

import (
	"bytes"
	"flag"
	"os"

	"github.com/OliverosA/indexer/client"
	"github.com/OliverosA/indexer/files"
)

func main() {
	inboxFlag := flag.Bool("inbox", false, "Untar inbox directory only")
	flag.Parse()

	var tarPath string

	if !*inboxFlag {
		tarPath = os.Args[1]
	} else {

		tarPath = os.Args[2]
	}

	ExecuteApp(tarPath, *inboxFlag)

}

func ExecuteApp(tarPath string, inboxFlag bool) {
	processedFileData := files.ProcessFiles(tarPath, inboxFlag)
	indexingProcessResponse := client.DocumentsBulkCreation(processedFileData)

	println(bytes.NewBuffer(indexingProcessResponse).String())
}
