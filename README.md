# wxbot

基于 Go 语言开发的微信个人号机器人框架，支持自动回复、消息转发、群管理等功能，通过插件化架构实现灵活扩展。

## 功能特性

- **插件化架构**：通过 `plugins.yaml` 管理插件，支持热插拔式的功能扩展
- **多框架适配**：支持 Dean (DaenWxHook) 和 VLW 两种微信 Hook 框架
- **丰富的消息类型**：支持文本、图片、视频、文件、音乐、小程序、名片、表情等多种消息发送
- **灵活的匹配规则**：支持前缀、后缀、命令、正则、关键词、全文匹配等多种触发方式
- **群管理与权限控制**：支持插件按群启停、用户封禁、全局屏蔽、管理员指令等
- **事件通道**：支持异步事件监听（Next/Repeat），实现多轮对话等场景
- **Web 管理界面**：内置 Vue.js 单页应用，提供可视化的插件管理
- **数据持久化**：基于 SQLite (纯 Go 实现，无需 CGO) 存储插件状态和消息记录
- **本地图片服务**：将本地图片映射为 HTTP 静态资源，支持在微信中展示

## 架构概览

```
[微信客户端]  <-->  [Hook 框架 (Dean/VLW)]  <-->  [wxbot 核心引擎]  <-->  [插件系统]
                                                      |
                                                 [Web 管理界面]
```

核心分层：

| 层级 | 说明 |
|------|------|
| `framework/` | 微信 Hook 框架适配层，对接 Dean 或 VLW 的 HTTP API |
| `engine/robot/` | 核心引擎：事件循环、匹配器管道、HTTP 服务、消息上下文 |
| `engine/control/` | 控制层：插件注册、启停管理、用户封禁、消息记录 |
| `engine/plugins/` | 插件导入中心，由 `plugins.yaml` 自动生成 `plugins.go` |
| `plugins/` | 具体插件实现，如天气查询、ChatGPT 对话、新闻推送等 |
| `web/` | 嵌入式 Web 管理界面（Vue.js SPA） |

### 事件处理流程

```
HTTP 回调 → 解析事件 → 缓冲队列 → 按优先级匹配 → 前置过滤 → 规则匹配 → 中间过滤 → 执行处理 → 后置处理
```

1. 微信 Hook 框架通过 HTTP POST 将事件发送到 `/wxbot/callback`
2. 框架适配器解析原始 JSON 为 `Event` 结构体
3. 事件进入缓冲队列，按配置的延迟出队（默认 1 秒）
4. 依次匹配已注册的规则（按优先级排序），命中后执行对应的 Handler
5. 支持 `Block`（停止匹配链）和 `Break`（退出匹配循环）控制

## 快速开始

### 环境要求

- Go 1.20+
- Windows 系统（用于运行微信客户端及 Hook 框架）
- 对应的微信 Hook 框架（Dean 或 VLW）

### 编译运行

```bash
# 克隆项目
git clone <repo-url> wxbot
cd wxbot

# 安装依赖
go mod tidy

# 生成插件导入代码（如已存在 plugins.yaml）
cd engine/plugins && go generate

# 编译
cd ../.. && go build -o wxbot.exe

# 运行（会自动生成 config.yaml 模板）
./wxbot.exe
```

### 配置

首次运行会自动生成 `config.yaml` 模板，根据实际情况修改：

```yaml
botWxId: "wxid_xxxxxx"          # 机器人微信 ID
botNickname: "小明同学"           # 机器人昵称
superUsers:                      # 管理员微信 ID 列表
  - "wxid_xxxxxx"
commandPrefix: "/"               # 管理员指令前缀
wakeUpRequire: ""                # 群聊唤醒方式，"at" 表示需要 @机器人
serverPort: 8867                 # HTTP 服务端口
serverAddress: "http://127.0.0.1:8867"  # 本地图片静态服务地址
framework:
  name: "Dean"                   # 框架类型：Dean / VLW
  apiUrl: "http://127.0.0.1:8866"  # Hook 框架 API 地址
  apiToken: ""                   # API Token（VLW 需要）
```

### 插件管理

插件通过 `plugins.yaml` 声明，运行 `go generate` 自动生成导入代码：

```yaml
# plugins.yaml 示例
plugins:
  - weather
  - chatgpt
  - zaobao
  - manager
```

## 内置插件

| 插件名称 | 目录 | 功能描述 |
|----------|------|----------|
| chatgpt | `plugins/chatgpt/` | AI 对话，支持上下文记忆和图片生成 |
| manager | `plugins/manager/` | 群管理，支持定时任务和提醒 |
| weather | `plugins/weather/` | 天气查询（和风天气 API） |
| zaobao | `plugins/zaobao/` | 每日 60 秒新闻早报 |
| baidubaike | `plugins/baidubaike/` | 百度百科查询 |
| youdaofanyi | `plugins/youdaofanyi/` | 有道翻译 |
| pinyinsuoxie | `plugins/pinyinsuoxie/` | 拼音缩写翻译 |
| wordcloud | `plugins/wordcloud/` | 聊天词云生成 |
| crazykfc | `plugins/crazykfc/` | 疯狂星期四文案 |
| moyuban | `plugins/moyuban/` | 摸鱼办倒计时 |
| choose | `plugins/choose/` | 选择困难助手 |
| sjyjyy | `plugins/sjyjyy/` | 随机语录 |
| sjyztx | `plugins/sjyztx/` | 随机头像 |
| xzys | `plugins/xzys/` | 星座运势 |
| love | `plugins/love/` | 恋爱语录 |
| hb | `plugins/hb/` | 模拟抢红包游戏 |
| friendadd | `plugins/friendadd/` | 自动通过好友申请 |
| ghmonitor | `plugins/ghmonitor/` | 公众号消息监控转发 |
| localimage | `plugins/localimage/` | 本地图片服务 |
| localimagespider | `plugins/localimagespider/` | 图片爬取 |
| memepicture | `plugins/memepicture/` | 表情包检索 |
| acgimg | `plugins/acgimg/` | ACG 图片 |
| news | `plugins/news/` | 新闻查询 |
| plmm | `plugins/plmm/` | 漂亮妹妹图片 |
| tbbi | `plugins/tbbi/` | 淘宝相关 |
| xddq | `plugins/xddq/` | 寻道大千 |
| chaid | `plugins/chaid/` | 微信 ID 查档 |

