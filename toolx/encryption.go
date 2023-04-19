package toolx

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/google/uuid"
)

/** 加密方式 **/

func Md5ByString(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		panic(err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

func Md5ByBytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

func Sha1(data []byte) string {
	_sha1 := sha1.New()
	_, _ = _sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}

func GenUUID() string {
	id := uuid.New()
	return strings.ReplaceAll(id.String(), "-", "")
}
