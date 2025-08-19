package barcode

import (
	"context"
	"fmt"
	"scanner-backend/config"
	"time"
)

func Create(ctx context.Context, data BarcodeRequest) error {
	id, err := config.Rdb.Incr(ctx, "barcodeInc").Result()
	if err != nil {
		return err
	}
	produto := map[string]any{
		"produto":  data.Produto,
		"barcode":  data.Barcode,
		"datahora": time.Now().Format("02/01/2006 15:04"),
	}
	_, err = config.Rdb.HSet(ctx, fmt.Sprintf("produto:%d", id), produto).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetAll(ctx context.Context) ([]BarcodeSimples, error){
	var cursor uint64
	produtos := []BarcodeSimples{}

	for {
		keys, nextCursor, err := config.Rdb.Scan(ctx, cursor, "produto:*", 100).Result()
		if err != nil {
			return nil, err
		}
		cursor = nextCursor

		for _, key := range keys {
			data, _ := config.Rdb.HGetAll(ctx, key).Result()
			produto := BarcodeSimples{
				Produto:  data["produto"],
				Code:     data["barcode"],
				Datahora: data["datahora"],
			}
			produtos = append(produtos, produto)
		}

		if cursor == 0 {
			break
		}
	}

	return produtos, nil
}
