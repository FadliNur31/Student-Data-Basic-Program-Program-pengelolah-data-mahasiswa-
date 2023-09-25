package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	awal "tugasBesar/Data_Awal"
)

const Nmax_akun int = 1000

type AkunMahasiswa struct {
	Nama     string
	Username string
	Password string
}

type ArrAkunMahasiswa [awal.NMAX]AkunMahasiswa

func UbahInput(a string) string {
	b := strings.ReplaceAll(a, "_", " ")
	return b
}

// aryo//
func templateAwal() {
	fmt.Print(
		"=============================================================================\n",
		"---------------------$ Program Pengelola Data Mahasiswa $--------------------\n",
		"=============================================================================\n",
		"------------------------------------OPSI-------------------------------------\n",
		"1. Registrasi\n",
		"2. Login\n",
		"3. Admin\n",
		"4. Keluar\n",
		"=============================================================================\n")
}
func templateSecondl() {
	fmt.Print(
		"==========================================\n",
		"***************** PROGRAM ****************\n",
		"==========================================\n",
	)
}
func Menu_mashasiswa() {
	fmt.Print(
		"------------------OPSI-------------------\n",
		"1. Informasi Mata Kuliah \n",
		"2. Informasi Mahasiswa\n",
		"3. Kembali\n",
		"=========================================\n",
	)

}

func Menu_mashasiswa_Informasi_Mahasiswa() {
	fmt.Print(
		"------------------OPSI--------------------\n",
		"1. Cari Berdasarkan Nama \n",
		"2. Cari Berdasarkan Nim \n",
		"3. Kembali \n",
		"==========================================\n",
	)

}

func Menu_Admin() {
	fmt.Print(
		"----------------Halo Admin----------------\n",
		"1. Tambah Data Mahasiswa \n",
		"2. Edit Data Mahasiswa \n",
		"3. Pengurutan Data Mahasiswa \n",
		"4. Transkrip Nilai Mahasiswa \n",
		"5. Kembali \n",
		"==========================================\n",
	)

}
func MenuEditDataMahasiswa() {
	fmt.Print(
		"----------------Menu Edit Data------------\n",
		"1. Tambah Data Tertentu\n",
		"2. Ubah Data Tertentu\n",
		"3. Hapus Data Tertentu\n",
		"4. Kembali\n",
		"==========================================\n",
	)
}

func MenuEditDataMahasiswa_Tambah() {
	fmt.Print(
		"----------------Menu Tambah Data----------\n",
		"1. Tambah Matkul\n",
		"2. Tambah Nilai Kuis Matkul Tertentu\n",
		"3. Kembali\n",
		"==========================================\n",
	)
}

func MenuEditDataMahasiswa_Ubah() {
	fmt.Print(
		"----------------Menu Ubah Data------------\n",
		"1. Ubah Data Mahasiswa\n",
		"2. Ubah Mata Kuliah Mahasiswa\n",
		"3. Kembali\n",
		"==========================================\n",
	)
}

func MenuEditDataMahasiswa_Hapus() {
	fmt.Print(
		"----------------Menu Ubah Data----------------------\n",
		"1. Hapus Seluruh Data Mahasiswa Yang Dipilih\n",
		"2. Hapus Mata Kuliah Mahasiswa Yang Dpilih\n",
		"3. Hapus Nilai Kuis Mahasiswa Yang Dpilih\n",
		"4. Hapus Nilai UTS Dan UAS Mahasiswa Yang Dpilih\n",
		"5. Kembali\n",
		"====================================================\n",
	)
}

func MenuEditDataMahasiswa_Ubah_Data_Mahasiswa() {
	fmt.Print(
		"----------------Menu Ubah Data-------------\n",
		"1. Ubah Nama\n",
		"2. Ubah Kelas\n",
		"3. Ubah Nim\n",
		"4. Kembali\n",
		"==========================================\n",
	)
}
func MenuEditDataMahasiswa_Ubah_Mata_Kuliah() {
	fmt.Print(
		"----------------Menu Ubah Data-------------\n",
		"1. Ubah Nama Mata Kuliah\n",
		"2. Ubah SKS Matkul\n",
		"3. Ubah Nilai Kuis\n",
		"4. Ubah Nilai UTS dan UAS\n",
		"5. Kembali\n",
		"==========================================\n",
	)
}

func MenuPengurutanDataMahasiswa() {
	fmt.Print(
		"----------------Menu Pengurutan Data------------------\n",
		"1. Berdasarkan SKS\n",
		"2. Berdasar Total Nilai\n",
		"3. Kembali\n",
		"======================================================\n",
	)
}

func MenuPengurutanDataMahasiswaLanjut() {
	fmt.Print(
		"----------------Menu Pengurutan Data------------------\n",
		"1. Ascending\n",
		"2. Descending\n",
		"3. Kembali\n",
		"======================================================\n",
	)
}

func Clear_Screen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func return_i_nama(ArrDataMhs awal.Arr, n_dataMhs int, b string) int {
	for i := 0; i < n_dataMhs; i++ {
		if b == ArrDataMhs[i].Nama {
			return i
		}
	}
	return -1
}

func return_i_matkul(ArrDataMhs awal.Arr, n_dataMhs int, n int, b string) int {
	for j := 0; j < ArrDataMhs[n].N_Matkul; j++ {
		if b == ArrDataMhs[n].Data_Mahasiswa[j].Nama_Matkul {
			return j
		}
	}
	return -1

}

