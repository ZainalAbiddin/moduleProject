//kasus : portal sekolah SMA sederhana

package portalSMA

import "fmt"

type Nilaiipa struct {
	fisika, biologi, kimia float64
}

type Nilaimurid struct {
	Namamurid   string
	kelas       int
	Rataannilai []Nilaiipa
}

// fungsi
func Perkenalan() {
	fmt.Println("Selamat datang di Portal Madrasah Aliyah MA'HAD ALQUR'AN WAL HADIS (MA MAQDIS) Bogor \n")
}

// fungsi return value
func Ujinama(user string) bool {
	aktif := false
	if user == "admin" {
		aktif = true
	} else {
		aktif = false
	}
	return aktif
}

// fungsi sebagai parameter
func Verifikasi(user string, Uji func(string) bool) {
	if Uji(user) {
		fmt.Printf("Hallo %v , kamu memiliki akses untuk melihat nilai \n", user)
	} else {
		fmt.Printf("Hallo %v \n", user)
	}
}

// fungsi multiple return value
func VisiMisi() (string, string) {
	visi := "Terwujudnya Madrasah Yang Berkualitas, Berprestasi Dan Berakhlaqul Karimah Berdasarkan Iman Dan Taqwa \n"
	misi := "Menyelenggarakan pendidikan yang berkualitas dalam pencapaian prestasi akademik dan non-akademik,Mewujudkan peserta didik yang berprestasi, baik di bidang akademik maupun nonakademik seraya mampu mengembangkan semua potensinya sebagai bekal untuk melanjutkan ke pendidikan yang lebih tinggi dan atau hidup mandiri \n"
	return visi, misi
}

// fungsi dengan parameter
func Ekstrakulikuler(kegiatan [4]string) {
	fmt.Println("kegiatan Ekstrakulikuler di MAQDIS adalah : ")
	for _, ekskul := range kegiatan {
		fmt.Println(ekskul)

	}
}

// interface

// struct method
func (n Nilaimurid) Nilai() string {
	fmt.Println(n.Namamurid)
	fmt.Println("kelas :", n.kelas)
	for _, nilai := range n.Rataannilai {
		fmt.Println("===================")
		fmt.Println(nilai.fisika)
		fmt.Println(nilai.biologi)
		fmt.Println(nilai.kimia)
	}
	return "----------------------"
}
