# encryptLR

A minimal Go module for decoding an encoded sequence of relative digit relationships using characters `'L'`, `'R'`, and `'='`.

## 📦 Description

The `encryptLR` package provides a function `decodeLR(encoded string) string` that takes a string representing a sequence of transitions between digits, and decodes it into the lexicographically smallest digit string that satisfies those transitions and has the minimum sum of digits.

### Input Characters
- `'L'`: The previous digit must be **greater than** the next digit.
- `'R'`: The previous digit must be **less than** the next digit.
- `'='`: The previous digit must be **equal to** the next digit.

### Example
```go
decodeLR("LLRR=") // returns "210122"
```

---

## 📦 Installation

```bash
go get github.com/andreaharris-g/encryptLR
```

---

## 📄 Usage

```go
package main

import (
	"fmt"
	"encryptLR"
)

func main() {
	fmt.Println(encryptLR.DecodeLR("RRL=R")) // Output: 012001
}
```
---
## 🧠 Algorithm
- Uses dynamic programming to track all possible digit combinations at each position.
- Among valid sequences, selects the one with the lowest digit sum.
- Maintains only valid transitions (L, R, =) between digits.
---
## 🧪 Running Unit Tests
```shell
go test ./...
```
---
## 📜 License

MIT License

---
## 🙋‍♂️ Author

Developed by [Duckdev84](https://linktr.ee/duckdev84)
