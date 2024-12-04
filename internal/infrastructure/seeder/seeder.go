package seeder

// import (
// 	"encoding/json"
// 	"log"
// 	"user_service/internal/domain"
// 	"user_service/internal/infrastructure/config"

// 	"github.com/dgraph-io/badger/v4"
// 	"github.com/google/uuid"
// )

// type Seeder struct {
// 	db *badger.DB
// }

// func NewSeeder(cfg *config.Config) *Seeder {
// 	opts := badger.DefaultOptions(cfg.BadgerPath)
// 	db, err := badger.Open(opts)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return &Seeder{db: db}
// }

// func (s *Seeder) Seed() {
// 	// Definir usuarios de prueba
// 	users := []domain.User{
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Juan Pérez",
// 			Email:    "juan.perez@example.com",
// 			Nickname: "@juanito",
// 			Bio:      "Desarrollador de software",
// 			Avatar:   "https://picsum.photos/200/200?random=1",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "María Gómez",
// 			Email:    "maria.gomez@example.com",
// 			Nickname: "@mary",
// 			Bio:      "Diseñadora gráfica",
// 			Avatar:   "https://picsum.photos/200/200?random=2",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Carlos Ramírez",
// 			Email:    "carlos.ramirez@example.com",
// 			Nickname: "@carlitos",
// 			Bio:      "Ingeniero mecánico",
// 			Avatar:   "https://picsum.photos/200/200?random=3",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Ana Fernández",
// 			Email:    "ana.fernandez@example.com",
// 			Nickname: "@anita",
// 			Bio:      "Médico pediatra",
// 			Avatar:   "https://picsum.photos/200/200?random=4",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Pedro López",
// 			Email:    "pedro.lopez@example.com",
// 			Nickname: "@pedrito",
// 			Bio:      "Arquitecto de soluciones",
// 			Avatar:   "https://picsum.photos/200/200?random=5",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Sofía Martínez",
// 			Email:    "sofia.martinez@example.com",
// 			Nickname: "@sofi",
// 			Bio:      "Chef profesional",
// 			Avatar:   "https://picsum.photos/200/200?random=6",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Luis Hernández",
// 			Email:    "luis.hernandez@example.com",
// 			Nickname: "@luisito",
// 			Bio:      "Abogado corporativo",
// 			Avatar:   "https://picsum.photos/200/200?random=7",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Laura Castro",
// 			Email:    "laura.castro@example.com",
// 			Nickname: "@lau",
// 			Bio:      "Psicóloga clínica",
// 			Avatar:   "https://picsum.photos/200/200?random=8",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Miguel Ángel",
// 			Email:    "miguel.angel@example.com",
// 			Nickname: "@mike",
// 			Bio:      "Artista plástico",
// 			Avatar:   "https://picsum.photos/200/200?random=9",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Carmen Díaz",
// 			Email:    "carmen.diaz@example.com",
// 			Nickname: "@carmi",
// 			Bio:      "Escritora independiente",
// 			Avatar:   "https://picsum.photos/200/200?random=10",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "José Torres",
// 			Email:    "jose.torres@example.com",
// 			Nickname: "@joseto",
// 			Bio:      "Ingeniero civil",
// 			Avatar:   "https://picsum.photos/200/200?random=11",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Isabel Sánchez",
// 			Email:    "isabel.sanchez@example.com",
// 			Nickname: "@isa",
// 			Bio:      "Enfermera",
// 			Avatar:   "https://picsum.photos/200/200?random=12",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Diego Ruiz",
// 			Email:    "diego.ruiz@example.com",
// 			Nickname: "@dieguito",
// 			Bio:      "Profesor de historia",
// 			Avatar:   "https://picsum.photos/200/200?random=13",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Valeria Morales",
// 			Email:    "valeria.morales@example.com",
// 			Nickname: "@vale",
// 			Bio:      "Fotógrafa profesional",
// 			Avatar:   "https://picsum.photos/200/200?random=14",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Sebastián Navarro",
// 			Email:    "sebastian.navarro@example.com",
// 			Nickname: "@sebas",
// 			Bio:      "Analista financiero",
// 			Avatar:   "https://picsum.photos/200/200?random=15",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Gabriela Vargas",
// 			Email:    "gabriela.vargas@example.com",
// 			Nickname: "@gabi",
// 			Bio:      "Marketing digital",
// 			Avatar:   "https://picsum.photos/200/200?random=16",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Manuel Ortiz",
// 			Email:    "manuel.ortiz@example.com",
// 			Nickname: "@manu",
// 			Bio:      "Ingeniero de software",
// 			Avatar:   "https://picsum.photos/200/200?random=17",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Camila Pérez",
// 			Email:    "camila.perez@example.com",
// 			Nickname: "@cami",
// 			Bio:      "Consultora de negocios",
// 			Avatar:   "https://picsum.photos/200/200?random=18",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Rodrigo Rojas",
// 			Email:    "rodrigo.rojas@example.com",
// 			Nickname: "@rodrigo",
// 			Bio:      "Empresario",
// 			Avatar:   "https://picsum.photos/200/200?random=19",
// 		},
// 		{
// 			ID:       uuid.NewString(),
// 			Name:     "Natalia Vega",
// 			Email:    "natalia.vega@example.com",
// 			Nickname: "@naty",
// 			Bio:      "Investigadora científica",
// 			Avatar:   "https://picsum.photos/200/200?random=20",
// 		},
// 	}

// 	// delete all keys
// 	err := s.db.DropAll()
// 	if err != nil {
// 		log.Fatalf("Error al eliminar datos: %s", err)
// 	}

// 	for _, user := range users {
// 		err := s.db.Update(func(txn *badger.Txn) error {
// 			data, err := json.Marshal(user)
// 			if err != nil {
// 				return err
// 			}
// 			return txn.Set([]byte(user.ID), data)
// 		})
// 		if err != nil {
// 			log.Fatalf("Error al insertar usuario: %s", err)
// 		}
// 	}

// 	s.db.Close()
// }
