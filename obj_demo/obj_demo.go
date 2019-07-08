package main

import "fmt"

// SavingBox : demo object
type SavingBox struct {
	money int
}

// NewBox Init SavingBox
func NewBox() *SavingBox {
	return new(SavingBox)
}

// Income add money to SavingBox
func (s *SavingBox) Income(amount int) {
	s.money += amount
}

// Break break SavingBox and get all money
func (s *SavingBox) Break() int {
	lastMoney := s.money
	s.money = 0
	return lastMoney
}

func main() {
	box := NewBox()
	box.Income(100)
	box.Income(200)
	box.Income(500)

	fmt.Printf("貯金箱を壊したら%d円出てきました。", box.Break())
}
