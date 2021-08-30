package whdbv1

import (
	"time"

	"github.com/LompeBoer/wh-notifs/internal/database"
)

func (d *Database) SelectOpenOrders() ([]string, error) {
	query := `
		SELECT m1.Symbol
		FROM PositionState m1 LEFT JOIN PositionState m2
		ON (m1.Symbol = m2.Symbol AND m1.Datetime < m2.Datetime)
		WHERE m2.Datetime IS NULL
		AND (m1.Status = 'Open' OR m1.Status = 'InitOpening' OR m1.Status = 'TPLimitPlacing' OR m1.Status = 'DCAOpening');
	`
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var i string
		if err := rows.Scan(
			&i,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (d *Database) SelectPositionStates() ([]database.PositionState, error) {
	rows, err := d.db.Query("SELECT LaunchId,Datetime,Symbol,Status,Side,BuyCount,Quantity,AveragePrice,TakeProfitPrice,StopLossPrice,TakeProfitLimitPrice,Reason FROM PositionState ORDER BY Datetime ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []database.PositionState
	for rows.Next() {
		var i database.PositionState
		if err := rows.Scan(
			&i.LaunchID,
			&i.DateTime,
			&i.Symbol,
			&i.Status,
			&i.Side,
			&i.BuyCount,
			&i.Quantity,
			&i.AveragePrice,
			&i.TakeProfitPrice,
			&i.StopLossPrice,
			&i.TakeProfitLimitPrice,
			&i.Reason,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (d *Database) SelectLatestTrades(since time.Time) ([]database.PositionState, error) {
	stmt, err := d.db.Prepare("SELECT LaunchId,Datetime,Symbol,Status,Side,BuyCount,Quantity,AveragePrice,TakeProfitPrice,StopLossPrice,TakeProfitLimitPrice,Reason FROM PositionState WHERE Datetime > ? ORDER BY Datetime ASC")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(since.Format("2006-01-02T15:04:05.999999999Z07:00"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []database.PositionState
	for rows.Next() {
		var i database.PositionState
		if err := rows.Scan(
			&i.LaunchID,
			&i.DateTime,
			&i.Symbol,
			&i.Status,
			&i.Side,
			&i.BuyCount,
			&i.Quantity,
			&i.AveragePrice,
			&i.TakeProfitPrice,
			&i.StopLossPrice,
			&i.TakeProfitLimitPrice,
			&i.Reason,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
