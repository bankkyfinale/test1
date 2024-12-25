package receivertest

import (
	"fmt"
	"sync"
)

// Cart represents a shopping cart.
type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value++
}

func CountGoroutine() error {
	//
	var wg sync.WaitGroup
	// สร้าง instance ของ Counter
	myCounter := Counter{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			myCounter.Increment() // เรียกใช้ method Increment ในแต่ละ goroutine
		}()
	}

	// รอ goroutines ทั้งหมดเสร็จ
	wg.Wait()

	// ผลลัพธ์: 5 เพราะมีการ increment ค่า 5 ครั้ง
	fmt.Println("Counter value:", myCounter.Value)

	return nil
}
