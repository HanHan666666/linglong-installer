// Package ui provides simple i18n helpers for built-in UI strings.
package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

var uiTranslations = map[string]map[string]string{
	"en": {
		"button.back":             "Go Back",
		"button.cancel":           "Cancel",
		"button.continue":         "Continue",
		"button.install":          "Install",
		"button.finish":           "Finish",
		"button.exit":             "Exit",
		"button.close":            "Close",
		"button.browse":           "Browse...",
		"dialog.cancel.title":     "Cancel Installation",
		"dialog.cancel.msg":       "Are you sure you want to cancel the installation?",
		"dialog.select.dir":       "Select Directory",
		"dialog.select.file":      "Select File",
		"dialog.validation.title": "Validation Error",
		"dialog.error.title":      "Error",
		"title.welcome":           "Welcome to %s",
		"title.license":           "License Agreement",
		"title.directory":         "Select Installation Directory",
		"title.options":           "Installation Options",
		"title.form":              "Configuration",
		"title.progress":          "Installing...",
		"title.detect":            "Check System",
		"title.summary":           "Installation Summary",
		"title.complete":          "Installation Complete",
		"title.info":              "Information",
		"desc.welcome":            "This will install %s on your computer.",
		"desc.welcome.version":    "This will install %s version %s on your computer.",
		"desc.directory":          "Choose the folder where you want to install the application.",
		"desc.progress.wait":      "Please wait while the installation completes.",
		"desc.summary.success":    "%s has been installed on your computer.",
		"footer.welcome":          "Click 'Continue' to proceed with the installation.",
		"label.install.dir":       "Installation Directory:",
		"label.required.space":    "Required space: ",
		"label.plan":              "Planned actions:",
		"label.errors":            "Errors:",
		"label.logfile":           "Log file: %s",
		"label.installed.to":      "Installed to:",
		"label.launch":            "Launch application after closing",
		"label.accept":            "I accept the terms of the license agreement",
		"label.admin":             "(admin)",
		"status.prepare":          "Preparing...",
		"status.failed":           "Installation Failed",
		"status.complete":         "Installation Complete",
		"status.no_tasks":         "No tasks to run",
		"msg.no_tasks":            "No tasks configured for installation.",
		"msg.ready":               "Review the details before installing.",
		"msg.content.file":        "Content would be loaded from: %s",
		"msg.license.file":        "License content would be loaded from: %s",
		"msg.license.empty":       "License content is empty.",
		"msg.scroll_end":          "Please scroll to the end of the license to continue.",
		"msg.accept_license":      "You must accept the license agreement to continue.",
		"msg.success":             "✓ The installation completed successfully!",
		"msg.failure":             "✗ The installation encountered errors.",
		"msg.detect.running":      "Detecting...",
		"msg.detect.failed":       "Detection failed.",
		"msg.detect.in_progress":  "Detection is still running",
		"msg.field.required":      "%s is required",
		"msg.dir.required":        "Please select an installation directory.",
		"msg.dir.create":          "Cannot create installation directory: %v",
		"msg.dir.parent":          "Parent path is not a directory.",
		"msg.task.queue_fail":     "Failed to queue task: %v",
		"msg.task.start":          "Starting: %s",
		"msg.task.complete":       "Completed: %s",
		"msg.task.error":          "Error in %s: %v",
		"msg.install.failed":      "Installation failed: %v",
		"msg.install.in_progress": "Installation is still in progress",
		"footer.close":            "Click 'Close' to exit the installer.",
	},
	"zh": {
		"button.back":             "返回",
		"button.cancel":           "取消",
		"button.continue":         "继续",
		"button.install":          "安装",
		"button.finish":           "完成",
		"button.exit":             "退出",
		"button.close":            "关闭",
		"button.browse":           "浏览...",
		"dialog.cancel.title":     "取消安装",
		"dialog.cancel.msg":       "确定要取消安装吗？",
		"dialog.validation.title": "校验错误",
		"dialog.error.title":      "错误",
		"title.welcome":           "欢迎使用 %s",
		"title.license":           "许可协议",
		"title.directory":         "选择安装目录",
		"title.options":           "安装选项",
		"title.form":              "配置",
		"title.progress":          "正在安装...",
		"title.detect":            "系统检测",
		"title.summary":           "安装摘要",
		"title.complete":          "安装完成",
		"title.info":              "信息",
		"desc.directory":          "请选择要安装应用的目录。",
		"footer.welcome":          "点击“继续”开始安装。",
		"label.install.dir":       "安装目录：",
		"label.required.space":    "所需空间：",
		"label.plan":              "计划执行：",
		"label.errors":            "错误：",
		"label.logfile":           "日志文件：%s",
		"label.installed.to":      "安装位置：",
		"label.launch":            "关闭后启动应用",
		"label.accept":            "我已阅读并同意许可协议",
		"status.prepare":          "准备中...",
		"status.failed":           "安装失败",
		"status.complete":         "安装完成",
		"status.no_tasks":         "没有可执行的任务",
		"msg.no_tasks":            "未配置安装任务。",
		"msg.scroll_end":          "请先滚动到协议末尾。",
		"msg.accept_license":      "请先勾选同意许可协议。",
		"msg.success":             "✓ 安装已成功完成！",
		"msg.failure":             "✗ 安装过程中出现错误。",
		"msg.detect.running":      "检测中...",
		"msg.detect.failed":       "检测失败。",
		"msg.detect.in_progress":  "检测仍在进行中",
		"footer.close":            "点击“关闭”退出安装程序。",
	},
}

func tr(ctx *core.InstallContext, key, fallback string) string {
	locale := resolveLocale(ctx)
	if bundle, ok := uiTranslations[locale]; ok {
		if val, ok := bundle[key]; ok {
			return val
		}
	}
	return fallback
}

func resolveLocale(ctx *core.InstallContext) string {
	if ctx != nil {
		if v, ok := ctx.Get("meta.lang"); ok {
			return normalizeLocale(v)
		}
		if v, ok := ctx.Get("meta.locale"); ok {
			return normalizeLocale(v)
		}
	}
	if env := os.Getenv("LANG"); env != "" {
		return normalizeLocale(env)
	}
	return "en"
}

func normalizeLocale(value any) string {
	raw := strings.ToLower(strings.TrimSpace(strings.Split(fmt.Sprintf("%v", value), ".")[0]))
	if strings.HasPrefix(raw, "zh") {
		return "zh"
	}
	return "en"
}
