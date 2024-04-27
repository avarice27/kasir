package main

import (
    "example/belanjaan/kasir"
)

func main() {
    // Membuat kasir baru
    var k kasir.Keranjang = &kasir.Kasir{}

    // Barang-barang yang akan dibeli oleh pelanggan
    barangBelanja := []string{"fanta", "cokelat", "parfum", "minyak goreng", "mie instan"}

    // Menambahkan barang-barang ke keranjang belanja
    for _, nama := range barangBelanja {
        // Mendapatkan harga barang dari daftar harga
        harga, ok := kasir.DaftarHarga(nama)
        if !ok {
            continue // Skip barang jika tidak ada di daftar harga
        }
        // Tambahkan barang dengan harga dan quantity 1
        k.TambahBarang(nama, harga, 1)
    }

    // Mencetak struk belanjaan
    k.CetakStruk()
}
