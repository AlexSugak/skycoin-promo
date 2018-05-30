package geo

// Country contains country code and name
type Country struct {
	ISO  string `db:"ISO"`
	Name string `db:"Name"`
}
