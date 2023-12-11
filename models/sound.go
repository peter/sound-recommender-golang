package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type Sound struct {
	ID    int
	Title string
	// TODO: need to handle text array with pq.StringArray or use JSONB
	// https://stackoverflow.com/questions/63256680/adding-an-array-of-integers-as-a-data-type-in-a-gorm-model
	// https://stackoverflow.com/questions/65825144/migrating-problem-using-psql-array-for-pq-to-pgx
	Genres            pq.StringArray `gorm:"type:text[]"`
	BPM               int
	DurationInSeconds int
	// JSONB data
	// https://pkg.go.dev/gorm.io/datatypes#section-readme
	// https://jorgequiterio.medium.com/how-to-manage-postgresql-json-data-with-go-golang-fb15883ef11e
	// https://gist.github.com/yanmhlv/d00aa61082d3b8d71bed
	// https://stackoverflow.com/questions/65434115/how-to-insert-data-in-jsonb-field-of-postgres-using-gorm
	Meta datatypes.JSONMap
	// openai_embedding vector(1536),
	Description     string
	DescriptionMini string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
