package mysql

import (
	"linqumate/domain/repositoryinterface"

	"github.com/jinzhu/gorm"
)

// UserRepository : /api/repository/mysql配下のuser.goはmysqlを使って永続化を行うための処理を書くため、それに必要な*gorm.DBをもつ構造体を定義する
// UserRepository構造体に紐づくメソッド(レシーバー) は、このSqlHandlerを使ってmysqlに対してデータの保存や読み込みを行う
type UserRepository struct {
	SQLHandler *gorm.DB
}

// NewMysqlUserRepository : mysqlを使ってデータの永続化を行うuserレポジトリのインスタンスを作成する
func NewMysqlUserRepository(db *gorm.DB) repositoryinterface.UserRepository {
	return &UserRepository{
		SQLHandler: db,
	}
}

// Create : ユーザーをmysqlに登録
func (m *UserRepository) Create(uid string, email string, name string) (int, error) {
	// ポイントレシーバーを使ってデータを永続化
	return 0, nil
}