func PilihanMenuAdmin(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	Clear_Screen()
	var a int
	for i := 1; i <= awal.NMAX; i++ {
		Menu_Admin()
		fmt.Print("Masukkan Pilihan Menu : ")
		fmt.Scan(&a)
		switch a {
		case 1:
			Admin_TambahDataMahasiswa(ArrDataMhs, n_dataMhs)
		case 2:
			Admin_EditDataMahasiswa(ArrDataMhs, n_dataMhs)
		case 3:
			Admin_PengurutanDataMahasiswa(ArrDataMhs, n_dataMhs)
		case 4:
			Admin_TranskripNilai(ArrDataMhs, n_dataMhs)
		case 5:
			i += 2000
		default:
			fmt.Println("Invalid Input")

		}
	}
}

func ConvertgradetoNilai(ArrDataMhs *awal.Arr, n_dataMhs *int, a string) float64 {
	var b float64
	switch a {
	case "A":
		b = 4.0
	case "AB":
		b = 3.5
	case "B":
		b = 3
	case "BC":
		b = 2.5
	case "C":
		b = 2
	case "D":
		b = 1
	case "E":
		b = 0
	}
	return b
}

func Admin_TranskripNilai(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	Reset_Nilai(ArrDataMhs, n_dataMhs)
	getTotalNilai(ArrDataMhs, n_dataMhs)
	var a string
	var i int
	fmt.Print("Masukkan Nama Mahasiswa Yang Ingin Ditampilkan Transkrip Nilainya : ")
	fmt.Scan(&a)
	a = UbahInput(a)
	i = return_i_nama(*ArrDataMhs, *n_dataMhs, a)
	switch i {
	case -1:
		fmt.Print("Maaf Mahasiswa Tidak Ditemukan")
	default:
		fmt.Print(
			"=========================================================================\n",
			"Nama Mahasiswa                        : ", ArrDataMhs[i].Nama, "\n",
			"Kelas                                 : ", ArrDataMhs[i].Kelas, "\n",
			"NIM                                   : ", ArrDataMhs[i].Nim, "\n",
			"SKS                                   : ", ArrDataMhs[i].SKS_total, "\n",
			"Jumlah Matkul Yang Diambil            : ", ArrDataMhs[i].N_Matkul, "\n",
			"=========================================================================\n",
			"Mata Kuliah Yang Diambil : \n",
			"----------------------------------------\n",
			"Nama Mata Kuliah           |         Grade\n",
			"...........................|............\n")
		for x := 0; x < ArrDataMhs[i].N_Matkul; x++ {
			fmt.Printf(
				"%21s      |         %v\n",
				ArrDataMhs[i].Data_Mahasiswa[x].Nama_Matkul,
				ArrDataMhs[i].Data_Mahasiswa[x].Grade_Matkul,
			)
		}
		fmt.Print(
			"........................................\n\n")
		fmt.Printf(
			"%21s      |         %v\n",
			"IPK",
			ArrDataMhs[i].IPK,
		)
	}

}

func Admin_PengurutanDataMahasiswa(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	Clear_Screen()
	for i := 0; i < awal.NMAX; i++ {
		var a int
		MenuPengurutanDataMahasiswa()
		fmt.Print("Masukkan Pilihan Menu : ")
		fmt.Scan(&a)
		switch a {
		case 1:
			PengurutanDataMahasiswa_SKS(ArrDataMhs, n_dataMhs)
		case 2:
			PengurutanDataMahasiswa_TotalNilai(ArrDataMhs, n_dataMhs)
		case 3:
			i += 10000
		default:
			fmt.Print("Invalid Input \n")
		}
	}

}

