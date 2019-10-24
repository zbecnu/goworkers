package slave

import "runtime"

type slave struct {
	ch chan interface{}
}

func (s *slave) close (){
	close(s.ch)
}

func NewSlave(w func(interface{}))(s slave){

	s.ch = make(chan interface{},1)

	go func() {

		var job interface{}
		for job= range s.ch{
			w(job)
		}
	}()

	return s
}

type Pool struct {
	sv [] slave
	n int
}

func NewPool(workers int ,w func(interface{}))( p Pool){
			if w == nil{
		return
	}

	if workers <=0{

		workers= runtime.GOMAXPROCS(0)
	}

	p.n=workers
	p.sv = make([]slave,p.n,p.n)

	for i := 0; i < p.n; i++ {
		p.sv[i] = NewSlave(w)
	}
	return
}


func(p *Pool) Serve(w interface{}){
	i:=0
	for{
		//遍历所有所有的channal 直到任务可以发送给所有空闲的channel
		select {
		case p.sv[i].ch<-w:
			return
		default:
			i++
			if i==p.n{
				i=0
			}

		}
	}
}

func (p *Pool) ServeNonStop( w interface{}) bool  {
	i:=0

	for i<p.n{
		select {
		case p.sv[i].ch<-w:
			return true
		default:
			i++

		}
	}

	return false
}

func(p *Pool) Close(){
	for i := 0; i<p.n;i++  {
		p.sv[i].close()

	}
}