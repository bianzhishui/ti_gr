# problems2

第二轮练习（今后刷题）的代码目录。

- `problems1/` —— 第一轮练习，已归档，不再改动
- `problems2/` —— 第二轮练习，新题目都放这里

## 目录约定

沿用 `problems1/` 的组织方式，每题一个子目录：

```
problems2/
└── t_NNN_<题号>_<题目名>/
    ├── README.md     # 题面
    ├── xxx.go        # 解法
    └── xxx_test.go   # 测试
```

## 跨包引用

本仓库模块路径为 `github.com/bianzhishui/ti_gr`，第二轮内部互相引用时写：

```go
import trs "github.com/bianzhishui/ti_gr/problems2/t_001_704_binary_search"
```

引用第一轮已归档的解法时写：

```go
import trs "github.com/bianzhishui/ti_gr/problems1/t_034_344_reverse_string"
```
