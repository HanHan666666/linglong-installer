# Linglong Installer - Agent 指南

玲珑商店社区版安装器，基于 Go + Tk GUI 的配置驱动 Linux 安装工具。

## 重点（极其重要）
- 在接到用户的任务的时候，先不要着急开始修改代码，要先分析需求，分析代码，列举解决方案，
- 详细的向用户说明你的思路，和你打算如何实现这个需求。
- 要分析整个项目的架构，一切都要从整个项目的角度入手，不能直接看完一个文件就写代码。
- 先问清楚、绝对不允许猜测：遇到需求或现状不确定时，先明确提问，不要主观假设；方案需先得到用户确认再开工。
- 每一处代码修改都要有必要的注释
- 先方案后编码：先梳理背景/现状 → 列备选方案（含改动面、影响范围、取舍理由）→ 让用户确认 → 再动手。**只有在用户确认你的方案后，才开始动手写代码, 不然你很快就会被关机，更换下一个AI，一定要小心。**
- 统一入口：能收敛的业务逻辑要集中封装（如卸载流程用 `useAppUninstall`），避免在多个页面/组件里写重复弹窗或副作用。
- 变更记录：完成功能后，将关键经验和约定同步到本指南，方便后续遵循。


## 代码要求
1. 代码要求结构清晰，不应付事情，长远维护考虑，遵循设计模式最佳实践，遵循项目代码风格。
2. 保证代码逻辑严谨，整洁，结构清晰，容易理解和维护，不要过度设计增加系统复杂性
3. 工程优化，以工程化，能安全正常使用不出错为主，考虑周全，遵循越复杂越容易出错，越简单越容易可控原则，一个健康的系统 越简单越可控
4. 遵循合理的组件化设计原则，要考虑组件复用性的可能。
5. 在你发现架构不合理的时候，要及时的提出来。
6. 编写代码的过程中，必须牢记以下几个原则：
    - 开闭原则（Open Closed Principle，OCP）
    - 单一职责原则（Single Responsibility Principle, SRP）
    - 里氏代换原则（Liskov Substitution Principle，LSP）
    - 依赖倒转原则（Dependency Inversion Principle，DIP）
    - 接口隔离原则（Interface Segregation Principle，ISP）
    - 合成/聚合复用原则（Composite/Aggregate Reuse Principle，CARP）
    - 最少知识原则（Least Knowledge Principle，LKP）或者迪米特法则（Law of  Demeter，LOD）


## 快速命令

| 操作 | 命令 |
|------|------|
| 运行 | `go run ./` |
| 运行（指定配置） | `go run ./ --config installer.yaml` |
| Headless/CLI 模式 | `go run ./ --headless` |
| 构建 | `go build -o linglong-installer ./` |
| 测试 | `go test ./...` |
| Docker 冒烟测试 | `./docker-smoke.sh` |

**CLI flags**: `--config`, `--action install|uninstall`, `--validate`, `--version`, `--headless`, `--verbose`, `--accept-license`, `--install-dir`, `--install-type`, `--privilege sudo|pkexec|none`, `--set key=value`

## 项目架构

```
linglong-installer/
├── main.go                    # 入口：CLI 解析、配置加载、工作流构建
├── embed.go                   # 内嵌 installer.yaml、logo、脚本到二进制
├── installer.yaml             # 安装流程配置（步骤、守卫、任务）
├── task_distro_script.go      # 发行版检测与脚本匹配任务
├── screen_linglong_version.go # 自定义 GUI 屏幕
├── scripts/
│   ├── common.sh              # 共享函数库（提权、ll-cli 调用等）
│   ├── distros/               # 各发行版安装脚本（ID_VERSION.sh 命名）
│   └── user/                  # 用户态脚本
└── vendor/
    ├── github.com/HanHan666666/go-pkg-installer/  # 核心框架
    │   └── pkg/ui/           
    └── modernc.org/tk9.0/     # 纯 Go Tk GUI（无需系统 Tcl/Tk）
```

**核心依赖**：`go-pkg-installer` 提供流程引擎、GUI 层、配置 schema；`tk9.0` 提供纯 Go Tk 绑定。

## 开发环境注意事项

### ⚠️ go.mod replace 指令

[go.mod](go.mod) 包含本地 replace 指令，新开发者需要处理：

```go
// 第 61 行 - 可能导致构建失败
replace modernc.org/tk9.0 => ../tk9.0

// 第 12 行 - 已注释，本地联调时取消注释
// replace github.com/HanHan666666/go-pkg-installer => /home/han/go-pkg-installer
```

**解决方案**：确保 `~/tk9.0` 目录存在，或注释掉 replace 行使用 vendor 中的版本。

### Go 版本要求

Go 1.24+（`go 1.24.0` + `toolchain go1.24.11`）

