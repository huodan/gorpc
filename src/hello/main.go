package main

import (
	"fmt"
	"utils/common"
)

func main() {

	var args = common.Args{16, 8}
	var result = common.Result{}
	var math_service common.MathService

	math_service.Add(&args, &result)
	fmt.Println("调用结果：", result.Value)

	math_service.Divide(&args, &result)
	fmt.Println("调用结果：", result.Value)

}
