package main

import (
	"io"
	"log"
	"os"
)

// CopyFile = Copy file
func CopyFile(src, dst string) {
	from, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}

	defer from.Close()

	to, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	to.Truncate(0)
	to.Seek(0, 0)
	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}

// RemoveFile
func RemoveFile(dst string) error {
	return os.Remove(dst)
}

// CheckIfConfigExist =
func CheckIfConfigExist(name string) bool {
	if _, err := os.Stat("./" + PathName + "/" + name + ".json"); os.IsNotExist(err) {
		return false
	}
	return true
}
