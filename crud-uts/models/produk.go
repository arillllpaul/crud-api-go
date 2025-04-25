// models/produk.go
package models

type Produk struct {
	ID         int    `json:"id"`
	Nama       string `json:"nama"`
	Harga      int    `json:"harga"`
	KategoriID int    `json:"kategori_id"`
}

type ProdukWithKategori struct {
	Produk
	KategoriNama string `json:"kategori_nama"`
}