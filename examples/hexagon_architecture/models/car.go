package models

type Car struct {
	ID           int64  `db:"id" json:"id"`
	LicensePlate string `db:"license_place" json:"license_plate"`
	Make         string `db:"make" json:"make"`
	Model        string `db:"model" json:"model"`
	Year         int    `db:"year" json:"year"`
}
