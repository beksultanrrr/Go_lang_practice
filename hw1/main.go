package main

import "fmt"

type Employee interface {
	GetPosition() string
	SetPosition(string)
	GetSalary() float64
	SetSalary(float64)
	GetAddress() string
	SetAddress(string)
}

type Manager struct {
	position string
	salary   float64
	address  string
}

func NewManager(position string, salary float64, address string) *Manager {
	return &Manager{position, salary, address}
}

func (m *Manager) GetPosition() string {
	return m.position
}

func (m *Manager) SetPosition(position string) {
	m.position = position
}

func (m *Manager) GetSalary() float64 {
	return m.salary
}

func (m *Manager) SetSalary(salary float64) {
	m.salary = salary
}

func (m *Manager) GetAddress() string {
	return m.address
}

func (m *Manager) SetAddress(address string) {
	m.address = address
}

type Developer struct {
	position string
	salary   float64
	address  string
}

func NewDeveloper(position string, salary float64, address string) *Developer {
	return &Developer{position, salary, address}
}

func (d *Developer) GetPosition() string {
	return d.position
}

func (d *Developer) SetPosition(position string) {
	d.position = position
}

func (d *Developer) GetSalary() float64 {
	return d.salary
}

func (d *Developer) SetSalary(salary float64) {
	d.salary = salary
}

func (d *Developer) GetAddress() string {
	return d.address
}

func (d *Developer) SetAddress(address string) {
	d.address = address
}

func main() {
	manager := NewManager("Manager", 50000, "123 Main St")
	developer := NewDeveloper("Developer", 60000, "456 Elm St")

	fmt.Println("Manager:")
	fmt.Println("Position:", manager.GetPosition())
	fmt.Println("Salary:", manager.GetSalary())
	fmt.Println("Address:", manager.GetAddress())

	fmt.Println("\nDeveloper:")
	fmt.Println("Position:", developer.GetPosition())
	fmt.Println("Salary:", developer.GetSalary())
	fmt.Println("Address:", developer.GetAddress())
}
