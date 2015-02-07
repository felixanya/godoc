// Author: Felix
// Email:  sugarmanman@163.com
// 正式接口

package game

// Go 的声明句法允许编组。单一的文注解可以引出一组相联的常量或变量。因为整组声明一起展现，注解可以很粗略
const (
	ldigits = "0123456789abcdef"
	udigits = "0123456789ABCDEF"
)

// Integer 测试函数返回一个Int
//	@Param   a 随时整形
func Integer(a int64) int {
     return 1
}
