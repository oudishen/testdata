package pkg

import "time"

type CommonField struct {
	ID int64     `db:"id" json:"id"` // ID is the primary key.
	CT time.Time `db:"ct" json:"ct"` // CT is the create time.
	MT time.Time `db:"mt" json:"mt"` // MT is the modify time.
	DT int64     `db:"dt" json:"dt"` // DT is the delete time.
	V  int64     `db:"v" json:"v"`   // V is the version.
}
