package easycgo

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func currentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	////
	if i < 0 {
		return "", errors.New(fmt.Sprint(path))
	}
	return string(path[0 : i+1]), nil
}

func md5s(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

func InstallLib(fileName string, buf []byte, isLodeCache bool) (sharedLibrary *EasyCgo) {
	tmpDir := os.TempDir()
	runDir, _ := os.Getwd()
	current, _ := currentPath()

	key := md5s(current)

	mainDllPath := filepath.Join(tmpDir, key, fileName)

	if isLodeCache {

		possible := []string{
			mainDllPath,
			filepath.Join(runDir, fileName),
			filepath.Join(current, fileName),
		}

		var err error
		for i := range possible {
			// log.Println("possible[i]:", possible[i])
			dllabs, _ := filepath.Abs(possible[i])
			if pathExists(dllabs) {
				pathstr := os.Getenv("PATH")
				os.Setenv("PATH", pathstr+";"+filepath.Dir(possible[i]))
				sharedLibrary, err = Load(possible[i])
				if err != nil {
					log.Println("加载失败:", err, possible[i])
					continue
				} else {
					log.Println("加载成功:", dllabs)
				}
				return
			}
		}

	}
	if sharedLibrary == nil {
		os.MkdirAll(filepath.Dir(mainDllPath), os.ModeDir)
		pathstr := os.Getenv("PATH")
		os.Setenv("PATH", pathstr+";"+filepath.Dir(mainDllPath))
		ioutil.WriteFile(mainDllPath, buf, 0644)
	}
	sharedLibrary, _ = Load(mainDllPath)
	return
}

func InstallLibZip(dir string, buf []byte) {

}
