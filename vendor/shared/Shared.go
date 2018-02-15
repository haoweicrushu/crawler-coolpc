package shared

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// GetFileNamesFromDir 抓dir底下副檔名是ext的檔名(需加. e.g. .txt)
// 如果有加contain 就抓出檔名包含 contain的檔案名
// 如果ext 帶空值 會抓出資料夾
func GetFileNamesFromDir(dir string, ext string, contain string) ([]string, error) {

	var files []string
	SetDir(&dir)
	filesInDir, err := ioutil.ReadDir(dir)
	if err != nil {
		return files, err
	}
	// 把符合的檔案丟進 files 陣列
	for _, fileInDir := range filesInDir {
		if contain != "" {
			if strings.ToLower(filepath.Ext(fileInDir.Name())) == strings.ToLower(ext) && strings.Contains(strings.ToLower(fileInDir.Name()), strings.ToLower(contain)) {
				files = append(files, dir+fileInDir.Name())
			}
		} else {
			if strings.ToLower(filepath.Ext(fileInDir.Name())) == strings.ToLower(ext) {
				files = append(files,
					dir+fileInDir.Name())
			}
		}
	}
	return files, nil
}

// SetDir 如果傳進來的資料夾 最後不是/ 就加上/
func SetDir(dir *string) {
	if len(*dir) > 0 && (*dir)[len(*dir)-1:] != "/" {
		*dir += "/"
	}
}
