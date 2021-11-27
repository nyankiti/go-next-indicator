package models

type Product struct {
	// 以下のようにModelを書くことで、全てのModelに共通の要素をまとめて扱える
	// ※Modelは同階層のgo ファイルのstructのため、importする必要がない
	Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}
