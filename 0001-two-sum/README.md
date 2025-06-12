### 📄 `two_sum/README.md`


# 🧮 Two Sum - LeetCode #1

This folder contains the Go solution for [LeetCode Problem #1: Two Sum](https://leetcode.com/problems/two-sum/), as part of my personal [LeetCode journey project](https://github.com/0x716/MyLeetcodeJourney).

## ✍️ Problem Description

Given an array of integers `nums` and an integer `target`, return **indices** of the two numbers such that they add up to `target`.

You may assume that each input would have **exactly one solution**, and you may not use the same element twice.

You can return the answer in **any order**.

### Example:

```

Input: nums = \[2, 7, 11, 15], target = 9
Output: \[0, 1]
Explanation: nums\[0] + nums\[1] == 9

````

---

## ✅ Solution Explanation

We use a hash map to store the value-to-index mapping as we iterate through the list.  
For each number, we calculate its complement and check if it already exists in the map.

### Time Complexity: `O(n)`  
### Space Complexity: `O(n)`

---

## 🧪 How to Test

This solution uses [`testify/assert`](https://github.com/stretchr/testify) for testing.

To run the tests: `go test`

---

## 📂 Files

| File              | Description                   |
| ----------------- | ----------------------------- |
| `two_sum.go`      | Core solution implementation  |
| `two_sum_test.go` | Unit tests using Go + Testify |
| `README.md`       | This file                     |

---

## 👨‍💻 Author

**Lin Jun** (GitHub: [@0x716](https://github.com/0x716))
This project is part of my [MyLeetcodeJourney](https://github.com/0x716/MyLeetcodeJourney), a personal learning archive for deepening my Go skills through solving algorithm problems.

---

## 📜 Disclaimer

All solutions are written for educational and personal learning purposes.
Please **do not copy directly** for submission or academic use without fully understanding the code.

---
