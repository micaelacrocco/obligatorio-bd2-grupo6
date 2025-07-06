package repositories

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"database/sql"
)

type departmentMySQLRepo struct {
	db *sql.DB
}

func NewDepartmentRepository(db *sql.DB) interfaces.DepartmentRepository {
	return &departmentMySQLRepo{db: db}
}

func (r *departmentMySQLRepo) GetAll() ([]models.Department, error) {
	rows, err := r.db.Query("SELECT id, name FROM DEPARTMENTS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []models.Department
	for rows.Next() {
		var d models.Department
		if err := rows.Scan(&d.ID, &d.Name); err != nil {
			return nil, err
		}
		departments = append(departments, d)
	}
	return departments, nil
}

func (r *departmentMySQLRepo) Add(dept models.Department) (*models.Department, error) {
	result, err := r.db.Exec("INSERT INTO DEPARTMENTS(name) VALUES (?)", dept.Name)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	dept.ID = int(id)
	return &dept, nil
}

func (r *departmentMySQLRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM DEPARTMENTS WHERE id = ?", id)
	return err
}
