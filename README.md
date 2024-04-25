# File Zipper

## Installation

```bash
go get -u github.com/metadiv-io/file_zip
```

## Zip files

```go
zipper := file_zip.New()

zipper.Add("file1.txt", bytes.NewReader([]byte("file1 content")))
zipper.Add("file2.txt", bytes.NewReader([]byte("file2 content")))

zipBytes := zipper.Zip()
```
