package data

import (
	"encoding/binary"
	"hash/crc32"
)

type LogRecordType = byte

const (
	LogRecordNormal LogRecordType = iota
	LogRecordDeleted
	LogRecordTxnFinished
)

// 暂存的事务的相关数据
type Transactionrecord struct {
	Record *LogRecord
	Pos    *LogRecordPos
}

// LogRecord的头部大小
const maxLogRecordHeaderSize = binary.MaxVarintLen32*2 + 5

// 写入到数据文件的记录 数据文件的数据是追加写入的，类似日志的格式，所以叫日志
type LogRecord struct {
	Key   []byte
	Value []byte
	Type  LogRecordType
}

// LogRecord的头部信息
type logRecordHeader struct {
	crc        uint32        //crc校验值
	recordType LogRecordType //标识 LogRecord 的类型
	keySize    uint32        //key 的长度
	valueSize  uint32        //value 的长度
}

// LogRecordPos 数据内存索引，描述数据在磁盘上的位置
type LogRecordPos struct {
	Fid    uint32 //文件 id，表示数据存储到哪个文件当中
	Offset int64  //偏移，表示数据存储到文件的哪个位置
	Size   uint32 //标识数据在磁盘上的大小
}

// 对logRecord进行编码，返回字节数组及长度
//
//	+-------------+-------------+-------------+--------------+-------------+--------------+
//	| crc 校验值  |  type 类型   |    key size |   value size |      key    |      value   |
//	+-------------+-------------+-------------+--------------+-------------+--------------+
//	    4字节          1字节        变长（最大5）   变长（最大5）     变长           变长
func EncodeLogRecord(logRecord *LogRecord) ([]byte, int64) {
	//初始化一个header部分发字节数组
	header := make([]byte, maxLogRecordHeaderSize)
	//从第5个字节开始存储type
	header[4] = logRecord.Type
	var index = 5
	//5字节之后，存储的是key和value的长度信息，使用变长类型，节省空间
	index += binary.PutVarint(header[index:], int64(len(logRecord.Key)))
	index += binary.PutVarint(header[index:], int64(len(logRecord.Value)))

	var size = index + len(logRecord.Key) + len(logRecord.Value)
	encBytes := make([]byte, size)
	//将header部分的内容拷贝过来
	copy(encBytes[:index], header[:index])
	//将key和value数据拷贝到字节数组中
	copy(encBytes[index:], logRecord.Key)
	copy(encBytes[index+len(logRecord.Key):], logRecord.Value)
	//对整个LogRecord数据进行CRC校验
	crc := crc32.ChecksumIEEE(encBytes[4:])
	binary.LittleEndian.PutUint32(encBytes[:4], crc)
	return encBytes, int64(size)
}

// 对字节数组中的 Header 信息进行解码
func decodeLogRecordHeader(buf []byte) (*logRecordHeader, int64) {
	if len(buf) <= 4 {
		return nil, 0
	}
	header := &logRecordHeader{
		crc:        binary.LittleEndian.Uint32(buf[:4]),
		recordType: buf[4],
	}

	var index = 5

	//取出实际的keysize
	keySize, n := binary.Varint(buf[index:])
	header.keySize = uint32(keySize)
	index += n

	//取出实际的valuesize
	valueSize, n := binary.Varint(buf[index:])
	header.valueSize = uint32(valueSize)
	index += n

	return header, int64(index)
}

func getLogRecordCRC(lr *LogRecord, header []byte) uint32 {
	if lr == nil {
		return 0
	}

	crc := crc32.ChecksumIEEE(header[:])
	crc = crc32.Update(crc, crc32.IEEETable, lr.Key)
	crc = crc32.Update(crc, crc32.IEEETable, lr.Value)

	return crc
}

// 对位置信息进行编码
func EncodeLogRecordPos(pos *LogRecordPos) []byte {
	buf := make([]byte, binary.MaxVarintLen32*2+binary.MaxVarintLen64)
	var index = 0
	index += binary.PutVarint(buf[index:], int64(pos.Fid))
	index += binary.PutVarint(buf[index:], pos.Offset)
	index += binary.PutVarint(buf[index:], int64(pos.Size))
	return buf[:index]
}

// 对位置信息进行解码
func DecodeLogRecordPos(buf []byte) *LogRecordPos {
	var index = 0
	fileId, n := binary.Varint(buf[index:])
	index += n
	offset, n := binary.Varint(buf[index:])
	index += n
	size, _ := binary.Varint(buf[index:])
	return &LogRecordPos{
		uint32(fileId),
		offset,
		uint32(size),
	}
}
