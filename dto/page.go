package dto

import "github.com/Lucifer07/e-wallet/entity"

type Page struct {
	Data     []entity.HistoryTransaction `json:"data"`
	PageInfo PaginateInfo                `json:"pages"`
}
