package kv_project

import "os"

type Options struct {
	DirPath             string      //数据库数据目录
	DataFileSize        int64       //数据文件的大小
	SyncWrites          bool        //每次写数据是否持久化
	BytesPerSync        uint        //累计写到多少字节后进行持久化
	IndexType           IndexerType // 索引类型
	MMapAtStartup       bool        //启动时是否使用MMap加载数据
	DataFileMergeRation float32     //数据文件合并的阈值
}

// 批量写配置项
type WriteBatchOptions struct {
	//一个批次当中最大的数据量
	MaxBatchNum uint
	//提交时是否 sync 持久化
	SyncWrites bool
}

// 索引迭代器配置项
type IteratorOptions struct {
	//遍历前缀为指定值的key,默认为空
	Prefix []byte
	//是否反向遍历，默认 false ，是正向
	Reverse bool
}

type IndexerType = int8

const (
	//BTree 索引
	BTree IndexerType = iota + 1

	//ART 自适应基数索引
	ART

	//B+树索引，将索引存储到磁盘上
	BPlusTree
)

var DefaultOptions = Options{
	DirPath:             os.TempDir(),
	DataFileSize:        256 * 1024 * 1024, //256mb
	SyncWrites:          false,
	BytesPerSync:        0,
	IndexType:           BTree,
	MMapAtStartup:       true,
	DataFileMergeRation: 0.5,
}

var DefaultIteratorOptions = IteratorOptions{
	Prefix:  nil,
	Reverse: false,
}

var DefaultWriteBatchOptions = WriteBatchOptions{
	MaxBatchNum: 10000,
	SyncWrites:  true,
}
