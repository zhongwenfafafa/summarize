# 排序算法
* 排序算法的好坏指标
    - 时间复杂度（最好、平均、最坏）
    - 是否为原地排序算法（空间复杂度为O(1)）
    - 是否为稳定排序算法（排序后不会改变相同元素原来的顺序）    
    
* 排序算法有很多种，但是常用的就那么几种，根据时间复杂度把排序算法分成三类
    - 冒泡排序、插入排序、选择排序 它们的时间复杂度都是O(n2)
    - 归并排序、快速排序 它们的时间复杂度是O(n*logN)
    - 桶排序、计数排序、基数排序 它们的时间复杂度O(n)
    ![alert 图片](https://static001.geekbang.org/resource/image/fb/cd/fb8394a588b12ff6695cfd664afb17cd.jpg)    

## 时间复杂度为O(n2)的常见排序算法
* 在实际使用中，插入排序是使用最多的O(n2)时间复杂度的排序算法，其主要原因是它在每次交换过程中，交换的元素最少
   
### 冒泡排序
* 冒泡排序是一种基于比较的排序算法，每次冒泡的过程，就是比较相邻两个元素的大小，决定是否交换位置      
    - 最好时间复杂度是O(n)、最坏时间复杂度是O(n2)、平均时间复杂度O(n)
    - 是原地排序算法
    - 是稳定排序算法
```
func BubbleSort(a []int, n int) {
    if  n <= 1 {
        return
    }
    
    for i := 0; i < n; i++ {
        flag := false
        for j := 0; j < n-i-1; j++ {
            if  a[j+1] > a[j] {
                a[j+1], a[j] = a[j], a[j+1]
                flag = true
            }
        }
        
        if !flag {
            return
        }
    }
}
```

### 插入排序
* 插入排序也是基于比较的排序算法，它把要排序的元素分成两个区间，一个有序区间、一个无序区间
* 每次排序的过程就是拿无序区间的元素到有序区间比较，并插入到合适的位置（需要移动有序区间元素）
    - 最好时间复杂度是O(n)、最坏时间复杂度是O(n2)、平均时间复杂度O(n)
    - 是原地排序算法
    - 是稳定排序算法
```$xslt
func InsertSort(a []int, n int) {
    if  n <= 1 {
        return
    }
    
    for i := 1; i < n; i++ {
        v := a[i]
        j := i - 1
        for ; j >= 0; j-- {
            if  a[j] < v {
                a[j+1] = a[j]
            } else {
                break
            }
        }
        
        a[j+1] = v
    }
}
```

### 选择排序
* 选择排序也是基于比较的排序算法，它和插入排序一样，把要排序的元素分成两个区间，有序区和无序区
* 排序的过程是每次将无序区的元素依次交换选出最小的一个放在有序区的末尾
    - 最好时间复杂度是O(n)、最坏时间复杂度是O(n2)、平均时间复杂度O(n)
    - 是原地排序算法
    - 不是稳定排序算法

```$xslt
func SelectSort(a []int, n int) {
    if  n <= 1 {
        return
    }
    
    for i := 0; i < n; i++ {
        minIndex := i
        for j := i + 1; j < n; j++ {
            if  a[j] < a[minIndex] {
                minIndex = j
            }
        }
        a[minIndex], a[i] = a[i], a[minIndex] 
    }
}
```    
