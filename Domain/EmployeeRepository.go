package Domain

type EmployeeRepository interface {
	// Получить сотруднка по его имени
	GetEmployeeByName(name string) *Employee

	// Получить всех сотрудников
	GetAllEmployees() []*Employee

	// Создать сотрудника
	CreateEmployee(emp *Employee)

	// Удалить сотрудника по ID
	DeleteEmployee(id int)
}
