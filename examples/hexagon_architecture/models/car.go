package models

type Car struct {
	ID           int64  `db:"id" json:"id"`
	LicensePlate string `db:"license_place" json:"license_plate"`
	Make         string `db:"make" json:"make"`
	Year         int    `db:"year" json:"year"`
	Crashed      bool   `db:"crashed" json:"crashed" valid:"type(bool)"`
}
