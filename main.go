package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	// 目标目录，默认使用当前目录
	path string
	// 生成的 Gitbook Summary 文件名，默认 ./SUMMARY.md
	output string
	// 生成的 Gitbook Summary 文件使用的 Title，默认使用父目录名
	summaryTitle string
)

const defaultSummaryTitle = "SUMMARY"

func init() {
	flag.StringVar(&path, "p", ".", "Path to directory")
	flag.StringVar(&summaryTitle, "t", defaultSummaryTitle, "Summary title")
	flag.StringVar(&output, "f", "./SUMMARY.md", "Summary output file")
	flag.Parse()

	if summaryTitle == defaultSummaryTitle {
		absPath, _ := filepath.Abs(path)

		wd := filepath.Base(absPath)

		wdName := strings.Trim(filepath.Base(wd), "/")
		if wdName != "" {
			summaryTitle = wdName
		}
	}
}

func main() {
	summary := generateSummary(path)
	err := os.WriteFile(output, []byte(summary), 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Generated SUMMARY.md")
	}
}

// generateSummary 生成 Summary 内容
func generateSummary(p string) string {
	summary := "# " + summaryTitle + "\n"
	rootDepth := strings.Count(p, "/")
	err := filepath.Walk(p, func(walkPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if isTargetPath(walkPath, info) {
			// 将目录添加到 SUMMARY.md 中
			depth := strings.Count(walkPath, "/") - rootDepth
			name := filepath.Base(walkPath)
			indent := strings.Repeat("  ", depth)
			summary += fmt.Sprintf("\n%s- [%s](%s/README.md)", indent, name, walkPath)
			// 将文件添加到 SUMMARY.md 中
			files, _ := os.ReadDir(walkPath)
			var hasReadme bool
			var targetFiles int
			for _, file := range files {
				if file.Name() == "README.md" {
					hasReadme = true
					continue
				}
				if isTargetFile(file) {
					targetFiles++
					indent := strings.Repeat("  ", depth+1)
					summary += fmt.Sprintf("\n%s- [%s](%s/%s)", indent, file.Name(), walkPath, file.Name())
				}
			}
			// 判断目录是否存在 README.md 文件，不存在并且包含目标文件则自动生成
			if !hasReadme && targetFiles > 0 {
				os.WriteFile(walkPath+"/README.md", []byte("# "+name+"\n"), 0644)
			}
			summary += "\n"
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return summary
}

// isTargetPath 判断是否是目标目录
// 1. 目录；2. 非当前目录；3. 不能是特殊目录，例如：.git；4. 目录名不能包含空格
func isTargetPath(p string, info os.FileInfo) bool {
	if !info.IsDir() {
		return false
	}

	if p == "." {
		return false
	}

	if strings.HasPrefix(p, ".") {
		return false
	}

	if strings.Contains(p, " ") {
		return false
	}
	return true
}

// isTargetFile 判断是否是目标文件
// 1. 非目录；2. 文件扩展名为 .md；3. 文件名不能包含空格
func isTargetFile(file os.DirEntry) bool {
	if file.IsDir() {
		return false
	}
	if strings.Contains(file.Name(), " ") {
		return false
	}
	if !strings.HasSuffix(file.Name(), ".md") {
		return false
	}
	return true
}
