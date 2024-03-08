package database

import (
	"ApiRestAct1/src/asignatures/domain/entities"
	"ApiRestAct1/src/core"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatal("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(asignature entities.Asignature) (entities.Asignature, error) {
	query := "INSERT INTO Asignature (name_asignature, description_asignature) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, asignature.Name, asignature.Description)
	if err != nil {
		return entities.Asignature{}, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	id, err := result.LastInsertId() // Obtener el ID generado
	if err != nil {
		return entities.Asignature{}, fmt.Errorf("error al obtener el ID insertado: %v", err)
	}

	asignature.ID = int(id) // Asignamos el ID generado

	log.Printf("[MySQL] - Asignatura guardada con ID: %d", asignature.ID)

	return asignature, nil // Ahora devolvemos la asignatura guardada con su ID
}

func (mysql *MySQL) GetAll() ([]entities.Asignature, error) {
	query := "SELECT id_asignature, name_asignature, description_asignature FROM Asignature"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener filas: %v", err)
	}
	defer rows.Close()

	var asignatures []entities.Asignature

	for rows.Next() {
		var asignature entities.Asignature
		if err := rows.Scan(&asignature.ID, &asignature.Name, &asignature.Description); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %v", err)
		}
		asignatures = append(asignatures, asignature)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %v", err)
	}

	return asignatures, nil
}

func (mysql *MySQL) GetById(id int) (entities.Asignature, error) {
	query := "SELECT id_asignature, name_asignature, description_asignature FROM Asignature WHERE id_asignature = ?"
	rows, err := mysql.conn.FetchRows(query, id)
	if err != nil {
		return entities.Asignature{}, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	var asignature entities.Asignature
	if rows.Next() {
		if err := rows.Scan(&asignature.ID, &asignature.Name, &asignature.Description); err != nil {
			return entities.Asignature{}, fmt.Errorf("error al escanear la asignatura: %v", err)
		}
		return asignature, nil
	}

	return entities.Asignature{}, fmt.Errorf("no se encontró la asignatura con ID %d", id)
}

func (mysql *MySQL) Update(asignature entities.Asignature) error {
	query := "UPDATE Asignature SET name_asignature = ?, description_asignature = ? WHERE id_asignature = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, asignature.Name, asignature.Description, asignature.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar la asignatura con ID %d: %v", asignature.ID, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ninguna asignatura con ID %d para actualizar", asignature.ID)
	}

	log.Printf("[MySQL] - Asignatura con ID %d actualizada", asignature.ID)
	return nil
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM Asignature WHERE id_asignature = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar la asignatura con ID %d: %v", id, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ninguna asignatura con ID %d para eliminar", id)
	}

	log.Printf("[MySQL] - Asignatura con ID %d eliminada", id)
	return nil
}
