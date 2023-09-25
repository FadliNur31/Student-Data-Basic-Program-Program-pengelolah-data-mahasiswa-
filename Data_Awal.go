package awal

const NMAX = 1000

type Arr [NMAX]Data_Mahasiswa

type nilaiQuiz [16]float64

type Data_kuliah [16]Matakuliah

type Matakuliah struct {
	Nama_Matkul        string
	SKS                int
	N_Quiz             int
	Quiz               nilaiQuiz
	Total_Nilai_Matkul float64
	Grade_Matkul       string
	IP_Matkul          float64
	Total_Nilai_Kuis   float64
	UTS                int
	UAS                int
}
type Data_Mahasiswa struct {
	Nama            string
	Nim             string
	Kelas           string
	N_Matkul        int
	SKS_total       int
	Total_Nilai     float64
	Grade           string
	IP_Total_Matkul float64
	IPK             float64
	Data_Mahasiswa  Data_kuliah
}

func Data_Awal(a *Arr) Arr {
	a[0] = Data_Mahasiswa{Nama: "Nabil Arrahman", Nim: "2211604112", Kelas: "IT-46-10", N_Matkul: 8,
		Data_Mahasiswa: Data_kuliah{Matakuliah{Nama_Matkul: "Algoritma Pemrograman", SKS: 4, Quiz: nilaiQuiz{30, 30, 30, 30, 30, 30, 39, 30}, UTS: 30, UAS: 30, N_Quiz: 8},
			Matakuliah{Nama_Matkul: "Aljabar Linear", SKS: 3, Quiz: nilaiQuiz{30, 30, 30, 30, 30, 30, 39, 30}, UTS: 30, UAS: 30, N_Quiz: 8},
			Matakuliah{Nama_Matkul: "Agama Islam", SKS: 2, Quiz: nilaiQuiz{30, 30, 30, 30, 30, 30, 39, 30}, UTS: 30, UAS: 30, N_Quiz: 8},
			Matakuliah{Nama_Matkul: "Matematika Diskrit", SKS: 2, Quiz: nilaiQuiz{30, 30, 30, 30, 30, 30, 39, 30}, UTS: 30, UAS: 30, N_Quiz: 8},
			Matakuliah{Nama_Matkul: "Sistem Digital", SKS: 3, Quiz: nilaiQuiz{30, 30, 30, 30, 30, 30, 39, 30}, UTS: 30, UAS: 30, N_Quiz: 8},
			Matakuliah{Nama_Matkul: "Bahasa Inggris", SKS: 2, Quiz: nilaiQuiz{30, 30, 30, 30, 30, 30, 39, 30}, UTS: 30, UAS: 30, N_Quiz: 8},
			Matakuliah{Nama_Matkul: "Kewirausahaan", SKS: 2, Quiz: nilaiQuiz{30, 30, 30, 30, 30, 30, 39, 30}, UTS: 30, UAS: 30, N_Quiz: 8},
			Matakuliah{Nama_Matkul: "Kewarganegaraan", SKS: 2, Quiz: nilaiQuiz{30, 30, 30, 30, 30, 30, 39, 30}, UTS: 30, UAS: 30, N_Quiz: 8},
		}}
	a[1] = Data_Mahasiswa{Nama: "Bryne Brimstone", Nim: "1303224401", Kelas: "IT-46-10", N_Matkul: 8,
		Data_Mahasiswa: Data_kuliah{Matakuliah{Nama_Matkul: "Kalkulus", SKS: 3, Quiz: nilaiQuiz{90, 90, 90, 90, 90, 90, 90, 90}, UTS: 90, UAS: 90, N_Quiz: 8},
			Matakuliah{Nama_Matkul: "Algoritma Pemrograman", SKS: 4, Quiz: nilaiQuiz{90, 90, 90, 90, 90, 90, 90, 90}, UTS: 80, UAS: 70, N_Quiz: 8},
			Matakuliah{Nama_Matkul: "Agama Islam", SKS: 2, Quiz: nilaiQuiz{90, 90, 90, 90, 90, 90, 90, 90}, UTS: 70, UAS: 90, N_Quiz: 8},
			Matakuliah{Nama_Matkul: "Matematika Diskrit", SKS: 3, Quiz: nilaiQuiz{90, 90, 90, 90, 90, 90, 90, 90}, UTS: 90, UAS: 100, N_Quiz: 8},
			Matakuliah{Nama_Matkul: "Sistem Digital", SKS: 3, Quiz: nilaiQuiz{90, 90, 90, 90, 90, 90, 90, 90}, UTS: 75, UAS: 90, N_Quiz: 8},
			Matakuliah{Nama_Matkul: "Bahasa Inggris", SKS: 2, Quiz: nilaiQuiz{90, 90, 90, 90, 90, 90, 90, 90}, UTS: 80, UAS: 85, N_Quiz: 8},
			Matakuliah{Nama_Matkul: "Kewirausahaan", SKS: 2, Quiz: nilaiQuiz{90, 90, 90, 90, 90, 90, 90, 90}, UTS: 90, UAS: 95, N_Quiz: 8},
			Matakuliah{Nama_Matkul: "Kewarganegaraan", SKS: 2, Quiz: nilaiQuiz{90, 90, 90, 90, 90, 90, 90, 90}, UTS: 77, UAS: 81, N_Quiz: 8},
		}}
	a[0].SKS_total = 20
	a[1].SKS_total = 21
	return *a

}
