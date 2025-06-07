package main

import "fmt"

type app struct {
	judul          string
	platform       string
	kategori       string
	tanggalPosting string
	jamPosting     string
	jumlahLike     int
	jumlahKomentar int
	jumlahShare    int
}

const maxKonten = 100
func panjang(s string) int {
	var i int
	i = 0
	for range s {
		i++
	}
	return i
}

func contains(s, sub string) bool {
	var i, j, countS, countSub int
	var cocok bool

	countS = panjang(s)
	countSub = panjang(sub)

	for i = 0; i <= countS - countSub; i++ {
		cocok = true
		for j = 0; j < countSub; j++ {
			if s[i+j] != sub[j] {
				cocok = false
			}
		}
		if cocok {
			return true
		}
	}
	return false
}
func main() {
	var dataKonten [maxKonten]app
	var n int = 0
	var pilihan int
	var repeat bool = false

	repeat = true
	for repeat {
		fmt.Println("\n====== SELAMAT DATANG DI APP MANAJER ======")
		fmt.Println("\n================ MENU UTAMA ===============")
		fmt.Println("1. Tambah Konten")
		fmt.Println("2. Ubah Konten")
		fmt.Println("3. Hapus Konten")
		fmt.Println("4. Tampilkan Semua Konten")
		fmt.Println("5. Jadwalkan Konten (Reschedule)")
		fmt.Println("6. Analisis Engagement Tertinggi")
		fmt.Println("7. Urutkan berdasarkan Tanggal")
		fmt.Println("8. Urutkan berdasarkan Engagement")
		fmt.Println("9. Cari Ide Konten")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			tambahKonten(&dataKonten, &n)
		} else if pilihan == 2 {
			ubahKonten(&dataKonten, n)
		} else if pilihan == 3 {
			hapusKonten(&dataKonten, &n)
		} else if pilihan == 4 {
			tampilkanKonten(dataKonten, n)
		} else if pilihan == 5 {
			jadwalkanKonten(&dataKonten, n)
		} else if pilihan == 6 {
			analisisEngagementTertinggi(dataKonten, n)
		} else if pilihan == 7 {
			urutkanTanggalSelection(&dataKonten, n)
		} else if pilihan == 8 {
			urutkanEngagementInsertion(&dataKonten, n)
		} else if pilihan == 9 {
			var opsiCari int
			var input string

			fmt.Println("\n================ Menu Pencarian Ide Konten ===============")
			fmt.Println("1. Cari Berdasarkan Keyword Judul")
			fmt.Println("2. Cari Berdasarkan Kategori")
			fmt.Println("3. Cari Berdasarkan Judul Lengkap")
			fmt.Println("Pilih Jenis Pencarian : ")
			fmt.Scanln(&opsiCari)

			if opsiCari == 1 {
				fmt.Print("Masukan Keyword Yang Dicari : ")
				fmt.Scanln(&input)
				searchByKeyword(dataKonten, n, input)
			} else if opsiCari == 2 {
				fmt.Print("Masukan Kategori Yang Dicari : ")
				fmt.Scanln(&input)
				searchByKeyword(dataKonten, n, input)
			} else if opsiCari == 3 {
				var index int
				
				fmt.Print("Masukan Judul Lengkap (Pastikan Sudah Terurut) : ")
				fmt.Scanln(&input)

				index = binarySearchJudul(dataKonten, n, input)

				if index != -1 {
					fmt.Println("Judul : ", dataKonten[index].judul)
					fmt.Println("Kategori : ", dataKonten[index].kategori)
					fmt.Println("Platform : ", dataKonten[index].platform)
				} else {
					fmt.Println("!!!!! Judul Tidak Ditemukan !!!!!")
				}
			}
		} else if pilihan == 0 {
			fmt.Println("Terima kasih! Keluar dari program...")
			repeat = false
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahKonten(data *[maxKonten]app, n *int) {
	var kontenBaru app
	var pilihan int

	if *n >= maxKonten {
		fmt.Println("Data sudah penuh, tidak bisa menambah konten lagi.")
		return
	}

	fmt.Println("\n=== Tambah Konten Baru ===")
	fmt.Print("Judul: ")
	fmt.Scanln(&kontenBaru.judul)

	fmt.Println("Pilih Platform:")
	fmt.Println("1. Instagram")
	fmt.Println("2. TikTok")
	fmt.Println("3. YouTube")
	fmt.Print("Pilihan (1-3): ")
	fmt.Scanln(&pilihan)
	if pilihan == 1 {
		kontenBaru.platform = "Instagram"
	} else if pilihan == 2 {
		kontenBaru.platform = "TikTok"
	} else if pilihan == 3 {
		kontenBaru.platform = "YouTube"
	} else {
		fmt.Println("Pilihan tidak valid, default ke 'Instagram'")
		kontenBaru.platform = "Instagram"
	}

	fmt.Println("Pilih Kategori:")
	fmt.Println("1. Edukasi")
	fmt.Println("2. Hiburan")
	fmt.Println("3. Promosi")
	fmt.Print("Pilihan (1-3): ")
	fmt.Scanln(&pilihan)
	if pilihan == 1 {
		kontenBaru.kategori = "Edukasi"
	} else if pilihan == 2 {
		kontenBaru.kategori = "Hiburan"
	} else if pilihan == 3 {
		kontenBaru.kategori = "Promosi"
	} else {
		fmt.Println("Pilihan tidak valid, default ke 'Edukasi'")
		kontenBaru.kategori = "Edukasi"
	}

	fmt.Print("Tanggal Posting (DD-MM-YYYY): ")
	fmt.Scanln(&kontenBaru.tanggalPosting)

	fmt.Print("Jam Posting (HH:MM): ")
	fmt.Scanln(&kontenBaru.jamPosting)

	fmt.Print("Jumlah Like: ")
	fmt.Scanln(&kontenBaru.jumlahLike)

	fmt.Print("Jumlah Komentar: ")
	fmt.Scanln(&kontenBaru.jumlahKomentar)

	fmt.Print("Jumlah Share: ")
	fmt.Scanln(&kontenBaru.jumlahShare)

	data[*n] = kontenBaru
	*n++
	fmt.Println("Konten berhasil ditambahkan!")
}

func ubahKonten(data *[maxKonten]app, n int) {
	var judulCari string
	var i, pilihan int
	var ditemukan bool = false
	
	if n == 0 {
		fmt.Println("\n=======================================================================")
		fmt.Println("Tidak Ada Data Yang Tersimpan, Silahkan Tambahkan Data Terlebih Dahulu.")
		fmt.Println("\n=======================================================================")
	} else {
		fmt.Print("Masukkan judul konten yang ingin diubah: ")
		fmt.Scanln(&judulCari)
		for i = 0; i < n; i++ {
			if data[i].judul == judulCari && !ditemukan {
				fmt.Print("Judul baru: ")
				fmt.Scanln(&data[i].judul)

				fmt.Println("Pilih Platform baru:")
				fmt.Println("1. Instagram")
				fmt.Println("2. TikTok")
				fmt.Println("3. YouTube")
				fmt.Print("Pilihan (1-3): ")
				fmt.Scanln(&pilihan)
				if pilihan == 1 {
					data[i].platform = "Instagram"
				} else if pilihan == 2 {
					data[i].platform = "TikTok"
				} else if pilihan == 3 {
					data[i].platform = "YouTube"
				} else {
					fmt.Println("Pilihan tidak valid, default ke 'Instagram'")
					data[i].platform = "Instagram"
				}

				fmt.Println("Pilih Kategori baru:")
				fmt.Println("1. Edukasi")
				fmt.Println("2. Hiburan")
				fmt.Println("3. Promosi")
				fmt.Print("Pilihan (1-3): ")
				fmt.Scanln(&pilihan)
				if pilihan == 1 {
					data[i].kategori = "Edukasi"
				} else if pilihan == 2 {
					data[i].kategori = "Hiburan"
				} else if pilihan == 3 {
					data[i].kategori = "Promosi"
				} else {
					fmt.Println("Pilihan tidak valid, default ke 'Edukasi'")
					data[i].kategori = "Edukasi"
				}

				fmt.Print("Tanggal baru: ")
				fmt.Scanln(&data[i].tanggalPosting)

				fmt.Print("Jam baru: ")
				fmt.Scanln(&data[i].jamPosting)

				fmt.Print("Jumlah Like: ")
				fmt.Scanln(&data[i].jumlahLike)

				fmt.Print("Jumlah Komentar: ")
				fmt.Scanln(&data[i].jumlahKomentar)

				fmt.Print("Jumlah Share: ")
				fmt.Scanln(&data[i].jumlahShare)

				fmt.Println("Konten berhasil diubah!")
				ditemukan = true
			}
		}
		if !ditemukan {
			fmt.Println("Konten tidak ditemukan.")
		}
	}
}

func hapusKonten(data *[maxKonten]app, n *int) {
	if *n == 0 {
		fmt.Println("\n=======================================================================")
		fmt.Println("Tidak Ada Data Yang Tersimpan, Silahkan Tambahkan Data Terlebih Dahulu.")
		fmt.Println("\n=======================================================================")
	} else {
		var judulCari string
		var i, j int
		var ditemukan bool = false

		fmt.Print("Masukkan judul konten yang ingin dihapus: ")
		fmt.Scanln(&judulCari)

		for i = 0; i < *n; i++ {
			if data[i].judul == judulCari && !ditemukan {
				for j = i; j < *n-1; j++ {
					data[j] = data[j+1]
				}
				*n--
				fmt.Println("Konten berhasil dihapus!")
				ditemukan = true
			}
		}

		if !ditemukan {
			fmt.Println("Konten tidak ditemukan.")
		}
	}
}

func tampilkanKonten(data [maxKonten]app, n int) {
	if n == 0 {
		fmt.Println("\n=======================================================================")
		fmt.Println("Tidak Ada Data Yang Tersimpan, Silahkan Tambahkan Data Terlebih Dahulu.")
		fmt.Println("=======================================================================")
	} else {
		fmt.Println("\n=== Daftar Semua Konten ===")
		for i := 0; i < n; i++ {
			fmt.Printf("%d. %s - %s [%s] pada %s %s | Like: %d, Komentar: %d, Share: %d\n",
				i+1, data[i].judul, data[i].platform, data[i].kategori,
				data[i].tanggalPosting, data[i].jamPosting,
				data[i].jumlahLike, data[i].jumlahKomentar, data[i].jumlahShare)
		}
	}
}

func binarySearchJudul(data [maxKonten]app, n int, target string) int {
	var kiri, kanan, tengah int
	kiri = 0
	kanan = n - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if data[tengah].judul == target {
			return tengah
		} else if data[tengah].judul < target {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func searchByKeyword(data [maxKonten]app, n int, keyword string) {
	var ditemukan bool = false
	fmt.Println("Hasil pencarian berdasarkan kata kunci di judul:")
	for i := 0; i < n; i++ {
		if contains(data[i].judul, keyword) {
			fmt.Printf("- [%d] %s (Kategori: %s, Platform: %s)\n", i+1, data[i].judul, data[i].kategori, data[i].platform)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ada konten dengan kata kunci tersebut.")
	}
}

func searchByKategori(data [maxKonten]app, n int, kategori string) {
	var ditemukan bool = false
	fmt.Println("Hasil pencarian berdasarkan kategori:")
	for i := 0; i < n; i++ {
		if data[i].kategori == kategori {
			fmt.Printf("- [%d] %s (Platform: %s, Jadwal: %s %s)\n", i+1, data[i].judul, data[i].platform, data[i].tanggalPosting, data[i].jamPosting)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ada konten dalam kategori tersebut!")
	}
}

func selectionSortJudul(data *[maxKonten]app, n int) {
	var i, j, min int
	var temp app
	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if data[j].judul < data[min].judul {
				min = j
			}
		}
		temp = data[i]
		data[i] = data[min]
		data[min] = temp
	}
}


func jadwalkanKonten(data *[maxKonten]app, n int) {
	if n == 0 {
		fmt.Println("\n=======================================================================")
		fmt.Println("Tidak Ada Data Yang Tersimpan, Silahkan Tambahkan Data Terlebih Dahulu.")
		fmt.Println("\n=======================================================================")
	} else {
		fmt.Println("===================== Jadwal Ulang Konten ===============================")
		var judulCari string
		var i int
		var ditemukan bool = false

		fmt.Print("Masukkan judul konten yang ingin dijadwalkan ulang: ")
		fmt.Scanln(&judulCari)

		for i = 0; i < n; i++ {
			if data[i].judul == judulCari && !ditemukan {
				fmt.Print("Masukkan tanggal baru (DD-MM-YYYY): ")
				fmt.Scanln(&data[i].tanggalPosting)

				fmt.Print("Masukkan jam baru (HH:MM): ")
				fmt.Scanln(&data[i].jamPosting)

				fmt.Println("Konten Berhasil Dijadwalkan Ulang!")
				ditemukan = true
			}
		}

		if !ditemukan {
			fmt.Println("Konten Tidak Ditemukan.")
		}
	}
}

func analisisEngagementTertinggi(data [maxKonten]app, n int) {
	var tanggalMulai, tanggalAkhir string
	var i int
	var maxIndex int = -1
	var engagementSaatIni int = -1
	var engagementTertinggi int = -1

	if n == 0 {
		fmt.Println("\n=======================================================================")
		fmt.Println("Tidak Ada Data Yang Tersimpan, Silahkan Tambahkan Data Terlebih Dahulu.")
		fmt.Println("\n=======================================================================")
	} else {
		fmt.Print("Masukkan tanggal mulai (DD-MM-YYYY): ")
		fmt.Scanln(&tanggalMulai)
		fmt.Print("Masukkan tanggal akhir (DD-MM-YYYY): ")
		fmt.Scanln(&tanggalAkhir)

		for i = 0; i < n; i++ {
			if data[i].tanggalPosting >= tanggalMulai && data[i].tanggalPosting <= tanggalAkhir {
				engagementSaatIni = data[i].jumlahLike + data[i].jumlahKomentar + data[i].jumlahShare
				if engagementSaatIni > engagementTertinggi {
					engagementTertinggi = engagementSaatIni
					maxIndex = i
				}
			}
		}

		if maxIndex != -1 {
			fmt.Println("\nKonten dengan engagement tertinggi:")
			fmt.Printf("Judul: %s | Platform: %s | Tanggal: %s | Jam: %s\n",
				data[maxIndex].judul, data[maxIndex].platform, data[maxIndex].tanggalPosting, data[maxIndex].jamPosting)
			fmt.Printf("Like: %d | Komentar: %d | Share: %d | Total: %d\n",
				data[maxIndex].jumlahLike, data[maxIndex].jumlahKomentar, data[maxIndex].jumlahShare, engagementTertinggi)
		} else {
			fmt.Println("Tidak ada konten dalam periode tersebut.")
		}
	}
}

func urutkanTanggalSelection(dataKonten *[maxKonten]app, n int) {
	var i, j, idxMin int
	var temp app

	for i = 0; i < n-1; i++ {
		idxMin = i
		for j = i + 1; j < n; j++ {
			if dataKonten[j].tanggalPosting < dataKonten[idxMin].tanggalPosting {
				idxMin = j
			}
		}
		if idxMin != i {
			temp = dataKonten[i]
			dataKonten[i] = dataKonten[idxMin]
			dataKonten[idxMin] = temp
		}
	}
	fmt.Println("Data berhasil diurutkan berdasarkan tanggal.")
}

func hitungEngagement(k app) int {
	return k.jumlahLike + k.jumlahKomentar + k.jumlahShare
}

func urutkanEngagementInsertion(dataKonten *[maxKonten]app, n int) {
	var i, j int
	var temp app

	for i = 1; i < n; i++ {
		temp = dataKonten[i]
		j = i - 1
		for j >= 0 && hitungEngagement(dataKonten[j]) < hitungEngagement(temp) {
			dataKonten[j+1] = dataKonten[j]
			j = j - 1
		}
		dataKonten[j+1] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan engagement.")
}
