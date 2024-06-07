package main

import (
	"fmt"
)

type User struct {
	ID        int
	nama      string
	pass      string
	saldo     int
	perizinan bool
	admin     bool
}

type Transfer struct {
	fromID string
	toID   string
	jumlah int
}

type Transaction struct {
	ID     int
	jumlah int
	jenis  string
}

const maxUser int = 1000
const maxTransactions int = 10000

var activeUser [maxUser]User

var listtransaksi [maxTransactions]Transaction

var totalpengguna int = 5
var totaltransaksi int
var belumsetuju int
var registerAkun [maxUser]User

// pengguna yang telah melakukan registrasi dan akunnya aktif

func PenggunaAktif() {
	activeUser[0] = User{ID: 100000, nama: "John", pass: "John123", saldo: 10000, perizinan: true, admin: false}
	activeUser[1] = User{ID: 100001, nama: "Imelda", pass: "3535", saldo: 15000, perizinan: true, admin: false}
	activeUser[2] = User{ID: 100002, nama: "Zayn", pass: "Zayn01", saldo: 1500000, perizinan: true, admin: false}
	activeUser[3] = User{ID: 100003, nama: "Angel", pass: "2203", saldo: 8500000, perizinan: true, admin: true}
	activeUser[4] = User{ID: 100004, nama: "Nico", pass: "0407", saldo: 27062023, perizinan: true, admin: false}
}

func main() {
	PenggunaAktif()
	menu()
}

func menu() {
	var pilih int
	for pilih != 3 {
		fmt.Println("-----------------------------------------------")
		fmt.Println("Halo, selamat datang di Swift Finance")
		fmt.Println("Unggul dalam Layanan, Aman dalam Bertransaksi!")
		fmt.Println("-----------------------------------------------")
		fmt.Println("Silakan pilih menu dibawah ini:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Keluar")
		fmt.Print("Masukan menu yang ingin diakses: ")
		fmt.Scan(&pilih)
		fmt.Println("-----------------------------------------------")
		fmt.Println()
		if pilih == 1 {
			login()
		} else if pilih == 2 {
			registrasi()
			loginAdmin()
		}
	}
}

func menuAdmin(user, pass string) {
	var pilih int
	if cekAkunAdmin(user, pass) {
		for pilih != 4 {
			fmt.Println("-----------------------------------------------")
			fmt.Println("Pilih menu yang ingin diakses:")
			fmt.Println("1. Akun yang telah disetujui")
			fmt.Println("2. Akun yang belum disetujui")
			fmt.Println("3. Hapus Akun")
			fmt.Println("4. Keluar")
			fmt.Print("Pilih menu yang ingin diakses: ")
			fmt.Scan(&pilih)
			fmt.Println("-----------------------------------------------")
			fmt.Println()
			if pilih == 1 {
				cetaksemuaAkun()
			} else if pilih == 2 {
				persetujuanAkun()
			} else if pilih == 3 {
				hapusAkun()
			}
		}
	}
}

func menuUser(username, pass string) {
	var pilih int
	idx := cariAkun(username)
	for pilih != 5 {
		fmt.Println("-----------------------------------------------")
		fmt.Println("Pilih menu dibawah ini:")
		fmt.Println("1. Transfer")
		fmt.Println("2. Bayar")
		fmt.Println("3. Informasi Saldo")
		fmt.Println("4. Riwayat Transaksi")
		fmt.Println("5. Kembali ke halaman depan")
		fmt.Print("Pilih menu yang ingin diakses: ")
		fmt.Scan(&pilih)
		fmt.Println("-----------------------------------------------")
		fmt.Println()
		if pilih == 1 {
			kirimUang(idx)
		} else if pilih == 2 {
			pembayaran()
		} else if pilih == 3 {
			cetakSaldo(idx)
		} else if pilih == 4 {
			cetakTransaksi()
		}
	}
}

