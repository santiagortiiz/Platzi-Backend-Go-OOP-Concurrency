package main

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id int
		dni string
		mockFunc func()
		expectedEmployee FullTimeEmployee
	} {
		{
			id: 1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee {
						Id: 1,
						Position: "CEO",
					}, nil
				}

				GetPersonByDni = func(id string) (Person, error) {
					return Person{
						Name: "John Doe",
						Age: 35,
						DNI: "1",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee {
				Person: Person{
					Age:  35,
					DNI:  "1",
					Name: "John Doe",
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		}
	}
	// Store original functions to preserve them
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDni := GetPersonByDNI
	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)

		if err != nil {
			t.Errorf("Error when getting Employee")
		}

		// reflect.DeepEqual() is also useful to compare the whole object at once
		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error, got %d expected %d", ft.Age, test.expectedEmployee.Age)
		}
	}

	// Original functions restoring to avoid the use of the mock functions later
	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDni
}