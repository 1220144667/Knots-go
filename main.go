/*
 * @Author: 洪陪 hp2022a@163.com
 * @Date: 2025-04-19 15:57:11
 * @LastEditors: 洪陪 hp2022a@163.com
 * @LastEditTime: 2025-04-19 16:03:22
 * @FilePath: /knots-go/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

package main

import (
	"github.com/knots/knots-go/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
