package ospackage

import (
	"fmt"
	"os"
)

func Exist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ExistFile() {
	filePath := "basics.txt"

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Name:", fileInfo.Name())
	fmt.Println("Size:", fileInfo.Size())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Println("Permissions:", fileInfo.Mode().Perm())
	fmt.Println("Mod Time:", fileInfo.ModTime())
}