func cetaksemuaAkun() {
	for i := 0; i < totalpengguna; i++ {
		fmt.Println(activeUser[i].nama, activeUser[i].ID)
	}
}

func persetujuanAkun() {
	var ID int
	var persetujuan string
	var temp User
	for i := 0; i < belumsetuju; i++ {
		fmt.Println(registerAkun[i].nama, registerAkun[i].ID)
	}
	fmt.Print("Masukan ID yang ingin disetujui:")
	fmt.Scan(&ID)
	for i := 0; i < maxUser; i++ {
		if ID == totalpengguna+100000 {
			fmt.Print("Setuju/Tidak? ")
			fmt.Scan(&persetujuan)
			if persetujuan == "Setuju" {
				activeUser[totalpengguna] = registerAkun[0]
				activeUser[totalpengguna].perizinan = true
				totalpengguna++
				belumsetuju--
			}
			temp = registerAkun[ID%10]
			for ID > 0 && temp.ID > registerAkun[(ID%10)-1].ID {
				registerAkun[0] = registerAkun[(ID%10)-1]
				ID = ID - 1
			}
		}
	}
}

func registrasi() {
	var id int
	fmt.Println("Pengguna diminta untuk mengisi Username dan Password untuk mendaftar akun")
	id = belumsetuju
	registerAkun[id].ID = totalpengguna + 100000
	fmt.Println("ID akun Anda:", registerAkun[id].ID)
	fmt.Print("Masukan username: ")
	fmt.Scan(&registerAkun[id].nama)
	fmt.Print("Masukan password: ")
	fmt.Scan(&registerAkun[id].pass)
	registerAkun[id].saldo = 0
	belumsetuju++
	fmt.Println()
	cetakAkunRegister(id)
}

func cetakSaldo(id int) {
	var jawab string
	idx := id % 100000
	fmt.Println("-----------------------------------------------")
	fmt.Printf("Halo, Pengguna %s, sisa saldo Anda sekarang:\n", activeUser[(idx)].nama)
	fmt.Printf("Rp.%d, -\n", activeUser[(idx)].saldo)
	fmt.Print("Apakah Anda ingin melakukan Top-Up? (Iya/Tidak)? ")
	fmt.Scan(&jawab)
	fmt.Println("-----------------------------------------------")
	if jawab == "Iya" {
		topUp(id)
	}
}

func cetakAkunRegister(id int) {
	fmt.Println("Detail Akun:")
	fmt.Printf("ID: %d\n", registerAkun[id].ID)
	fmt.Printf("Nama: %s\n", registerAkun[id].nama)
	fmt.Printf("Saldo: %d\n", registerAkun[id].saldo)
	fmt.Println()
}

func cetakAkunTerdaftar(id int) {
	fmt.Println("Detail Akun:")
	fmt.Printf("ID: %d\n", activeUser[id].ID)
	fmt.Printf("Nama: %s\n", activeUser[id].nama)
	fmt.Printf("Saldo: %d\n", activeUser[id].saldo)
	fmt.Println()
}

func cekAkun(nama string, pass string) {
	// Sequential Search
	for i := 0; i < totalpengguna; i++ {
		if activeUser[i].nama == nama && activeUser[i].pass == pass && activeUser[i].perizinan && !activeUser[i].admin {
			fmt.Println("--- Login Berhasil ---")
		} else {
			if cekAkunAdmin(nama, pass) {
				menuAdmin(nama, pass)
			}
		}
	}
}

func cekAkunAdmin(user, pass string) bool {
	if user == activeUser[3].nama && pass == activeUser[3].pass {
		return true
	}
	return false
}

func login() {
	var username string
	var pass string
	fmt.Println("-----------------------------------------------")
	fmt.Println("Pengguna diminta untuk mengisi username dan password dengan benar")
	fmt.Print("Masukan username Anda: ")
	fmt.Scan(&username)
	fmt.Print("Masukan password Anda: ")
	fmt.Scan(&pass)
	fmt.Println("-----------------------------------------------")
	cekAkun(username, pass)
	menuUser(username, pass)
}

