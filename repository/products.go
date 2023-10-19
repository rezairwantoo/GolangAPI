package repository

import (
	"case2/helpers"
	"case2/model"
	"case2/model/constant"
	"context"
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

func (p *PostgresRepository) Create(ctx context.Context, req model.CreateRequest) error {
	query := `INSERT INTO products (item, price, is_active, created_at, updated_at)
		VALUES (:item, :price, :is_active, :created_at, :updated_at)`

	now, _ := helpers.GetNow()

	args := map[string]any{
		"item":       req.Item,
		"price":      req.Price,
		"is_active":  true,
		"created_at": now,
		"updated_at": now,
	}

	_, err := ExecStatementContext(ctx, p.Conn, query, args)
	return err
}

func (p *PostgresRepository) GetByID(ctx context.Context, productID int64) (*model.Products, error) {
	query := `SELECT id, item, price, is_active, created_at, updated_at, deleted_at
		FROM products
		WHERE id = :product_id`

	args := map[string]interface{}{
		"product_id": productID,
	}

	result, err := GetStatementContext[model.Products](ctx, p.Conn, query, args)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, constant.ErrMsgNotFoundDefault)
		}
		return nil, errors.Wrap(err, "failed get influencee")
	}

	return &result, err
}

func (p *PostgresRepository) List(ctx context.Context, req *model.ListRequest) ([]model.Products, error) {
	query := `SELECT id, item, price, is_active, created_at, updated_at, deleted_at
	FROM products
	Limit :limit OFFSET :offset`

	if req.Search != "" {
		query = `SELECT id, item, price, is_active, created_at, updated_at, deleted_at
		FROM products 
		WHERE item iLIKE :search
		Limit :limit OFFSET :offset`
	}

	args := map[string]interface{}{
		"limit":  req.Limit,
		"offset": req.Offset,
		"search": req.Search,
	}

	products, err := SelectStatementContext[model.Products](ctx, p.Conn, query, args)
	if err != nil {
		if len(products) > 0 {
			return nil, errors.Wrap(err, constant.ErrMsgNotFoundDefault)
		}
		return nil, errors.Wrap(err, constant.ErrGetData)
	}

	return products, nil
}

func (p *PostgresRepository) CountProductBySearch(ctx context.Context, req *model.ListRequest) (int64, error) {
	var (
		remisierTotal model.ProductsTotalData
		err           error
	)

	query := `SELECT count(id) as total  
		FROM products`

	if req.Search != "" {
		query = `SELECT count(id) as total 
		FROM products 
		WHERE item iLIKE :search`
	}

	args := map[string]interface{}{
		"search": req.Search,
	}

	remisierTotal, err = GetStatementContext[model.ProductsTotalData](ctx, p.Conn, query, args)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.Wrap(err, constant.ErrMsgNotFoundDefault)
		}
		return 0, errors.Wrap(err, constant.ErrGetData)
	}

	return remisierTotal.Total, nil
}

func (p *PostgresRepository) CountAllProduct(ctx context.Context) (int64, error) {
	log.Println("repo")
	log.Println("dsfdsfdsfdsf")
	log.Println(p.Conn)
	var (
		remisierTotal model.ProductsTotalData
		err           error
	)

	query := `SELECT count(id) as total  
		FROM products`
	remisierTotal, err = GetStatementContext[model.ProductsTotalData](ctx, p.Conn, query, nil)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.Wrap(err, constant.ErrMsgNotFoundDefault)
		}
		return 0, errors.Wrap(err, constant.ErrGetData)
	}

	return remisierTotal.Total, nil
}

func (p *PostgresRepository) Update(ctx context.Context, req model.UpdateRequest) error {
	query := `update products set item = :item, price = :price, is_active = :is_active, updated_at = :updated_at where id = :product_id`
	now, _ := helpers.GetNow()

	args := map[string]any{
		"item":       req.Item,
		"price":      req.Price,
		"is_active":  req.IsActive,
		"updated_at": now,
		"product_id": req.ProductID,
	}
	if !req.IsActive {
		query = `update products set item = :item, price = :price, is_active = :is_active, updated_at = :updated_at, deleted_at = :deleted_at where id = :product_id`
		args = map[string]any{
			"item":       req.Item,
			"price":      req.Price,
			"is_active":  req.IsActive,
			"updated_at": now,
			"deleted_at": now,
			"product_id": req.ProductID,
		}
	}

	_, err := ExecStatementContext(ctx, p.Conn, query, args)
	return err
}
