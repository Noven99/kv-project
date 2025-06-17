# KV-Project

基于 **Bitcask** 模型，兼容 Redis 数据结构和协议的高性能 KV 存储引擎。

---

## 📌 项目概述

**KV-Project** 是一个高性能的键值存储系统，采用轻量级设计，专注于快速、稳定和高效的数据存储与检索。该项目兼容 Redis 的协议，具备高吞吐量和低资源占用的特点。

---

## ✨ 设计细节

1. **数据模型**：
   - 采用 Key/Value 的数据模型，支持快速存储和检索。
   - 提供高效的键值操作，保证系统的稳定性和高性能。

2. **存储模型**：
   - 基于 Bitcask 存储模型，具备高吞吐量和低读写放大的特性。

3. **数据持久化**：
   - 支持数据持久化，确保数据的可靠性和可恢复性。

4. **索引设计**：
   - 采用多种内存索引结构，实现高效、快速的数据访问。

5. **并发控制**：
   - 使用锁机制，确保数据的一致性和并发访问的正确性。

6. **编程语言**：
   - 项目使用 **Go** 

---

## 🚀 核心功能

1. **高性能存储**：
   - 高速的键值对读写操作，满足大规模数据场景需求。

2. **兼容性**：
   - 支持 Redis 协议，可无缝与现有 Redis 客户端集成。

3. **持久化支持**：
   - 提供可靠的数据持久化机制，防止数据丢失。

4. **低资源消耗**：
   - 采用轻量级设计，适合资源受限的环境。

---

## 🛒 项目结构
<pre>
kv-project/
├── benchmark/                  # 基准测试模块
│   ├── bench_test.go           # 基准测试文件，测试性能（Put、Get、Delete 操作的吞吐量）
├── data/                       # 数据管理模块
│   ├── data_file.go            # 数据文件管理，负责文件的读写和存储
│   ├── data_file_test.go       # 测试数据文件读写的单元测试
│   ├── log_record.go           # 日志记录的编码与解码，实现数据的序列化与反序列化
│   ├── log_record_test.go      # 测试日志记录功能的单元测试
├── examples/                   # 示例模块
│   ├── basic_operation.go      # 示例代码，展示基本的数据库操作（如 Put、Get、Delete）
├── fio/                        # 文件 I/O 抽象模块
│   ├── file_io.go              # 标准文件 I/O 实现
│   ├── file_io_test.go         # 测试标准文件 I/O 的单元测试
│   ├── io_manager.go           # 抽象 I/O 接口，支持不同类型的 I/O 实现
│   ├── mmap.go                 # 内存映射文件 I/O 实现
│   ├── mmap_test.go            # 测试内存映射文件 I/O 的单元测试
├── http/                       # HTTP 服务模块
│   ├── main.go                 # HTTP 服务的入口文件，提供 HTTP API 支持
├── index/                      # 索引模块
│   ├── art.go                  # 自适应基数树（ART）索引实现
│   ├── art_test.go             # 测试 ART 索引的单元测试
│   ├── bptree.go               # B+Tree 索引实现
│   ├── bptree_test.go          # 测试 B+Tree 索引的单元测试
│   ├── btree.go                # BTree 索引实现
│   ├── btree_test.go           # 测试 BTree 索引的单元测试
│   ├── index.go                # 索引接口与工厂方法，支持多种索引类型
├── redis/                      # Redis 数据结构支持模块
│   ├── cmd/                    # Redis 协议支持（命令解析与服务端实现）
│   │   ├── client.go           # Redis 客户端命令解析
│   │   ├── server.go           # Redis 服务端实现
│   ├── generic.go              # 通用 Redis 数据操作
│   ├── meta.go                 # Redis 元数据管理
│   ├── types.go                # 支持的 Redis 数据结构（如 String、Hash、Set、List、ZSet）
│   ├── types_test.go           # 测试 Redis 数据结构的单元测试
├── utils/                      # 工具模块
│   ├── file.go                 # 文件操作工具（如目录大小、磁盘空间计算）
│   ├── file_test.go            # 测试文件工具的单元测试
│   ├── float.go                # 浮点数与字节数组的转换工具
│   ├── rand_kv.go              # 随机键值生成工具，用于性能测试
│   ├── rand_kv_test.go         # 测试随机键值生成工具的单元测试
├── batch.go                    # 批量写入模块，支持事务操作
├── batch_test.go               # 测试批量写入功能的单元测试
├── db.go                       # 数据库核心实现，管理存储引擎的核心逻辑
├── db_test.go                  # 测试数据库核心功能的单元测试
├── errors.go                   # 错误类型定义
├── iterator.go                 # 索引迭代器实现，用于遍历索引
├── iterator_test.go            # 测试索引迭代器功能的单元测试
├── merge.go                    # 数据文件合并逻辑，清理无效数据
├── merge_test.go               # 测试数据文件合并的单元测试
├── options.go                  # 配置模块，定义存储引擎的配置选项
├── go.mod                      # Go 模块定义文件
├── go.sum                      # Go 模块依赖文件
├── 测试问题.txt                 # 测试中遇到的问题记录
</pre>
