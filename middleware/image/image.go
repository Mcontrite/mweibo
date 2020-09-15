package image

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"mweibo/conf"
	"mweibo/middleware/file"
	"mweibo/middleware/logging"

	"mweibo/utils"
)

func GetImagePath() string {
	return conf.Imageconfig.ImageSavePath
}

func GetImageFullUrl(name string) string {
	return "/static/upload" + GetImagePath() + name
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.MD5(fileName)
	return fileName + ext
}

func GetImageFullPath() string {
	return conf.Imageconfig.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range conf.Imageconfig.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}
	return size <= conf.Imageconfig.ImageMaxSize
}

func CheckImageSize2(f *multipart.FileHeader) bool {
	return int(f.Size/1024/1024) <= conf.Imageconfig.ImageMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}
	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	return nil
}
