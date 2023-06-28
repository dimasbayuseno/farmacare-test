package models

import (
	"gorm.io/gorm"
	"log"
)

type FightResult struct {
	ID          int    `json:"-"`
	PokemonName string `json:"pokemon_name"`
	Score       int    `json:"score"`
	BattleID    int    `json:"-"`
}

type BattleInfo struct {
	ID             int    `json:"id"`
	BattleSchedule string `json:"battle_schedule"`
}

type FoughtPokemon struct {
	PokemonName string `json:"pokemon_name"`
	ID          int    `json:"id"`
}

type BattlePayload struct {
	BattleInfo   string        `json:"battle_schedule"`
	FightResults []FightResult `json:"fight_results"`
}

func CreateBattleInfo(db *gorm.DB, battleSchedule string) (int, error) {
	battleInfo := BattleInfo{
		BattleSchedule: battleSchedule,
	}
	if err := db.Create(&battleInfo).Error; err != nil {
		log.Println("Failed to create battle info:", err)
		return 0, err
	}
	return battleInfo.ID, nil
}

func CreateFightResults(db *gorm.DB, fightResults []FightResult, battleID int) error {
	for i, fightResult := range fightResults {
		fightResult.BattleID = battleID
		fightResult.Score = i + 1

		if err := db.Create(&fightResult).Error; err != nil {
			log.Println("Failed to create fight result:", err)
			return err
		}
	}
	return nil
}

func GetScoreList(db *gorm.DB) ([]FightResult, error) {
	scoreList := []FightResult{}
	err := db.Select("pokemon_name, SUM(score) as score").
		Group("pokemon_name").
		Order("score desc").
		Find(&scoreList).Error
	if err != nil {
		log.Println("Failed to retrieve score list:", err)
		return nil, err
	}

	return scoreList, nil
}

func UpdateScores(db *gorm.DB, battleID int) error {
	rowsToUpdate := []FightResult{}
	err := db.Table("fight_results").
		Where("battle_id = ?", battleID).
		Order("score DESC").
		Find(&rowsToUpdate).Error
	if err != nil {
		log.Println("Failed to retrieve rows to update:", err)
		return err
	}
	tx := db.Begin()

	for i := len(rowsToUpdate) - 1; i >= 0; i-- {
		row := &rowsToUpdate[i]
		if row.Score == 5 {
			row.Score = 1
		} else {
			row.Score = row.Score + 1
		}
		err = db.Save(row).Error
		if err != nil {
			tx.Rollback()
			log.Println("Failed to update row:", err)
			return err
		}
	}
	tx.Commit()

	return nil
}

func GetAllBattleInfo(db *gorm.DB) ([]BattleInfo, error) {
	battleInfoList := []BattleInfo{}

	err := db.Find(&battleInfoList).Error
	if err != nil {
		return battleInfoList, err
	}

	return battleInfoList, nil
}

func GetFightResultsByBattleID(db *gorm.DB, battleID int) ([]FightResult, error) {
	fightResults := []FightResult{}

	err := db.Where("battle_id = ?", battleID).Find(&fightResults).Error
	if err != nil {
		return fightResults, err
	}

	return fightResults, nil
}