func loginAdmin() {
	var jawab, user, pass string
	fmt.Println("Akun Anda sedang dalam tahap verifikasi, mohon tunggu dalam 24 jam")
	fmt.Println("Anda dapat meninggalkan halaman ini selama proses verifikasi berlangsung")
	fmt.Println("Jika selama pembuatan akun Anda dibantu oleh Admin, tolong serahkan device Anda ke Admin")
	fmt.Println()
	fmt.Print("Apakah Anda berada di kantor pusat? (Ya/Tidak)? ")
	fmt.Scan(&jawab)
	if jawab == "Ya" {
		fmt.Println()
		fmt.Print("Masukan username Anda: ")
		fmt.Scan(&user)
		fmt.Print("Masukan password Anda: ")
		fmt.Scan(&pass)
		cekAkunAdmin(user, pass)
		menuAdmin(user, pass)
	}
}

func cariAkun(user string) int {
	for i := 0; i < totalpengguna; i++ {
		if activeUser[i].nama == user {
			return i // Mengembalikan indeks pengguna jika ditemukan
		}
	}
	return -1
}

func hapusAkun() {
	var ID, idx int
	fmt.Print("Masukan ID akun yang ingin dihapus: ")
	fmt.Scan(&ID)
	for i := 0; i < totalpengguna; i++ {
		if activeUser[i].ID == ID {
			idx = i
		}
	}
	cetakAkunTerdaftar(idx)
	for i := idx; i < totalpengguna-1; i++ {
		activeUser[i] = activeUser[i+1]
	}
	totalpengguna--
}

func kirimUang(id int) {
	var username string
	var jawab string
	var tf int
	fmt.Println("-----------------------------------------------")
	fmt.Print("Masukan Username yang ingin ditransfer: ")
	fmt.Scan(&username)
	fmt.Print("Masukan jumlah uang yang ingin ditransfer: ")
	fmt.Scan(&tf)
	for tf > activeUser[id].saldo {
		fmt.Println()
		fmt.Print("Saldo Anda tidak mencukupi, mohon ketik kembali nominal yang ingin ditransfer: ")
		fmt.Scan(&tf)
	}
	idx := cariAkun(username)
	activeUser[id].saldo = activeUser[id].saldo - tf
	activeUser[idx].saldo = activeUser[idx].saldo + tf
	fmt.Println("Apakah Anda ingin mencetak riwayat Transfer (Ya/Tidak)? ")
	fmt.Scan(&jawab)
	fmt.Println("-----------------------------------------------")
	fmt.Println()
	if jawab == "Ya" {
		fmt.Println("-----------------------------------------------")
		fmt.Printf("Transfer dari %d atas nama %s\n", activeUser[id].ID, activeUser[id].nama)
		fmt.Println("ke", activeUser[idx].nama, "sebanyak:")
		fmt.Printf("Rp. %d, -\n", tf)
		fmt.Println("-----------------------------------------------")
		fmt.Println()
	}
}

func topUp(id int) {
	var jumlahTopUp int
	fmt.Print("Masukkan jumlah saldo yang ingin ditambahkan: ")
	fmt.Scan(&jumlahTopUp)
	activeUser[id].saldo += jumlahTopUp
	fmt.Printf("Saldo berhasil ditambahkan. Saldo sekarang: Rp.%d,-\n", activeUser[id].saldo)
}

