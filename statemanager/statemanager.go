package statemanager

import (
	"fmt"

	cm "github.com/9500073161/coursemanagement/common"
	"github.com/9500073161/coursemanagement/mysqldbmodels"
	dbmodels "github.com/9500073161/coursemanagement/mysqldbmodels"
)

// Goal is to keep tracking of open positions
// Extension would be to support for multi-day strategy
type StateManager struct {
	db *dbmodels.DBClient
}

func InitStateManager() (*StateManager, error) {
	sm := &StateManager{}
	var err error
	if sm.db, err = dbmodels.InitializeDatabase(); err != nil {
		return sm, fmt.Errorf("error occured while creating database connection,err: %s", err.Error())
	}
	return sm, nil
}

func (sm *StateManager) CreateCourseEntry(c1 cm.Course) error {
	return sm.db.CreateCourseRow(c1)
}

func (sm *StateManager) CreateStudentEntry(s1 cm.Student) error {
	return sm.db.CreateStudentRow(s1)
}

func (sm *StateManager) CreateTeacherEntry(t1 cm.Teacher) error {
	return sm.db.CreateTeacherRow(t1)
}

func (sm *StateManager) CreateEntrollmentEntry(e1 cm.Entrollment) error {
	return sm.db.CreateEntrollmentRow(e1)
}

func (sm *StateManager) GetAllCourses() ([]mysqldbmodels.Course, error) {
	return sm.db.GetCourseRaw()
}

func (sm *StateManager) GetAllStudents() ([]mysqldbmodels.Student, error) {
	return sm.db.GetStudentRaw()
}

func (sm *StateManager) GetAllTeachers() ([]mysqldbmodels.Teacher, error) {
	return sm.db.GetTeacherRaw()
}

func (sm *StateManager) GetAllEntrollments() ([]mysqldbmodels.Entrollment, error) {
	return sm.db.GetEntrollmentRaw()
}
