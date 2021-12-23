package csv

import (
	"encoding/csv"
	"os"
)

type File struct {
	Path    string
	File    *os.File
	streams [][]string
}

func NewFile() *File {
	return &File{}
}

func (f *File) CreateRow(strings []string) {
	f.streams = append(f.streams, strings)
}

func (f *File) SaveCsv(filename string) (err error) {
	f.Path = filename
	f.File, err = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.File.Close()
	// 写入UTF-8 BOM，防止中文乱码
	f.File.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f.File)
	return w.WriteAll(f.streams)
}
