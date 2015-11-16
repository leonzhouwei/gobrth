package common

import  (
	"os"
)

const GET = "GET"

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}
