package repo

import (
	"L0/dbconnection"
	"L0/dbconnection/entity"
	"encoding/json"
	"log"
)

func InsertOrder(input entity.Order) (string, error) {
	const q = `INSERT INTO orders (order_uid,
		track_number,entry, delivery, payment, items,
		locale, internal_signature, customer_id, delivery_service,
		shardkey, sm_id, date_created, oof_shard)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	deliveryJSON, err := json.Marshal(input.Delivery)
	if err != nil {
		log.Panic(err)
	}
	paymentJSON, err := json.Marshal(input.Payment)
	if err != nil {
		log.Panic(err)
	}
	itemsJSON, err := json.Marshal(input.Items)
	if err != nil {
		log.Panic(err)
	}
	args := []interface{}{
		input.OrderUid,
		input.TrackNumber,
		input.Entry,
		deliveryJSON,
		paymentJSON,
		itemsJSON,
		input.Locale,
		input.InternalSignature,
		input.CustomerId,
		input.DeliveryService,
		input.ShardKey,
		input.SmId,
		input.DateCreated,
		input.OofShard,
	}

	db, err := dbconnection.DbConnection()
	if err != nil {
		log.Panic(err)
	}
	db.Exec(q, args...)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	return input.OrderUid, nil
}
