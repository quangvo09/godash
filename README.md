# godash

### Group by a slice integer
```
import (
  "fmt"

  group_by "github.com/quangvo09/godash/group_by"
)

func main() {
  input := []int{1, 2, 3, 4, 2, 3, 4}
  output := group_by.SliceInt(input)
  fmt.Println(output)
}
```

```
map[1:[1] 2:[2 2] 3:[3 3] 4:[4 4]]
```

### Group by a slice string
```
import (
  "fmt"

  group_by "github.com/quangvo09/godash/group_by"
)

func main() {
  input := []string{"X", "Y", "Z", "Y", "Z", "Z"}
  output := group_by.SliceString(input)
  fmt.Println(output)
}
```

```
map[X:[X] Y:[Y Y] Z:[Z Z Z]]
```

### Group by a slice of struct
```
import (
  "fmt"

  group_by "github.com/quangvo09/godash/group_by"
)

type Student struct {
  Name string
  Age  int
}

func main() {
  intput := []Student{
    Student{"A", 1},
    Student{"B", 1},
    Student{"C", 2},
    Student{"D", 2},
  }

  output := group_by.SliceStable(intput, func(student interface{}) interface{} {
    return student.(Student).Age
  })
  
  fmt.Println(output)
}
```

```
map[1:[{A 1} {B 1}] 2:[{C 2} {D 2}]]
```