# 🚀 LeetCode Solutions in Go

<div align="center">

[English](#english) | [中文](#中文)

</div>

---

## English

Go implementations of LeetCode problems. **276 problems solved** (97 Easy / 136 Medium / 43 Hard).

### 📊 Stats
- **Total**: 276 problems
- **Difficulty**: 97 Easy / 136 Medium / 43 Hard (15.6%)
- **Ranking**: Top 51.77% globally
- **Contest Rating**: 1498
- **Submissions**: 698 in the past year
- **Monthly Badges**: 6

### 📁 Project Structure
```
go-leetcode-solutions/
├── easy/          # Easy problems (97)
├── medium/        # Medium problems (136)
├── hard/          # Hard problems (43)
└── topics/        # Categorized by topics
    ├── dp/        # Dynamic Programming
    ├── tree/      # Binary Tree
    ├── backtrack/ # Backtracking
    └── graph/     # Graph Theory
```

### 🔥 Hard Problems Highlights (43 problems)
| #    | Problem                        | Category                       | Code                                        |
| ---- | ------------------------------ | ------------------------------ | ------------------------------------------- |
| 42   | Trapping Rain Water            | Two Pointers / Monotonic Stack | [Code](hard/trapping_rain_water.go)         |
| 23   | Merge k Sorted Lists           | Heap / Divide & Conquer        | [Code](hard/merge_k_sorted_lists.go)        |
| 84   | Largest Rectangle in Histogram | Monotonic Stack                | [Code](hard/largest_rectangle_histogram.go) |

### 🧪 Run Tests
```bash
# Install dependencies
go mod tidy

# Run all tests
go test ./...

# Run specific package tests
go test ./medium
go test ./hard
```

### 📚 Topics Coverage
| Topic               | Count | Representative Problems                     |
| ------------------- | ----- | ------------------------------------------- |
| Dynamic Programming | 35    | Knapsack, Stock Trading, LIS                |
| Backtracking        | 28    | Permutations, N-Queens, Combination Sum     |
| Binary Tree         | 30    | Level Order, Path Sum, LCA                  |
| Graph Theory        | 20    | Shortest Path, Topological Sort, Union Find |
| Greedy              | 25    | Interval Scheduling, Jump Game              |
| Others              | 138   | Array, Linked List, String, Hash Table      |

### 📝 Code Standards
- ✅ Follow Go official style (gofmt + golangci-lint)
- ✅ Unit tests with time/space complexity analysis
- ✅ Reusable algorithm templates

### 📧 Contact
- LeetCode: [@KaiserKun](https://leetcode.cn/u/KaiserKun/)
- GitHub: [@KaiserKun](https://github.com/KaiserKun)
- Email: coderxiao@126.com

### 📄 License
[MIT License](LICENSE)

---

## 中文

使用 Go 语言实现 LeetCode 算法题，已完成 **276 道**（简单 97 / 中等 136 / 困难 43）。

### 📊 数据统计
- **总题数**：276 道
- **难度分布**：简单 97 / 中等 136 / 困难 43（占比 15.6%）
- **全国排名**：Top 51.77%
- **竞赛分**：1498
- **年提交次数**：698 次
- **月度徽章**：6 次

### 📁 项目结构
```
go-leetcode-solutions/
├── easy/          # 简单题（97 道）
├── medium/        # 中等题（136 道）
├── hard/          # 困难题（43 道）
└── topics/        # 按主题分类
    ├── dp/        # 动态规划
    ├── tree/      # 二叉树
    ├── backtrack/ # 回溯
    └── graph/     # 图论
```

### 🔥 困难题精选（43 道）
| 题号 | 题目               | 分类          | 代码                                        |
| ---- | ------------------ | ------------- | ------------------------------------------- |
| 42   | 接雨水             | 双指针/单调栈 | [代码](hard/trapping_rain_water.go)         |
| 23   | 合并 K 个升序链表  | 堆/分治       | [代码](hard/merge_k_sorted_lists.go)        |
| 84   | 柱状图中最大的矩形 | 单调栈        | [代码](hard/largest_rectangle_histogram.go) |

### 🧪 运行测试
```bash
# 安装依赖
go mod tidy

# 运行所有测试
go test ./...

# 运行指定包测试
go test ./medium
go test ./hard
```

### 📚 主题分类
| 分类     | 题数 | 代表题目                           |
| -------- | ---- | ---------------------------------- |
| 动态规划 | 35   | 背包问题、股票买卖、最长递增子序列 |
| 回溯     | 28   | 全排列、N 皇后、组合总和           |
| 二叉树   | 30   | 层序遍历、路径总和、最近公共祖先   |
| 图论     | 20   | 最短路径、拓扑排序、并查集         |
| 贪心     | 25   | 区间调度、跳跃游戏                 |
| 其他     | 138  | 数组、链表、字符串、哈希表等       |

### 📝 代码规范
- ✅ 遵循 Go 官方风格（gofmt + golangci-lint）
- ✅ 每题配单元测试与复杂度分析
- ✅ 核心算法模板沉淀为工具函数

### 📧 联系方式
- LeetCode：[@KaiserKun](https://leetcode.cn/u/KaiserKun/)
- GitHub：[@KaiserKun](https://github.com/KaiserKun)
- 邮箱：coderxiao@126.com

### 📄 开源协议
[MIT License](LICENSE)
