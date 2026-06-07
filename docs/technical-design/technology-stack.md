# 英语学习平台技术设计

## 1. 技术选型结论

本项目建议采用以下总体技术栈：

```text
Wails + Go + React + TypeScript + PostgreSQL + pgvector + AI Provider Adapter
```

核心判断：

- 项目目标是 macOS 桌面应用。
- 开发者主语言是 Go。
- 产品需要复杂 UI，包括阅读器、写作编辑器、AI 反馈面板、音频交互和学习记录。
- 第一版需要可上线 MVP，同时保留后续接入题库、语音实时对话、多模型、RAG、向量检索和云同步能力的空间。

因此，不建议使用纯 Go GUI 框架作为主方案，也不建议第一版使用 Electron。推荐使用 Wails 作为桌面壳，Go 作为业务后端，React/TypeScript 作为前端界面。

## 2. 总体架构

```text
macOS Desktop App
├── Wails Desktop Shell
├── Frontend
│   ├── React
│   ├── TypeScript
│   ├── Vite
│   ├── Tailwind CSS
│   ├── Zustand
│   └── TanStack Query
├── Go Backend
│   ├── User/Profile Service
│   ├── ListeningSpeaking Service
│   ├── Reading Service
│   ├── Writing Service
│   ├── AI Orchestration Service
│   ├── Audio Service
│   ├── Content Service
│   └── Learning Record Service
├── Data Layer
│   ├── PostgreSQL
│   ├── pgvector
│   ├── Full Text Search
│   └── Local Audio/File Storage
└── External AI Providers
    ├── LLM
    ├── Speech-to-Text
    ├── Text-to-Speech
    └── Realtime Speech
```

## 3. 桌面应用框架

### 3.1 推荐方案：Wails

Wails 适合本项目的原因：

- 后端使用 Go，贴合开发者技术栈。
- 前端可以使用 React/TypeScript 构建复杂桌面工作台。
- 使用系统 WebView，整体比 Electron 更轻。
- 适合 macOS 桌面应用打包。
- Go 后端服务可以直接暴露给前端调用。

## 4. 前端技术栈

前端视觉规范以 `docs/design-system/frontend-style-guide.md` 为准。整体方向是白色主题、克制、高级、去 AI 味的桌面学习工作台。

### 4.1 基础技术

- React：构建复杂模块化界面。
- TypeScript：提高前端类型可靠性。
- Vite：开发构建工具。
- Tailwind CSS：快速构建清晰、统一的桌面 UI。

### 4.2 状态和请求管理

- Zustand：管理前端 UI 状态，例如当前模块、侧边栏状态、编辑器状态。
- TanStack Query：管理异步请求、缓存、加载状态和错误状态。

### 4.3 编辑器和阅读器

- TipTap 或 Lexical：用于写作编辑器。
- 自研阅读器组件：用于文章段落展示、划词、划句、双语对照、AI 解释面板。
- Web Audio API / MediaRecorder：用于录音采集和音频播放。

## 5. Go 后端技术栈

### 5.1 核心后端

- Go：业务服务、AI 编排、数据访问、桌面端能力桥接。
- Wails Service：前后端桥接。
- PostgreSQL：主数据库，承载用户、内容、训练记录、AI 调用日志等核心数据。
- pgvector：承载 embedding 和向量检索，支持后续 RAG、相似文章、相似错题、表达推荐。
- PostgreSQL Full Text Search：支持文章、笔记、生词、表达、写作草稿搜索。
- goose：数据库 migration。
- sqlc：生成类型安全的数据访问代码。
- zap 或 zerolog：结构化日志。
- HTTP/API 测试：优先用于验证登录注册、用户角色、学习记录、AI 工作流等后端行为。
- Go 单元测试：仅用于验证纯工具逻辑，例如解析、校验、hash、模型辅助函数。

### 5.2 后端架构分层

```text
internal/
  app/
    app.go
  platform/
    config/
    logger/
    errors/
    events/
    filesystem/
    security/
  models/
    dto/
    db/
    common/
  core/
    profile/
    learning_records/
    content/
    ai/
    audio/
    vocabulary/
    expressions/
  modules/
    listening_speaking/
    reading/
    writing/
    question_bank/
  adapters/
    ai/
    speech/
    storage/
    desktop/
```

