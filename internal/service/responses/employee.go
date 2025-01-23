package responses

import (
	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"github.com/recovery-flow/organization-storage/resources"
	"github.com/sirupsen/logrus"
)

func Employee(employee models.Employee) resources.Employee {
	ver := "false"
	if employee.Verified {
		ver = "verified"
	}
	logrus.Infof("Employee: %v", employee)
	return resources.Employee{
		Data: resources.EmployeeData{
			Id:   employee.UserID.String(),
			Type: resources.EmployeeType,
			Attributes: resources.EmployeeDataAttributes{
				FirstName:   employee.FirstName,
				SecondName:  employee.SecondName,
				ThirdName:   employee.ThirdName,
				DisplayName: employee.DisplayName,
				Position:    employee.Position,
				Verified:    ver,
				Desc:        employee.Desc,
				Role:        string(employee.Role),
				CreatedAt:   employee.CreatedAt,
			},
		},
	}
}
