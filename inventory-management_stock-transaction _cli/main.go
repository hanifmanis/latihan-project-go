package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var history []Transaction

type Item struct {
	ID    string //Kode unik barang (contoh: "BRG01", "BRG02").
	Name  string
	Stock int //Jumlah ketersediaan stok saat ini.
	Price int //Harga satuan barang.
}

type Transaction struct {
	ItemID    string    // Merujuk pada ID barang yang ditransaksikan.
	Type      string    // Jenis transaksi, bernilai "IN" (stok masuk) atau "OUT" (stok keluar).
	Quantity  int       // Jumlah barang yang dimasukkan atau dikeluarkan.
	Timestamp time.Time // Waktu transaksi dilakukan.
}

func AddStock(items map[string]Item, id string, qty int) {
	// history := []Transaction{}

	for index := range items {
		if items[index].Name == id {
			fmt.Printf("Menambahkan %d %s ke dalam stok\n", qty, id)
			item := items[index]
			item.Stock += qty
			items[index] = item

			RecordTransaction(history, item.ID, item.Name, item.Stock)
		}
	}
	// handling kalau angka stock ga valid -> ulangi dari pilih item dari awal

}

func SellItem(items map[string]Item, id string, qty int) error {
	for index := range items {
		if items[index].Name == id {
			item := items[index]

			if item.Stock < qty {
				return fmt.Errorf("qty tidak cukup")
			}

			fmt.Printf("Menjual %d %s dari stok\n", qty, id)
			item.Stock -= qty
			items[index] = item
		}
	}
	return nil

	//belum handle error kalau stock kurang- DONE
}

func RecordTransaction(history []Transaction, itemID string, tType string, qty int) {
	history = append(history, Transaction{
		ItemID:    itemID,
		Type:      tType,
		Quantity:  qty,
		Timestamp: time.Now(),
	})
	fmt.Println(history)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inventory := map[string]Item{
		"BRG01": {ID: "BRG01", Name: "laptop", Stock: 10, Price: 8500000},
		"BRG02": {ID: "BRG02", Name: "mouse", Stock: 25, Price: 150000},
		"BRG03": {ID: "BRG03", Name: "keyboard", Stock: 15, Price: 450000},
		"BRG04": {ID: "BRG04", Name: "monitor", Stock: 8, Price: 2100000},
		"BRG05": {ID: "BRG05", Name: "headset", Stock: 12, Price: 350000},
	}

	// history := map[string]Transaction{}

	// Contoh pengecekan data awal
	fmt.Println("=== Data Inisial Inventaris Gudang ===")
	for id, item := range inventory {
		fmt.Printf("[%s] %s | Stok: %d | Harga: Rp%d\n", id, item.Name, item.Stock, item.Price)
	}

	for {
		fmt.Println("Ketik 1 untuk tambah stok, 2 Jual barang, 3 keluar/rekap")
		opsi, _ := reader.ReadString('\n')
		opsi = strings.TrimSpace(strings.ToLower(opsi))

		if opsi != "1" && opsi != "2" && opsi != "3" {
			fmt.Println("Opsi invalid, ketik 1,2 atau 3 saja")
			continue
		}

		if opsi == "3" {
			break
		}

		if opsi == "1" {

			for {
				fmt.Println("Stock apa yang mau kamu tambah?")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(strings.ToLower(input))

				var itemDitemukan bool

				for _, item := range inventory {
					if item.Name == input {
						itemDitemukan = true
						break
					}
				}

				if !itemDitemukan {
					fmt.Println("Item tidak ditemukan, add stock lain!")
					continue
				}

				fmt.Printf("Berapa banyak %s yang ingin ditambah? ", input)
				qtyInput, _ := reader.ReadString('\n')
				qtyInput = strings.TrimSpace(qtyInput)

				qty, err := strconv.Atoi(qtyInput)
				if err != nil || qty <= 0 {
					fmt.Println("Jumlah tidak valid, masukkan angka lebih dari 0.")
					continue
				}
				AddStock(inventory, input, qty)
				break
			}
		}

		if opsi == "2" {
			for {
				fmt.Println("Stock apa yang terjual?")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(strings.ToLower(input))

				var itemDitemukan bool

				for _, item := range inventory {
					if item.Name == input {
						itemDitemukan = true
						break
					}
				}

				if !itemDitemukan {
					fmt.Println("Item tidak ditemukan, pilih stock lain!")
					continue
				}

				fmt.Printf("Berapa banyak %s yang ingin dikurangi? ", input)
				qtyInput, _ := reader.ReadString('\n')
				qtyInput = strings.TrimSpace(qtyInput)

				qty, err := strconv.Atoi(qtyInput)
				if err != nil || qty <= 0 {
					fmt.Println("Jumlah tidak valid, masukkan angka lebih dari 0.")
					continue
				}

				// SellItem(inventory, input, qty)
				error := SellItem(inventory, input, qty)
				if error != nil {
					// If an error is returned, print the error message
					fmt.Println("Error:", error) // Output: Error: ""
					continue
				}
				break

			}
		}

	}

	fmt.Println("\n=== Data Gudang Terbaru ===")
	for id, item := range inventory {
		fmt.Printf("[%s] %s | Stok: %d | Harga: Rp%d\n", id, item.Name, item.Stock, item.Price)
	}
	fmt.Println(history)
}