设计原则：

- `platform/` 放全局基础设施，例如配置、日志、错误、事件、文件系统和安全能力。
- `models/` 放跨层共享的数据结构，重点包括前后端 DTO、数据库模型、通用枚举和分页结构。
- `core/` 放跨模块核心能力，例如 AI 编排、学习记录、内容系统、生词本、表达库和音频能力。
- `modules/` 放听说、阅读、写作、题库等业务模块。模块内部先不强制设计第三级目录，等功能复杂后再按能力拆分。
- `adapters/` 放外部适配器，例如 AI 厂商、语音服务、PostgreSQL、文件存储和 Wails 桌面桥接。

## 6. AI 技术栈

### 6.1 AI 层设计原则

AI 层不建议直接绑定某一家模型，也不建议第一版引入过重的 Agent 框架。建议自研轻量 AI 编排层，并通过 Provider Adapter 接入不同模型服务。

核心原则：

- 模块服务不直接调用模型厂商。
- 所有模型调用统一经过 AI Orchestration Service。
- AI 输出尽量使用结构化 JSON。
- Prompt 模板集中管理。
- 每次 AI 调用记录用途、输入摘要、输出、耗时、token 用量和成本。

### 6.2 AI 架构

```text
Module Service
  -> AI Orchestration Service
    -> Prompt Template Registry
    -> AI Provider Adapter
      -> OpenAI / DeepSeek / Qwen / Doubao / GLM / Other Providers
    -> Structured Output Parser
    -> Learning Record Service
```

### 6.3 Provider 接口

建议在 Go 中定义统一接口：

```go
type AIProvider interface {
    Generate(ctx context.Context, req GenerateRequest) (*GenerateResponse, error)
    GenerateStructured(ctx context.Context, req StructuredRequest) (*StructuredResponse, error)
    Stream(ctx context.Context, req StreamRequest) (<-chan StreamEvent, error)
}
```

模块只依赖内部接口，不依赖具体厂商 SDK。

### 6.4 推荐模型接入顺序

第一阶段建议：

- 主力文本模型：DeepSeek 或 Qwen。
- 备用文本模型：OpenAI。
- 中文环境与成本优先：DeepSeek、Qwen、豆包、GLM。
- 语音能力：可先使用厂商 STT/TTS API，后续再接 Realtime Speech。

第二阶段建议：

- 增加多模型路由。
- 按任务选择模型，例如阅读解析用便宜模型，写作评分用更强模型。
- 增加调用成本统计。
- 增加 Prompt A/B 测试。

## 7. 中国低成本 AI 模型建议

国内可重点关注以下模型供应商：

### 7.1 DeepSeek

适合：

- 阅读解析
- 写作批改
- 问题生成
- 学习报告
- 结构化输出

特点：

- 成本通常较低。
- 推理和文本能力强。
- 适合作为 MVP 文本模型主力。

### 7.2 通义千问 Qwen / 阿里云百炼

适合：

- 中文解释
- 英文学习辅导
- 阅读翻译
- 写作建议
- 多模型规格选择

特点：

- 国内云服务生态成熟。
- 模型种类丰富。
- 适合企业级稳定接入。

### 7.3 豆包 / 火山方舟

适合：

- 低成本高频对话
- 通用文本生成
- 学习陪练
- 内容生成

特点：

- 成本和吞吐通常有优势。
- 适合高频用户交互场景。

### 7.4 智谱 GLM

适合：

- 通用对话
- 文本分析
- 中文解释
- 结构化反馈

特点：

- 国内生态较完整。
- 可作为备用或特定任务模型。

### 7.5 选型建议

MVP 可以这样配置：

```text
默认文本模型：DeepSeek
中文解释/备用模型：Qwen
高频低成本对话备用：豆包
国际模型备用：OpenAI
```

不建议第一版同时深接太多模型。建议先实现 Provider Adapter，再接入 1 个主模型和 1 个备用模型。

## 8. 语音技术栈

### 8.1 MVP 语音方案

