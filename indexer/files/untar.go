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
	println("Triying to open file: ", source)
	tarFile, tarFileOpenErr := os.Open(source)

	if tarFileOpenErr != nil {
		return tarFileOpenErr
	}

	println("Triying to decompress with gz...")
	decompressedFile, decompressionErr := gzip.NewReader(tarFile)

	if decompressionErr != nil {
		return decompressionErr
	}

	defer decompressedFile.Close()

	println("Triying to decompress with tar...")
	tarBall := tar.NewReader(decompressedFile)

	println("Decompressing Files, This could take a while...")
	for {
		header, nextHeaderErr := tarBall.Next()

		switch {
		case nextHeaderErr == io.EOF: 
			return nil
		case nextHeaderErr != nil: 
			return nextHeaderErr
		case header == nil: 
			continue

		}

		if inboxFlag {
			
			inboxDirectory := false
			headerCheck := strings.Split(header.Name, "/")
			if len(headerCheck) >= 4 {
				for _, name := range headerCheck {
					if name == "inbox" {
						inboxDirectory = true
					}
				}
			}

			if !inboxDirectory {
				continue
			}
		}

		target := filepath.Join(".", header.Name)

		switch header.Typeflag {
		case tar.TypeDir: 
			if _, err := os.Stat(target); err != nil { 

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

		case tar.TypeReg: 
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(0750))
			if err != nil {
				return err
			}

			// using this causes the files to be closed until they are all complete
			//defer f.Close()

			if _, err := io.Copy(f, tarBall); err != nil {
				return err
			}

			f.Close()
		}
	}
}
