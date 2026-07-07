package service

import (
	"time"
	"github.com/google/uuid"
	"qpic-server/model"
)

func GetAllGames() ([]model.Game, error) {
	dummyGames := []model.Game{
		{
			ID:          uuid.New(),
			Title:       "Qpic Action Adventure",
			CreatorID:   "user-001",
			Genre:       "Action",
			Tags:        "2D, Action, SinglePlayer",
			Explanation: "ダミーデータ",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.New(),
			Title:       "Block Puzzle 3D",
			CreatorID:   "user-002",
			Genre:       "Puzzle",
			Tags:        "3D, Puzzle",
			Explanation: "ダミーデータ",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	return dummyGames, nil
}