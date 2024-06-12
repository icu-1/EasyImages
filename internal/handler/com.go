package handler

import (
	"crypto/md5"
	"easy-images/internal/common"
	"encoding/hex"
	"os"
	"path/filepath"
	"time"
)

var (
	storagePath = "images/"
)

func Handler(buffer []byte, filename string) (string, error) {
	hasher := md5.New()
	hasher.Write(buffer)
	fileMD5 := hex.EncodeToString(hasher.Sum(nil))

	timePath := time.Now().Format("2006/01/02")
	filePath := filepath.Join(storagePath, timePath, fileMD5+filepath.Ext(filename))
	if common.FileExists(filePath) {
		return filePath, nil
	}

	if !common.FileExists(storagePath) {
		_ = os.MkdirAll(storagePath, 0744)
	}

	localFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}

	defer localFile.Close()
	_, err = localFile.Write(buffer)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
