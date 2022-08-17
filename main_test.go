package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	tables := []struct{
		id int 
		dni string
		mockFunc func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id: 1,
			dni: "1",
			mockFunc: func(){
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id: 1,
						Position: "CEO",
					}, nil
				}
				GetPersonByDNI = func(id string) (Person, error) {
					return Person{
						DNI: "1",
						Name: "Hugo",
						Age: 34,
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Age: 34,
					Name: "Hugo",
					DNI: "1",
				},
				Employee: Employee{
					Id: 1,
					Position: "CEO",
				},
			},
		},
	}
	
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, test := range tables {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)

		if err != nil {
			t.Errorf("Error when getting Employee")
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error, got %d expected %d", ft.Age, test.expectedEmployee.Age)

		}

		GetEmployeeById = originalGetEmployeeById
		GetPersonByDNI = originalGetPersonByDNI

	}
}
