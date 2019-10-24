package main

import "fmt"

func main() {
	//var funcs map[string]interface{}
	//funcs = make(map[string]interface{})
	//
	//for k, v := range funcs {
	//	fmt.Println(k)
	//	fmt.Println(v)
	//
	//}

	ch := make(chan interface{})

	go func() {
		for j := range ch {
			if v,ok:= j.(string);ok{
				fmt.Println(v)
			}
			fmt.Println("zhangbign")
		}
	}()


	ch <- "wangdachui"
	ch <- 22
	close(ch)


}



func tt()(map[string]interface{},error){
	return nil,nil
}