func opsiRiwayatTransaksi() {
	var jawab int
	for jawab != 4 {
		fmt.Println("-----------------------------------------------")
		fmt.Println("Data di atas telah di-update secara real time")
		fmt.Println("Pilih menu di bawah ini:")
		fmt.Println("1. Tampilkan dari nominal terbesar ke terkecil")
		fmt.Println("2. Tampilkan dari nominal terkecil ke terbesar")
		fmt.Println("3. Hapus riwayat Transaksi")
		fmt.Println("4. Keluar dari halaman ini")
		fmt.Print("Masukan jawaban Anda: ")
		fmt.Scan(&jawab)
		fmt.Println("-----------------------------------------------")
		fmt.Println()
		if jawab == 1 {
			sortTransaksiDescend()
			cetakTransaksi()
		} else if jawab == 2 {
			sortTransaksiAscend()
			cetakTransaksi()
		} else if jawab == 3 {
			hapusRiwayat()
			cetakTransaksi()
		}
	}
}

func cetakTransaksi() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("Riwayat Transaksi:")
	// Mencetak daftar riwayat transaksi
	for i := 0; i < totaltransaksi; i++ {
		transaction := listtransaksi[i]
		fmt.Printf("Kode Transaksi: %d\n", transaction.ID+1)
		fmt.Printf("Jumlah: Rp.%d,-\n", transaction.jumlah)
		fmt.Printf("Jenis: %s\n", transaction.jenis)
		fmt.Println("-----------------------------------------------")
		fmt.Println()
	}
	opsiRiwayatTransaksi()
}

func sortTransaksiDescend() {
	// Insertion Sort
	var temp Transaction
	i := 1
	for i <= len(listtransaksi)-1 {
		j := i
		temp = listtransaksi[j]
		for j > 0 && temp.jumlah > listtransaksi[j-1].jumlah {
			listtransaksi[j] = listtransaksi[j-1]
			j = j - 1
		}
		listtransaksi[j] = temp
		i = i + 1
	}
}

func sortTransaksiAscend() {
	// Selection Sort
	var idx int
	n := len(listtransaksi)
	for pass := 1; pass <= n; pass++ {
		idx = pass - 1
		for j := 0; j < n; j++ {
			if listtransaksi[idx].jumlah > listtransaksi[j].jumlah {
				idx = j
			}
		}
		temp := listtransaksi[idx]
		listtransaksi[idx] = listtransaksi[pass-1]
		listtransaksi[pass-1] = temp
	}

}

func hapusRiwayat() {
	var kode int
	fmt.Println("-----------------------------------------------")
	fmt.Print("Masukan kode transaksi yang ingin dihapus: ")
	fmt.Scan(&kode)
	for i := 0; i < totaltransaksi; i++ {
		if listtransaksi[i].ID+1 == kode {
			listtransaksi[i] = listtransaksi[i+1]
			totaltransaksi--
		}
	}
	fmt.Println("-----------------------------------------------")
}

