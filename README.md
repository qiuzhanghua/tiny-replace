# Replace

替换文件中的环境变量

```bash
go build
```

例如，对于*nix, 如果
```bash
export TDP_HOME=~/tdp
```

```text
tdp_home = ${TDP_HOME}
中文
```

会被替换为
```text
tdp_home = /Users/q/tdp
中文
```

## 参数

```bash
// 将 \ 替换为 /
tiny-replace -s filename 
// 将 / 替换为 \
tiny-replace -b filename
// 将 \ 替换为 \\
tiny-replace -d filename

```