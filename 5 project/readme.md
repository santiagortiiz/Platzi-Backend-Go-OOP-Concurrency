# Project description
Create a web server that will generate jobs
that a worker pool will be consuming.

# Hints
- El dispatcher recibe todos los jobs, se puede decir que es como el componente global
- Cada worker tiene su canal de jobs, y saben cual es el canal del disptacher, es decir el workerpool es el mismo canal para todos los workers.
- Cada worker esta enviando su canal al canal del dispatcher
- En la medida que el dispatcher recibe jobs este los va repartiendo entre los workers a través de sus canales


-El Dispacher obtendra una cola, que contedra una cola de cada uno de los workers.
-Una vez que el dispacher reciba un trabajo, esta agarrara de su cola a un worker, y mandara la tarea por ese medio
-El worker, procesara el trabajao y volvera a agregar su cola de trabajo a la cola del dispacher


# Challenge to send multiple requests
func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}