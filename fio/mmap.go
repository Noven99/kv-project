package fio

import (
	"golang.org/x/exp/mmap"
	"os"
)

// IO，内存文件映射
type MMap struct {
	readerAT *mmap.ReaderAt
}

// 初始化MMap IO
func NewMMapIOManager(fileName string) (*MMap, error) {
	/*_, err := os.OpenFile(fileName, os.O_CREATE, DataFilePerm)
	if err != nil {
		return nil, err
	}
	readerAT, err := mmap.Open(fileName)
	if err != nil {
		return nil, err
	}
	return &MMap{readerAT}, nil*/

	// 修正：显式关闭文件句柄
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, DataFilePerm)
	if err != nil {
		return nil, err
	}
	_ = f.Close()  // 关键修复：立即关闭临时创建的文件句柄

	readerAT, err := mmap.Open(fileName)
	if err != nil {
		return nil, err
	}
	return &MMap{readerAT}, nil
}

func (mmap *MMap) Read(b []byte, offset int64) (int, error) {
	return mmap.readerAT.ReadAt(b, offset)
}

func (mmap *MMap) Write([]byte) (int, error) {
	panic("not implemented")
}

func (mmap *MMap) Sync() error {
	panic("not implemented")
}

func (mmap *MMap) Close() error {
	return mmap.readerAT.Close()
}

func (mmap *MMap) Size() (int64, error) {
	return int64(mmap.readerAT.Len()), nil
}