func PengurutanDataMahasiswa_TotalNilai(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	for i := 0; i < awal.NMAX; i++ {
		var a int
		MenuPengurutanDataMahasiswaLanjut()
		fmt.Print("Masukkan Pilihan Menu : ")
		fmt.Scan(&a)
		switch a {
		case 1:
			AscendingIsort(ArrDataMhs, n_dataMhs)
			i += 10000
		case 2:
			DescendingISsort(ArrDataMhs, n_dataMhs)
			i += 10000
		case 3:
			i += 10000
		default:
			fmt.Print("Invalid Input \n")
		}
	}
}
func ConvertGrade(ArrDataMhs *awal.Arr, n_dataMhs *int, a float64) string {
	var b string
	if a > 80 {
		b = "A"
	} else if a <= 80 && a > 70 {
		b = "AB"
	} else if a <= 70 && a > 65 {
		b = "B"
	} else if a <= 65 && a > 60 {
		b = "BC"
	} else if a <= 60 && a > 50 {
		b = "C"
	} else if a <= 50 && a > 40 {
		b = "D"
	} else if a <= 40 {
		b = "E"
	}
	return b
}
func Nilai_Kuis(ArrDataMhs *awal.Arr, n_dataMhs *int, i, x int) {
	for j := 0; j < ArrDataMhs[i].Data_Mahasiswa[x].N_Quiz; j++ {
		ArrDataMhs[i].Data_Mahasiswa[x].Total_Nilai_Kuis += ArrDataMhs[i].Data_Mahasiswa[x].Quiz[j]
	}
}
func getTotalNilai(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	var Total_Nilai_Kuis, Total_Nilai_UasDanUTS, b float64
	for i := 0; i < *n_dataMhs; i++ {
		for x := 0; x < ArrDataMhs[i].N_Matkul; x++ {
			Nilai_Kuis(ArrDataMhs, n_dataMhs, i, x)
			Total_Nilai_Kuis = (ArrDataMhs[i].Data_Mahasiswa[x].Total_Nilai_Kuis / float64(ArrDataMhs[i].Data_Mahasiswa[x].N_Quiz)) * 0.6
			Total_Nilai_UasDanUTS = (float64(ArrDataMhs[i].Data_Mahasiswa[x].UAS) + float64(ArrDataMhs[i].Data_Mahasiswa[x].UTS)) * 0.2
			b = (Total_Nilai_Kuis + Total_Nilai_UasDanUTS)
			ArrDataMhs[i].Data_Mahasiswa[x].Total_Nilai_Matkul = b * float64(ArrDataMhs[i].Data_Mahasiswa[x].SKS)
			ArrDataMhs[i].Data_Mahasiswa[x].Grade_Matkul = ConvertGrade(ArrDataMhs, n_dataMhs, b)
			ArrDataMhs[i].Data_Mahasiswa[x].IP_Matkul = ConvertgradetoNilai(ArrDataMhs, n_dataMhs, ArrDataMhs[i].Data_Mahasiswa[x].Grade_Matkul)
			ArrDataMhs[i].IP_Total_Matkul += ArrDataMhs[i].Data_Mahasiswa[x].IP_Matkul
			ArrDataMhs[i].Total_Nilai += ArrDataMhs[i].Data_Mahasiswa[x].Total_Nilai_Matkul
		}

		ArrDataMhs[i].IPK = ArrDataMhs[i].IP_Total_Matkul / float64(ArrDataMhs[i].N_Matkul)
		fmt.Print(4)
		ArrDataMhs[i].Total_Nilai = ArrDataMhs[i].Total_Nilai / float64(ArrDataMhs[i].SKS_total)
		ArrDataMhs[i].Grade = ConvertGrade(ArrDataMhs, n_dataMhs, ArrDataMhs[i].Total_Nilai)
	}

}
func Reset_Nilai(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	for i := 0; i < *n_dataMhs; i++ {
		for x := 0; x < ArrDataMhs[i].N_Matkul; x++ {
			ArrDataMhs[i].Data_Mahasiswa[x].Total_Nilai_Kuis = 0
			ArrDataMhs[i].Data_Mahasiswa[x].Total_Nilai_Matkul = 0
			ArrDataMhs[i].Total_Nilai = 0
			ArrDataMhs[i].IP_Total_Matkul = 0
			ArrDataMhs[i].Data_Mahasiswa[x].IP_Matkul = 0
		}
	}

}

func AscendingIsort(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	Clear_Screen()
	Reset_Nilai(ArrDataMhs, n_dataMhs)
	getTotalNilai(ArrDataMhs, n_dataMhs)
	var pass, i int
	pass = 1
	for pass < *n_dataMhs {
		i = pass
		temp := ArrDataMhs[pass].Total_Nilai
		for i > 0 && temp < ArrDataMhs[i-1].Total_Nilai {
			ArrDataMhs[i] = ArrDataMhs[i-1]
			i--
		}
		ArrDataMhs[i].Total_Nilai = temp
		pass++
	}
	for i := 0; i < *n_dataMhs; i++ {
		fmt.Printf("%v. %v Dengan Total Nilai : %v \n", i+1, ArrDataMhs[i].Nama, ArrDataMhs[i].Total_Nilai)
	}
	ArrDataMhs[0].Data_Mahasiswa[0].Total_Nilai_Kuis = 0
}

func DescendingISsort(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	Clear_Screen()
	Reset_Nilai(ArrDataMhs, n_dataMhs)
	getTotalNilai(ArrDataMhs, n_dataMhs)
	var pass, i int
	pass = 1
	for pass < *n_dataMhs {
		i = pass
		temp := ArrDataMhs[pass].Total_Nilai
		for i > 0 && temp > ArrDataMhs[i-1].Total_Nilai {
			ArrDataMhs[i] = ArrDataMhs[i-1]
			i--
		}
		ArrDataMhs[i].Total_Nilai = temp
		pass++
	}
	for i := 0; i < *n_dataMhs; i++ {
		fmt.Printf("%v. %v Dengan Total Nilai : %v \n", i+1, ArrDataMhs[i].Nama, ArrDataMhs[i].Total_Nilai)
	}

}

func PengurutanDataMahasiswa_SKS(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	for i := 0; i < awal.NMAX; i++ {
		var a int
		MenuPengurutanDataMahasiswaLanjut()
		fmt.Print("Masukkan Pilihan Menu : ")
		fmt.Scan(&a)
		switch a {
		case 1:
			AscendingSsort(ArrDataMhs, n_dataMhs)
		case 2:
			DescendingSsort(ArrDataMhs, n_dataMhs)
		case 3:
			i += 10000
		default:
			fmt.Print("Invalid Input \n")
		}
	}
}

func AscendingSsort(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	Clear_Screen()
	var pass, idx, i int
	pass = 1
	for pass < *n_dataMhs {
		idx = pass - 1
		i = pass
		for i < *n_dataMhs {
			if ArrDataMhs[idx].SKS_total > ArrDataMhs[i].SKS_total {
				idx = i
			}
			i++
		}
		temp := ArrDataMhs[pass-1]
		ArrDataMhs[pass-1] = ArrDataMhs[idx]
		ArrDataMhs[idx] = temp
		pass++
	}
	for i := 0; i < *n_dataMhs; i++ {
		fmt.Printf("%v. %v Dengan SKS : %v \n", i+1, ArrDataMhs[i].Nama, ArrDataMhs[i].SKS_total)
	}
}

