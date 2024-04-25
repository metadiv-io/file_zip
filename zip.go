package file_zip

import (
	"archive/zip"
	"bytes"
)

func New() *zipper {
	return &zipper{}
}

type zipper struct {
	names []string
	files [][]byte
}

func (z *zipper) Add(filename string, file []byte) {
	z.names = append(z.names, filename)
	z.files = append(z.files, file)
}

func (z *zipper) Zip() []byte {
	buf := new(bytes.Buffer)

	zipWriter := zip.NewWriter(buf)

	for i := range z.names {
		zipFile, err := zipWriter.Create(z.names[i])
		if err != nil {
			panic(err)
		}

		_, err = zipFile.Write(z.files[i])
		if err != nil {
			panic(err)
		}
	}

	err := zipWriter.Close()
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}
