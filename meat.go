package main

import (
	"errors"
	"time"
)

var (
	MeatTypeErr = errors.New("meat type error")
)

const (
	MeatTypeBeef = iota
	MeatTypePork
	MeatTypeChicken
)

type Meat struct {
	Type int
}

func NewMeat(meatType int) *Meat {
	return &Meat{Type: meatType}
}

// Name 名稱
func (m *Meat) Name() string {
	switch m.Type {
	case MeatTypeBeef:
		return "牛肉"
	case MeatTypePork:
		return "豬肉"
	case MeatTypeChicken:
		return "雞肉"
	default:
		panic(MeatTypeErr)
	}
}

// Process 加工
func (m *Meat) Process() {
	switch m.Type {
	case MeatTypeBeef:
		time.Sleep(1 * time.Second)
	case MeatTypePork:
		time.Sleep(2 * time.Second)
	case MeatTypeChicken:
		time.Sleep(3 * time.Second)
	default:
		panic(MeatTypeErr)
	}
}
