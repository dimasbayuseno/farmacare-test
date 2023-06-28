package handlers

import (
	"github.com/dimasbayuseno/farmacare-test/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func CreateBattle(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		payload := &models.BattlePayload{}

		if err := c.Bind(payload); err != nil {
			return c.String(http.StatusBadRequest, "Invalid request body")
		}

		tx := db.Begin()
		battleID, err := models.CreateBattleInfo(db, payload.BattleInfo)
		if err != nil {
			log.Println("Failed to create battle info:", err)
			return c.String(http.StatusInternalServerError, "Failed to create battle info")
		}

		err = models.CreateFightResults(db, payload.FightResults, battleID)
		if err != nil {
			tx.Rollback()
			return c.String(http.StatusInternalServerError, "Failed to create fight results")
		}

		tx.Commit()

		return c.String(http.StatusCreated, "Battle info and fight results created successfully")
	}
}

func GetScoreListHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		scoreList, err := models.GetScoreList(db)
		if err != nil {
			log.Println("Failed to retrieve score list:", err)
			return c.String(http.StatusInternalServerError, "Failed to retrieve score list")
		}

		return c.JSON(http.StatusOK, scoreList)
	}
}

func UpdateScoresHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		battleIDParam := c.QueryParam("battle_id")
		battleID, err := strconv.Atoi(battleIDParam)
		if err != nil {
			log.Println("Invalid battle ID:", err)
			return c.String(http.StatusBadRequest, "Invalid battle ID")
		}

		err = models.UpdateScores(db, battleID)
		if err != nil {
			log.Println("Failed to update scores:", err)
			return c.String(http.StatusInternalServerError, "Failed to update scores")
		}

		return c.String(http.StatusOK, "Scores updated successfully")
	}
}

func GetAllBattleInfoHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		battleInfoList, err := models.GetAllBattleInfo(db)
		if err != nil {
			log.Println("Failed to retrieve battle info:", err)
			return c.String(http.StatusInternalServerError, "Failed to retrieve battle info")
		}

		battlePayloadList := make([]models.BattlePayload, 0)
		for _, battleInfo := range battleInfoList {
			fightResults, err := models.GetFightResultsByBattleID(db, battleInfo.ID)
			if err != nil {
				log.Println("Failed to retrieve fight results:", err)
				return c.String(http.StatusInternalServerError, "Failed to retrieve fight results")
			}

			battlePayload := models.BattlePayload{
				BattleInfo:   battleInfo.BattleSchedule,
				FightResults: fightResults,
			}
			battlePayloadList = append(battlePayloadList, battlePayload)
		}

		return c.JSON(http.StatusOK, battlePayloadList)
	}
}
