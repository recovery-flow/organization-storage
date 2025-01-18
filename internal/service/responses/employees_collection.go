package responses

import (
	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"github.com/recovery-flow/organization-storage/resources"
)

func EmployeesCollection(employees []models.Employee) resources.EmployeeCollection {
	var res []resources.Employee
	for _, el := range employees {
		ver := "false"
		if el.Verified {
			ver = "verified"
		}
		res = append(res, resources.Employee{
			Data: resources.EmployeeData{
				Id:   el.UserID.String(),
				Type: resources.EmployeeType,
				Attributes: resources.EmployeeDataAttributes{
					FirstName:   el.FirstName,
					SecondName:  el.SecondName,
					ThirdName:   el.ThirdName,
					DisplayName: el.DisplayName,
					Position:    el.Position,
					Verified:    ver,
					Desc:        el.Desc,
					Role:        string(el.Role),
					UpdatedAt:   el.UpdatedAt,
					CreatedAt:   el.CreatedAt,
				},
			},
		})
	}

	return resources.EmployeeCollection{
		Data: res,
	}
}
