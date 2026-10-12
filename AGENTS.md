# AGENTS.md

本文件是本仓库的约定说明，供 AI 编码代理（以及新加入的人）在动手前阅读。

## 项目是什么

LeetCode 刷题仓库，**纯 Go**，**零第三方依赖**。

| 目录 | 用途 | 状态 |
|---|---|---|
| `problems1/` | 第一轮练习，107 个题目包 | **已归档**，除明确要求外不要改动 |
| `problems2/` | 第二轮练习 | **新题目一律写在这里** |

- 模块路径：`github.com/bianzhishui/ti_gr`（见 `go.mod`）
- Go 版本：`1.25.5`
- 无 `go.sum`，`go.mod` 里没有 `require`

> 历史：本仓库曾同时包含 Rust 实现（`Cargo.toml`、`*.rs`，且代码放在 `src/` 下）。这些内容已全部移除，**请勿重新引入**。

## 目录布局

每题一个独立目录，**目录名即包名**：

```
problems<N>/
└── t_<三位序号>_<题号或题集>_<snake_case 题目名>/
    ├── README.md      # 题面：LeetCode 链接 + 原题描述
    ├── xxx.go         # 解法
    └── xxx_test.go    # 测试，与解法文件同名 + _test
```

真实例子：

```
problems1/t_001_704_binary_search/
├── README.md
├── search.go
└── search_test.go
```

命名细节：

- **三位序号是顺序编号，不是题号**：`t_001_704_binary_search` = 第 1 题、LeetCode 704。
- **序号每轮独立编号**：`problems2/` 从 `t_001` **重新开始**，不与 `problems1/` 连续。因此两轮可能出现形近同名的目录（如 `problems1/t_001_704_binary_search` 与 `problems2/t_001_......`），**靠 `problemsN/` 目录区分**，在根 README 索引中登记时注意标注轮次。
- 题号用 LeetCode 官方数字；剑指 Offer 类用题集标识，如 `t_036_offer05_...`、`t_038_offer58ii_...`。
- 题目名用 snake_case。
- `t_999_` 前缀留给非 LeetCode 的杂项（如 `problems1/t_999_lru`）。
- 根 `README.md` 的索引里前两条写作 `t_01` / `t_02`，与实际目录 `t_001` / `t_002` 不一致——**以目录为准**，新条目按三位补零书写。

## 代码约定

- **包名 = 目录名**，例如 `package t_001_704_binary_search`。
  唯一例外：`problems1/t_999_lru` 是 `package main` 的可执行示例。
- 解法函数**默认不导出**（首字母小写），如 `search`、`minSubArrayLen`。
- **一题多解**：同一题的不同思路并列写在同一文件里，按由浅入深排列，**最基础的实现不带后缀**，后续用数字后缀：

  ```go
  func search(nums []int, target int) int { /* 循环遍历 */ }
  func search1(nums []int, target int) int { /* 二分：左闭右闭 */ }
  func search2(nums []int, target int) int { /* 二分：左闭右开 */ }
  ```

  仓库现有 60 多个函数采用这种 `xxx1` / `xxx2` / `xxx3` 后缀。
- **方法名本身以数字结尾时，后缀改用下划线分隔**，避免 `combinationSum41` 这种易被误读成另一道题的写法：

  ```go
  func combinationSum4(nums []int, target int) int   { /* 解法 1 */ }
  func combinationSum4_1(nums []int, target int) int { /* 解法 2 */ }
  func combinationSum4_2(nums []int, target int) int { /* 解法 3 */ }
  ```

  测试函数同名：`TestCombinationSum4`、`TestCombinationSum4_1`。**不带数字的名字仍用 `xxx1` / `xxx2`。**
