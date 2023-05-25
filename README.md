# gitbook-summary-generator

`gitbook-summary-generator` 是一款自动生成 Gitbook SUMMARY.md 的工具。

## 下载

```bash
go install github.com/poneding/gitbook-summary-generator@latest
```

## 使用

### 查看帮助信息

```bash
$ gitbook-summary-generator -h
gitbook-summary-generator is a tool to generate summary from Gitbook context files. version: v1.1.0

Usage:
  gitbook-summary-generator [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  start       Start to generate summary from Gitbook context files
  version     Version

Flags:
  -h, --help     help for gitbook-summary-generator
  -t, --toggle   Help message for toggle

Use "gitbook-summary-generator [command] --help" for more information about a command.

$ gitbook-summary-generator start -h
Start to generate summary from Gitbook context files.

Usage:
  gitbook-summary-generator start [flags]

Flags:
  -d, --dir string             Gitbook directory (default ".")
  -f, --force                  Force update summary file.
  -h, --help                   help for start
      --ignored-dirs strings   Ignore directories, multiple directories seperated by comma.
      --readme-title string    Gitbook readme title
      --summary-file string    Gitbook summary file (default "./SUMMARY.md")
      --summary-title string   Gitbook summary title
```

### 命令自动补全

```bash
# bash
source <(gitbook-summary-generator completion bash)

# zsh
source <(gitbook-summary-generator completion zsh)
```

### 查看版本信息

```bash
gitbook-summary-generator version
```

### 生成 SUMMARY

```bash
gitbook-summary-generator start -f
```

参数说明：

1. 使用 `-f` 或 `--force` 强制更新 SUMMARY.md。当文件不存在时，默认直接生成，文件存在时，默认不更新；
2. 使用 `-d` 或 `--dir` 指定读取目录，默认读取当前目录；
3. 使用 `-summary-file` 指定 Summary 生成文件，默认在当前目录下生成 `SUMMARY.md`；
4. 使用 `-readme-title` 指定 Summary 文件标题，默认为读取目录名；
5. 使用 `-summary-title` 指定 Readme 文件标题，默认为读取目录名；
6. 使用 `-ignored-dirs` 指定忽略的目录，多个目录使用 `,` 分隔。

```bash
gitbook-summary-generator start -f \
  -d ./Notes \
  --summary-file ./Notes/SUMMARY.md \
  --summary-title Notes \
  --readme-title Notes \
  --ignored-dirs draft,tmp
```
