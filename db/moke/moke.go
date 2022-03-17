package moke

import (
	"github.com/Aazarias/ApiGomasterclass/db"
	"github.com/Aazarias/ApiGomasterclass/model"

	"errors"

	"github.com/google/uuid"
)

var _ db.StorageGames = &Moke{}

type Moke struct {
	listGames map[string]*model.Games
}

func New() *db.Storage {
	return &db.Storage{
		Games: &Moke{
			listGames: make(map[string]*model.Games),
		},
	}
}

//Get a game information by ID
func (m *Moke) GetByID(id string) (*model.Games, error) {
	u, ok := m.listGames[id]
	if !ok {
		return nil, errors.New("db user: not found")
	}
	return u, nil
}

// Delete a game by the ID
func (m *Moke) DeleteByID(id string) error {
	_, ok := m.listGames[id]
	if !ok {
		return errors.New("db user: not found")
	}
	delete(m.listGames, id)
	return nil
}

//Add a game
func (m *Moke) Create(u *model.Games) (*model.Games, error) {
	u.ID = uuid.New().String()
	m.listGames[u.ID] = u
	return u, nil
}

//Update information of a game
func (m *Moke) Update(id string, data map[string]interface{}) (*model.Games, error) {
	u, ok := m.listGames[id]
	if !ok {
		return nil, errors.New("db game: not found")
	}
	if value, ok := data["games_name"]; ok {
		u.Name = value.(string)
	}
	return nil, nil
}

//Get the list of all game
func (m *Moke) GetAll() ([]model.Games, error) {
	us := make([]model.Games, len(m.listGames))
	var i int
	for k := range m.listGames {
		if m.listGames[k] != nil {
			us[i] = *m.listGames[k]
		}
		i++
	}
	return us, nil
}
