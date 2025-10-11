# 🚀 LeetCode Solutions in Go

<div align="center">

[English](#english) | [中文](README_CN.md)

</div>

---

## English

Go implementations of LeetCode problems. **12 problems solved** (4 Easy / 7 Medium / 1 Hard).

### 📊 Stats

- **Total**: 12 problems
- **Difficulty**: 4 Easy / 7 Medium / 1 Hard (8.3%)
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

| #   | Problem                        | Category                       | Code                                        |
| --- | ------------------------------ | ------------------------------ | ------------------------------------------- |
| 42  | Trapping Rain Water            | Two Pointers / Monotonic Stack | [Code](hard/trapping_rain_water.go)         |
| 23  | Merge k Sorted Lists           | Heap / Divide & Conquer        | [Code](hard/merge_k_sorted_lists.go)        |
| 84  | Largest Rectangle in Histogram | Monotonic Stack                | [Code](hard/largest_rectangle_histogram.go) |

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
