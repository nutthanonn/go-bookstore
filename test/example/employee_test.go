package example

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	ir "github.com/nutthanonn/go-clean-architecture/pkg/interface/repository"
	ur "github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
	"github.com/nutthanonn/go-clean-architecture/testutil"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (r ur.EmployeeRepository, mock sqlmock.Sqlmock, teardown func()) {
	db, mock, _ := testutil.NewDB(t)

	r = ir.NewEmployeeRepository(db)

	return r, mock, func() {
		sqlDB, err := db.DB()

		if err != nil {
			panic(err)
		}

		sqlDB.Close()

	}
}

func TestGetEmployee(t *testing.T) {
	r, mock, teardown := setup(t)
	defer teardown()

	cases := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, u []*entities.Employees, err error)
	}{
		"create_employee": {
			arrange: func(t *testing.T) {
				uid := uuid.New()
				time_now := time.Now()
				emp := entities.Employees{
					Employee_id:   uid,
					First_name:    "nutthanonn",
					Last_name:     "tho",
					Email:         "nutthanon.tho@gmail.com",
					Phone:         "081-123-4567",
					Job_title:     "developer",
					Date_of_birth: "1990-01-01",
					Salary:        100000,
					Hire_date:     time_now,
					Create_at:     time_now,
					Update_at:     time_now,
				}
				rows := sqlmock.NewRows([]string{"employee_id", "first_name", "last_name", "email",
					"phone", "job_title", "date_of_birth", "salary", "hire_date", "create_at", "update_at"}).
					AddRow(emp.Employee_id, emp.First_name, emp.Last_name, emp.Email, emp.Phone, emp.Job_title, emp.Date_of_birth, emp.Salary, emp.Hire_date, emp.Create_at, emp.Update_at)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT *`)).WillReturnRows(rows)
			},
			assert: func(t *testing.T, u []*entities.Employees, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 1, len(u))
				assert.NotEmpty(t, u[0].Employee_id)
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			c.arrange(t)

			u, err := r.ReadEmployee()

			c.assert(t, u, err)
		})
	}
}
