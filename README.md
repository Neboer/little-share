# Little Share

一个现代、高效的局域网内 HTTP 文件分享服务，提供简洁、匿名、快速的文件传输体验。

## 特性

- 🚀 **现代化框架**: 基于最新的 Go 1.21 和 Gin 框架
- ⚙️ **灵活配置**: 完全可配置的服务器参数和存储设置
- 📝 **结构化日志**: 使用标准库 slog 的结构化日志输出
- 🔧 **简易部署**: 单一可执行文件，开箱即用
- 📦 **无限存储**: 文件存储时间为无限（仅受服务器最大存储空间限制）
- 🌐 **灵活地址**: 支持 IPv4、IPv6 和域名绑定
- 📱 **前端支持**: 内置简洁的网页前端，支持文件上传和下载

## 快速开始

### 安装

```bash
# 克隆仓库
git clone https://github.com/Neboer/little-share.git
cd little-share

# 编译
go build -o little-share main.go

# 运行
./little-share
```

### 默认配置

首次运行时，程序会在工作目录下自动生成默认配置文件 `little-share-conf.json`：

```json
{
  "host": "0.0.0.0",
  "port": 19999,
  "frontend_path": "front",
  "file_store_path": "files",
  "max_storage_size": 10737418240
}
```

请根据需要修改配置文件，然后重新启动程序。

### 命令行使用

```bash
# 使用默认配置文件位置（./little-share-conf.json）
./little-share

# 指定自定义配置文件位置
./little-share -config /path/to/custom-config.json
```

## 配置文件说明

### 配置参数

| 参数 | 类型 | 说明 | 默认值 |
|------|------|------|--------|
| `host` | string | 服务器监听地址，支持 IPv4、IPv6 和域名 | `0.0.0.0` |
| `port` | integer | 服务器监听端口 | `19999` |
| `frontend_path` | string | 前端文件所在目录的相对或绝对路径 | `front` |
| `file_store_path` | string | 文件存储目录的相对或绝对路径 | `files` |
| `max_storage_size` | integer | 最大存储空间（字节） | `10737418240` (10 GB) |

### 配置示例

#### 示例 1: 局域网 IPv4 绑定

```json
{
  "host": "192.168.1.100",
  "port": 19999,
  "frontend_path": "front",
  "file_store_path": "files",
  "max_storage_size": 53687091200
}
```

#### 示例 2: IPv6 支持

```json
{
  "host": "::",
  "port": 19999,
  "frontend_path": "front",
  "file_store_path": "files",
  "max_storage_size": 107374182400
}
```

#### 示例 3: 域名绑定

```json
{
  "host": "share.example.com",
  "port": 8080,
  "frontend_path": "front",
  "file_store_path": "files",
  "max_storage_size": 1099511627776
}
```

### 存储大小参考

| 大小 | 字节数 |
|------|--------|
| 1 GB | 1073741824 |
| 5 GB | 5368709120 |
| 10 GB | 10737418240 |
| 50 GB | 53687091200 |
| 100 GB | 107374182400 |
| 1 TB | 1099511627776 |

## API 端点

### 获取可用空间

```bash
GET /maxSpace
```

响应: 可用存储空间（字节数）

```
5368709120
```

### 获取文件列表

```bash
GET /files
```

响应: JSON 数组，包含所有已上传的文件信息

```json
[
  {
    "name": "document.pdf",
    "size": 1048576
  },
  {
    "name": "image.jpg",
    "size": 2097152
  }
]
```

### 上传文件

```bash
POST /upload
Content-Type: multipart/form-data

file=@/path/to/file
```

响应: 上传成功消息

```json
{
  "message": "File document.pdf uploaded successfully"
}
```

### 删除文件

```bash
DELETE /files/:filename
```

响应: 删除成功消息

```
File document.pdf deleted successfully
```

### 下载文件

```bash
GET /files/:filename
```

直接下载指定文件

## 使用示例

### 网页界面

访问 `http://localhost:19999`，使用内置的网页界面上传、下载和管理文件。

### 命令行上传

```bash
# 使用 curl 上传文件
curl -F "file=@/path/to/file.pdf" http://localhost:19999/upload

# 使用 wget 上传文件
wget --post-file=/path/to/file.pdf http://localhost:19999/upload
```

### 命令行列表

```bash
# 获取文件列表
curl http://localhost:19999/files | jq

# 获取可用空间
curl http://localhost:19999/maxSpace
```

### 命令行删除

```bash
# 删除文件
curl -X DELETE http://localhost:19999/files/document.pdf
```

## 存储管理

- **无限保留时间**: 文件上传后将永久存储，直到手动删除或存储空间满
- **空间管理**: 当存储空间不足时，上传将被拒绝，返回 413 错误
- **自动创建**: 文件存储目录会在首次使用时自动创建

## 项目结构

```
little-share/
├── main.go                 # 程序入口
├── config/
│   └── config.go          # 配置管理
├── lib/
│   ├── fileManager.go     # 文件操作
│   └── checkAndDelete.go  # 清理相关说明
├── temp/
│   ├── front/             # 前端文件
│   └── files/             # 文件存储目录
├── go.mod                 # Go 模块定义
├── go.sum                 # 依赖包校验
├── Dockerfile             # Docker 容器配置
└── README.md              # 本文件
```

## Docker 支持

### 构建镜像

```bash
docker build -t little-share .
```

### 运行容器

```bash
# 基础运行
docker run -p 19999:19999 \
  -v /host/data:/data \
  little-share

# 指定配置文件
docker run -p 19999:19999 \
  -v /host/config.json:/app/little-share-conf.json \
  -v /host/data:/data \
  little-share -config /app/little-share-conf.json
```

## 开发

### 环境要求

- Go 1.21 或更高版本
- Git

### 编译

```bash
go build -o little-share main.go
```

### 运行测试

```bash
go test ./...
```

### 代码风格

本项目遵循标准的 Go 代码风格和最佳实践：
- 使用 `go fmt` 进行格式化
- 使用 `goimports` 管理导入
- 使用 `golint` 检查代码质量

## 变更日志

### v2.0.0 (现代化版本)

- ✨ 升级到 Go 1.21，使用最新的 Gin 框架
- 🔧 完全重写配置系统，支持 JSON 配置文件
- 📝 使用结构化日志（slog）替代旧日志系统
- ♻️ 移除 filesMaxKeepTime.json 和 maxSpaceUsage 文件依赖
- 🌐 支持灵活的地址绑定（IPv4、IPv6、域名）
- 📚 编写详细的文档和 API 说明
- 🚀 改进错误处理和日志输出

## 许可证

本项目采用 MIT 许可证。详见 LICENSE 文件。

## 贡献

欢迎提交 Issue 和 Pull Request！

## 相关链接

- [Gin Web Framework](https://gin-gonic.com/)
- [Go 官方文档](https://golang.org/doc/)
- [结构化日志 (slog)](https://pkg.go.dev/log/slog)