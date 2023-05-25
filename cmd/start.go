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

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/poneding/gitbook-summary-generator/pkg/generator"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start to generate summary from Gitbook context files",
	Long:  `Start to generate summary from Gitbook context files.`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	ignoredDirsAbs := []string{}
	for _, ignoredDir := range ignoredDirs {
		if strings.HasPrefix(ignoredDir, "~") {
			ignoredDir = strings.Replace(ignoredDir, "~", os.Getenv("HOME"), 1)
		}
		ignoredDirAbs, _ := filepath.Abs(ignoredDir)
		ignoredDirsAbs = append(ignoredDirsAbs, ignoredDirAbs)
	}

	se := generator.NewSummaryGenerator(&generator.GeneratorOption{
		Path:         dir,
		SummaryTitle: summaryTitle,
		ReadmeTitle:  readmeTitle,
		IgnoredDirs:  ignoredDirsAbs,
	})
	summary := se.Generate()

	// 如果不存在 SUMMARY.md 文件，则创建并写入
	// 如果存在 SUMMARY.md 文件，则根据 --force 参数决定是否覆盖
	writeToSummaryFile := true
	if _, err := os.Stat(summaryFile); err == nil {
		if !*force {
			writeToSummaryFile = false
		}
	}
	if writeToSummaryFile {
		err := os.WriteFile(summaryFile, []byte(summary), 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println("Generated " + summaryFile)
		}
	} else {
		fmt.Println("Generated SUMMARY:")
		fmt.Println(summary)
	}
}

var (
	// 目标目录，默认使用当前目录
	dir string
	// 生成的 Gitbook Summary 文件名，默认 ./SUMMARY.md
	summaryFile string
	// 生成的 Gitbook Summary 文件使用的 Title，默认使用父目录名
	summaryTitle string
	// Gitbook 中 README.md 的 Title，Title，默认使用父目录名
	readmeTitle string
	// 忽略的目录，多目录以逗号分隔
	ignoredDirs []string
	// 当存在 summaryFile 时，是否直接更新文件
	force *bool = func() *bool {
		var val = false
		return &val
	}()
)

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startCmd.Flags().StringVarP(&dir, "dir", "d", ".", "Gitbook directory")
	startCmd.Flags().StringVar(&summaryTitle, "summary-title", "", "Gitbook summary title")
	startCmd.Flags().StringVar(&readmeTitle, "readme-title", "", "Gitbook readme title")
	startCmd.Flags().StringVar(&summaryFile, "summary-file", "./SUMMARY.md", "Gitbook summary file")
	startCmd.Flags().StringSliceVar(&ignoredDirs, "ignored-dirs", []string{}, "Ignore directories, multiple directories seperated by comma.")
	startCmd.Flags().BoolVarP(force, "force", "f", *force, "Force update summary file.")
}