第一版听说模块建议采用非实时链路：

```text
前端录音
-> Go 保存音频
-> STT 转写
-> LLM 分析/回复
-> TTS 生成语音
-> 前端播放
-> 会话复盘
```

优点：

- 实现复杂度可控。
- 容易保存音频、转写和反馈。
- 适合 MVP。

### 8.2 后续增强方案

后续可以加入实时语音对话：

```text
前端 WebRTC
-> Realtime Speech API
-> 实时语音输入输出
-> Go 保存会话摘要和训练记录
```

这会显著提升口语陪练体验，但第一版不建议作为唯一实现路径。

## 9. 数据存储设计

### 9.1 PostgreSQL 优先

第一版建议使用 PostgreSQL 作为主数据库，而不是 SQLite。原因是本项目后续会涉及大量 AI 相关能力，包括 embedding、向量检索、训练记录分析、内容包、题库、学习画像和可能的云同步。PostgreSQL 更适合长期扩展。

- 用户资料存在 PostgreSQL。
- 文章、题目、写作草稿存在 PostgreSQL。
- AI 反馈、训练记录、能力标签存在 PostgreSQL。
- embedding 向量使用 pgvector 存储。
- 音频文件存在本地文件目录。
- 数据库只保存音频路径和元数据。
- 桌面前端不直接连接数据库，统一通过 Go 后端服务访问数据。

需要注意：如果未来要做云同步或多设备，PostgreSQL 应部署在服务端。macOS App 通过 Go 后端接口或远端 API 访问数据，不应让桌面前端直接暴露数据库连接。

### 9.2 主要数据表方向

```text
users
profiles
learning_records
ability_tags
articles
reading_sessions
reading_questions
speaking_scenarios
speaking_sessions
writing_tasks
writing_drafts
writing_reviews
saved_words
saved_expressions
ai_call_logs
```

### 9.3 搜索

使用 PostgreSQL Full Text Search 和 pgvector 支持：

- 文章全文搜索
- 阅读笔记搜索
- 生词搜索
- 表达库搜索
- 写作草稿搜索
- 相似文章推荐
- 相似错题推荐
- 相似表达推荐
- 基于用户弱点的内容召回

## 10. 三大模块技术实现

### 10.1 听说模块

核心组件：

- AudioRecorder
- AudioPlayer
- SpeakingSessionService
- SpeechToTextProvider
- TextToSpeechProvider
- ConversationCoach
- PronunciationReviewer
- SessionReviewGenerator

核心流程：

```text
用户选择场景
-> 前端开始录音
-> Go 保存音频
-> STT 转写
-> AI 生成回复或反馈
-> TTS 播放
-> 会话结束后生成复盘
-> 保存学习记录和弱点标签
```

### 10.2 阅读模块

核心组件：

- ReadingLibrary
- ArticleReader
- TextSelectionMenu
- TranslationPanel
- ArticleAnalyzer
- QuestionGenerator
- AnswerReviewService
- VocabularyService

核心流程：

```text
用户选择文章
-> 阅读器展示原文
-> 用户划词/划句
-> AI 返回解释
-> 用户生成题目
-> 用户答题
-> AI 讲解答案
-> 保存生词、笔记和阅读记录
```

### 10.3 写作模块

核心组件：

- WritingTaskLibrary
- WritingEditor
- DraftService
- OutlineAssistant
- WritingReviewer
- ParagraphReviewer
- RevisionComparator
- ExpressionService

核心流程：

```text
用户选择题目
-> AI 辅助构思
-> 用户写作
-> 用户提交草稿
-> AI 评分和批改
-> 用户二次修改
-> AI 对比一稿和二稿
-> 保存表达、错误和能力标签
```

## 11. 推荐项目目录