func DescendingSsort(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	Clear_Screen()
	var pass, idx, i int
	pass = 1
	for pass < *n_dataMhs {
		idx = pass - 1
		i = pass
		for i < *n_dataMhs {
			if ArrDataMhs[idx].SKS_total < ArrDataMhs[i].SKS_total {
				idx = i
			}
			i++
		}
		temp := ArrDataMhs[pass-1]
		ArrDataMhs[pass-1] = ArrDataMhs[idx]
		ArrDataMhs[idx] = temp
		pass++
	}
	for i := 0; i < *n_dataMhs; i++ {
		fmt.Printf("%v. %v Dengan SKS : %v \n", i+1, ArrDataMhs[i].Nama, ArrDataMhs[i].SKS_total)
	}
}

func Admin_TambahDataMahasiswa(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	Clear_Screen()
	var a, b string
	templateSecondl()
	for i := 2; i < awal.NMAX; i++ {
		if i > 2 {
			fmt.Print("Lanjut Untuk Mengisi data (Yes/No))\n")
			fmt.Scan(&a, "\n")
		}
		switch a {
		case "No":
			i += 99999999
		default:
			fmt.Print("Ganti spasi ganti dengan ( _ )\n")
			*n_dataMhs++
			fmt.Print("Masukkan Nama Mahasiswa              : ")
			fmt.Scan(&b)
			ArrDataMhs[i].Nama = UbahInput(b)
			fmt.Print("Masukkan Nim Mahasiswa               : ")
			fmt.Scan(&ArrDataMhs[i].Nim, "\n")
			fmt.Print("Masukkan Kelas Mahasiswa             : ")
			fmt.Scan(&ArrDataMhs[i].Kelas, "\n")
			InputMatkulQuiz(ArrDataMhs, i)
		}

	}
}
func Admin_EditDataMahasiswa(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	Clear_Screen()
	var a int
	MenuEditDataMahasiswa()
	fmt.Print("Masukkan Pilihan Menu : ")
	fmt.Scan(&a)
	switch a {
	case 1:
		TambahTertentu(ArrDataMhs, n_dataMhs)
	case 2:
		UbahTertentu(ArrDataMhs, n_dataMhs)
	case 3:
		HapusTertentu(ArrDataMhs, n_dataMhs)

	}
}

func HapusTertentu(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	Clear_Screen()
	for i := 0; i < awal.NMAX; i++ {
		var a, b int
		var nama string
		MenuEditDataMahasiswa_Hapus()
		fmt.Print("Masukkan Pilihan Menu : ")
		fmt.Scan(&a)
		if a == 5 {
			Clear_Screen()
			i += 20000
		} else {
			fmt.Print("Ganti Spasi dengan (_)\n")
			fmt.Print("Masukkan Nama Mahasiswa Yang Ingin Dihapus Datanya : ")
			fmt.Scan(&nama)
			nama = UbahInput(nama)
			b = return_i_nama(*ArrDataMhs, *n_dataMhs, nama)
			if b != -1 {
				switch a {
				case 1:
					Hapus_Seluruh_Data_Mahasiswa_Yang_Dipilih(ArrDataMhs, n_dataMhs, b)
				case 2:
					Hapus_Mata_Kuliah_Mahasiswa_Yang_Dpilih(ArrDataMhs, n_dataMhs, b)
				case 3:
					Hapus_Nilai_Kuis_Mahasiswa_Yang_Dpilih(ArrDataMhs, n_dataMhs, b)
				case 4:
					Hapus_Nilai_UTS_Dan_UAS_Mahasiswa_Yang_Dpilih(ArrDataMhs, n_dataMhs, b)
				}
			} else {
				Clear_Screen()
				fmt.Print("Nama Siswa Tidak Ditemukan\n")
			}
		}
	}
}

func Hapus_Seluruh_Data_Mahasiswa_Yang_Dipilih(ArrDataMhs *awal.Arr, n_dataMhs *int, b int) {
	var i, found int
	var c string
	fmt.Print("Anda Yakin Menghapus Semua Data Mahasiswa Ini (Y/N) : ")
	fmt.Scan(&c)
	if c == "Y" {
		found = b
		i = found
		for i <= *n_dataMhs {
			ArrDataMhs[i] = ArrDataMhs[i+1]
			i++
		}
		*n_dataMhs--
		Clear_Screen()
		fmt.Print("Data Mahasiswa Berhasil Dihapus\n")
	} else {
		Clear_Screen()
	}
}

func Hapus_Mata_Kuliah_Mahasiswa_Yang_Dpilih(ArrDataMhs *awal.Arr, n_dataMhs *int, b int) {
	var i, found, x int
	var c, matkul string
	fmt.Print("Masukkan Nama Matkul Yang Ingin Dihapus : ")
	fmt.Scan(&matkul)
	matkul = UbahInput(matkul)
	x = return_i_matkul(*ArrDataMhs, *n_dataMhs, b, matkul)
	if x != -1 {
		fmt.Print("Anda Yakin Menghapus Semua Data Matkul Ini (Y/N) : ")
		fmt.Scan(&c)
		if c == "Y" {
			found = x
			i = found
			ArrDataMhs[b].SKS_total -= ArrDataMhs[b].Data_Mahasiswa[i].SKS
			for i <= ArrDataMhs[b].N_Matkul {
				ArrDataMhs[b].Data_Mahasiswa[i] = ArrDataMhs[b].Data_Mahasiswa[i+1]
				i++
			}
			ArrDataMhs[b].N_Matkul--
			Clear_Screen()
			fmt.Print("Data Matkul Berhasil Dihapus\n")
		} else {
			Clear_Screen()
		}
	} else {
		Clear_Screen()
		fmt.Print("Matkul Tidak Ditemukan")
	}
}

