package repository

import (
	"database/sql"
	"github.com/ismailash/be-enigma-laundry/utils/model_util"
	"log"
	"time"

	"github.com/ismailash/be-enigma-laundry/model/entity"
)

type BillRepository interface {
	Create(billReq entity.Bill) (entity.Bill, error)
	Get(id string) (entity.Bill, error)
	GetWithPagination(paging model_util.Paging) ([]entity.Bill, error)
}

type billRepository struct {
	db *sql.DB
}

func NewBillRepository(db *sql.DB) BillRepository {
	return &billRepository{db: db}
}

func (r *billRepository) Create(billReq entity.Bill) (entity.Bill, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return entity.Bill{}, err
	}

	var bill entity.Bill
	query := `INSERT INTO bills (bill_date, customer_id, user_id, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, bill_date, created_at, updated_at`
	err = tx.QueryRow(query, time.Now(), billReq.Customer.Id, billReq.User.Id, time.Now()).Scan(
		&bill.Id,
		&bill.BillDate,
		&bill.CreatedAt,
		&bill.UpdatedAt,
	)
	if err != nil {
		return entity.Bill{}, tx.Rollback()
	}

	var billDetails []entity.BillDetail
	for _, v := range billReq.BillDetails {
		var billDetail entity.BillDetail
		query := `INSERT INTO bill_details (bill_id, product_id, qty, price, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, qty, price, created_at, updated_at`
		err := tx.QueryRow(query, bill.Id, v.Product.Id, v.Qty, v.Price, time.Now()).Scan(
			&billDetail.Id,
			&billDetail.Qty,
			&billDetail.Price,
			&billDetail.CreatedAt,
			&billDetail.UpdatedAt,
		)
		if err != nil {
			return entity.Bill{}, tx.Rollback()
		}
		billDetail.Product = v.Product
		billDetails = append(billDetails, billDetail)
	}

	bill.Customer = billReq.Customer
	bill.User = billReq.User
	bill.BillDetails = billDetails
	if err := tx.Commit(); err != nil {
		return entity.Bill{}, err
	}

	return bill, nil
}

func (r *billRepository) Get(id string) (entity.Bill, error) {
	log.Println("REPO DISINIII BANGGGG >> ", id)
	var bill entity.Bill
	query := `
	SELECT
	    b.id,
		b.bill_date,
		c.id,
		c.name,
		c.phone_number,
		c.address,
		c.created_at,
		c.updated_at,
		u.id,
		u.name,
		u.email,
		u.username,
		u.role,
		u.created_at,
		u.updated_at,
		b.created_at,
		b.updated_at
	FROM bills b
	JOIN customers c ON c.id = b.customer_id
	JOIN users u ON u.id = b.user_id
	WHERE b.id = $1
	`
	err := r.db.QueryRow(query, id).Scan(
		&bill.Id,
		&bill.BillDate,
		&bill.Customer.Id,
		&bill.Customer.Name,
		&bill.Customer.PhoneNumber,
		&bill.Customer.Address,
		&bill.Customer.CreatedAt,
		&bill.Customer.UpdatedAt,
		&bill.User.Id,
		&bill.User.Name,
		&bill.User.Email,
		&bill.User.Username,
		&bill.User.Role,
		&bill.User.CreatedAt,
		&bill.User.UpdatedAt,
		&bill.CreatedAt,
		&bill.UpdatedAt,
	)

	if err != nil {
		return entity.Bill{}, err
	}

	var billDetails []entity.BillDetail
	query = `
				SELECT	bd.id,
						p.id,
						p.name,
						p.price,
						p.type,
						p.created_at,
						p.updated_at,
						bd.qty,
						bd.price,
						bd.created_at,
						bd.updated_at
						FROM 	bill_details bd
				JOIN 	bills b ON b.id = bd.bill_id
				JOIN 	products p ON p.id = bd.product_id
				WHERE 	b.id = $1
	`
	rows, err := r.db.Query(query, bill.Id)

	if err != nil {
		return entity.Bill{}, err
	}

	for rows.Next() {
		var billDetail entity.BillDetail
		rows.Scan(
			&billDetail.Id,
			&billDetail.Product.Id,
			&billDetail.Product.Name,
			&billDetail.Product.Price,
			&billDetail.Product.Type,
			&billDetail.Product.CreatedAt,
			&billDetail.Product.UpdatedAt,
			&billDetail.Qty,
			&billDetail.Price,
			&billDetail.CreatedAt,
			&billDetail.UpdatedAt,
		)

		billDetails = append(billDetails, billDetail)
	}

	bill.BillDetails = billDetails

	return bill, nil
}

func (r *billRepository) GetWithPagination(paging model_util.Paging) ([]entity.Bill, error) {
	var bills []entity.Bill

	query := `
        SELECT
            b.id,
            b.bill_date,
            c.id,
            c.name,
            c.phone_number,
            c.address,
            c.created_at,
            c.updated_at,
            u.id,
            u.name,
            u.email,
            u.username,
            u.role,
            u.created_at,
            u.updated_at,
            b.created_at,
            b.updated_at
        FROM bills b
        JOIN customers c ON c.id = b.customer_id
        JOIN users u ON u.id = b.user_id
        LIMIT $1 OFFSET $2
    `

	rows, err := r.db.Query(query, paging.Limit, paging.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bill entity.Bill
		err := rows.Scan(
			&bill.Id,
			&bill.BillDate,
			&bill.Customer.Id,
			&bill.Customer.Name,
			&bill.Customer.PhoneNumber,
			&bill.Customer.Address,
			&bill.Customer.CreatedAt,
			&bill.Customer.UpdatedAt,
			&bill.Id,
			&bill.User.Name,
			&bill.User.Email,
			&bill.User.Username,
			&bill.User.Role,
			&bill.User.CreatedAt,
			&bill.User.UpdatedAt,
			&bill.CreatedAt,
			&bill.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		bills = append(bills, bill)
	}

	return bills, nil
}
