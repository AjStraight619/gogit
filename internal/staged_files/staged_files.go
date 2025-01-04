package stagedfiles

import (
	"fmt"
	"os"
)

type FileState int

const (
	Added FileState = iota
	Modified
	Deleted
)

type FileMetadata struct {
	FilePath  string // Full file path
	FileName  string // Name of the file (basename)
	FileState FileState
	Size      int32  // File size in bytes
	Checksum  string // Hex-encoded checksum
}

type StagedFiles struct {
	files *map[string]FileMetadata
}

func (sf *StagedFiles) CreateFiles([]os.FileInfo) {

}

func (sf *StagedFiles) Add(file FileMetadata) {

}

func (sf *StagedFiles) load() {

}

func (sf *StagedFiles) Compare(filesInfo *[]FileMetadata) {

}

func (sf *StagedFiles) Print() {
	for _, file := range *sf.files {
		fmt.Println("Current file:")
		fmt.Printf("File name: %s\n", file.FileName)
		fmt.Printf("File size: %d\n", file.Size)
		fmt.Printf("File checksum: %s\n", file.Checksum)
		fmt.Printf("File state: %v\n", file.FileState)
		fmt.Println("---------------------")
	}
}
