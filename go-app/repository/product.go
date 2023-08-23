package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/marioheryanto/erajaya/go-app/constants"
	"github.com/marioheryanto/erajaya/go-app/helper"
	"github.com/marioheryanto/erajaya/go-app/model"
	"github.com/redis/go-redis/v9"
)

type ProductRepository struct {
	dbClient    *sql.DB
	redisClient *redis.Client
}

type ProductRepositoryInterface interface {
	PingDB() error
	CreateProduct(ctx context.Context, params model.Product) (interface{}, error)
	GetProducts(ctx context.Context) ([]model.Product, error)
}

func NewProductRepository(dbClient *sql.DB, redisClient *redis.Client) ProductRepositoryInterface {
	return &ProductRepository{
		dbClient:    dbClient,
		redisClient: redisClient,
	}
}

func (r *ProductRepository) PingDB() error {
	err := r.dbClient.Ping()
	if err != nil {
		return helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (r *ProductRepository) CreateProduct(ctx context.Context, params model.Product) (interface{}, error) {
	timeNow, err := helper.GetTimeNowWithLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	query, args, err := squirrel.Insert("products").Columns("name, price, description, quantity,created_at").
		Values(params.Name, params.Price, params.Description, params.Quantity, timeNow).
		ToSql()

	if err != nil {
		return 0, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	result, err := r.dbClient.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	params.PublishAt = timeNow

	// maintain redis data
	data, _ := r.redisClient.Get(ctx, constants.Redis_Key_Product).Result()
	if data != "" {
		products := []model.Product{}

		err := json.Unmarshal([]byte(data), &products)
		if err == nil {
			products = append(products, params)

			productsData, err := json.Marshal(products)
			if err == nil {
				r.redisClient.Set(ctx, constants.Redis_Key_Product, productsData, 0)
			}
		}
	}

	return result.LastInsertId()
}

func (r *ProductRepository) GetProducts(ctx context.Context) ([]model.Product, error) {
	products := []model.Product{}

	// get data from redis first, dont handle err redis.
	// we want to fallback to DB
	data, errRedis := r.redisClient.Get(ctx, constants.Redis_Key_Product).Result()
	if errRedis == nil {
		errRedis = json.Unmarshal([]byte(data), &products)
		if errRedis == nil {
			return products, nil
		}
	}

	// cache miss, then get from DB
	query, args, err := squirrel.Select("name, price, description, quantity, created_at").
		From("products").ToSql()
	if err != nil {
		return products, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	rows, err := r.dbClient.QueryContext(ctx, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return products, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	if err == sql.ErrNoRows {
		return products, helper.NewServiceError(http.StatusNotFound, "data not found")
	}

	defer rows.Close()

	for rows.Next() {
		product := model.Product{}

		err := rows.Scan(&product.Name, &product.Price, &product.Description, &product.Quantity, &product.PublishAt)
		if err != nil {
			return products, helper.NewServiceError(http.StatusInternalServerError, err.Error())
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return products, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	// set redis if nil
	if errRedis == redis.Nil && len(products) > 0 {
		productsData, err := json.Marshal(products)
		if err == nil {
			r.redisClient.Set(ctx, constants.Redis_Key_Product, productsData, 0).Result()
		}
	}

	return products, nil
}
