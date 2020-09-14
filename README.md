# Selection

Golang implementation of **Introselect(Introspective Selection)**, which can get the k-th smallest value of a slice.

# Usage

Download package with `go get github.com/leiysky/selection`, and make sure your Golang version is `>=1.9`.

Example:
```go
import "github.com/leiysky/selection"
import "fmt"

type SortBy []int

func (a SortBy) Len() { return len(a) }
func (a SortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func main() {
  data := SortBy{3, 2, 1, 4, 4}
  result := selection.Select(data, 1) // data[result] = 1
  result = selection.Select(data, 4) // data[result] = 4
  result = selection.Select(data, 5) // data[result] = 4
}
```

# Contact
Email: leiysky@outlook.com

If you have any question, please open an issue.