- **每道题默认建多个解法骨架**，数量按该题**实际可行的思路数**决定（通常 2–3 个）。不要为凑数加入不可行的思路（例如在 `target <= 1000` 的约束下写指数级暴力解）。
- **需要导出的例外**：设计类题目要按判题要求导出构造器与方法（`Constructor()`、`Constructor1()`、`NewStack()`、`GetLRUCache()`、`StrStr()` 等）；被其它包复用的工具函数也必须导出（`ReverseString`、`GetNext`）。
- 注释使用**中文**，说明思路与边界条件。可参考 `problems1/t_001_704_binary_search/search.go`。

## 跨包引用

复用别的题的解法时，写**完整模块路径**，不要用相对路径：

```go
import trs "github.com/bianzhishui/ti_gr/problems1/t_034_344_reverse_string"
```

`problems1` 中现有 5 处此类引用（t_035 / t_037 / t_038 / t_040 / t_044）。第二轮内部互相引用则用 `problems2/...`。

## 测试约定

- 测试文件与被测文件同名 + `_test.go`，测试函数名为 `TestXxx`。
- 仓库存在两种风格，**新增代码推荐 (b)**：

**(a) 传统单用例风格**（problems1 主流）——每个实现单独一个测试函数，`var (...)` 声明块 + `Logf`/`Fatalf`：

```go
func TestSearch1(t *testing.T) {
	var (
		nums     = []int{-1, 0, 3, 5, 9, 12, 18}
		target   = 9
		expected = 4
		actual   int
	)
	actual = search1(nums, target)
	if actual == expected {
		t.Logf("expected=%d, actual=%d", expected, actual)
	} else {
		t.Fatalf("expected=%d, actual=%d", expected, actual)
	}
	t.Log("done")
}
```

比较切片或结构体时用 `reflect.DeepEqual`。

**(b) 表驱动 + 子测试**（推荐）——一份用例表驱动全部实现，见 `problems1/t_106_376_wiggle_subsequence/wiggleMaxLength_test.go`：

```go
var cases = []struct {
	name     string
	nums     []int
	expected int
}{
	{"例1 完整摆动序列", []int{1, 7, 4, 9, 2, 5}, 6},
	// ...
}

func TestWiggleMaxLength(t *testing.T) { checkCases(t, wiggleMaxLength) }
func TestWiggleMaxLength2(t *testing.T) { checkCases(t, wiggleMaxLength2) }
```

- ⚠️ **`t.Logf` / `t.Fatalf` 必须传固定格式串**。Go 1.24+ 的 `printf` vet 检查会拒绝「已用 `Sprintf` 格式化好的变量」当格式串，直接导致整个包 `go test` 失败。已格式化的字符串请改用 `t.Log` / `t.Fatal`：

  ```go
  logprint := fmt.Sprintf("expected=%v, actual=%v", expected, actual)
  t.Log(logprint)    // ✅
  t.Fatalf(logprint) // ❌ vet: non-constant format string
  ```

- 用例要覆盖**题面全部示例**和边界情况（单元素、全相等、平台期、空输入等），不要只测第一个样例。

## 命令

```bash
gofmt -l .      # 格式检查，应无输出
gofmt -w .      # 格式化
go build ./...
go vet ./...    # 应零告警
go test ./...   # 应全绿
```

改动完成的**最低标准**：`gofmt -l .` 无输出、`go vet ./...` 零告警、`go test ./...` 全部通过。

> 沙箱提示：受限环境下 `go vet` / `go test` 可能因无法写默认构建缓存而失败（报 `operation not permitted`），把缓存指到工作区外即可：
> `GOCACHE=/tmp/gocache go test ./...`

## 新增一道题：协作流程

用户给出题目标题或 `https://leetcode.cn/problems/<slug>/` 链接后，**默认产出除解法实现之外的全部文件**，不要让用户再逐项交代：

1. **建目录**：`problems2/t_<序号>_<题号>_<snake_case 题目名>/`
   序号在 `problems2/` 内从 `t_001` 起顺序递增（每轮独立编号）。
