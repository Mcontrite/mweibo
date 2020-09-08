package file

import (
	"math/rand"
	timepkg "mweibo/middleware/time"
	"os"
	"strconv"
	"time"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreatePath(pathName string) (err error) {
	fileExist, _ := PathExists(pathName)
	if !fileExist {
		err = os.Mkdir(pathName, 0777)
	}
	return
}

func CreatePathInToday(pathName string) (pathInToday string, err error) {
	err = CreatePath(pathName)
	if err != nil {
		return
	}
	today := timepkg.TimeFormat("Ymd")
	pathInToday = pathName + "/" + today
	fileExist, _ := PathExists(pathInToday)
	if !fileExist {
		err = os.Mkdir(pathInToday, 0777)
	}
	return
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func MakeFileName(userid string, fileName string) (newFilename string) {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(10000) + 999
	newFilename = userid + "_" + strconv.Itoa(int(time.Now().UnixNano())) + strconv.Itoa(randNum) + GetExt(fileName)
	return newFilename
}
