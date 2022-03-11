//kasus : portal sekolah SMA sederhana

package portalSMA

import "fmt"

// versi v1.0.0
// fungsi
func Perkenalan() {
	fmt.Println("Selamat datang di Portal Madrasah Aliyah MA'HAD ALQUR'AN WAL HADIS (MA MAQDIS) Bogor \n")
}

// fungsi dengan parameter
func Pengguna(nama string) {
	if nama == "admin" {
		fmt.Printf("Selamat datang %s di Portal Maqdis \n", nama)
	}
}

// fungsi sebagai parameter

// versi v1.0.2
// fungsi return value
func Ekstrakulikuler(kegiatan [4]string) []string {
	fmt.Println("kegiatan Ekstrakulikuler di MAQDIS adalah :")
	for _, ekskul := range kegiatan {
		fmt.Println(ekskul)
		return ekskul
	}
}

// fungsi multiple return value
func VisiMisi() (string, string) {
	visi := "Terwujudnya Madrasah Yang Berkualitas, Berprestasi Dan Berakhlaqul Karimah Berdasarkan Iman Dan Taqwa \n"
	misi := "Menyelenggarakan pendidikan yang berkualitas dalam pencapaian prestasi akademik dan non-akademik,Mewujudkan peserta didik yang berprestasi, baik di bidang akademik maupun nonakademik seraya mampu mengembangkan semua potensinya sebagai bekal untuk melanjutkan ke pendidikan yang lebih tinggi dan atau hidup mandiri \n"
	return visi, misi
}

// fungsi anonymous

// struct
// type ProfilSekolah struct {
// 	namaSekolah, alamatSekolah string
// }

// interface

// anonymous struct

// struct method
