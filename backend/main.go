package main

import (
	"crud_usuarios/config"
	"crud_usuarios/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar .env
	err := godotenv.Load()

	if err != nil {
		log.Fatal(" Error cargando el archivo .env")
	}

	// Conectar a la base de datos
	config.ConnectDB()

	// Configurar servidor
	r := gin.Default()

	// Habilitar CORS correctamente ✅
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	// Configurar rutas
	routes.SetupRoutes(r)
	// Iniciar servidor
	log.Println(" Servidor corriendo en http://localhost:8080")
	r.Run(":8080")
}
