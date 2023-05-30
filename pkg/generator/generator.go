/*
Copyright 2023 Pone Ding <poneding@gmail.com>.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// summaryEntry summary 条目
type summaryEntry struct {
	Path              string
	SubSummaryEntries []*summaryEntry
	SubSummaryLinks   []summaryLink
}

// summaryLink -
type summaryLink struct {
	title string
	path  string
}

// newSummaryEntry -
func newSummaryLink(title, path string) summaryLink {
	return summaryLink{
		title: title,
		path:  path,
	}
}

// summaryGenerator summary 生成器
type summaryGenerator struct {
	summaryEntry  *summaryEntry
	originWorkDir string
	ReadmeTitle   string
	SummaryTitle  string
	ignoredDirs   []string
}

const (
	defaultSummaryTitle = "SUMMARY"
	defaultReadmeTitle  = "README"
)

// GeneratorOption SummaryGenerator 配置
type GeneratorOption struct {
	// Path Gitbook 项目根目录，默认为 "."
	Path string
	// SummaryTitle SUMMARY.md 标题，默认为 Gitbook 项目根目录名或 "SUMMARY"
	SummaryTitle string
	// ReadmeTitle README.md 标题，默认为 Gitbook 项目根目录名或 "README"
	ReadmeTitle string
	// IgnoredDirs 忽略的目录
	IgnoredDirs []string
}

// NewSummaryGenerator 创建 summary generator
func NewSummaryGenerator(opt *GeneratorOption) *summaryGenerator {
	if opt.Path == "" {
		opt.Path = "."
	}

	return &summaryGenerator{
		summaryEntry: &summaryEntry{
			Path:            opt.Path,
			SubSummaryLinks: make([]summaryLink, 0),
		},
		ReadmeTitle:  opt.ReadmeTitle,
		SummaryTitle: opt.SummaryTitle,
		ignoredDirs:  opt.IgnoredDirs,
	}
}

// enterWorkDir 进入工作目录
// 1. 切换工作目录; 2. Path 重置为 "."
func (g *summaryGenerator) enterWorkDir() {
	var err error
	g.originWorkDir, err = os.Getwd()
	if err != nil {
		log.Fatalf("init work dir error: %v", err)
	}
	err = os.Chdir(g.summaryEntry.pathAbs())
	if err != nil {
		log.Fatalf("change work dir error: %v", err)
	}
	g.summaryEntry.Path = "."
}

// exitWorkDir 退出工作目录
func (g *summaryGenerator) exitWorkDir() {
	err := os.Chdir(g.originWorkDir)
	if err != nil {
		log.Fatalf("exit work dir error: %v", err)
	}
}

// Generate 生成 SUMMARY.md 文件内容
func (g *summaryGenerator) Generate() string {
	g.enterWorkDir()
	defer g.exitWorkDir()

	setupSummaryEntries(g.summaryEntry, g.ignoredDirs)

	summary := "# " + g.summaryTitleOrDefault() + "\n"
	summary += fmt.Sprintf("\n- [%s](%s)", g.readmeTitleOrDefault("README"), "README.md")
	summary = g.summaryEntry.generate(summary)

	return summary
}

// summaryTitleOrDefault 获取 SUMMARY 标题
func (g *summaryGenerator) summaryTitleOrDefault() string {
	if g.SummaryTitle == "" {
		g.SummaryTitle = pathBaseName(g.summaryEntry.Path)
		if g.SummaryTitle == "/" {
			g.SummaryTitle = defaultSummaryTitle
		}
	}
	return g.SummaryTitle
}

// readmeTitleOrDefault 获取 README 标题
func (g *summaryGenerator) readmeTitleOrDefault(def string) string {
	if g.ReadmeTitle == "" {
		g.ReadmeTitle = pathBaseName(g.summaryEntry.Path)
		if g.ReadmeTitle == "/" {
			g.ReadmeTitle = defaultReadmeTitle
		}
	}
	return g.ReadmeTitle
}

// setupSummaryEntries 设置 summary entries
func setupSummaryEntries(se *summaryEntry, ignoredDirs []string) *summaryEntry {
	dirEntries, _ := os.ReadDir(se.Path)
	for _, dirEntry := range dirEntries {
		path := filepath.Join(se.Path, dirEntry.Name())
		if pathIgnored(ignoredDirs, path) {
			continue
		}
		if isEffectiveDir(dirEntry) {
			subSummaryEntry := setupSummaryEntries(&summaryEntry{
				Path:            path,
				SubSummaryLinks: make([]summaryLink, 0),
			}, ignoredDirs)
			se.SubSummaryEntries = append(se.SubSummaryEntries, subSummaryEntry)
			continue
		}
		if isEffectiveMdFile(dirEntry) {
			title := strings.TrimRight(dirEntry.Name(), ".md")
			se.SubSummaryLinks = append(se.SubSummaryLinks, newSummaryLink(title, path))
		}
	}

	return se
}

// generate 生成 summary entry 内容，递归
func (se *summaryEntry) generate(summary string) string {
	depth := strings.Count(se.Path, "/")
	indent := strings.Repeat("  ", depth)

	if se.hasInvaliadMdFile() && !isRootPath(se.Path) {
		// 编码路径中的空格
		url := strings.ReplaceAll(se.Path, " ", "%20")
		summary += fmt.Sprintf("\n%s- [%s](%s)", indent, pathBaseName(se.Path), filepath.Join(url, "README.md"))
	}

	for _, subSummaryEntry := range se.SubSummaryEntries {
		summary = subSummaryEntry.generate(summary)
	}

	for _, sl := range se.SubSummaryLinks {
		depth = strings.Count(sl.path, "/")
		indent = strings.Repeat("  ", depth)
		// 编码路径中的空格
		url := strings.ReplaceAll(sl.path, " ", "%20")
		summary += fmt.Sprintf("\n%s- [%s](%s)", indent, sl.title, url)
	}
	if se.hasInvaliadMdFile() && !strings.HasSuffix(summary, "\n") {
		summary += "\n"
	}

	return summary
}

// pathAbs 获取绝对路径
func (se *summaryEntry) pathAbs() string {
	if strings.HasPrefix(se.Path, "~") {
		se.Path = strings.Replace(se.Path, "~", os.Getenv("HOME"), 1)
	}
	if se.Path == "" {
		se.Path = "."
	}
	abs, _ := filepath.Abs(se.Path)
	return abs
}

// hasMdFile 判断是否有 .md 文件
func (se *summaryEntry) hasInvaliadMdFile() bool {
	return hasInvaliadMdFile(se.Path)
}