func Hapus_Nilai_Kuis_Mahasiswa_Yang_Dpilih(ArrDataMhs *awal.Arr, n_dataMhs *int, b int) {
	var i, found, x, kuis int
	var c, matkul string
	fmt.Print("Masukkan Nama Matkul : ")
	fmt.Scan(&matkul)
	matkul = UbahInput(matkul)
	x = return_i_matkul(*ArrDataMhs, *n_dataMhs, b, matkul)
	if x != -1 {
		fmt.Print("Kuis Minggu Keberapa Yang Ingin Anda Hapus : ")
		fmt.Scan(&kuis)
		for kuis > ArrDataMhs[b].Data_Mahasiswa[x].N_Quiz {
			fmt.Print("Kuis Minggu Tersebut Kosong : ")
			fmt.Scan(&kuis)
		}
		fmt.Print("Anda Yakin Menghapus Semua Data Matkul Ini (Y/N) : ")
		fmt.Scan(&c)
		if c == "Y" {
			found = kuis - 1
			i = found
			for i <= ArrDataMhs[b].Data_Mahasiswa[x].N_Quiz {
				ArrDataMhs[b].Data_Mahasiswa[x].Quiz[i] = ArrDataMhs[b].Data_Mahasiswa[x].Quiz[i+1]
				i++
			}
			ArrDataMhs[b].Data_Mahasiswa[x].N_Quiz--
			Clear_Screen()
			fmt.Print("Nilai Kuis Berhasil Dihapus\n")
		} else {
			Clear_Screen()
		}
	} else {
		Clear_Screen()
		fmt.Print("Matkul Tidak Ditemukan")
	}
}

func Hapus_Nilai_UTS_Dan_UAS_Mahasiswa_Yang_Dpilih(ArrDataMhs *awal.Arr, n_dataMhs *int, b int) {
	var x int
	var c, matkul string
	fmt.Print("Masukkan Nama Matkul Yang Nilai UTS Dan UAS Nya Ingin Dihapus : ")
	fmt.Scan(&matkul)
	matkul = UbahInput(matkul)
	x = return_i_matkul(*ArrDataMhs, *n_dataMhs, b, matkul)
	if x != -1 {
		fmt.Print("Anda Yakin Menghapus Nilai UTS Dan UAS Matkul Ini (Y/N) : ")
		fmt.Scan(&c)
		if c == "Y" {
			ArrDataMhs[b].Data_Mahasiswa[x].UAS = 0
			ArrDataMhs[b].Data_Mahasiswa[x].UTS = 0
			Clear_Screen()
			fmt.Print("Nilai UTS Dan UAS Berhasil Dihapus\n")
		} else {
			Clear_Screen()
		}
	} else {
		Clear_Screen()
		fmt.Print("Matkul Tidak Ditemukan")
	}
}

func UbahTertentu(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	var nama, matkul string
	fmt.Print("Masukkan Nama Mahasiswa Yang Ingin Diubah Datanya : \n")
	fmt.Scan(&nama)
	nama = UbahInput(nama)
	j := return_i_nama(*ArrDataMhs, *n_dataMhs, nama)
	if j == -1 {
		Clear_Screen()
		fmt.Print("Maaf Nama tidak ditemukan\n")
	} else {
		Clear_Screen()
		for i := 0; i < awal.NMAX; i++ {
			var a int
			MenuEditDataMahasiswa_Ubah()
			fmt.Print("Masukkan Pilihan Menu : ")
			fmt.Scan(&a)
			fmt.Print("Ganti Spasi dengan (_)\n")
			switch a {
			case 1:
				Clear_Screen()
				UbahDataMahasiswa(ArrDataMhs, n_dataMhs, j)
			case 2:
				Clear_Screen()
				fmt.Print("Masukkan Nama Matkul yang ingin diubah datanya : \n")
				fmt.Scan(&matkul)
				matkul = UbahInput(matkul)
				x := return_i_matkul(*ArrDataMhs, *n_dataMhs, j, matkul)
				if x != -1 {
					Clear_Screen()
					UbahDataMatkul(ArrDataMhs, n_dataMhs, j, x)
				} else {
					Clear_Screen()
					fmt.Print("Maaf Matkul Tidak Ditemukan \n")
				}
			case 3:
				i = i + 20000
				Clear_Screen()
			}
		}
	}
}
func UbahDataMatkul(ArrDataMhs *awal.Arr, n_dataMhs *int, j, x int) {
	for i := 0; i < awal.NMAX; i++ {
		var a, j, idxKuis int
		MenuEditDataMahasiswa_Ubah_Mata_Kuliah()
		fmt.Print("Masukkan Pilihan Menu : ")
		fmt.Scan(&a)
		switch a {
		case 1:
			fmt.Print("Masukkan Nama Matkul Yang Baru : ")
			fmt.Scan(&ArrDataMhs[j].Data_Mahasiswa[x].Nama_Matkul)
		case 2:
			fmt.Print("Masukkan SKS Matkul Yang Baru : ")
			fmt.Scan(&ArrDataMhs[j].Data_Mahasiswa[x].SKS)
		case 3:
			fmt.Print("Kuis Minggu Berapa Yang Ingin Diubah : ")
			fmt.Scan(&idxKuis)
			fmt.Print("Masukkan Nilai Kuis Yang Baru: ")
			fmt.Scan(&ArrDataMhs[j].Data_Mahasiswa[x].Quiz[idxKuis-1])
			fmt.Print(ArrDataMhs[j].Data_Mahasiswa[x].Quiz)
		case 4:
			fmt.Print("Masukkan Nilai UTS Yang Baru : ")
			fmt.Scan(&ArrDataMhs[j].Data_Mahasiswa[x].UTS)
			fmt.Print("Masukkan Nilai UAS Yang Baru : ")
			fmt.Scan(&ArrDataMhs[j].Data_Mahasiswa[x].UAS)
		case 5:
			i += 20000
			Clear_Screen()
		}
	}
}

