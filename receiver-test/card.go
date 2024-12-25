package receivertest

import "fmt"

// Cart represents a shopping cart.
type Cart struct {
	Items []string
}

// AddItem adds an item to the cart.
func (c *Cart) AddItem(item string) error {
	if item == "" {
		return fmt.Errorf("cannot add an empty item to the cart")
	}
	c.Items = append(c.Items, item)
	return nil
}

// PrintItems เป็น receiver function สำหรับแสดงรายการสินค้าทั้งหมดในตะกร้า
func (c *Cart) PrintItems() {
	fmt.Println("Items in the cart:")
	for _, item := range c.Items {
		fmt.Println("- ", item)
	}
}

func CartX() error {
	// สร้าง instance ของ Cart
	myCart := Cart{}

	// เพิ่มสินค้าในตะกร้า
	myCart.AddItem("Apple")
	myCart.AddItem("Banana")
	myCart.AddItem("Orange")

	// แสดงรายการสินค้าในตะกร้า
	myCart.PrintItems()

	return nil
}
