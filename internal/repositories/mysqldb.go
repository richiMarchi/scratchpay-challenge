package repositories

import (
	"database/sql"

	"github.com/richiMarchi/scratchpay-challenge/internal/core/domain"

	_ "github.com/go-sql-driver/mysql"
)

type mysqldb struct {
	dbConn *sql.DB
}

func NewMySqlDb(user, password, dbname string) *mysqldb {
	db, err := sql.Open("mysql", user+":"+password+"@/"+dbname)
	if err != nil {
		panic(err)
	}

	verifyDbInit(db)

	return &mysqldb{
		dbConn: db,
	}
}

func (db *mysqldb) Create(user domain.User) error {
	stmtIns, err := db.dbConn.Prepare("INSERT INTO users VALUES( ?, ? )")
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(user.Id, user.Name)
	if err != nil {
		return err
	}

	return nil
}

func (db *mysqldb) Get(id uint) (domain.User, error) {
	stmt, err := db.dbConn.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		return domain.User{}, err
	}
	defer stmt.Close()

	user := domain.User{}
	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (db *mysqldb) GetAll() ([]domain.User, error) {
	result, err := db.dbConn.Query("SELECT * FROM users")
	if err != nil {
		return []domain.User{}, err
	}
	defer result.Close()

	var users []domain.User
	for result.Next() {
		user := domain.User{}
		result.Scan(&user.Id, &user.Name)
		users = append(users, user)
	}

	return users, nil
}

func verifyDbInit(db *sql.DB) {
	result, err := db.Query("CREATE TABLE IF NOT EXISTS users (id INT PRIMARY KEY, name VARCHAR(255) NOT NULL)")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
}
