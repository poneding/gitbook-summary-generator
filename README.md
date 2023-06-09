# gitbook-summary-generator

English | [中文](./README_zh.md)

`gitbook-summary-generator`  is a tool that automatically generates Gitbook SUMMARY.md.

## Download

```bash
go install github.com/poneding/gitbook-summary-generator@latest
```

## Usage

### View help information

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

### Command auto-completion

```bash
# bash
source <(gitbook-summary-generator completion bash)

# zsh
source <(gitbook-summary-generator completion zsh)
```

### View version

```bash
gitbook-summary-generator version
```

### Generate SUMMARY

```bash
gitbook-summary-generator start -f
```

Arguments description:

1. Use `-f` or `--force` to force an update of the SUMMARY.md. When the file does not exist, it is generated directly by default, and when the file exists, it is not updated by default;
2. Use `-d` or `--dir` to specify the read directory, and the current directory is read by default;
3. Use `-summary-file` to specify the Summary generated file, and 1 is generated in the current directory by default;
4. Use `-readme-title` to specify the Summary file title, which defaults to reading the directory name;
5. Use `-summary-title` to specify the Readme file title, which defaults to reading directory names;
6. Use `-ignored-dirs` to specify ignored directories, and multiple directories to separate by `,` .

```bash
gitbook-summary-generator start -f \
  -d ./Notes \
  --summary-file ./Notes/SUMMARY.md \
  --summary-title Notes \
  --readme-title Notes \
  --ignored-dirs draft,tmp
```
