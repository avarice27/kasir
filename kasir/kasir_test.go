package kasir

import (
    "testing"
)

func TestKasir(t *testing.T) {
    kasir := &Kasir{}

    // Tambahkan barang-barang ke dalam keranjang belanja
    kasir.TambahBarang("fanta", 5000, 2)
    kasir.TambahBarang("cokelat", 10000, 1)
    kasir.TambahBarang("parfum", 150000, 1)
    kasir.TambahBarang("minyak goreng", 12000, 1)
    kasir.TambahBarang("mie instan", 3000, 1)

    // Struktur yang diharapkan setelah penambahan barang
    expected := []Item{
        {Name: "fanta", Price: 5000, Quantity: 2},
        {Name: "cokelat", Price: 10000, Quantity: 1},
        {Name: "parfum", Price: 150000, Quantity: 1},
        {Name: "minyak goreng", Price: 12000, Quantity: 1},
        {Name: "mie instan", Price: 3000, Quantity: 1},
    }

    // Periksa apakah struk belanja sesuai dengan yang diharapkan
    for i, item := range kasir.Struk {
        if item.Name != expected[i].Name || item.Price != expected[i].Price || item.Quantity != expected[i].Quantity {
            t.Errorf("Barang %s: harga atau jumlah tidak sesuai, got: %v, want: %v", item.Name, item, expected[i])
        }
    }
}