func pembayaran() {
	var pilih int
	var harusbayar int
	var cetak, mau string
	for pilih != 5 {
		fmt.Println("-----------------------------------------------")
		fmt.Println("Pilih menu di bawah ini:")
		fmt.Println("1. Bayar Pulsa")
		fmt.Println("2. Bayar Listrik")
		fmt.Println("3. Bayar Makanan")
		fmt.Println("4. Bayar BPJS")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu pembayaran yang ingin diakses: ")
		fmt.Scan(&pilih)
		fmt.Println("-----------------------------------------------")
		fmt.Println()
		if pilih == 1 {
			fmt.Println("-----------------------------------------------")
			fmt.Print("Masukan jumlah uang yang harus dibayar: ")
			fmt.Scan(&harusbayar)
			if activeUser[totalpengguna-1].saldo > harusbayar {
				activeUser[totalpengguna-1].saldo -= harusbayar
				listtransaksi[totaltransaksi].ID = totaltransaksi
				listtransaksi[totaltransaksi].jumlah = harusbayar
				listtransaksi[totaltransaksi].jenis = "Pulsa"
				totaltransaksi++
				fmt.Print("Pembayaran Berhasil, apakah Anda ingin mencetak transaksi (Iya/Tidak)? ")
				fmt.Scan(&cetak)
				if cetak == "Iya" {
					cetakTransaksi()
				}
			} else {
				fmt.Println("Saldo Anda tidak mencukupi, mohon untuk melakukan Top-Up terlebih dahulu")
				fmt.Print("Apakah Anda ingin melakukan Top-Up (Ya/Tidak)? ")
				fmt.Scan(&mau)
				if mau == "Ya" {
					topUp(activeUser[totalpengguna-1].ID)
				}
			}
		} else if pilih == 2 {
			fmt.Println("-----------------------------------------------")
			fmt.Print("Masukan jumlah uang yang harus dibayar: ")
			fmt.Scan(&harusbayar)
			if activeUser[totalpengguna-1].saldo > harusbayar {
				activeUser[totalpengguna-1].saldo -= harusbayar
				listtransaksi[totaltransaksi].ID = totaltransaksi
				listtransaksi[totaltransaksi].jumlah = harusbayar
				listtransaksi[totaltransaksi].jenis = "Listrik"
				totaltransaksi++
				fmt.Print("Pembayaran Berhasil, apakah Anda ingin mencetak transaksi (Iya/Tidak)? ")
				fmt.Scan(&cetak)
				if cetak == "Iya" {
					cetakTransaksi()
				}
			} else {
				fmt.Println("Saldo Anda tidak mencukupi, mohon untuk melakukan Top-Up terlebih dahulu")
				fmt.Print("Apakah Anda ingin melakukan Top-Up (Ya/Tidak)? ")
				fmt.Scan(&mau)
				if mau == "Ya" {
					topUp(activeUser[totalpengguna-1].ID)
				}
			}
		} else if pilih == 3 {
			fmt.Println("-----------------------------------------------")
			fmt.Print("Masukan jumlah uang yang harus dibayar: ")
			fmt.Scan(&harusbayar)
			if activeUser[totalpengguna-1].saldo > harusbayar {
				activeUser[totalpengguna-1].saldo -= harusbayar
				listtransaksi[totaltransaksi].ID = totaltransaksi
				listtransaksi[totaltransaksi].jumlah = harusbayar
				listtransaksi[totaltransaksi].jenis = "Makanan"
				totaltransaksi++
				fmt.Print("Pembayaran Berhasil, apakah Anda ingin mencetak transaksi (Iya/Tidak)? ")
				fmt.Scan(&cetak)
				if cetak == "Iya" {
					cetakTransaksi()
				}
			} else {
				fmt.Println("Saldo Anda tidak mencukupi, mohon untuk melakukan Top-Up terlebih dahulu")
				fmt.Print("Apakah Anda ingin melakukan Top-Up (Ya/Tidak)? ")
				fmt.Scan(&mau)
				if mau == "Ya" {
					topUp(activeUser[totalpengguna-1].ID)
				}
			}
		} else if pilih == 4 {
			fmt.Println("-----------------------------------------------")
			fmt.Print("Masukan jumlah uang yang harus dibayar: ")
			fmt.Scan(&harusbayar)
			if activeUser[totalpengguna-1].saldo > harusbayar {
				activeUser[totalpengguna-1].saldo -= harusbayar
				listtransaksi[totaltransaksi].ID = totaltransaksi
				listtransaksi[totaltransaksi].jumlah = harusbayar
				listtransaksi[totaltransaksi].jenis = "BPJS"
				totaltransaksi++
				fmt.Print("Pembayaran Berhasil, apakah Anda ingin mencetak transaksi (Iya/Tidak)? ")
				fmt.Scan(&cetak)
				if cetak == "Iya" {
					cetakTransaksi()
				}
			} else {
				fmt.Println("Saldo Anda tidak mencukupi, mohon untuk melakukan Top-Up terlebih dahulu")
				fmt.Print("Apakah Anda ingin melakukan Top-Up (Ya/Tidak)? ")
				fmt.Scan(&mau)
				if mau == "Ya" {
					topUp(activeUser[totalpengguna-1].ID)
				}
			}
		}
	}
}
