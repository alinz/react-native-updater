package crypto

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math"
	"os"
)

const filechunk = 8192

//HashFile accepting a filename and convert it into secure hash value
func HashFile(filename string) string {
	file, err := os.Open("utf8.txt")
	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	info, _ := file.Stat()
	filesize := info.Size()

	blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))

	hash := sha1.New()

	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(filesize-int64(i*filechunk))))
		buf := make([]byte, blocksize)

		file.Read(buf)
		io.WriteString(hash, string(buf))
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}