// aryo//
func UbahDataMahasiswa(ArrDataMhs *awal.Arr, n_dataMhs *int, j int) {
	for i := 0; i < awal.NMAX; i++ {
		var a int
		MenuEditDataMahasiswa_Ubah_Data_Mahasiswa()
		fmt.Print("Masukkan Pilihan Menu : ")
		fmt.Scan(&a)
		switch a {
		case 1:
			fmt.Print("Masukkan Nama Mahasiswa Yang Baru : ")
			fmt.Scan(&ArrDataMhs[j].Nama)
		case 2:
			fmt.Print("Masukkan Kelas Mahasiswa Yang Baru : ")
			fmt.Scan(&ArrDataMhs[j].Kelas)
		case 3:
			fmt.Print("Masukkan Nim Mahasiswa Yang Baru : ")
			fmt.Scan(&ArrDataMhs[j].Nim)
		case 4:
			i += 20000
			Clear_Screen()
		}
	}
}
func Tambah_Nilai_Kuis_Matkul_Tertentu(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	var b string
	for j := 0; j < awal.NMAX; j++ {
		fmt.Print("Ketik (-) Untuk Berhenti\n")
		fmt.Print("Masukkan Nama Mahasiswa : ")
		fmt.Scan(&b)
		b = UbahInput(b)
		switch return_i_nama(*ArrDataMhs, *n_dataMhs, b) {
		case -1:
			Clear_Screen()
			fmt.Print("Maaf Mahasiswa Tidak Ditemukan \n")
			j += 1000
		default:
			tambahKuis(ArrDataMhs, n_dataMhs, return_i_nama(*ArrDataMhs, *n_dataMhs, b))
		}
	}
}

func Tambah_Matkul(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	var b string
	for j := 0; j < awal.NMAX; j++ {
		fmt.Print("Ketik (-) Untuk Berhenti\n")
		fmt.Print("Masukkan Nama Mahasiswa Yang Ingin Ditambah Matkulnya : ")
		fmt.Scan(&b)
		b = UbahInput(b)
		switch return_i_nama(*ArrDataMhs, *n_dataMhs, b) {
		case -1:
			Clear_Screen()
			fmt.Print("Maaf Mahasiswa Tidak Ditemukan \n")
			j += 1000
		default:
			InputMatkulQuiz(ArrDataMhs, return_i_nama(*ArrDataMhs, *n_dataMhs, b))
		}
	}
}

func TambahTertentu(ArrDataMhs *awal.Arr, n_dataMhs *int) {
	Clear_Screen()
	for i := 0; i < awal.NMAX; i++ {
		var a int
		MenuEditDataMahasiswa_Tambah()
		fmt.Print("Masukkan Pilihan Menu : ")
		fmt.Scan(&a)
		fmt.Print("Ganti Spasi dengan (_)\n")
		switch a {
		case 1:
			Tambah_Matkul(ArrDataMhs, n_dataMhs)
		case 2:
			Tambah_Nilai_Kuis_Matkul_Tertentu(ArrDataMhs, n_dataMhs)
		case 3:
			i = i + 20000
			Clear_Screen()
		}
	}
}

func Registrasi(Ac_mhs *ArrAkunMahasiswa, n *int) {
	templateSecondl()
	fmt.Print("*** REGISTRASI ***\n")
	fmt.Print("Username : ")
	fmt.Scan(&Ac_mhs[*n].Username)
	fmt.Print("Password : ")
	fmt.Scan(&Ac_mhs[*n].Password)
	*n++
	fmt.Println("Registrasi Berhasil :) ")
}

func Login(Ac_mhs ArrAkunMahasiswa, n int) bool {
	var a, b string
	var c bool = false
	templateSecondl()
	for x := 1; x <= 3; x++ {
		fmt.Print("Masukkan Username : ")
		fmt.Scan(&a)
		fmt.Print("Masukkan Password : ")
		fmt.Scan(&b)
		for i := 0; i <= n; i++ {
			if a == Ac_mhs[i].Username && b == Ac_mhs[i].Password {
				fmt.Println("Login Berhasil :) ")
				c = true
				x = x + 4
				i = i + n
			}

		}
		if !c {
			fmt.Println("Username atau Password salah")
		}
	}
	return c
}