## 匹配规则说明

插件通过 `control.Register()` 注册后，可使用以下内置匹配器：

| 匹配器 | 说明 | 示例 |
|--------|------|------|
| `OnMessage` | 匹配任意消息 | - |
| `OnPrefix` | 前缀匹配 | `天气 北京` |
| `OnSuffix` | 后缀匹配 | - |
| `OnCommand` | 指令匹配 | `/help` |
| `OnRegex` | 正则匹配 | `\d+天气` |
| `OnKeyword` | 关键词匹配 | `天气` |
| `OnFullMatch` | 完全匹配 | `帮助` |

内置过滤规则：

| 规则 | 说明 |
|------|------|
| `AdminPermission` | 仅超级管理员 |
| `OnlyGroup` / `OnlyPrivate` | 仅群聊/仅私聊 |
| `OnlyAtMe` / `AtMeOrReference` | 需要 @机器人 或引用 |
| `UserOrGroupAdmin` | 私聊或群管理员 |

## 管理指令

超级管理员支持以下系统指令（默认前缀 `/`）：

| 指令 | 说明 |
|------|------|
| `/沉默 [wxid]` | 沉默指定群（机器人不回复） |
| `/响应 [wxid]` | 恢复指定群响应 |
| `/启用 [插件]` | 在当前群启用插件 |
| `/禁用 [插件]` | 在当前群禁用插件 |
| `/全局启用 [插件]` | 全局启用插件 |
| `/全局禁用 [插件]` | 全局禁用插件 |
| `/ban [插件] [用户]` | 封禁用户使用指定插件 |
| `/unban [插件] [用户]` | 解封用户 |

## 开发插件

### 示例：创建一个简单插件

```go
package myplugin

import (
    "wxbot/engine/control"
    "wxbot/engine/robot"
)

func init() {
    engine := control.Register("myplugin", &control.Options{
        Alias: "我的插件",
        Help:  "这是一个示例插件",
    })

    // 匹配关键词 "你好"
    engine.OnKeyword("你好").SetBlock(true).Handle(func(ctx *robot.Ctx) {
        ctx.ReplyText("你好呀！有什么可以帮你的？")
    })
}
```

### 关键 API

```go
// 消息回复
ctx.ReplyText("文本消息")
ctx.ReplyImage("local://path/to/image.png")  // 发送本地图片
ctx.ReplyImage("https://example.com/img.jpg") // 发送网络图片

// 获取事件信息
ctx.Event.FromWxId      // 发送者微信 ID
ctx.Event.FromGroup     // 来源群 ID（群聊时有效）
ctx.Event.MessageString // 消息文本内容

// 异步事件
ch := ctx.Channel.Next(robot.OnText)
go func() {
    next := <-ch
    // 处理下一条文本消息
}()
```

## 项目结构

```
wxbot/
├── main.go                  # 入口文件
├── config.yaml              # 配置文件
├── plugins.yaml             # 插件声明
├── engine/
│   ├── control/             # 插件控制层
│   ├── plugins/             # 插件导入中心（自动生成）
│   ├── pkg/                 # 工具包（日志、数据库、加密、JWT 等）
│   └── robot/               # 核心引擎
├── framework/
│   ├── dean/                # Dean 框架适配
│   └── vlw/                 # VLW 框架适配
├── plugins/                 # 插件实现（26+）
├── web/                     # Web 管理界面
│   ├── web.go               # 嵌入指令
│   └── dist/                # 编译后的 SPA 产物
├── examples/                # 示例代码
│   ├── batchSend/           # 批量发送
│   ├── conversation/        # 多轮对话
│   ├── friendAddReq/        # 好友请求处理
│   └── invite/              # 群邀请
└── docs/                    # 文档图片
```

## 主要依赖

| 依赖 | 用途 |
|------|------|
| [Gin](https://github.com/gin-gonic/gin) | HTTP 服务器 |
| [GORM](https://gorm.io/) + [SQLite](https://github.com/glebarez/sqlite) | ORM 与数据库（纯 Go） |
| [Viper](https://github.com/spf13/viper) | 配置管理 |
| [Logrus](https://github.com/sirupsen/logrus) | 日志记录 |
| [req/v3](https://github.com/imroc/req) | HTTP 客户端 |
| [gjson](https://github.com/tidwall/gjson) | JSON 快速解析 |
| [gocron](https://github.com/go-co-op/gocron) | 定时任务 |
| [go-openai](https://github.com/sashabaranov/go-openai) | OpenAI API 客户端 |

## 许可证

[Apache License 2.0](LICENSE)
