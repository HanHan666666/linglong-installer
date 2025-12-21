# Detect 屏幕需求文档（草案）

## 目标
在 go-pkg-installer 中新增一个内置、可配置的 `detect` screen，用于在进入页面时执行当前步骤的任务，并在页面上仅显示简洁文本（无进度条、无日志），同时支持基于上下文渲染内容。

## 背景
现有 `richtext` screen 渲染一次后不会在任务执行完成后自动更新。专门的 `detect` screen 可以解决以下问题：
- 进入步骤即执行预检类任务
- 在 UI 中展示简短的确认信息
- 任务未完成前禁止继续

## 范围
包含：
- 新增 screen 类型：`detect`
- 可配置文本（title/content/description）
- 自动执行 `step.tasks`
- 任务完成后刷新内容
- 失败处理与阻止继续

不包含：
- 现有 screen 的改造或替换
- 流程引擎或任务注册机制的改造

## 交互与体验
- 进入时展示“检测中…”类提示（可配置）。
- 任务执行期间不展示日志、不展示进度条。
- 成功后展示基于上下文渲染的内容。
- 失败后显示错误提示并阻止继续（沿用已有错误弹窗逻辑）。

## 配置示例（YAML）
```
steps:
  - id: detect
    title: "Check System"
    screen:
      type: detect
      title: "Check System"
      description: "Detecting environment..."
      content: |
        Detected info: ${env.some_field}
        Please confirm and continue.
    tasks:
      - type: go:someTask
        param: "value"
```

说明：
- `description` 用于任务执行中展示。
- `content` 在任务结束后展示，支持上下文变量。

## 规则与校验
- `screen.type` 必须为 `detect`。
- `content` 必填。
- `description` 可选，默认提示为“Detecting…”。
- 允许并默认执行 `step.tasks`。

## 执行模型
- `detect` screen 内部创建 TaskRunner 并执行当前步骤任务。
- 任务完成后更新文本区域。
- 任务失败时设置失败状态并阻止下一步。

## 失败处理
- 页面内显示简短失败提示。
- 点击“继续”时沿用已有错误弹窗机制。

## 最小改动清单（库内）
- 新增 `pkg/ui/screen_detect.go`。
- 在 `pkg/ui/window.go` 注册 `detect` screen。
- 在 `installer.schema.json` 增加 `detect` screen 的 schema，并允许 title/description/content。
- 文档中补充 `detect` screen 的说明。
