package service

import (
	"errors"

	"github.com/devfullcycle/imersao10-consolidacao/internal/domain/entity"
)

func ChoosePlayers(myTeam entity.MyTeam, players []entity.Player) error {
	totalCoast := 0.0
	totalEarned := 0.0

	for _, player := range players {
		if playerInMyTeam(player, myTeam) && !playerInPlayersList(player, &players) {
			totalEarned += player.Price
		}
		if !playerInMyTeam(player, myTeam) && playerInPlayersList(player, &players) {
			totalCoast += player.Price
		}
	}

	if totalCoast > myTeam.Score+totalEarned {
		return errors.New("not enough money")
	}

	for _, player := range players {
		myTeam.Players = append(myTeam.Players, player.ID)
	}

	return nil
}

func playerInMyTeam(player entity.Player, myTeam entity.MyTeam) bool {
	for _, playerID := range myTeam.Players {
		if player.ID == playerID {
			return true
		}
	}
	return false
}

func playerInPlayersList(player entity.Player, players *[]entity.Player) bool {
	for _, p := range *players {
		if player.ID == p.ID {
			return true
		}
	}
	return false
}