func Admin() bool {
	templateSecondl()
	var a, b string
	var c bool = false
	for x := 1; x <= 3; x++ {
		fmt.Print("Masukkan Username : ")
		fmt.Scan(&a)
		fmt.Print("Masukkan Password : ")
		fmt.Scan(&b)
		if a == "admin" && b == "admin" {
			fmt.Println("Login Berhasil :) ")
			c = true
			x = x + 4
		} else {
			fmt.Println("Username atau Passsword salah")
		}

	}
	return c
}

func PilihanMenuAwal(a int, Akun_mhs *ArrAkunMahasiswa, n_akun *int, ArrDataMhs *awal.Arr, n_dataMhs *int) {
	switch a {
	case 1:
		Registrasi(Akun_mhs, n_akun)
	case 2:
		if Login(*Akun_mhs, *n_akun) {
			templateSecondl()
			PilihanMenuMahasiswa(*ArrDataMhs, *n_dataMhs)
		}
	case 3:
		if Admin() {
			PilihanMenuAdmin(ArrDataMhs, n_dataMhs)
		}
	case 4:
		fmt.Print("TERIMA KASIH :) ")
		os.Exit(0)

	default:
		fmt.Println("Invalid Input")

	}
}
func PilihanMenuMahasiswa(ArrDataMhs awal.Arr, n_dataMhs int) {
	Clear_Screen()
	var a, b int
	var str1, Nama_Matkul, Nama_Mahasiswa, Nim_Mahasiswa string
	for i := 1; i <= awal.NMAX; i++ {
		Menu_mashasiswa()
		fmt.Print("Masukkan Pilihan Menu : ")
		fmt.Scan(&a)
		switch a {
		case 1:
			fmt.Print("Ganti spasi ganti dengan ( _ )\n")
			fmt.Print("Masukkan Nama Matkul : ")
			fmt.Scan(&str1)
			Nama_Matkul = UbahInput(str1)
			Mencari_Mahasiswa_Mengambil_Suatu_Matkul(ArrDataMhs, n_dataMhs, Nama_Matkul)
		case 2:
			for x := 1; x <= awal.NMAX; x++ {
				Menu_mashasiswa_Informasi_Mahasiswa()
				fmt.Print("Masukkan Pilihan Menu : ")
				fmt.Scan(&b)
				switch b {
				case 1:
					fmt.Print("Ganti spasi ganti dengan ( _ )\n")
					fmt.Print("Masukkan Nama Mahasiswa : ")
					fmt.Scan(&str1)
					Nama_Mahasiswa = UbahInput(str1)
					Mencari_Matkul_YangDiambil_Mahasiswa_Nama(ArrDataMhs, n_dataMhs, Nama_Mahasiswa)
				case 2:
					fmt.Print("Masukkan Nim Mahasiswa : ")
					fmt.Scan(&Nim_Mahasiswa)
					Mencari_Matkul_YangDiambil_Mahasiswa_Nim(ArrDataMhs, n_dataMhs, Nim_Mahasiswa)
				case 3:
					x = x + 200000
				default:
					fmt.Println("Invalid Input")

				}
			}
		case 3:
			i = i + 200000
		default:
			fmt.Println("Invalid Input")

		}
	}
}

func Mencari_Mahasiswa_Mengambil_Suatu_Matkul(ArrDataMhs awal.Arr, n_dataMhs int, str string) {
	Clear_Screen()
	var pass bool = false
	fmt.Print("Nama Mahasiswa Yang Mengambil Matkul Tersebut : \n")
	for i := 0; i < n_dataMhs; i++ {
		for x := 0; x < ArrDataMhs[i].N_Matkul; x++ {
			if str == ArrDataMhs[i].Data_Mahasiswa[x].Nama_Matkul {
				fmt.Println(ArrDataMhs[i].Nama)
				pass = true
			}
		}

	}
	if !pass {
		fmt.Println("-----")
	}

}

