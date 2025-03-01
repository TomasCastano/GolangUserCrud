package controllers

import (
	"crud_usuarios/config"
	"crud_usuarios/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Obtener todos los usuarios
func GetUsers(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Println("Error al obtener usuarios:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}

	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Println("Error al escanear usuario:", err)
			continue
		}

		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

// Crear un usuario
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	_, err := config.DB.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)

	if err != nil {
		log.Println("Error al insertar usuario:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al insertar usuario"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado"})
}

// Eliminar un usuario
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	result, err := config.DB.Exec("DELETE FROM users WHERE id = $1", id)

	if err != nil {
		log.Println("Error al eliminar usuario:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar usuario " + id})
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar usuario " + id})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
}
