package Repository

import (
	"awesomeProject2/Domain"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/microsoft/go-mssqldb"
)

type EmployeeRepositoryDB struct {
	db *sql.DB
}

func (repo *EmployeeRepositoryDB) GetAllEmployees() []*Domain.Employee {
	var query = "select * from EmployeesDB.dbo.Employees"
	var rows, err = repo.db.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			panic("Нет данных о рабочих")
		}
	}

	var emps []*Domain.Employee

	for rows.Next() {
		var emp Domain.Employee
		if err := rows.Scan(&emp.Id, &emp.FirstName, &emp.LastName, &emp.Phone, &emp.CompanyId, &emp.PassportType, &emp.PassportNumber, &emp.DepartmentName, &emp.DepartmentPhone); err != nil {
			panic("Ошибка при считывании работника")
		}
		emps = append(emps, &emp)
	}

	return emps
}

func (repo *EmployeeRepositoryDB) CreateEmployee(emp *Domain.Employee) {
	var query = "insert into EmployeesDB.dbo.Employees (FirstName, LastName, Phone, CompanyId, PassportType, PassportNumber, DepartmentName, DepartmentPhone) values (?, ?, ?, ?, ?, ?, ?, ?)"

	var _, err = repo.db.Exec(query, emp.FirstName, emp.LastName, emp.Phone, emp.CompanyId, emp.PassportType, emp.PassportNumber, emp.DepartmentName, emp.DepartmentPhone)

	if err != nil {
		panic("Ошибка при добавлении сотрудника - " + err.Error())
	}

}

func (repo *EmployeeRepositoryDB) DeleteEmployee(id int) {
	var query = "delete from EmployeesDB.dbo.Employees where Id = ?"

	var _, err = repo.db.Exec(query, id)

	if err != nil {
		panic("Ошибка при удалнии сотрудника")
	}
}

func NewEmployeeRepositoryDB(connectionString string) *EmployeeRepositoryDB {
	db, _ := sql.Open("mssql", connectionString)
	return &EmployeeRepositoryDB{db}
}

func (repo *EmployeeRepositoryDB) GetEmployeeByName(name string) *Domain.Employee {
	var query = "select Id, FirstName, LastName, Phone, CompanyId, PassportType, PassportNumber, DepartmentName, DepartmentPhone FROM EmployeesDB.dbo.Employees  WHERE FirstName = ?"
	var row = repo.db.QueryRow(query, name)

	var emp Domain.Employee
	var err = row.Scan(&emp.Id, &emp.FirstName, &emp.LastName, &emp.Phone, &emp.CompanyId, &emp.PassportType, &emp.PassportNumber, &emp.DepartmentName, &emp.DepartmentPhone)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			panic(fmt.Sprintf("Нет данных о работнике с именем %s", name))
			return nil
		}

	}

	return &emp
}
