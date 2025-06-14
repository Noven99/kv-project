package index

import (
	"github.com/stretchr/testify/assert"
	"kv-project/data"
	"os"
	"path/filepath"
	"testing"
)

func TestBPlusTree_Put(t *testing.T) {
	//测试前先手动删除该路径（windows系统）
	path := filepath.Join("C:\\Users\\86182\\Desktop\\tmp")

	defer func() {
		_ = os.RemoveAll(path)
	}()

	tree := NewBPlusTree(path, false)

	res1 := tree.Put([]byte("aac"), &data.LogRecordPos{Fid: 123, Offset: 999})
	assert.Nil(t, res1)
	res2 := tree.Put([]byte("abc"), &data.LogRecordPos{Fid: 123, Offset: 999})
	assert.Nil(t, res2)
	res3 := tree.Put([]byte("acc"), &data.LogRecordPos{Fid: 123, Offset: 999})
	assert.Nil(t, res3)

	res4 := tree.Put([]byte("acc"), &data.LogRecordPos{Fid: 7744, Offset: 883})
	assert.Equal(t, uint32(123), res4.Fid)
	assert.Equal(t, int64(999), res4.Offset)

}

func TestBPlusTree_Get(t *testing.T) {
	path := filepath.Join("C:\\Users\\86182\\Desktop\\tmp")

	defer func() {
		_ = os.RemoveAll(path)
	}()
	tree := NewBPlusTree(path, false)

	pos := tree.Get([]byte("not exist"))
	assert.Nil(t, pos)

	tree.Put([]byte("aac"), &data.LogRecordPos{Fid: 123, Offset: 999})
	pos1 := tree.Get([]byte("aac"))
	assert.NotNil(t, pos1)

	tree.Put([]byte("aac"), &data.LogRecordPos{Fid: 9884, Offset: 1232})
	pos2 := tree.Get([]byte("aac"))
	assert.NotNil(t, pos2)
}

func TestBPlusTree_Delete(t *testing.T) {
	path := filepath.Join(os.TempDir(), "bptree-delete")
	_ = os.MkdirAll(path, os.ModePerm)
	defer func() {
		_ = os.RemoveAll(path)
	}()
	tree := NewBPlusTree(path, false)

	res1, ok1 := tree.Delete([]byte("not exist"))
	assert.False(t, ok1)
	assert.Nil(t, res1)

	tree.Put([]byte("aac"), &data.LogRecordPos{Fid: 123, Offset: 999})
	res2, ok2 := tree.Delete([]byte("aac"))
	assert.True(t, ok2)
	assert.Equal(t, uint32(123), res2.Fid)
	assert.Equal(t, int64(999), res2.Offset)

	pos1 := tree.Get([]byte("aac"))
	assert.Nil(t, pos1)
}

func TestBPlusTree_Size(t *testing.T) {
	path := filepath.Join("C:\\Users\\86182\\Desktop\\tmp")

	defer func() {
		_ = os.RemoveAll(path)
	}()
	tree := NewBPlusTree(path, false)

	tree.Put([]byte("aac"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("abc"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("acc"), &data.LogRecordPos{Fid: 123, Offset: 999})

	assert.Equal(t, 3, tree.Size())
}

func TestBPlusTree_Iterator(t *testing.T) {
	path := filepath.Join(os.TempDir(), "bptree-iter")
	_ = os.MkdirAll(path, os.ModePerm)
	defer func() {
		_ = os.RemoveAll(path)
	}()
	tree := NewBPlusTree(path, false)

	tree.Put([]byte("caac"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("bbca"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("acce"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("ccec"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("bbba"), &data.LogRecordPos{Fid: 123, Offset: 999})

	iter := tree.Iterator(true)
	for iter.Rewind(); iter.Valid(); iter.Next() {
		assert.NotNil(t, iter.Key())
		assert.NotNil(t, iter.Value())
	}
}
