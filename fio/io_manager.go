package fio

const DataFilePerm = 0644

type FileIOType = byte

const (
	//标准文件IO
	StandardFIO FileIOType = iota
	//内存文件映射
	MemoryMap
)

// IOMANAGER　抽象，接口，可以接入不同类型的　IO，这里是标准　IO
type IOmanager interface {
	Read([]byte, int64) (int, error) //从文件的给定位置读取对应的数据
	Write([]byte) (int, error)       //写入字节数组到文件中
	Sync() error                     //持久化数据，存入到磁盘中
	Close() error                    //关闭文件
	Size() (int64, error)            //获取到文件大小
}

// 初始化 IOManager ，目前只支持 标准FileIO
func NewIOManager(fileName string, ioType FileIOType) (IOmanager, error) {
	switch ioType {
	case StandardFIO:
		return NewFileIOManager(fileName)
	case MemoryMap:
		return NewMMapIOManager(fileName)
	default:
		panic("unsupported io type")
	}
}
