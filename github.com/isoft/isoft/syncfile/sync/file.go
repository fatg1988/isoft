package sync

import (
	"encoding/xml"
	"github.com/astaxie/beego/logs"
	"github.com/isoft/isoft/common/fileutil"
	"io/ioutil"
	"log"
	"os"
)

type SyncFile struct {
	XMLName xml.Name `xml:"syncfile"` // 指定最外层的标签为 syncfile
	Source  string   `xml:"source"`   // 读取source配置项,并将结果保存到Source变量中
	Targets []Target `xml:"target"`   // 读取target标签下的内容,以结构方式
}

type Target struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

func ReadSyncFile(filepath string) (syncFile SyncFile) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	err = xml.Unmarshal(content, &syncFile)
	if err != nil {
		log.Fatal(err)
	}
	return syncFile
}

// 同步所有目录
func StartAllSyncFile(syncFile SyncFile, filterTargetName string) {
	source := syncFile.Source
	targets := syncFile.Targets
	for _, target := range targets {
		if filterTargetName == "" || (filterTargetName != "" && filterTargetName == target.Name) {
			_, err := StartOneSyncFile(source, target.Value)
			if err != nil {
				logs.Error(err.Error())
			}
		}

	}
}

// 开始同步一个目录
func StartOneSyncFile(source, target string) (bool, error) {
	// 判断源文件夹是否存在
	if exist, err := fileutil.PathExists(source); exist == false {
		return false, err
	}
	// 判断目标文件夹是否存在,存在则进行删除
	if exist, err := fileutil.PathExists(target); exist == true {
		if err = os.RemoveAll(target); err != nil {
			return false, err
		} else {
			log.Println("clean dir %s", target)
		}
	}
	// 拷贝文件
	err := fileutil.CopyDir(source, target)
	if err != nil {
		return false, err
	} else {
		log.Println("copy dir %s to %s", source, target)
	}
	return true, nil
}