2. **`README.md`** —— 题面，格式与 `problems1` 完全一致：第一个代码块放题目链接，第二个代码块放「题号 + 题目名 + 难度 + 描述 + **全部**示例 + 提示 + 进阶」原文（保留原文的序号与排版）。
   - 必须保留原文的**缩进与层级**：属于上一条的补充说明、列表项、解释行要按官方页面那样缩进，不要全部拍平到顶格（参考 `problems1/t_106_376_wiggle_subsequence/README.md`）
3. **`<方法名>.go`** —— **多个解法骨架**（默认行为）：
   - 只写包声明、每个解法的完整签名、`// 解法 N：<思路名>` 注释、`TODO: 实现`，函数体返回零值
   - 注释只写「思路名 + 一句话要点」，**不要写依赖具体实现才成立的复杂度结论**（例如「空间 O(min(m,n))」在没写明先交换两串时是错的）
   - 骨架数量按该题**实际可行的思路数**决定（通常 2–3 个）；只有一种可行思路就只建一个，**不要为凑数加入不可行思路**（例如指数级暴力解）
   - 命名遵循上文「一题多解」，方法名以数字结尾时用下划线（`combinationSum4` / `combinationSum4_1`）
4. **`<方法名>_test.go`** —— 表驱动子测试：
   - 一份 `cases` 表由全部解法共用；每个解法一个 `Test<方法名>`（后缀规则同解法命名），通过共享的 helper 跑表
   - 必须覆盖题面**全部示例** + 边界（单元素、不可达、约束上限等），总数不少于 6 个
   - **表格规模要让每一个骨架都能从容跑完**（整个包目标 1 秒内）：若某个骨架渐进复杂度明显更差，上限规模用例取到它能从容处理的大小，并在测试文件顶部注释里说明该取舍
   - **期望值必须独立核对**：用临时程序算（写在 `/tmp` 下，不要进仓库），不要凭记忆填；并确认答值不超过题目声明的数值范围
5. **验证脚手架本身**：`gofmt -l .` 无输出、`go build ./...` 与 `go vet ./...` 通过。
   此时 `go test` **必然失败**（解法是桩），这是**预期状态**——不要为了让测试变绿而替用户填实现。
6. **登记索引**：在根 `README.md` 的 `problems2` 分区补一行。
7. **独立验证**（批量生成时必做）：由另一个 agent 把包复制到 `/tmp`，自写一份正确的参考实现替换骨架，确认 `go test` 全绿——以此证明测试期望值本身正确，而不是写了一个永远失败的测试。

**不要代替用户实现解法**，除非用户明确要求。

## 不要做的事

- ❌ **不要引入第三方依赖**，保持零依赖。
- ❌ **不要重新引入 Rust / Cargo** 相关文件。
- ❌ **不要改动 `problems1/`** 里已归档的解法，除非任务明确要求。
- ❌ 不要在 `_test.go` 里留 `func main()` 做临时调试。（历史遗留：`problems1/t_008_844_backspace_string_compare/backspaceCompare_test.go` 和 `problems1/t_010_977_squares_of_a_sorted_array/sortedSquares_test.go` 里有这种代码，它只会编译成非 main 包里一个永不执行的函数。新代码不要模仿。）
- ❌ 不要把 `target/` 提交进来：它是本机外部 Rust 工具链（flycheck）自动生成的目录，内含 `Cargo.toml does not exist` 报错，**不是仓库内容**。

## 提交风格

提交信息用英文，采用 conventional 前缀（历史提交已统一）：`feat:` / `fix:` / `test:` / `style:` / `chore:` / `docs:`。

例：`feat: implement and test 376 wiggle subsequence`

## 已知遗留问题

| 位置 | 问题 |
|---|---|
| `problems1/t_999_lru/lru.go` | 全仓库唯一的 `package main` 命令（含 `func main`），无 README、无测试 |
| 根 `README.md` | 索引仅 105 条，缺 `t_106_376_wiggle_subsequence` 与 `t_999_lru`；前两条写作 `t_01` / `t_02` 与目录不符 |
| 上述 2 个测试文件 | 残留 `func main()` 调试代码 |
| 定期出现 | 工作区可能被外部工具重建出空的 `target/` 目录 |
