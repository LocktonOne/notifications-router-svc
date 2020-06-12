package pg

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/notifications/notifications-router-svc/internal/data"
)

const deliveriesTableName = "deliveries"

func NewDeliveriesQ(db *pgdb.DB) data.DeliveriesQ {
	return &deliveriesQ{
		db:  db.Clone(),
		sql: sq.Select("*").From(deliveriesTableName),
	}
}

type deliveriesQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func (q *deliveriesQ) New() data.DeliveriesQ {
	return NewDeliveriesQ(q.db)
}

func (q *deliveriesQ) Get() (*data.Delivery, error) {
	var result data.Delivery
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *deliveriesQ) Select() ([]data.Delivery, error) {
	var result []data.Delivery
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *deliveriesQ) Transaction(fn func(q data.DeliveriesQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *deliveriesQ) FilterByNotificationID(ids ...int64) data.DeliveriesQ {
	q.sql = q.sql.Where(sq.Eq{"notification_id": ids})
	return q
}

func (q *deliveriesQ) FilterByDestination(destinations ...string) data.DeliveriesQ {
	q.sql = q.sql.Where(sq.Eq{"destination": destinations})
	return q
}

func (q *deliveriesQ) FilterByDestinationType(destinationTypes ...string) data.DeliveriesQ {
	q.sql = q.sql.Where(sq.Eq{"destination_type": destinationTypes})
	return q
}
