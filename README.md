# distributed-file-downloader_60
📦 Distributed File Downloader 60

面向分布式文件系统的多协议下载器（学号后两位：60）

本项目实现一个支持 HTTP / HTTPS / FTP / IPFS / CNS+Path 的分布式文件下载工具，能够在本机 KV 存储系统中记录下载元数据，并提供 RESTful API、目录浏览、压缩下载、域名解析映射等功能。
项目采用 Go 语言 + Gin 框架 + IPFS DAG 思想 设计，结构清晰、扩展性强。

✨ 功能特性
多协议支持：HTTP、HTTPS、FTP、IPFS、CNS+Path
单机 KV 存储：用于保存元数据、CID 映射、域名解析缓存
基于 IPFS 理念的 DAG 构建：支持文件切片与哈希链构建
内容上传与下载 API：包括注册域名、解析 CID、下载文件
目录浏览 & 压缩下载：支持文件夹级打包
可扩展的 Resolver 体系：可挂接链上智能合约作为域名解析

模块化设计：Config / Storage / DAG / Resolver / API
跨平台运行：Linux / macOS / Windows

📁 项目结构
distributed-file-downloader_60
├── cmd/
│   └── server/main.go         # 程序入口
├── internal/
│   ├── api/                   # Gin 接口层
│   ├── config/                # 配置管理
│   ├── dag/                   # DAG 构建与 CID 生成
│   ├── resolver/              # 域名解析
│   ├── storage/               # 本地 KV 存储
│   ├── downloader/            # 多协议下载器
│   └── utils/                 # 通用工具集
├── docs/
│   ├── architecture.md        # 架构设计文档
│   ├── api_doc.md             # API 文档
│   └── design_report.md       # 设计报告（含学号 60）
├── README.md
└── go.mod

🏗️ 架构设计
       ┌───────────────────────┐
       │       HTTP API        │
       │   (Gin Framework)     │
       └───────────┬───────────┘
                   │
        ┌──────────▼───────────┐
        │     Service Layer     │
        │  Upload / Download /  │
        │  Register / Resolve   │
        └──────────┬───────────┘
                   │
  ┌────────────────▼────────────────┐
  │                                 │
  │   Internal Core Architecture    │
  │                                 │
  │  ┌────────┬────────┬────────┐  │
  │  │  DAG   │Storage │Resolver│  │
  │  │ (CID)  │  KV    │ DNS/CNS│  │
  │  └────────┴────────┴────────┘  │
  │                                 │
  └────────────────┬────────────────┘
                   │
        ┌──────────▼──────────┐
        │   Downloader Core    │
        │ HTTP/FTP/IPFS/CNS    │
        └──────────┬──────────┘
                   │
           ┌───────▼───────┐
           │  File System   │
           └───────────────┘

📦 模块说明
1. downloader/

支持以下协议：
http, https，ftp，ipfs://CID，cns://domain/path（基于 Resolver 自动解析 CID）
内部采用分块下载，可扩展为多线程。

2. resolver/
Resolver 体系负责将：
域名 → CID
CID → 文件路径
默认是本地 KV 解析，也可扩展为：以太坊智能合约域名服务，去中心化 DNS，内网自定义 Name Service

3. storage/
提供 KV 存储接口（BoltDB / BadgerDB）。
用途：保存下载记录，保存 CID 映射，保存域名解析缓存

4. dag/
模拟 IPFS Merkle-DAG：，文件分片，生成每片哈希，构建链式 DAG，组合成最终 CID
可后续扩展 chunk-level deduplication。

🚀 安装与运行
1. 克隆项目
git clone https://github.com/nono719/distributed-file-downloader_60.git
cd distributed-file-downloader_60

2. 下载依赖
go mod tidy

3. 运行服务
go run ./cmd/server

默认启动在：
http://localhost:8080

🔗 API 使用说明
1. 上传文件
POST /upload
返回：

{
  "cid": "Qmxxxx"
}

2. 基于域名注册 CID
POST /register


body：

{
  "domain": "abc.test",
  "cid": "Qm123..."
}

3. CNS+path 下载文件
GET /download?path=cns://abc.test/file.txt

4. 以 CID 下载文件
GET /download?path=ipfs://Qmxxx

5. 浏览目录
GET /ls?path=cns://abc.test


## 分支管理规范

本项目采用 “Feature-branch Workflow” 进行版本管理与协作：

- `main` 分支为稳定主分支，仅用于存放已测试通过的代码。  
- 所有新功能、修改或 bug 修复都应创建新的 feature 分支，例如 `feature-xxx`。  
- 在 feature 分支完成开发后，应推送至远端，并在 GitHub 上发 Pull Request (PR)，说明变更内容。  
- PR 经测试／审核无误后合并到 main，合并后可删除 feature 分支。  
- 合并过程中若发生冲突，应手动解决冲突后再合并/推送。  

这样可以保证主干稳定，并且便于多人协作，防止功能混杂、代码混乱。
