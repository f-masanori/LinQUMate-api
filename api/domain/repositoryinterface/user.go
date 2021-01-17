package repositoryinterface

// UserRepository : UserRepositoryのインターフェイス
// usecase・repositoryによる実装はここで定義するインターフェイスに依存する
// /api/repository内に実際の実装を書く。それによって、BDに依存しない設計になる
type UserRepository interface {
	Create(string, string, string) (int, error)
}
