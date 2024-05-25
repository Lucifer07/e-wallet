package repository

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"strings"

	"github.com/Lucifer07/e-wallet/entity"
	"github.com/Lucifer07/e-wallet/util"
)

type HistoryRepositoryDb struct {
	db *sql.DB
}

type HistoryRepository interface {
	CreateHistory(ctx context.Context, history entity.HistoryTransaction) (*entity.HistoryTransaction, error)
	MyTransactions(ctx context.Context, wallet entity.Wallet, params map[string]string) (*[]entity.HistoryTransaction, error)
	GetWalletNumbers(ctx context.Context, historyId int) (*[]entity.Wallet, error)
}

func NewHistoryRepository(db *sql.DB) *HistoryRepositoryDb {
	return &HistoryRepositoryDb{
		db: db,
	}
}
func (r *HistoryRepositoryDb) CreateHistory(ctx context.Context, history entity.HistoryTransaction) (*entity.HistoryTransaction, error) {
	var historyData entity.HistoryTransaction
	statment := `insert into history_transactions
	(recipient_wallet_id,sender_wallet_id,amount,source_of_fund,description)
	values($1,$2,$3,$4,$5) 
	returning id,description,source_of_fund,recipient_wallet_id,sender_wallet_id,amount,created_at,updated_at;`
	db := util.GetQueryRunner(ctx, r.db)
	err := db.QueryRowContext(ctx, statment, history.RecipientWalletId, history.SenderWalletId, history.Amount, history.TransactionMethod, history.Description).Scan(&historyData.Id, &historyData.Description, &historyData.TransactionMethod, &historyData.RecipientWalletId, &historyData.SenderWalletId, &historyData.Amount, &historyData.CreatedAt, &historyData.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &historyData, nil
}
func (r *HistoryRepositoryDb) MyTransactions(ctx context.Context, wallet entity.Wallet, params map[string]string) (*[]entity.HistoryTransaction, error) {
	var statment strings.Builder
	var histories []entity.HistoryTransaction
	args := []interface{}{}
	pageInt := 1
	page, okPage := params["page"]
	if okPage {
		pageInt = validatePage(page)
	}
	limitInt := 10
	limit, okLimit := params["limit"]
	if okLimit {
		limitInt = validateLimit(limit)
	}
	statment.WriteString(`select id,recipient_wallet_id,sender_wallet_id,amount,source_of_fund,description,created_at,updated_at,deleted_at from history_transactions `)
	if params["type"] == "all" {
		statment.WriteString("where ((recipient_wallet_id=$1 or recipient_wallet_id=$2) or ((sender_wallet_id=$1 or sender_wallet_id=$2) and source_of_fund = 'wallet'))")
	} else if params["type"] == "topup" {
		statment.WriteString("where ((recipient_wallet_id=$1 or recipient_wallet_id=$2) and source_of_fund != 'wallet')")
	}else{
	statment.WriteString("where (((recipient_wallet_id=$1 or recipient_wallet_id=$2) or (sender_wallet_id=$1 or sender_wallet_id=$2)) and source_of_fund = 'wallet')")
	}
	args = append(args, wallet.Id)
	args=append(args, wallet.WalletNumber)
	if params != nil {
		if description := params["description"]; description != "" {
			statment.WriteString("and description ILIKE $3")
			args = append(args, "%"+description+"%")
		}
		from := params["from"]
		to := params["to"]
		if from != "" && to != "" {
			statment.WriteString(" and created_at between " + from + " and " + to)
		}
		sortBy := params["sortBy"]
		sortOrder := params["sortOrder"]
		if sortBy != "" && sortOrder != "" {
			statment.WriteString(" ORDER BY " + sortBy + " " + sortOrder)
		}
	}
	offset := (pageInt - 1) * limitInt
	statment.WriteString(" LIMIT $" + strconv.Itoa(len(args)+1) + " OFFSET $" + strconv.Itoa(len(args)+2))
	args = append(args, limitInt, offset)
	log.Println(statment.String())
	rows, err := r.db.QueryContext(ctx, statment.String(), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var history entity.HistoryTransaction
		err := rows.Scan(&history.Id, &history.RecipientWalletId, &history.SenderWalletId, &history.Amount, &history.TransactionMethod, &history.Description, &history.CreatedAt, &history.UpdatedAt, &history.DeletedAt)
		if err != nil {
			return nil, err
		}
		histories = append(histories, history)
	}
	log.Println(&histories)
	return &histories, nil
}
func validateLimit(limit string) int {
	limitInt := 10
	if num, err := strconv.Atoi(limit); err == nil && num > 0 {
		limitInt = num
	}
	return limitInt
}

func validatePage(page string) int {
	pageInt := 1
	if num, err := strconv.Atoi(page); err == nil && num > 0 {
		pageInt = num
	}
	return pageInt
}
func (r *HistoryRepositoryDb) GetWalletNumbers(ctx context.Context, historyId int) (*[]entity.Wallet, error) {
	var wallets []entity.Wallet
	var recipient, sender entity.Wallet
	statment := `select 
	wr.wallet_number,ws.wallet_number 
	from history_transactions ht 
	join wallets wr on wr.id=ht.recipient_wallet_id 
	join wallets ws on ws.id=ht.sender_wallet_id where ht.id=$1;`
	db := util.GetQueryRunner(ctx, r.db)
	err := db.QueryRowContext(ctx, statment, historyId).Scan(&recipient.WalletNumber, &sender.WalletNumber)
	if err != nil {
		return nil, err
	}
	wallets = append(wallets, recipient, sender)
	return &wallets, nil
}
