package main
import (
	model "github.com/txzdream/agenda-go/entity/model"
	service "github.com/txzdream/agenda-go/entity/service"
	"fmt"
)

func main() {
	var k service.Service
	service.StartAgenda(&k)
	fmt.Println(k)
	fmt.Println(k.AgendaStorage)
	k.AgendaStorage.CurrentUser = model.User{"sads", "SDfd", "asds", "sadsd"}
	fmt.Println(k.AgendaStorage)

	var m service.Service
	service.StartAgenda(&m)
	fmt.Println(m)
	fmt.Println(m.AgendaStorage)
	// fmt.Println(k.CurrentUser.UserName == "")
	//  ------------------------------------
	a := []int{1, 2, 3, 4, 5, 6, 7, 8,9}
	for index, t := range a {
		if t == 9 {
			a = append(a[:index], a[index+1:]...)
			break
		}
	}
	fmt.Println(a)
}
// 开启读取数据--》获取当前用户名--》是否登陆--》再获取之后的命令--》进行之后的操作--》退出并写入文件