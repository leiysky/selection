package selection

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

type testSlice []int

func (a testSlice) Len() int           { return len(a) }
func (a testSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a testSlice) Less(i, j int) bool { return a[i] < a[j] }

func TestSelection(t *testing.T) {
	s := testSlice{1, 2, 3, 4, 5}
	index := Select(s, 3)
	if s[index] != 3 {
		t.Fail()
	}
}
func TestSelectionWithDuplicate(t *testing.T) {
	s := testSlice{1, 2, 3, 3, 5}
	index := Select(s, 3)
	if s[index] != 3 {
		t.Fail()
	}
	index = Select(s, 5)
	if s[index] != 5 {
		t.Fail()
	}
}

func TestSelectionWithRandomCase(t *testing.T) {
	data := randomTestCase(100)
	index := Select(data, 50)
	value := data[index]
	sort.Stable(data)
	expected := data[49]
	if value != expected {
		t.Errorf("Expected: %v, Actual: %v\n%v", expected, value, data)
	}
}

func randomTestCase(size int) testSlice {
	data := make(testSlice, 0, size)
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		data = append(data, rand.Int()%100)
	}
	return data
}

// func TestSomething(t *testing.T) {
// 	data := testSlice{3, 5, 4, 3, 3, 1}
// 	fmt.Println(data, Select(data, 3))
// 	fmt.Println(data, Select(data, 2))
// 	sort.Sort(data)
// 	fmt.Printf("%v\n", data)
// }
