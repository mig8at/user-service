package seeder

import (
	"fmt"
	"log"
	"user_service/internal/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Seeder es responsable de precargar datos en la base de datos.
type Seeder struct {
	db *gorm.DB
}

// NewSeeder crea una nueva instancia de Seeder.
func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{db: db}
}

func (s *Seeder) Seed() {
	// Definir usuarios de prueba
	users := []models.User{
		{
			ID:       "2a42c7ae-7f78-4e36-8358-902342fe23f1",
			Name:     "Juan Pérez",
			Email:    "juan.perez@example.com",
			Nickname: "@juanito",
			Bio:      "Desarrollador de software",
			Avatar:   "https://picsum.photos/200/200?random=1",
		},
		{
			ID:       uuid.NewString(),
			Name:     "María Gómez",
			Email:    "maria.gomez@example.com",
			Nickname: "@mary",
			Bio:      "Diseñadora gráfica",
			Avatar:   "https://picsum.photos/200/200?random=2",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Carlos Ramírez",
			Email:    "carlos.ramirez@example.com",
			Nickname: "@carlitos",
			Bio:      "Ingeniero mecánico",
			Avatar:   "https://picsum.photos/200/200?random=3",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Ana Fernández",
			Email:    "ana.fernandez@example.com",
			Nickname: "@anita",
			Bio:      "Médico pediatra",
			Avatar:   "https://picsum.photos/200/200?random=4",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Pedro López",
			Email:    "pedro.lopez@example.com",
			Nickname: "@pedrito",
			Bio:      "Arquitecto de soluciones",
			Avatar:   "https://picsum.photos/200/200?random=5",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Sofía Martínez",
			Email:    "sofia.martinez@example.com",
			Nickname: "@sofi",
			Bio:      "Chef profesional",
			Avatar:   "https://picsum.photos/200/200?random=6",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Luis Hernández",
			Email:    "luis.hernandez@example.com",
			Nickname: "@luisito",
			Bio:      "Abogado corporativo",
			Avatar:   "https://picsum.photos/200/200?random=7",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Laura Castro",
			Email:    "laura.castro@example.com",
			Nickname: "@lau",
			Bio:      "Psicóloga clínica",
			Avatar:   "https://picsum.photos/200/200?random=8",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Miguel Ángel",
			Email:    "miguel.angel@example.com",
			Nickname: "@mike",
			Bio:      "Artista plástico",
			Avatar:   "https://picsum.photos/200/200?random=9",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Carmen Díaz",
			Email:    "carmen.diaz@example.com",
			Nickname: "@carmi",
			Bio:      "Escritora independiente",
			Avatar:   "https://picsum.photos/200/200?random=10",
		},
		{
			ID:       uuid.NewString(),
			Name:     "José Torres",
			Email:    "jose.torres@example.com",
			Nickname: "@joseto",
			Bio:      "Ingeniero civil",
			Avatar:   "https://picsum.photos/200/200?random=11",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Isabel Sánchez",
			Email:    "isabel.sanchez@example.com",
			Nickname: "@isa",
			Bio:      "Enfermera",
			Avatar:   "https://picsum.photos/200/200?random=12",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Diego Ruiz",
			Email:    "diego.ruiz@example.com",
			Nickname: "@dieguito",
			Bio:      "Profesor de historia",
			Avatar:   "https://picsum.photos/200/200?random=13",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Valeria Morales",
			Email:    "valeria.morales@example.com",
			Nickname: "@vale",
			Bio:      "Fotógrafa profesional",
			Avatar:   "https://picsum.photos/200/200?random=14",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Sebastián Navarro",
			Email:    "sebastian.navarro@example.com",
			Nickname: "@sebas",
			Bio:      "Analista financiero",
			Avatar:   "https://picsum.photos/200/200?random=15",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Gabriela Vargas",
			Email:    "gabriela.vargas@example.com",
			Nickname: "@gabi",
			Bio:      "Marketing digital",
			Avatar:   "https://picsum.photos/200/200?random=16",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Manuel Ortiz",
			Email:    "manuel.ortiz@example.com",
			Nickname: "@manu",
			Bio:      "Ingeniero de software",
			Avatar:   "https://picsum.photos/200/200?random=17",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Camila Pérez",
			Email:    "camila.perez@example.com",
			Nickname: "@cami",
			Bio:      "Consultora de negocios",
			Avatar:   "https://picsum.photos/200/200?random=18",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Rodrigo Rojas",
			Email:    "rodrigo.rojas@example.com",
			Nickname: "@rodrigo",
			Bio:      "Empresario",
			Avatar:   "https://picsum.photos/200/200?random=19",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Natalia Vega",
			Email:    "natalia.vega@example.com",
			Nickname: "@naty",
			Bio:      "Investigadora científica",
			Avatar:   "https://picsum.photos/200/200?random=20",
		},
	}

	// Insertar usuarios si no existen
	for _, user := range users {
		var existingUser models.User
		result := s.db.Where("email = ?", user.Email).First(&existingUser)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				if err := s.db.Create(&user).Error; err != nil {
					log.Printf("Error al insertar usuario %s: %v", user.Email, err)
				} else {
					log.Printf("Usuario %s insertado correctamente.", user.Email)
				}
			} else {
				log.Printf("Error al buscar usuario %s: %v", user.Email, result.Error)
			}
		} else {
			log.Printf("Usuario %s ya existe. Skipping.", user.Email)
		}
	}

	// Definir relaciones de seguidores de prueba
	followers := []models.Follower{
		{
			ID:         uuid.NewString(),
			UserID:     users[0].ID, // Juan Pérez
			FollowerID: users[1].ID, // María Gómez
		},
		{
			ID:         uuid.NewString(),
			UserID:     users[0].ID, // Juan Pérez
			FollowerID: users[2].ID, // Carlos López
		},
		{
			ID:         uuid.NewString(),
			UserID:     users[1].ID, // María Gómez
			FollowerID: users[0].ID, // Juan Pérez
		},
	}

	// Insertar relaciones de seguidores si no existen
	for _, follower := range followers {
		var existingFollower models.Follower
		result := s.db.Where("user_id = ? AND follower_id = ?", follower.UserID, follower.FollowerID).First(&existingFollower)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				if err := s.db.Create(&follower).Error; err != nil {
					log.Printf("Error al insertar follower: %v", err)
				} else {
					log.Printf("Follower de %s insertado correctamente.", follower.UserID)
					// Actualizar conteos de seguidores y siguiendo
					if err := s.db.Model(&models.User{}).Where("id = ?", follower.UserID).
						UpdateColumn("followers", gorm.Expr("followers + ?", 1)).Error; err != nil {
						log.Printf("Error al actualizar followers para user %s: %v", follower.UserID, err)
					}
					if err := s.db.Model(&models.User{}).Where("id = ?", follower.FollowerID).
						UpdateColumn("following", gorm.Expr("following + ?", 1)).Error; err != nil {
						log.Printf("Error al actualizar following para follower %s: %v", follower.FollowerID, err)
					}
				}
			} else {
				log.Printf("Error al buscar follower: %v", result.Error)
			}
		} else {
			log.Printf("Follower de %s ya existe. Skipping.", follower.UserID)
		}
	}

	fmt.Println("Seeding completado.")
}
