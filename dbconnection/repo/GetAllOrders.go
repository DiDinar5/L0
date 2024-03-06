package repo

import (
	"L0/dbconnection"
	"L0/dbconnection/entity"
	"encoding/json"
	"log"
)

func GetAllOrders() ([]entity.Order, error) {
	db, err := dbconnection.DbConnection()
	if err != nil {
		log.Panic(err)
	}
	var order entity.Order

	rows, err := db.Query(`SELECT order_uid,
	track_number,entry, delivery, payment, items,
	locale, internal_signature, customer_id, delivery_service,
	shardkey, sm_id, date_created, oof_shard FROM orders`)
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()
	var deliveryJSON, paymentJSON, itemsJSON []byte
	var orders []entity.Order
	for rows.Next() {
		err = rows.Scan(
			&order.OrderUid,
			&order.TrackNumber,
			&order.Entry,
			&deliveryJSON,
			&paymentJSON,
			&itemsJSON,
			&order.Locale,
			&order.InternalSignature,
			&order.CustomerId,
			&order.DeliveryService,
			&order.ShardKey,
			&order.SmId,
			&order.DateCreated,
			&order.OofShard,
		)
		if err != nil {
			log.Panic(err)
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	err = json.Unmarshal(deliveryJSON, &order.Delivery)
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(paymentJSON, &order.Payment)
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(itemsJSON, &order.Items)
	if err != nil {
		log.Panic(err)
	}
	jsonData, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonData, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
