package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		in = stage(in)
	}

	outChanel := make(Bi)
	go func() {
		defer close(outChanel)
		for {
			select {
			case <-done:
				return
			case value, ok := <-in:
				if !ok {
					return
				}
				outChanel <- value
			}
		}
	}()
	return outChanel
}
