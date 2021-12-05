package repository

import (
	"gorm.io/gorm"
)

type Repo struct {
	Users *userLayer
	Tasks *taskLayer
}

func NewRepository(db *gorm.DB) Repo {
	return Repo{
		Users: newUserRepoLayer(db),
		Tasks: newTaskRepoLayer(db),
		// Transactions:      newTransactionLayer(db),
		// Wallets:           newWalletLayer(db),
		// TransactionEvents: newEventLayer(db),
	}

}
