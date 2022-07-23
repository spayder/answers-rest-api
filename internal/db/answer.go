package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/spayder/answers-rest-api/internal/answer"
)

type AnswerRow struct {
	ID    string
	Key   sql.NullString
	Value sql.NullString
}

func convertAnswerRowToAnswer(ar AnswerRow) answer.Answer {
	return answer.Answer{
		ID:    ar.ID,
		Key:   ar.Key.String,
		Value: ar.Value.String,
	}
}

func (d *Database) GetAnswer(ctx context.Context, uuid string) (answer.Answer, error) {
	var ansRow AnswerRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, key, value FROM answers WHERE id = $1;`,
		uuid,
	)

	err := row.Scan(&ansRow.ID, &ansRow.Key, &ansRow.Value)
	if err != nil {
		return answer.Answer{}, fmt.Errorf("could not fetch the answer with given id: %w", err)
	}

	return convertAnswerRowToAnswer(ansRow), nil
}

func (d *Database) PostAnswer(ctx context.Context, ans answer.Answer) (answer.Answer, error) {
	ans.ID = uuid.NewV4().String()
	postRow := AnswerRow{
		ID:    ans.ID,
		Key:   sql.NullString{String: ans.Key, Valid: true},
		Value: sql.NullString{String: ans.Value, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO answers (id, key, value) VALUES (:id, :key, :value)`,
		postRow,
	)

	if err != nil {
		return answer.Answer{}, fmt.Errorf("could not insert the answer to the database: %w", err)
	}

	if err = rows.Close(); err != nil {
		return answer.Answer{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return ans, nil
}

func (d *Database) DeleteAnswer(ctx context.Context, uuid string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM answers WHERE id = $1`,
		uuid,
	)

	if err != nil {
		return fmt.Errorf("failed to delete answer: %w", err)
	}

	return nil
}