```text
English-platform/
  docs/
    prd/
    technical-design/
    architecture/
    decisions/
    api/

  app/
    desktop/
      main.go
      wails.json
      bindings/
      lifecycle/

  frontend/
    src/
      main.tsx
      App.tsx
      routes/
      pages/
      components/
      api/
      store/
      hooks/
      lib/

  internal/
    app/
    platform/
      config/
      logger/
      errors/
      events/
      filesystem/
      security/
    models/
      dto/
      db/
      common/
    core/
      ai/
      audio/
      content/
      profile/
      learning_records/
      vocabulary/
      expressions/
    modules/
      listening_speaking/
      reading/
      writing/
      question_bank/
    adapters/
      ai/
        openai/
        deepseek/
        qwen/
        doubao/
        glm/
      speech/
      storage/
        postgres/
        files/
      desktop/
        wails/

  migrations/
    postgres/

  prompts/
    shared/
    listening-speaking/
    reading/
    writing/
    question-bank/

  content-packs/
    base/
    exams/
      ielts/
      toefl/
    scenes/
      business/
      interview/

  assets/
  tests/
    unit/
    integration/
    e2e/
    http/

  scripts/
    dev/
    build/
    db/
```

目录设计说明：

- 第三级业务能力目录先不做过度设计。例如 `internal/modules/reading/` 下第一阶段可以只放少量入口文件，等阅读模块复杂后，再拆出 `library/`、`reader/`、`questions/` 等能力目录。
- `internal/models/` 是统一模型层，用于放 DTO、数据库模型、通用响应结构、枚举、分页参数等。这样可以避免每个模块重复定义相似模型。
- 模块目录不强制套用 `models/usecases/contracts/repos/services/handlers`。只有当某个模块或能力真的变复杂时，再按需要拆分。
- 前端采用标准 React 目录结构，先保持 `routes/`、`pages/`、`components/`、`api/`、`store/`、`hooks/`、`lib/` 这些通用目录，不预先按听说、阅读、写作拆死。
- PostgreSQL 相关实现放在 `internal/adapters/storage/postgres/`，业务模块通过核心服务或接口访问数据。
- AI 厂商接入放在 `internal/adapters/ai/`，业务模块不直接依赖 DeepSeek、Qwen、OpenAI 等 SDK。
- `content-packs/` 用于承载未来雅思、托福、商务英语、面试英语等训练包，不直接塞进业务代码。
- 后端功能测试优先放在 `tests/http/`，通过 HTTP 接口覆盖真实请求、响应、错误码和鉴权流程。

## 12. MVP 实施顺序建议

1. 搭建 Wails + Go + React 基础工程。
2. 建立 PostgreSQL、migration、基础数据表。
3. 建立 AI Provider Adapter。
4. 实现阅读模块，因为文本链路最稳定。
5. 实现写作模块，因为复用 AI 结构化输出。
6. 实现听说模块基础版，先做录音、转写、AI 回复、TTS。
7. 建立学习记录和能力标签。
8. 打包 macOS 应用。

## 13. 风险与规避

### 13.1 AI 成本风险

规避方式：

- Provider Adapter 支持多模型。
- 高频任务使用低成本模型。
- 高价值任务使用更强模型。
- 记录 token 和成本。
- 对阅读解析、写作 review 做缓存。

### 13.2 语音体验风险

规避方式：

- MVP 使用非实时语音链路。
- 后续再升级实时语音。
- 保存转写和音频，便于复盘。

### 13.3 Prompt 不稳定风险

规避方式：

- Prompt 模板集中管理。
- AI 输出使用 JSON Schema。
- 对评分、题目生成、批改结果做解析校验。
- 保留 AI 调用日志。

### 13.4 模块耦合风险

规避方式：

- 三大模块服务独立。
- 学习记录、AI 编排、内容系统作为平台级服务。
- 模块之间通过明确数据对象联动。
- 模块内部不提前创建过多空目录，随着能力复杂度逐步拆分。

## 14. 最终建议

本项目第一版推荐技术栈为：

```text
Wails
+ Go
+ React
+ TypeScript
+ Tailwind CSS
+ PostgreSQL
+ pgvector
+ 自研 AI Orchestration Layer
+ DeepSeek/Qwen 等国内低成本模型
+ OpenAI 作为国际模型备用
```

这个方案可以兼顾：

- Go 开发体验
- macOS 桌面应用交付
- 复杂学习 UI
- AI 能力扩展
- PostgreSQL 数据扩展
- embedding 和向量检索
- 后续题库和云同步扩展
