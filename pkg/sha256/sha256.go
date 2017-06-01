package sha256

import (
	"github.com/minio/sha256-simd"
)

const (
	Size      = sha256.Size
	BlockSize = sha256.BlockSize
)

var (
	New    = sha256.New
	Sum256 = sha256.Sum256
)
