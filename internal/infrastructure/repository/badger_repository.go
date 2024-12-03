package repository

import (
	"encoding/json"
	"user_service/internal/domain"
	"user_service/internal/infrastructure/config"
	"user_service/internal/ports"

	"github.com/dgraph-io/badger/v4"
)

type badgerRepository struct {
	db *badger.DB
}

func NewBadgerRepository(cfg *config.Config) ports.UserRepository {
	opts := badger.DefaultOptions(cfg.BadgerPath)
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}

	return &badgerRepository{db: db}
}

func (r *badgerRepository) Create(user *domain.User) error {
	return r.db.Update(func(txn *badger.Txn) error {
		data, err := json.Marshal(user)
		if err != nil {
			return err
		}
		return txn.Set([]byte(user.ID), data)
	})
}

func (r *badgerRepository) GetByID(id string) (*domain.User, error) {
	var user domain.User
	err := r.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return json.Unmarshal(val, &user)
		})
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *badgerRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			var user domain.User
			if err := item.Value(func(val []byte) error {
				return json.Unmarshal(val, &user)
			}); err != nil {
				return err
			}
			users = append(users, user)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return users, nil
}
