package fio

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

// 测试后清理掉原始路径文件
func destroyFile(name string) {
	if err := os.Remove(name); err != nil {
		panic(err)
	}
}

func TestNewFileIOManager(t *testing.T) {
	//这里测试写入的路劲需要自己创建  C:\Users\86182\Desktop\tmp
	path := filepath.Join("C:\\Users\\86182\\Desktop\\tmp", "a.data")
	fio, err := NewFileIOManager(path)

	defer destroyFile(path)
	defer fio.Close() // 先添加关闭文件的延迟调用

	assert.Nil(t, err)
	assert.NotNil(t, fio)
}

func TestFileIO_Write(t *testing.T) {
	path := filepath.Join("C:\\Users\\86182\\Desktop\\tmp", "a.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)
	defer fio.Close() // 先添加关闭文件的延迟调用

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	n, err := fio.Write([]byte("hello world"))
	assert.Equal(t, 11, n)
	assert.Nil(t, err)

	n, err = fio.Write([]byte("bitcask kv"))
	assert.Equal(t, 10, n)
	assert.Nil(t, err)

	n, err = fio.Write([]byte("storage"))
	assert.Equal(t, 7, n)
	assert.Nil(t, err)
}

func TestFileIO_Read(t *testing.T) {
	path := filepath.Join("C:\\Users\\86182\\Desktop\\tmp", "a.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)
	defer fio.Close() // 先添加关闭文件的延迟调用

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	_, err = fio.Write([]byte("key-a"))
	assert.Nil(t, err)

	_, err = fio.Write([]byte("key-b"))
	assert.Nil(t, err)

	b1 := make([]byte, 5)
	n, err := fio.Read(b1, 0)
	assert.Equal(t, 5, n)
	assert.Equal(t, []byte("key-a"), b1)

	b2 := make([]byte, 5)
	n, err = fio.Read(b2, 5)
	assert.Equal(t, 5, n)
	assert.Equal(t, []byte("key-b"), b2)
}

func TestFileIO_Sync(t *testing.T) {
	path := filepath.Join("C:\\Users\\86182\\Desktop\\tmp", "a.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)
	defer fio.Close() // 先添加关闭文件的延迟调用

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	err = fio.Sync()
	assert.Nil(t, err)

}

func TestFileIO_Close(t *testing.T) {
	path := filepath.Join("C:\\Users\\86182\\Desktop\\tmp", "a.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)
	defer fio.Close() // 先添加关闭文件的延迟调用

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	err = fio.Close()
	assert.Nil(t, err)
}
