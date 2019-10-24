package parallel

import "sync"

type Func func()error

func Run(functions ...Func) chan error{
	errs := make(chan error,len(functions))

	var wg sync.WaitGroup
	wg.Add(len(functions))

	go func(errs chan error) {
		wg.Wait()
		close(errs)
	}(errs)

	for _,fn:= range functions{

		go func(fn Func,errs chan error) {

			defer wg.Done()

			errs<- fn()


		}(fn,errs)
	}

	//wg.Wait()

	return errs
}