package kasir

import (
    "fmt"
)

// Keranjang adalah interface untuk keranjang belanja pelanggan.
type Keranjang interface {
    TambahBarang(nama string, harga float64, quantity int)
    HapusBarang(nama string, quantity int)
    CetakStruk()
}

// Item adalah struktur data untuk menyimpan informasi tentang barang.
type Item struct {
    Name     string
    Price    float64
    Quantity int
}

// Kasir adalah struct untuk menyimpan daftar belanjaan.
type Kasir struct {
    Struk []Item
}

// daftarBarang adalah map untuk menyimpan harga barang berdasarkan nama barang
var daftarBarang = map[string]float64{
    "fanta":         5000,
    "cokelat":       10000,
    "parfum":        150000,
    "minyak goreng": 12000,
    "mie instan":    3000,
}

// DaftarHarga mengembalikan harga barang berdasarkan nama barang.
func DaftarHarga(nama string) (float64, bool) {
    harga, ok := daftarBarang[nama]
    return harga, ok
}

// TambahBarang menambahkan item ke struk belanjaan.
func (k *Kasir) TambahBarang(nama string, harga float64, quantity int) {
    if _, ok := daftarBarang[nama]; !ok {
        fmt.Printf("Barang %s tidak ditemukan\n", nama)
        return
    }

    for i := range k.Struk {
        if k.Struk[i].Name == nama {
            k.Struk[i].Quantity += quantity
            return
        }
    }
    k.Struk = append(k.Struk, Item{Name: nama, Price: harga, Quantity: quantity})
}

// HapusBarang mengurangi quantity item dari struk belanjaan.
func (k *Kasir) HapusBarang(nama string, quantity int) {
    for i := range k.Struk {
        if k.Struk[i].Name == nama {
            if k.Struk[i].Quantity > quantity {
                k.Struk[i].Quantity -= quantity
            } else {
                k.Struk = append(k.Struk[:i], k.Struk[i+1:]...)
            }
            return
        }
    }
}

// CetakStruk mencetak struk belanjaan.
func (k *Kasir) CetakStruk() {
    fmt.Println("===== Struk Belanjaan =====")
    var total float64
    for _, item := range k.Struk {
        total += item.Price * float64(item.Quantity)
        fmt.Printf("%s (Qty: %d) - Rp%.2f/pcs - Rp%.2f\n", item.Name, item.Quantity, item.Price, item.Price*float64(item.Quantity))
    }
    fmt.Printf("Total Belanja: Rp%.2f\n", total)
}
