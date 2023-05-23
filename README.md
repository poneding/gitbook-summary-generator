# gitbook-summary-generator

`gitbook-summary-generator` 是一款自动生成 Gitbook SUMMARY.md 的工具。

## 下载

```bash
go install github.com/poneding/gitbook-summary-generator@latest
```

## 使用

```bash
$ gitbook-summary-generator -h
Usage of gitbook-summary-generator:
  -f string
    	Summary output file (default "./SUMMARY.md")
  -p string
    	Path to directory (default ".")
  -t string
    	Summary title (default "SUMMARY")
```

> 使用 `-p` 指定读取目录，默认读取当前目录；
> 使用 `-f` 指定 Summary 生成文件，默认在当前目录下生成 `SUMMARY.md`；
> 使用 `-t` 指定 Summary 文件标题，默认为读取目录名；
> 当读目录或其自目录下包含有效的 `Markdown(.md)` 文件，并且不存在 `README.md` 文件时，将自动创建 `README.md` 文件。

## 示例

在当前目录生成 SUMMARY.md 文件

```bash
gitbook-summary-generator
```

指定参数生成 SUMMARY.md 文件

```bash
gitbook-summary-generator -p ./Notes -f ./Notes/SUMMARY.md -t Notes
```
