package payload

import (
	"github.com/Lucifer07/e-wallet/dto"
	"github.com/Lucifer07/e-wallet/entity"
)

func TransactionToPages(transactions []entity.HistoryTransaction, pageInfo dto.PaginateInfo) dto.Page {
	return dto.Page{
		Data:     transactions,
		PageInfo: pageInfo,
	}
}