func Mencari_Matkul_YangDiambil_Mahasiswa_Nama(ArrDataMhs awal.Arr, n_dataMhs int, str string) {
	Clear_Screen()
	var pass bool = false
	for i := 0; i < n_dataMhs; i++ {
		if str == ArrDataMhs[i].Nama {
			pass = true
			fmt.Print(
				"=========================================================================\n",
				"Nama Mahasiswa                        : ", ArrDataMhs[i].Nama, "\n",
				"Kelas                                 : ", ArrDataMhs[i].Kelas, "\n",
				"NIM                                   : ", ArrDataMhs[i].Nim, "\n",
				"SKS                                   : ", ArrDataMhs[i].SKS_total, "\n",
				"Jumlah Matkul Yang Diambil            : ", ArrDataMhs[i].N_Matkul, "\n",
				"=========================================================================\n",
				"Mata Kuliah Yang Diambil : \n",
				"----------------------------------------\n",
				"Nama Mata Kuliah           |         SKS\n",
				"...........................|............\n")
			for x := 0; x < ArrDataMhs[i].N_Matkul; x++ {
				fmt.Printf(
					"%21s      |         %v\n",
					ArrDataMhs[i].Data_Mahasiswa[x].Nama_Matkul,
					ArrDataMhs[i].Data_Mahasiswa[x].SKS,
				)
			}
			fmt.Print(
				"........................................\n\n")
		}

	}
	if !pass {
		fmt.Println("Nama Mahasiswa Tidak Ditemukan")
	}
}
func Mencari_Matkul_YangDiambil_Mahasiswa_Nim(ArrDataMhs awal.Arr, n_dataMhs int, str1 string) {
	Clear_Screen()
	var pass bool = false
	for i := 0; i < n_dataMhs; i++ {
		if str1 == ArrDataMhs[i].Nim {
			pass = true
			fmt.Print(
				"=========================================================================\n",
				"Nama Mahasiswa                        : ", ArrDataMhs[i].Nama, "\n",
				"Kelas                                 : ", ArrDataMhs[i].Kelas, "\n",
				"NIM                                   : ", ArrDataMhs[i].Nim, "\n",
				"SKS                                   : ", ArrDataMhs[i].SKS_total, "\n",
				"Jumlah Matkul Yang Diambil            : ", ArrDataMhs[i].N_Matkul, "\n",
				"=========================================================================\n",
				"Mata Kuliah Yang Diambil : \n",
				"----------------------------------------\n",
				"Nama Mata Kuliah           |         SKS\n",
				"...........................|............\n")
			for x := 0; x < ArrDataMhs[i].N_Matkul; x++ {
				fmt.Printf(
					"%21s      |         %v\n",
					ArrDataMhs[i].Data_Mahasiswa[x].Nama_Matkul,
					ArrDataMhs[i].Data_Mahasiswa[x].SKS,
				)
			}
			fmt.Print(
				"........................................\n\n")
		}

	}
	if !pass {
		fmt.Println("Nama Mahasiswa Tidak Ditemukan")
	}
}
func InputMatkulQuiz(ArrDataMhs *awal.Arr, i int) {
	var a, b string
	for x := ArrDataMhs[i].N_Matkul; x < 16; x++ {
		if x > 0 {
			fmt.Print("Tambah Mata Kuliah?(Yes/No) \n")
			fmt.Scan(&a, "\n")
		}
		switch a {
		case "No":
			x += 17
		default:
			ArrDataMhs[i].N_Matkul++
			fmt.Print("Ganti spasi ganti dengan ( _ )\n")
			fmt.Print("Masukkan Nama Mata Kuliah            : ")
			fmt.Scan(&b, "\n")
			ArrDataMhs[i].Data_Mahasiswa[x].Nama_Matkul = UbahInput(b)
			fmt.Printf("Masukkan SKS Mata Kuliah             : ")
			fmt.Scan(&ArrDataMhs[i].Data_Mahasiswa[x].SKS, "\n")
			ArrDataMhs[i].SKS_total += ArrDataMhs[i].Data_Mahasiswa[x].SKS
			fmt.Printf("Masukkan Nilai UTS Mata Kuliah       : ")
			fmt.Scan(&ArrDataMhs[i].Data_Mahasiswa[x].UTS, "\n")
			fmt.Printf("Masukkan Nilai UAS Mata Kuliah       : ")
			fmt.Scan(&ArrDataMhs[i].Data_Mahasiswa[x].UAS, "\n")
			fmt.Print("Isi Nilai Quiz Dengan (404) Untuk Berhenti \n")
			inputQuiz(ArrDataMhs, i, x)
		}
	}
}
func inputQuiz(ArrDataMhs *awal.Arr, i, x int) {
	for b := 0; b < 16; b++ {
		fmt.Printf("Nilai Quiz ke %v                      : ", b+1)
		ArrDataMhs[i].Data_Mahasiswa[x].N_Quiz++
		fmt.Scan(&ArrDataMhs[i].Data_Mahasiswa[x].Quiz[b], "\n")
		ArrDataMhs[i].Data_Mahasiswa[x].N_Quiz++
		switch ArrDataMhs[i].Data_Mahasiswa[x].Quiz[b] {
		case 404:
			b += 17
		}

	}
}

func tambahKuis(ArrDataMhs *awal.Arr, n_dataMhs *int, n int) {
	Clear_Screen()
	var matkul, r string
	var j int
	fmt.Print("Masukkan Nama Matkul Yang Ingin Ditambah Nilai Quiz nya : ")
	fmt.Scan(&matkul)
	r = UbahInput(matkul)
	j = return_i_matkul(*ArrDataMhs, *n_dataMhs, n, r)
	switch j {
	case 0:
		fmt.Print("Matkul Tidak Ditemukan \n")
	default:
		if ArrDataMhs[n].Data_Mahasiswa[j].N_Quiz >= 16 {
			fmt.Print("Data Kuis Sudah Penuh, Tidak Bisa Menambahkan Data\n")
		}
		for b := ArrDataMhs[n].Data_Mahasiswa[j].N_Quiz; b < 16; b++ {
			fmt.Printf("Nilai Quiz ke %v                      : ", b+1)
			ArrDataMhs[n].Data_Mahasiswa[j].N_Quiz++
			fmt.Scan(&ArrDataMhs[n].Data_Mahasiswa[j].Quiz[b], "\n")
			ArrDataMhs[n].Data_Mahasiswa[j].N_Quiz++
			switch ArrDataMhs[n].Data_Mahasiswa[j].Quiz[b] {
			case 404:
				b += 17
			}
		}
	}
}

func main() {
	var a, n_akun, n_dataMhs int
	var Akun_mhs ArrAkunMahasiswa
	n_dataMhs = 2
	T := awal.Data_Awal(&awal.Arr{})
	for i := 0; i <= awal.NMAX; i++ {
		templateAwal()
		fmt.Print("Masukkan Pilihan Menu : ")
		fmt.Scan(&a)
		PilihanMenuAwal(a, &Akun_mhs, &n_akun, &T, &n_dataMhs)
	}
}
