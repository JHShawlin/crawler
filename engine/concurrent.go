package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
	ItemChan  chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	//in := make(chan Request)
	out := make(chan ParseResult)
	c.Scheduler.Run()

	for i := 0; i < c.WorkCount; i++ {
		createWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
	}

	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			tmp := item
			go func() {
				c.ItemChan <- tmp
			}()
		}

		for _, request := range result.Requests {
			if !isDuplicate(request.Url) {
				c.Scheduler.Submit(request)
			}
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}
