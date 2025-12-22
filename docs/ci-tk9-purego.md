# CI 里生成 tk9.0 纯 Go 依赖库的方案说明

## 背景
Debian10 上运行时出现 `purego: fn is nil`，根因是 tk9.0 的 purego 模式需要加载内置的
`libtcl/libtk` 动态库，而这些库如果是在更新系统上编译，会依赖更高版本的 glibc/系统库，
导致 Debian10 无法加载。

## 解决方案
在 CI 的 Debian10 容器内，临时编译 tk9.0 的 purego 依赖库，然后再执行 `go build`：
1. CI 进入 Debian10 环境，安装构建依赖（gcc、make、X11/font 等开发包）。
2. 删除 `go.mod` 里本地替换的 `replace modernc.org/tk9.0 => ../tk9.0`，避免 CI 无法解析本地路径。
3. `go mod download modernc.org/tk9.0` 下载模块到 GOPATH 缓存。
4. 在模块缓存目录执行 `make lib_linux_purego`，生成 Debian10 兼容的 `libtcl/libtk`。
5. 用 `-mod=mod` 构建二进制，确保使用的是刚生成的库。

这样构建出的二进制包含 Debian10 编译的动态库，能在老系统上稳定运行，同时无需把 tk9.0
源码 vendoring 到项目里。

## 当前 CI 变更
见 `.github/workflows/build.yml`：
- 增加构建依赖安装
- 在构建前执行 `make lib_linux_purego`
- 构建时使用 `-mod=mod`

## 注意事项
- 该流程仅生成 **linux/amd64** 的 purego 库。如果要兼容 Debian10 的 **arm64**，
  需要在 arm64 构建环境执行同样流程。
- 每次升级 tk9.0 版本时，CI 会自动重新生成适配 Debian10 的库，无需手动提交产物。
