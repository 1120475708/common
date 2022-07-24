package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/1120475708/common/constant"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

func GetOffsetFromHeader(h http.Header) int64 {
	byteRange := h.Get("range")
	if len(byteRange) < 7 {
		return 0
	}
	if byteRange[:6] != "bytes=" {
		return 0
	}
	bytePos := strings.Split(byteRange[6:], "-")
	offset, _ := strconv.ParseInt(bytePos[0], 0, 64)
	return offset
}

func GetHashFromHeader(h http.Header) string {
	digest := h.Get("digest")
	if len(digest) < 9 {
		return ""
	}
	if digest[:8] != "SHA-256=" {
		return ""
	}
	return digest[8:]
}

func GetSizeFromHeader(h http.Header) int64 {
	size, _ := strconv.ParseInt(h.Get("content-length"), 0, 64)
	return size
}

func CalculateHash(r io.Reader) string {
	h := sha256.New()
	io.Copy(h, r)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

var once sync.Once

func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

func mkDir(path string) {
	_exist, _err := HasDir(constant.StoragePath)
	if _err != nil {
		log.Printf("获取文件夹异常 -> %v\n", _err)
		return
	}
	if _exist {
		log.Println("文件夹已存在！")
	} else {
		err := os.Mkdir(constant.StoragePath, os.ModePerm)
		if err != nil {
			log.Printf("创建目录异常 -> %v\n", err)
		} else {
			log.Println("创建成功!")
		}
	}
}

func GetPrefixPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println("Getwd err", err)
	}
	path := pwd + constant.StoragePath
	f := func() { mkDir(path) }

	once.Do(f)
	return constant.StoragePath
}
