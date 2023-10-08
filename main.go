package main

import (
	"awesomeProject2/Domain"
	"awesomeProject2/Repository"
	// Драйвер для SQL Server
)

func main() {
	var repo Domain.EmployeeRepository = Repository.NewEmployeeRepositoryDB("server=localhost;user id=admin;password=123;database=EmployeesDB")

	var _ = repo.GetEmployeeByName("retard")

	var _ = repo.GetAllEmployees()

	var empToCreate = &Domain.Employee{
		FirstName:       "Daniel",
		LastName:        "Balenko",
		Phone:           "228-1488",
		CompanyId:       2,
		PassportType:    "Passport",
		PassportNumber:  "1337-666-999",
		DepartmentName:  "Cloud",
		DepartmentPhone: "22222",
	}

	repo.CreateEmployee(empToCreate)

	repo.DeleteEmployee(1012)
}
