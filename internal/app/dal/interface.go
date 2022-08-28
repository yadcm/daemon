package dal

type dbName string

func (dbn dbName) String() string {
	return string(dbn) + ".db"
}

const (
	dbNameClient dbName = "client"
)

type DAL interface {
}
