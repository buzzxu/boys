package bytess

import (
	"github.com/buzzxu/boys/types"
	"net/http"
)

//是否是图片
func IsImage(buff []byte) (bool, string) {
	contentType := http.DetectContentType(buff)
	switch contentType {
	case "image/jpeg", "image/jpg", "image/webp", "image/gif", "image/png":
		return true, contentType
	default:
		return false, ""
	}
}

//图片base64前缀
func PrefixImageBase64(data *[]byte) (string, error) {
	if flag, contentType := IsImage(*data); flag {
		return "data:" + contentType + ";base64,", nil
	}
	return "", types.NewError(400, "非图片无法区分图片类型")
}
