package files

import (
	"archive/tar"   // decompress tar package
	"compress/gzip" // decompress gzip package
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Untar(decompress) the file
func Untar(source string, inboxFlag bool) error {
	// open the file
	println("Triying to open file: ", source)
	tarFile, tarFileOpenErr := os.Open(source)

	if tarFileOpenErr != nil {
		return tarFileOpenErr
	}

	// if success start decompressing the file with gz
	println("Triying to decompress with gz...")
	decompressedFile, decompressionErr := gzip.NewReader(tarFile)

	if decompressionErr != nil {
		return decompressionErr
	}

	// when all ends close the file
	defer decompressedFile.Close()

	// start decompressing with tar
	println("Triying to decompress with tar...")
	tarBall := tar.NewReader(decompressedFile)

	println("Decompressing Files, This could take a while...")
	for {
		header, nextHeaderErr := tarBall.Next()

		// FOSUM-D DOESNT HAVE INBOX
		switch {
		case nextHeaderErr == io.EOF: // IF END OF FILE
			return nil
		case nextHeaderErr != nil: // IF ERROR != NULL
			return nextHeaderErr
		case header == nil: // iF THE HEADER IS NULL, SKIP IT
			continue

		}

		// if inbox flag is true only inbox directory would be untared
		if inboxFlag {
			// cheking for inbox email files only
			inboxDirectory := false
			headerCheck := strings.Split(header.Name, "/")
			if len(headerCheck) >= 4 {
				for _, name := range headerCheck {
					if name == "inbox" { // inside inbox directory
						inboxDirectory = true
					}
				}
			}

			if !inboxDirectory {
				continue
			}
		}

		// making the target location with header name from compresed file
		// example: enron_mail_20110402\maildir\haedicke-m\inbox\339.
		target := filepath.Join(".", header.Name)

		// check the type file
		switch header.Typeflag {
		case tar.TypeDir: // type file = directory
			if _, err := os.Stat(target); err != nil { // if theres no error path file

				/*os.MkdirAll(path, FileMode = 0750 (permission))
				*
				*	Chmod 0750 (chmod a+rwx,g-w,o-rwx,ug-s,-t) sets permissions so that,
				*	 (U)ser / owner can read, can write and can execute.
				*	(G)roup can read, can't write and can execute.
				*	(O)thers can't read, can't write and can't execute.
				*
				*	the permission bitx was obtained from go documentation
				*	https://pkg.go.dev/os#MkdirAll
				 */

				if err := os.MkdirAll(target, 0750); err != nil {
					log.Println("error aca")
					return err
				}
			}

		case tar.TypeReg: // if es a file create it
			// file opened with create and read/write flags and permissioo 0750
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(0750))
			if err != nil {
				return err
			}

			// using this causes the files to be closed until they are all complete
			//defer f.Close()

			// copying content from the file
			if _, err := io.Copy(f, tarBall); err != nil {
				return err
			}

			// manually closing the file
			f.Close()
		}
	}
}
