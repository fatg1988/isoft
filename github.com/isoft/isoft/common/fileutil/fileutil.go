package fileutil

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

// golang判断文件或文件夹是否存在的方法为使用os.Stat()函数返回的错误值进行判断:
// 如果返回的错误为nil,说明文件或文件夹存在
// 如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
// 如果返回的错误为其它类型,则不确定是否在存在
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

// 拷贝文件夹
func CopyDir(sourceDir string, destDir string) error {
	// 遍历文件夹
	err := filepath.Walk(sourceDir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		// 是目录同时是子目录则进行目录拷贝
		if f.IsDir() {
			if sourceDir != path {
				err = CopyDir(path, filepath.Join(destDir, filepath.Base(path)))
				if err != nil {
					return err
				}
			}
		} else {
			// 是文件则进行文件拷贝
			err = CopyFile(path, filepath.Join(destDir, filepath.Base(path)))
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// 拷贝文件,要拷贝的文件路径,拷贝到哪里
func CopyFile(source, dest string) error {
	if source == "" || dest == "" {
		return errors.New("source file path or dest file path is empty.")
	}
	// 判断源文件是否存在
	if exist, _ := PathExists(source); exist == false {
		return errors.New("source file is exist.")
	}
	// 判断目标文件是否存在,不存在则创建
	if exist, _ := PathExists(dest); exist == false {
		os.MkdirAll(filepath.Dir(dest), os.ModePerm)
	}
	//打开文件资源
	source_open, err := os.Open(source)
	//养成好习惯.操作文件时候记得添加 defer 关闭文件资源代码
	if err != nil {
		return err
	}
	defer source_open.Close()
	//只写模式打开文件,如果文件不存在进行创建 并赋予 644的权限.详情查看linux 权限解释
	dest_open, err := os.Create(dest)
	if err != nil {
		return err
	}
	//养成好习惯.操作文件时候记得添加 defer 关闭文件资源代码
	defer dest_open.Close()
	//进行数据拷贝
	_, copy_err := io.Copy(dest_open, source_open)
	if copy_err != nil {
		return copy_err
	}
	return nil
}