### 构建时间

vendor 包含大量 C 库的纯 Go 翻译（libX11、libXft、libfreetype、libtcl、libtk 等），首次构建时间较长。

## 项目约定

### 配置驱动流程

[installer.yaml](installer.yaml) 定义整个安装流程：
- 步骤顺序（welcome → detect → linglong_version → confirm → install → finish）
- 页面类型和守卫条件
- 任务定义

Go 代码只注册自定义扩展点（tasks、screens、guards）。

### 上下文变量系统

任务通过 `ctx.Set("distro.xxx", ...)` 写入变量，页面模板通过 `${distro.xxx}` 渲染，实现任务与 UI 解耦。

### 发行版脚本匹配优先级

`task_distro_script.go` 中的匹配逻辑：
1. 精确匹配：`debian_13.sh`
2. Generic 兜底：`nixos_generic.sh`
3. 上游映射：Linux Mint → Ubuntu

### META 注释协议

脚本头部 `# META: key=value` 行被解析为结构化数据（`repo_name`、`repo_url`、`command`、`next_steps`），用于确认页展示。

## 代码注释要求

- 必须添加代码注释解释行为、约束和权衡，不能让意图隐晦。
- **vendor 下的 UI 代码**（`vendor/github.com/HanHan666666/go-pkg-installer/pkg/ui/`）特别重要 — 依赖更新时本地改动会被覆盖，改动必须有充分的内联注释。
- **markdown/链接渲染路径**需在代码中显式记录支持的子集。当前范围仅限 `[text](https://...)` 风格的 Markdown 链接和裸 `http(s)` URL。
- **Tk widget 限制的 workaround** 必须在代码附近说明原因。

## 变更记录

- 2026-03-13：`Text` 类富文本/日志/协议正文的主题样式统一收敛到 `vendor/github.com/HanHan666666/go-pkg-installer/pkg/ui/theme.go` 的 `applyTextStyle()`；优先使用 `flat + highlight ring` 的轻边框方案，不要在各个 screen 里单独堆叠 `solid` 黑边。
- 2026-03-13：主操作按钮（`Primary.TButton` / `Accent.TButton`）的禁用态文字保持白色，避免安装/完成这类蓝色强调按钮在禁用时出现黑字；灰色中性按钮继续沿用 muted 文案颜色，不要一刀切修改全部 disabled button 前景色。
- 2026-03-13：需要“某一步不能点上一部”时，优先在 `installer.yaml` 的 step 上声明 `allowBack: false`，不要在页面里零散写按钮禁用逻辑；同时注意在 `main.go` 构建 `core.Step` 时把 `Next/Prev/Branch/AllowBack/AllowJump/Route` 这类导航元数据完整拷入运行时结构。
- 2026-03-16：安装中的日志复制能力优先收敛到通用 `go:progress` screen，不要为单个项目复制一份专用进度页；复制内容应直接来源于当前日志 `Text` 组件对应的缓冲，保证用户复制到的内容与界面可见日志一致。
- 2026-03-17：`tools/github2gitee/github2gitee.js` 的发布链路必须先把 GitHub 的分支和 tags 镜像到 Gitee，再创建/更新 Gitee Release；不要依赖 Gitee 在 release 创建时为一个尚未同步的 commit 自动建 tag，这会触发“创建标签失败”。
- 2026-03-17：`tools/github2gitee/github2gitee.js` 里涉及网络传输的大型 `git clone/push` 不要吞掉子进程输出；应显式开启 `--progress` 并把 stdout/stderr 直接透传到终端，否则用户只能看到脚本停在某一行，误判为卡死。
- 2026-03-29：Evernight Vista 44 从 `ID=fedora` 切到 `ID=evernight` 后，发行版解析与 `install-linyaps-env.sh` 都必须显式映射到 Fedora 44；不要把 `evernight` 泛化成“任意版本都复用 Fedora”，当前只确认 Vista 44 这条链路可用。

## Git Commit 规范

- 提交内容的语言使用简体中文
- 提交信息必须使用 Conventional Commits 风格，常用前缀包括：`feat:`、`fix:`、`docs:`、`refactor:`、`test:`、`chore:`。
- `feat:` 用于新增功能；`fix:` 用于修复问题；不要随意混用前缀。
- 提交标题要简洁明确，优先写“做了什么”，不要写空泛描述。
- 示例：
  - `feat: add clickable markdown links to installer text`
  - `fix: handle missing distro metadata in confirm screen`
  - `docs: clarify local replace setup in AGENTS guide`

## 添加自定义 Screen

参考 [screen_linglong_version.go](screen_linglong_version.go)：实现 `ui.ScreenRenderer` 接口，在 `main.go` 中通过 `ui.RegisterScreen` 注册。
