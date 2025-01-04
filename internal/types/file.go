package types

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
