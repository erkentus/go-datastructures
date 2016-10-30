# go-datastructures

### Installation

The normal way you would use golang packages

```sh
$ go get github.com/erkentus/go-datastructures
```

### Data Structures

#### Segment Tree

https://en.wikipedia.org/wiki/Segment_tree

#### Example

```go
package main
import (
    "fmt"
    "github.com/erkentus/go-datastructures/segment_tree"
)

func main(){
    tree := segment_tree.New([]int{1,2,3,4,5,6})
    //minimum element in range [2..4]
    val, err := tree.RangeMinQuery(2,4)
    fmt.Println(val) //3
    
    //add 3 to all elements in range 1..5
    err := tree.RangeAdd(1,5,3)
    newVal, err := tree.RangeMinQuery(2,4)
    fmt.Println(newVal) //6
}
```