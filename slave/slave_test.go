package slave

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestServeNoneStop_SlavePool(t *testing.T){
	sp:= NewPool(2, func(_ interface{}) {
		time.Sleep(time.Second)
	})


	if sp.ServeNonStop(nil) ==false{
		t.Fatal("no return  values")
	}

	if sp.ServeNonStop(nil) ==false{
		t.Fatal("no return  values")
	}

	if sp.ServeNonStop(nil) ==false{
		t.Fatal("no return  values")
	}
}

func executeServe(p *Pool,rounds int){
	for i:=0;i<rounds ;i++ 	 {
		p.Serve(i)

	}
}

func  TestServe_SlavePool(t *testing.T)  {
	ch := make(chan int ,1)
	counter := uint32(0)

	locker := sync.Mutex{}
	sp:= NewPool(0, func(obj interface{}) {
		locker.Lock()


		counter++

		//c:= counter+1
		//if t:= atomic.AddUint32(&counter,1);t!=c{
		//	panic("does  not match ")
		//}

		locker.Unlock()
		ch<-obj.(int)


	})

	rounds :=10000

	go executeServe(&sp,rounds)


	i:=0
	for i<rounds{
		select {
			case <-ch:
				i++
		}
	}

	if i!= int(counter){

	}
	sp.Close()
}

func TestPanic(t *testing.T){
	fmt.Println("begin")
	defer func() {
		p:=recover()
		fmt.Println(p)
		fmt.Println(reflect.TypeOf(p))
	}()
	panic(errors.New("this is en error"))
	fmt.Println("end")

}