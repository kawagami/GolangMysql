package mysql

type MInterface interface {
	Get()
	Insert()
	// InsertAll()
	Exist()
}
