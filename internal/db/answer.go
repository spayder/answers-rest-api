package db

import (
	"context"
	"database/sql"
	"fmt"
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
