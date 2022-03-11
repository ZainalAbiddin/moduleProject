//kasus : portal sekolah SMA sederhana

package portalSMA

import "fmt"

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

func Ekstrakulikuler(kegiatan [4]string) {
	fmt.Println("kegiatan Ekstrakulikuler di MAQDIS adalah : ")
	for _, ekskul := range kegiatan {
		fmt.Println(ekskul)

	}
}

// fungsi return value
func NomorTelepon() int {
	noHP := 8622314
	fmt.Println("Nomor Telepon (0251) \n", noHP)
	return noHP
}

// fungsi multiple return value
func VisiMisi() (string, string) {
	visi := "Terwujudnya Madrasah Yang Berkualitas, Berprestasi Dan Berakhlaqul Karimah Berdasarkan Iman Dan Taqwa \n"
	misi := "Menyelenggarakan pendidikan yang berkualitas dalam pencapaian prestasi akademik dan non-akademik,Mewujudkan peserta didik yang berprestasi, baik di bidang akademik maupun nonakademik seraya mampu mengembangkan semua potensinya sebagai bekal untuk melanjutkan ke pendidikan yang lebih tinggi dan atau hidup mandiri \n"
	return visi, misi
}

// fungsi anonymous

// fungsi sebagai parameter

// struct
// type ProfilSekolah struct {
// 	namaSekolah, alamatSekolah string
// }

// interface

// anonymous struct

// struct method
