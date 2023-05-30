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
	"os"
	"path/filepath"
	"strings"
)

var hasMdFileMap = make(map[string]bool)

// hasMdFile 判断目录下是否有合格的 .md 文件
func hasInvaliadMdFile(dir string) bool {
	if b, ok := hasMdFileMap[dir]; ok {
		return b
	}
	dirEntries, _ := os.ReadDir(dir)

	var subDirs, subFiles []os.DirEntry
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			subDirs = append(subDirs, dirEntry)
		} else {
			subFiles = append(subFiles, dirEntry)
		}
	}

	for _, file := range subFiles {
		if filepath.Ext(file.Name()) == ".md" {
			hasMdFileMap[dir] = true
			return true
		}
	}

	for _, subDir := range subDirs {
		sub := filepath.Join(dir, subDir.Name())
		if hasInvaliadMdFile(sub) {
			hasMdFileMap[dir] = true
			hasMdFileMap[sub] = true
			return true
		}
	}
	// hasMdFileMap[dir] = false
	return hasMdFileMap[dir]
}

// isValidDir 判断是否是有效的目录
func isEffectiveDir(dir os.DirEntry) bool {
	return dir.IsDir()
}

// isEffectiveMdFile 判断是否是有效的 markdown 文件
// 1. 不是目录；2. 不是 README.md 或 SUMMARY.md；3. 是 .md 文件
func isEffectiveMdFile(file os.DirEntry) bool {
	if file.IsDir() {
		return false
	}
	if file.Name() == "README.md" || file.Name() == "SUMMARY.md" {
		return false
	}
	return filepath.Ext(file.Name()) == ".md"
}

// pathBaseName 返回目录的基础名称
func pathBaseName(path string) string {
	abs, _ := filepath.Abs(path)
	if abs == "/" {
		return ""
	}
	return filepath.Base(abs)
}

// pathIgnored 判断目录是否被忽略
func pathIgnored(ignoredDirs []string, path string) bool {
	pathAbs, _ := filepath.Abs(path)

	for _, ignoredDir := range ignoredDirs {
		trimed := strings.TrimPrefix(pathAbs, ignoredDir)
		if len(trimed) != len(pathAbs) {
			if len(trimed) == 0 || trimed[0] == '/' {
				return true
			}
		}
	}
	return false
}

// isRootPath 判断是否是 Gitbook 默认的根目录 "."
func isRootPath(path string) bool {
	return path == "."
}
