# Little Share 开发文档

## 项目架构

Little Share 采用模块化的设计，分为以下几个主要部分：

### 1. `config` 包
负责程序配置的管理，包括：
- 从 JSON 文件读取配置
- 生成默认配置
- 支持命令行参数指定配置文件路径

**关键类型和函数：**
- `Config`: 配置结构体
- `DefaultConfig()`: 返回默认配置
- `Load(configPath string)`: 加载指定路径的配置
- `LoadWithFlag()`: 从命令行参数加载配置

### 2. `lib` 包
包含所有文件操作和业务逻辑：

#### fileManager.go
处理文件上传、下载、删除和列表：
- `StoreToLocal()`: 处理文件上传
- `DeleteOneFile()`: 删除指定文件
- `GetFileList()`: 获取存储目录中的所有文件
- `GetCurrentTotalFileSize()`: 计算当前已用存储空间
- `GetAvailableSpace()`: 获取可用存储空间

#### checkAndDelete.go
提供关于文件生命周期管理的文档说明（现代版本中不再自动删除过期文件）

### 3. `main.go`
应用程序入口，负责：
- 加载配置
- 初始化 Gin 路由器
- 定义 API 端点
- 启动 HTTP 服务器

## 关键改进点

### 从旧版本升级的主要变化

1. **配置管理**
   - 旧：使用 `maxSpaceUsage` 文件和 `filesMaxKeepTime.json` 
   - 新：使用单一的 `little-share-conf.json` JSON 配置文件

2. **文件生命周期**
   - 旧：文件有保留时间限制，自动删除过期文件
   - 新：文件永久保存，直到手动删除或存储空间满

3. **日志系统**
   - 旧：使用 `log` 包，输出不规范
   - 新：使用 `log/slog` 结构化日志

4. **错误处理**
   - 旧：错误信息不一致，使用错误代码
   - 新：使用 Go 惯例返回 `error` 类型

5. **代码风格**
   - 旧：变量名混合中英文，不符合 Go 规范
   - 新：遵循 Go 命名规范（驼峰式，英文）

## API 端点详解

### GET /maxSpace
获取可用存储空间

**请求：**
```bash
GET /maxSpace
```

**响应：**
```
5368709120
```

**实现细节：** 
从 `max_storage_size` 减去当前使用的总大小

---

### GET /files
获取文件列表

**请求：**
```bash
GET /files
```

**响应：**
```json
[
  {
    "name": "file1.pdf",
    "size": 1048576
  }
]
```

**实现细节：**
遍历 `file_store_path` 目录，返回所有文件的信息

---

### POST /upload
上传文件

**请求：**
```bash
POST /upload
Content-Type: multipart/form-data

file=@/path/to/file
```

**响应：**
```json
{
  "message": "File file1.pdf uploaded successfully"
}
```

**错误处理：**
- 400: 请求中没有文件字段
- 413: 文件过大（超过可用空间）
- 500: 服务器错误

**实现细节：**
1. 获取上传的文件
2. 检查文件大小是否超过可用空间
3. 确保存储目录存在
4. 保存文件到磁盘

---

### DELETE /files/:file
删除指定文件

**请求：**
```bash
DELETE /files/file1.pdf
```

**响应：**
```
File file1.pdf deleted successfully
```

**错误处理：**
- 404: 文件不存在
- 500: 删除失败

**实现细节：**
检查文件存在性，然后从磁盘删除

---

### GET /files/:file
下载文件

**请求：**
```bash
GET /files/file1.pdf
```

**响应：**
文件内容二进制数据

**实现细节：**
通过 Gin 的静态文件中间件处理

## 添加新功能的指南

### 添加新的 API 端点

1. 在 `main.go` 的 `main()` 函数中添加新的路由：

```go
r.GET("/new-endpoint", func(c *gin.Context) {
    // Your handler logic here
    c.JSON(200, gin.H{"status": "ok"})
})
```

2. 如果需要文件操作，在 `lib/fileManager.go` 中添加新函数
3. 添加适当的日志记录：

```go
slog.Info("Operation completed", "detail", "value")
slog.Warn("Warning message", "key", "value")
slog.Error("Error occurred", "error", err)
```

### 添加新的配置项

1. 在 `config/config.go` 的 `Config` 结构体中添加新字段
2. 在 `DefaultConfig()` 中设置默认值
3. 更新 JSON 标签以匹配配置文件中的键名

## 运行和测试

### 构建

```bash
make build
# 或
go build -o little-share main.go
```

### 运行

```bash
make run
# 或
./little-share -config ./little-share-conf.json
```

### 代码检查

```bash
make fmt     # 格式化代码
make lint    # 检查代码质量
go vet ./... # 检查潜在问题
```

## 依赖项

- `github.com/gin-gonic/gin`: Web 框架
- `github.com/gin-contrib/static`: 静态文件服务中间件

标准库依赖：
- `encoding/json`: JSON 序列化/反序列化
- `flag`: 命令行参数解析
- `log/slog`: 结构化日志
- `os`: 文件系统操作
- `path/filepath`: 路径处理
- `fmt`: 格式化输出

## 常见问题

### Q: 如何修改默认监听地址？
A: 编辑 `little-share-conf.json`，修改 `host` 和 `port` 字段

### Q: 文件会自动删除吗？
A: 不会。在现代版本中，文件永久保存，直到手动删除或存储空间满

### Q: 如何支持 HTTPS？
A: 这需要在反向代理（如 Nginx）或负载均衡器中配置，Little Share 本身仅支持 HTTP

### Q: 如何限制上传文件大小？
A: 设置配置文件中的 `max_storage_size` 参数

## 代码风格指南

遵循标准 Go 代码风格：
- 函数和变量使用驼峰式（camelCase）命名
- 导出的函数首字母大写
- 导入语句按类别分组（标准库 → 外部库）
- 每行代码不超过 100 个字符（推荐）
- 使用 `fmt` 包中的 `Sprintf` 而不是字符串连接

## 贡献指南

1. Fork 仓库
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 打开 Pull Request

请确保：
- 代码遵循项目风格
- 添加适当的日志记录
- 更新文档
- 测试新功能
