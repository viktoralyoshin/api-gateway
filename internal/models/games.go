package models

import (
	"time"
)

type Game struct {
	// Внутренний ID (UUID)
	ID string `json:"id" db:"id"`

	// ID из IGDB (в API IGDB это число, но в Proto у тебя строка.
	// Если в базе хранишь как строку, оставляй string. Если int - поменяй тут)
	IGDBID string `json:"igdb_id" db:"igdb_id"`

	Name    string  `json:"name" db:"name"`
	Slug    string  `json:"slug" db:"slug"`
	Summary string  `json:"summary" db:"summary"`
	Rating  float64 `json:"rating" db:"rating"`
	Hypes   int     `json:"hypes" db:"hypes"` // uint32 в Go неудобен, лучше int

	// Дата может быть null, поэтому указатель или sql.NullString,
	// но для JSON лучше просто string
	FirstReleaseDate string `json:"first_release_date" db:"first_release_date"`

	// Массивы/Слайсы
	// В PostgreSQL они хранятся как TEXT[], в Go это []string
	ReleaseDates []string `json:"release_dates" db:"release_dates"`
	CoverURL     string   `json:"cover_url" db:"cover_url"`
	ArtworkURLs  []string `json:"artwork_urls" db:"artwork_urls"`
	Screenshots  []string `json:"screenshots" db:"screenshots"`
	Genres       []string `json:"genres" db:"genres"`
	Themes       []string `json:"themes" db:"themes"`
	Platforms    []string `json:"platforms" db:"platforms"`

	// Время
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
