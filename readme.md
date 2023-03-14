# Tarsus Go 微服务模块

## Intro

心血来潮想着把Go 的微服务模块也写下，写了一天发现语法是真的奇特，不能像以前写 TypeScript 和 Java那样 随心所欲了

- 启动服务

````GO
package main

import (
 "tarsus/go/src/pkg"
 "tarsus/go/src/registry"
)

// 直接启动 Tarsus 微服务
func main() {
 pkg.Tarsus()
}


````

- 关键函数
 拆分 @Tarsus/GateWay传入的数据

````GO
func getData(buf string) []any {
 ret := make([]any, 0)
 init := 0
 start := strings.Index(buf, size[init])

 // 相当于while
 for {
  next_init := init + 1
  next := indexOf(buf, size[next_init], start)
  if next == -1 {
   // 是否分割完
   if start+len(size[init]) == len(buf) {
    break
   }
   // 切片查看是否是没切完的
   sub_pkg := buf[start : start+6]
   is_un_pkg := sub_pkg == size[init]+size[0]
   if is_un_pkg {
    un_pkg := buf[start+3 : len(buf)-3]
    args := getData(un_pkg)
    ret = append(ret, args)
   } else {
    arg := buf[start+3 : len(buf)-3]
    ret = append(ret, arg)
   }
   break
  } else {
   isObject := buf[start:start+6] == size[init]+size[0]
   if isObject {
    curr_end_str := size[len(size)-1] + size[init+1]
    end := strings.Index(buf, curr_end_str)
    un_pkg := buf[start+3 : end+3]
    args := getData(un_pkg)
    ret = append(ret, args)
    start = end + 3
   } else {
    println(start, " --- ", next, " --- ", len(buf))
    arg := buf[start+3 : next]
    println(arg)
    ret = append(ret, arg)
    start = next
   }
  }
  init++
 }
 return ret
}


````
