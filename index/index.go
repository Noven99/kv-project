package index

import (
	"bytes"
	"github.com/google/btree"
	"kv-project/data"
)

// 抽象接口，用来接入不同的数据结构
type Indexer interface {
	Put(key []byte, pos *data.LogRecordPos) *data.LogRecordPos //存储键值对，键：一个字节切片，值：一个指向 data.LogRecordPos 结构体的指针，表示键对应的日志记录位置。返回值：bool，表示操作是否成功
	Get(key []byte) *data.LogRecordPos                         //根据键获取值
	Delete(key []byte) (*data.LogRecordPos, bool)              //删除键
	Size() int                                                 //索引中的数据量
	Iterator(reverse bool) Iterator                            //索引迭代器
	Close() error                                              //关闭索引
}

// 根据类型初始化索引
type IndexType = int8

const (
	//BTree 索引
	Btree IndexType = iota + 1

	//ART 自适应基数索引
	ART

	//B+树索引
	BPTree
)

func NewIndexer(typ IndexType, dirPath string, sync bool) Indexer {
	switch typ {
	case Btree:
		return NewBTree()
	case ART:
		return NewART()
	case BPTree:
		return NewBPlusTree(dirPath, sync)
	default:
		panic("unsupported index type")
	}
}

// 定义自己的 item
type Item struct {
	Key []byte
	Pos *data.LogRecordPos
}

func (ai *Item) Less(bi btree.Item) bool {
	return bytes.Compare(ai.Key, bi.(*Item).Key) == -1
}

// 通用索引迭代器接口
type Iterator interface {
	Rewind()         //重新回到迭代器的起点，即第一个数据
	Seek(key []byte) //根据传入的key，查找第一个 >= 这个key 的key，从这个地方开始遍历
	Next()           //跳转到下一个key
	Valid() bool     //是否已经遍历完了所有的key，用于退出遍历

	Key() []byte               //当前遍历位置key的数据
	Value() *data.LogRecordPos //当前遍历位置value的数据

	Close() //关闭迭代器，释放相应资源
}
