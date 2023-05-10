package db

import (
	"context"

	"github.com/google/uuid"
)

type ClockOutTxParams struct {
	EmployeeID uuid.UUID `json:"employee_id"`
	ClockInID  int32     `json:"clock_in_id"`
}

type ClockOutTxResult struct {
	ClockOut ClockOut `json:"clock_out"`
}

func (store *SQLStore) ClockOutTx(ctx context.Context, arg ClockOutTxParams) (ClockOutTxResult, error) {
	var result ClockOutTxResult

	// transaction to create clockout and update clockin
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.ClockOut, err = q.CreateClockOut(ctx, CreateClockOutParams(arg))
		if err != nil {
			return err
		}

		err = q.UpdateClockIn(ctx, UpdateClockInParams{
			ID:         arg.ClockInID,
			ClockedOut: true,
		})
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
