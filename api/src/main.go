package main    // 包名，同一个包的多个go文件组成一个程序单元。路径名需要与包名相同
 
import "fmt"    // 引入依赖
 
func main(){    // 每一个go项目都必须有一个main方法作为程序入口
	fmt.Printf("Hello World")
}