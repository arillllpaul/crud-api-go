package model

type Produk struct {
    ID       int     `json:"id"`
    Nama     string  `json:"nama"`
    Harga    float64 `json:"harga"`
    Kategori string  `json:"kategori"`
}
