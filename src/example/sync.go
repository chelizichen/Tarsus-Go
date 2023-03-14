package example

import (
	"log"
	"sync"
	"time"
)

var remine = 5
var mu sync.Mutex
var wg sync.WaitGroup

func seckill(){
	defer wg.Done()
	defer mu.Unlock()
	mu.Lock()
	if remine <= 0{
		return
	}
	time.Sleep(100*time.Millisecond)
	remine = remine -1;
	log.Println("seckill success")
}
// 协程秒杀案例
func main(){
	for i:=0;i<10;i++{
		wg.Add(1)
		go seckill()
	}
	wg.Wait()
	log.Println(remine)
}