# Sort包

## sort.Search

### 作用

前提: 一个已经排序的数组

该函数使用二分查找的方法，会从[0, n)中取出一个值index，index为[0, n)中最小的使函数f(index)为True的值，并且f(index+1)也为True。
如果无法找到该index值，则该方法为返回n。

### example

```go
func Search(n int, f func(int) bool) int
```

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    // 定义一个已排序的切片
    nums := []int{1, 3, 5, 7, 9}

    // 使用sort.Search函数查找元素5的索引
    index := sort.Search(len(nums), func(i int) bool {
        return nums[i] >= 5
    })

    // 输出结果
    if index < len(nums) && nums[index] == 5 {
        fmt.Printf("元素5的索引为：%d\n", index)
    } else {
        fmt.Println("未找到元素5")
    }
}
```

结果:

```powershell
元素5的索引为：2
```

