package database

import (
	"ApiRestAct1/src/core"
	"ApiRestAct1/src/students/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(student entities.Student) error {
	query := "INSERT INTO Student (name_student, last_name_student, matricule_student, age_student) VALUES (?, ?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, student.Name, student.LastName, student.Matricule, student.Age)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]entities.Student, error) {
	query := "SELECT id_student, name_student, last_name_student, matricule_student, age_student FROM Student"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener filas: %v", err)
	}
	defer rows.Close()

	var students []entities.Student

	for rows.Next() {
		var student entities.Student
		if err := rows.Scan(&student.ID, &student.Name, &student.LastName, &student.Matricule, &student.Age); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %v", err)
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %v", err)
	}

	return students, nil
}

func (mysql *MySQL) GetById(id int) (entities.Student, error) {
	query := "SELECT id_student, name_student, last_name_student, matricule_student, age_student FROM Student WHERE id_student = ?"
	rows, err := mysql.conn.FetchRows(query, id)

	if err != nil {
		return entities.Student{}, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	var student entities.Student
	if rows.Next() {
		if err := rows.Scan(&student.ID, &student.Name, &student.LastName, &student.Matricule, &student.Age); err != nil {
			return entities.Student{}, fmt.Errorf("error al escanear el estudiante: %v", err)
		}
		return student, nil
	}

	return entities.Student{}, fmt.Errorf("no se encontró el estudiante con ID %d", id)
}

func (mysql *MySQL) Update(student entities.Student) error {
	query := "UPDATE Student SET name_student = ?, last_name_student = ?, matricule_student = ?, age_student = ? WHERE id_student = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, student.Name, student.LastName, student.Matricule, student.Age, student.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar el estudiante con ID %d: %v", student.ID, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ningún estudiante con ID %d para actualizar", student.ID)
	}

	log.Printf("[MySQL] - Estudiante con ID %d actualizado", student.ID)
	return nil
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM Student WHERE id_student = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar el estudiante con ID %d: %v", id, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ningún estudiante con ID %d para eliminar", id)
	}

	log.Printf("[MySQL] - Estudiante con ID %d eliminado", id)
	return nil
}

func (mysql *MySQL) GetByAge(minAge int) ([]entities.Student, error) {
	query := "SELECT id_student, name_student, last_name_student, matricule_student, age_student FROM Student WHERE age_student >= ?"
	rows, err := mysql.conn.FetchRows(query, minAge)
	if err != nil {
		return nil, fmt.Errorf("error al obtener estudiantes mayores de %d años: %v", minAge, err)
	}
	defer rows.Close()

	var students []entities.Student

	for rows.Next() {
		var student entities.Student
		if err := rows.Scan(&student.ID, &student.Name, &student.LastName, &student.Matricule, &student.Age); err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando filas: %v", err)
	}

	return students, nil
}
