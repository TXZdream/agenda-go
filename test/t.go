package main
import (
	"github.com/txzdream/agenda-go/entity/service"
	"github.com/txzdream/agenda-go/entity/model"
	"fmt"
)

func main() {
	var k service.Service
	service.StartAgenda(&k)
	fmt.Println(k)
	fmt.Println(k.AgendaStorage)
	aa := []model.User{}
	aa = append(aa, model.User{"sads", "SDfd", "asds", "sadsd"})
	k.AgendaStorage.Users = aa
	fmt.Println(k.AgendaStorage)

	var m service.Service
	service.StartAgenda(&m)
	fmt.Println(m)
	fmt.Println(m.AgendaStorage)
	// fmt.Println(k.CurrentUser.UserName == "")
	//  ------------------------------------
	// a := []int{1, 2, 3, 4, 5, 6, 7, 8,9}
	// for index, t := range a {
	// 	if t == 9 {
	// 		a = append(a[:index], a[index+1:]...)
	// 		break
	// 	}
	// }
	// fmt.Println(a)
}
// 开启读取数据--》获取当前用户名--》是否登陆--》再获取之后的命令--》进行之后的操作--》退出并写入文件