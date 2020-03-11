/* TUGAS BESAR DASAR ALGORITMA DAN PEMROGRAMAN
   IF-43-05
   Kelompok 6
   Nama    : Farid Krida Mukti (1301194052)
   			 Manuel Benedict (1301194182)
   Topik   : Aplikasi Penerbangan
*/

package main 
import "fmt"
import "os"

const T=5 //jumlah maksimum penumpang dalam sekali reservasi
const N=3 //jumlah baris kursi bisnis
const M=7 //jumlah baris kursi ekonomi

type Reservasi struct {
	jumlah int
	kodpas[T] string
	nama[T] string
	usia[T] int
	dietary[T] bool
	kursi[T][2] int
}

var( kursi_Bisnis [N][4] bool
 	 kursi_Ekonomi [M][6] bool
 	 arrReservasiBis [12] Reservasi
 	 arrReservasiEko [42] Reservasi
 	 arrKodresBis[12] int
 	 arrKodresEko[42] int
 	 bookedBis int
 	 bookedEko int
	 terisiBisnis,terisiEkonomi int							
	)

func main () {
	var menu int
	for menu!=9 {
		fmt.Println("===========================================================")
		fmt.Println("RESERVASI PENERBANGAN MASKAPAI EMIRAT RUTE JAKARTA KE DUBAI")
		fmt.Println("                    Sabtu, 8 Desember 2019                 ")
		fmt.Println("                        Pukul 00.10                        ")
		fmt.Println("===========================================================")
		fmt.Println("                            MENU							")
		fmt.Println("1. Reservasi")
		fmt.Println("2. Check in")
		fmt.Println("3. Mencari penumpang dengan kode reservasi")
		fmt.Println("4. Mencari penumpang dengan nomor paspor")
		fmt.Println("5. Mencari orang tua dari bayi atau anak")
		fmt.Println("6. Menampilkan dietary penumpang")
		fmt.Println("7. Menmapilkan semua penumpang anak berdasarkan usia")
		fmt.Println("8. Menampilkan semua kursi kosong")
		fmt.Println("9. Keluar")
		fmt.Print("====================Pilih menu: ")
		fmt.Scan(&menu)
		fmt.Println()
		if menu==1 {
			fmt.Println("RESERVASI")
			fmt.Println("Kursi kelas bisnis yang terisi: ", terisiBisnis, "dari 12 kursi")
			fmt.Println("Kursi kelas ekonomi yang terisi: ", terisiEkonomi, "dari 42 kursi")
			fmt.Println("===========================================================")
			Reservation()	
		} else if menu==2 {
			fmt.Println("CHECK IN")
			checkIn()
		} else if menu==3 {
			searchPenumpangRes()
		} else if menu==4 {
			searchPenumpangPas()
		} else if menu==5 {
			searchParents()
		} else if menu==6 {
			showVegetarian()
		} else if menu==7 {
			showKidsandParents()
		} else if menu==8 {
			showKursiKosong()
		}
		fmt.Println()
		fmt.Println("===========================================================")
		fmt.Println("===========================================================")
		fmt.Println()
		}	
}

func cekKursiBis(t *bool) {
	if terisiBisnis<12 {
		*t=true
	} else {
		*t=false
	}
}

func cekKursiEko (t *bool){
	if terisiEkonomi<42 {
		*t=true
	} else {
		*t=false
	}
}

func Reservation () {
	var status,class,q,r int
	var t bool
	q=0
	r=0
	t=false
	fmt.Println("Pilih kelas kursi Anda")
	fmt.Println("1. Kelas bisnis")
	fmt.Println("2. Kelas ekonomi")
	fmt.Print("====================Kelas kursi: ")
	fmt.Scan(&class)
	if class==1 {
		cekKursiBis(&t)
		if t==true {
			fmt.Print("Masukkan jumlah penumpang: ")
			fmt.Scan(&arrReservasiBis[bookedBis].jumlah)
			for arrReservasiBis[bookedBis].jumlah < 1 || arrReservasiBis[bookedBis].jumlah>5{
				fmt.Print("Maaf, sekali reservasi hanya dapat memesan paling banyak 5 tiket.")
				fmt.Print(" Silahkan masukkan kembali jumlah penumpang: ")
				fmt.Scan(&arrReservasiBis[bookedBis].jumlah)
			}
			fmt.Println("Anak-anak dan bayi tidak dapat berangkat sendirian. Jika ada penumpang yang membawa anak-anaknya, maka pastikan data penumpang pertama adalah orang dewasa yang ikut bersamanya.")
			fmt.Println("===========================================================")
			for i:=0; i<arrReservasiBis[bookedBis].jumlah; i++ {
				fmt.Println("Penumpang ", i+1)
				fmt.Print("Masukkan nama: ")
				fmt.Scan(&arrReservasiBis[bookedBis].nama[i])
				fmt.Print("Usia: ")
				fmt.Scan(&arrReservasiBis[bookedBis].usia[i])
				for arrReservasiBis[bookedBis].usia[0]<18 {
					fmt.Println()
					fmt.Println("Pastikan data penumpang pertama minimal berusia 18 tahun. Silahkan masukkan data penumpang 1 dengan benar.")
					fmt.Println("Penumpang ", i+1)
					fmt.Print("Masukkan nama: ")
					fmt.Scan(&arrReservasiBis[bookedBis].nama[i])
					fmt.Print("Usia: ")
					fmt.Scan(&arrReservasiBis[bookedBis].usia[i])
				}
				if arrReservasiBis[bookedBis].usia[i]<2{
					fmt.Println("Karena usia anda dibawah 18 tahun, kode paspor Anda sama dengan orang dewasa sebelumnya")
				}else if arrReservasiBis[bookedBis].usia[i] < 18 {
					arrReservasiBis[bookedBis].kodpas[i] = arrReservasiBis[bookedBis].kodpas[i-1]
					fmt.Println("Karena usia anda dibawah 18 tahun, kode paspor Anda sama dengan orang dewasa sebelumnya")
					terisiBisnis=terisiBisnis+1
				} else {
					fmt.Print("Masukkan nomor paspor Anda: ")
					fmt.Scan(&arrReservasiBis[bookedBis].kodpas[i])
					terisiBisnis=terisiBisnis+1
				}
				fmt.Println("Apakah Anda vegetarian? ")
				fmt.Println("1. ya")
				fmt.Println("2. tidak")
				fmt.Print("====================Status dietary: ")
				fmt.Scan(&status)
				if status==1 {
					arrReservasiBis[bookedBis].dietary[i]=true
				} else if status==2 {
					arrReservasiBis[bookedBis].dietary[i]=false
				}
				var check bool
				check=true
				for h:=0; h<3 && check; h++ {
					if check {
						for j:=0; j<4 && check; j++ {
							if !kursi_Bisnis[h][j] {
								arrReservasiBis[bookedBis].kursi[i][0]=h
								arrReservasiBis[bookedBis].kursi[i][1]=j
								kursi_Bisnis[h][j] = !kursi_Bisnis[h][j]
								check=false
							}	
						}
					}	
				}
			fmt.Println("=======================================================")
			arrKodresBis[q]=1000+bookedBis+1
			q++
			}
			fmt.Println("Kode reservasi: ", arrKodresBis[bookedBis])
			bookedBis++
		}else {
			fmt.Println("Kursi kelas bisnis penuh")
		}
	} else if class==2 {
		cekKursiEko(&t)
		if t==true {
			fmt.Print("Masukkan jumlah penumpang: ")
			fmt.Scan(&arrReservasiEko[bookedEko].jumlah)
			for arrReservasiEko[bookedEko].jumlah < 1 || arrReservasiEko[bookedEko].jumlah>5{
				fmt.Print("Maaf, sekali reservasi hanya dapat memesan paling banyak 5 tiket.")
				fmt.Print(" Silahkan masukkan kembali jumlah penumpang: ")
				fmt.Scan(&arrReservasiEko[bookedEko].jumlah)
			}
			fmt.Println("Anak-anak dan bayi tidak dapat berangkat sendirian. Jika ada penumpang yang membawa anak-anaknya, maka pastikan data penumpang pertama adalah orang dewasa yang ikut bersamanya.")
			fmt.Println("===========================================================")
			for i:=0; i<arrReservasiEko[bookedEko].jumlah; i++ {
				fmt.Println("Penumpang ", i+1)
				fmt.Print("Masukkan nama: ")
				fmt.Scan(&arrReservasiEko[bookedEko].nama[i])
				fmt.Print("Usia: ")
				fmt.Scan(&arrReservasiEko[bookedEko].usia[i])
				for arrReservasiEko[bookedEko].usia[0]<18 {
					fmt.Println()
					fmt.Println("Pastikan data penumpang pertama minimal berusia 18 tahun. Silahkan masukkan data penumpang 1 dengan benar.")
					fmt.Println("Penumpang ", i+1)
					fmt.Print("Masukkan nama: ")
					fmt.Scan(&arrReservasiEko[bookedEko].nama[i])
					fmt.Print("Usia: ")
					fmt.Scan(&arrReservasiEko[bookedEko].usia[i])
				}
				if arrReservasiEko[bookedEko].usia[i]<2{
					fmt.Println("Karena usia anda dibawah 18 tahun, kode paspor Anda sama dengan orang dewasa sebelumnya")
				}else if arrReservasiEko[bookedEko].usia[i] < 18 {
					arrReservasiEko[bookedEko].kodpas[i] = arrReservasiEko[bookedEko].kodpas[i-1]
					fmt.Println("Karena usia anda dibawah 18 tahun, kode paspor Anda sama dengan orang dewasa sebelumnya")
					terisiEkonomi=terisiEkonomi+1
				} else {
					fmt.Print("Masukkan nomor paspor Anda: ")
					fmt.Scan(&arrReservasiEko[bookedEko].kodpas[i])
					terisiEkonomi=terisiEkonomi+1
				}
				fmt.Println("Apakah Anda vegetarian? ")
				fmt.Println("1. ya")
				fmt.Println("2. tidak")
				fmt.Print("====================Status dietary: ")
				fmt.Scan(&status)
				if status==1 {
					arrReservasiEko[bookedEko].dietary[i]=true
				} else if status==2 {
					arrReservasiEko[bookedEko].dietary[i]=false
				}
				var check bool
				check=true
				for h:=0; h<3 && check; h++ {
					if check {
						for j:=0; j<4 && check; j++ {
							if !kursi_Ekonomi[h][j] {
								arrReservasiEko[bookedEko].kursi[i][0]=h
								arrReservasiEko[bookedEko].kursi[i][1]=j
								kursi_Ekonomi[h][j] = !kursi_Ekonomi[h][j]
								check=false
							}	
						}
					}	
				}
			fmt.Println("=======================================================")
			arrKodresEko[r]=2000+bookedEko+1
			r++
			}
			fmt.Println("Kode reservasi: ", arrKodresEko[bookedEko])
			bookedEko++	
		}else {
			fmt.Println("Kursi kelas ekonomi penuh")
		}
	}
	fmt.Println("TERIMA KASIH TELAH MEMILIH MASKAPAI EMIRAT")
}

func checkIn () {
	var kodres,j,k,dietary int
	var kursi bool
	j=0
	k=0
	fmt.Print("Masukkan kode reservasi Anda: ")
	fmt.Scan(&kodres)
	fmt.Println()
	if kodres>1000&&kodres<2000 {
		for j<12 && kodres!=arrKodresBis[j]{
			j++	
		}
		for i:=0; i<arrReservasiBis[j].jumlah; i++ {
			fmt.Println("Penumpang ke", i+1)
			fmt.Println("Nama penumpang: ",arrReservasiBis[j].nama[i])
			fmt.Println("Usia penumpang: ",arrReservasiBis[j].usia[i])
			fmt.Println("Kode paspor: ", arrReservasiBis[j].kodpas[i])
			changeSeatBis(i,j,kursi)
			if arrReservasiBis[j].dietary[i]==true {
				fmt.Println("Status dietary Anda saat ini: vegetarian")
			} else {
				fmt.Println("Status dietary Anda saat ini: non-vegetarian")
			}
			fmt.Println("Apakah Anda ingin mengubah status dietary?")
			fmt.Println("1. ya")
			fmt.Println("2. tidak")
			fmt.Print("Silahkan pilih: ")
			fmt.Scan(&dietary)
			if dietary==1 {
				fmt.Print("Ketik true jika anda vegetarian atau false jika bukan: ")
				fmt.Scan(&arrReservasiBis[j].dietary[i])
				if arrReservasiBis[j].dietary[i]==true {
					fmt.Println("Vegetarian? ya")
				} else {
					fmt.Println("Vegetarian? tidak")	
				}
			} 
			fmt.Println()	
		}
	} else if kodres>2000&&kodres<3000 {
		for k<42&&kodres!=arrKodresEko[k] {
			k++	
		}
		for i:=0; i<arrReservasiEko[k].jumlah; i++ {
			fmt.Println("Penumpang ke", i+1)
			fmt.Println("Nama penumpang: ",arrReservasiEko[j].nama[i])
			fmt.Println("Usia penumpang: ",arrReservasiEko[j].usia[i])
			fmt.Println("Kode paspor: ", arrReservasiEko[j].kodpas[i])
			changeSeatEko(i,j,kursi)
			if arrReservasiEko[j].dietary[i]==true {
				fmt.Println("Status dietary Anda saat ini: vegetarian")
			} else {
				fmt.Println("Status dietary Anda saat ini: non-vegetarian")
			}
			fmt.Println("Apakah Anda ingin mengubah status dietary?")
			fmt.Println("1. ya")
			fmt.Println("2. tidak")
			fmt.Print("Silahkan pilih: ")
			fmt.Scan(&dietary)
			if dietary==1 {
				fmt.Print("Ketik true jika anda vegetarian atau false jika bukan: ")
				fmt.Scan(&arrReservasiEko[j].dietary[i])
				if arrReservasiEko[j].dietary[i]==true {
					fmt.Println("Vegetarian? ya")
				} else {
					fmt.Println("Vegetarian? tidak")	
				}
			}
			fmt.Print()
		}
	}
}

func searchPenumpangPas () {
	var kodpas string
	fmt.Print("Masukkan kode paspor Anda: ")
	fmt.Scan(&kodpas)
	for i:=0; i<bookedBis; i++ {
		for h:=0; h<arrReservasiBis[i].jumlah; h++ {
			if kodpas == arrReservasiBis[i].kodpas[h] {
				fmt.Println(arrReservasiBis[i].nama[h])
			}
		}
	}
	for i:=0; i<bookedEko; i++ {
		for h:=0; h<arrReservasiEko[i].jumlah; h++ {
			if kodpas == arrReservasiEko[i].kodpas[h] {
				fmt.Println(arrReservasiEko[i].nama[h])
			}
		}
	}
}

func searchPenumpangRes () {
	var kodres,j,k int
	j=0
	k=0
	fmt.Print("Masukkan kode reservasi Anda: ")
	fmt.Scan(&kodres)
	if kodres>1000&&kodres<2000 {
		for j<12&&kodres!=arrKodresBis[j] {
			j++	
		}
		for i:=0; i<arrReservasiBis[j].jumlah; i++ {
				fmt.Println(arrReservasiBis[j].nama[i])
		}
	} else if kodres>2000&&kodres<3000 {
		for k<42&&kodres!=arrKodresEko[k]{
			k++	
		}
		for i:=0; i<arrReservasiEko[k].jumlah; i++ {
				fmt.Println(arrReservasiEko[k].nama[i])
		}
	}
}

func changeSeatBis (i,j int, kursi bool) {
	var pilih int
	showKursiKosong()
	fmt.Println("Apakah Anda ingin mengubah posisi tempat duduk? ")
	fmt.Println("1. ya")
	fmt.Println("2. tidak")
	fmt.Print("Silahkan pilih: ")
	fmt.Scan(&pilih)
	if pilih==1 {
		kursi_Bisnis[arrReservasiBis[i].kursi[j][0]][arrReservasiBis[i].kursi[j][1]] = false
		fmt.Println("Silahkan tentukan posisi tempat duduk Anda yang belum ditempati orang lain.")
		fmt.Scan(&arrReservasiBis[i].kursi[j][0],&arrReservasiBis[i].kursi[j][1])
		kursi_Bisnis[arrReservasiBis[i].kursi[j][0]][arrReservasiBis[i].kursi[j][1]] = true	

	}	
}

func changeSeatEko (i,j int, kursi bool) {
	var pilih int
	showKursiKosong()
	fmt.Println("Apakah Anda ingin mengubah posisi tempat duduk? ")
	fmt.Println("1. ya")
	fmt.Println("2. tidak")
	fmt.Print("Silahkan pilih: ")
	fmt.Scan(&pilih)
	if pilih==1 {
		kursi_Ekonomi[arrReservasiEko[i].kursi[j][0]][arrReservasiEko[i].kursi[j][1]] = false
		fmt.Println("Silahkan tentukan posisi tempat duduk Anda yang belum ditempati orang lain.")
		fmt.Scan(&arrReservasiEko[i].kursi[j][0],&arrReservasiEko[i].kursi[j][1])
		kursi_Ekonomi[arrReservasiEko[i].kursi[j][0]][arrReservasiEko[i].kursi[j][1]] = true	
	}	
}

func searchParents () {
	var anak string
	var cari string
	fmt.Print("Masukkan nama anak atau bayi: ")
	fmt.Scan(&anak)
	for i:=0; i<bookedBis; i++ {
		for h:=0; h<arrReservasiBis[i].jumlah; h++ {
			if anak == arrReservasiBis[i].nama[h] {
				cari=arrReservasiBis[i].kodpas[h]
				var temu bool
				temu=true
				for g:=0; g<arrReservasiBis[i].jumlah; g++ {
					if arrReservasiBis[i].kodpas[g]==cari&&temu {
						fmt.Println("Orang tua dari anak tersebut adalah ",arrReservasiBis[i].nama[g])
						temu=false
					}
				}
			}
		}
	}
	for i:=0; i<bookedEko; i++ {
		for h:=0; h<arrReservasiEko[i].jumlah; h++ {
			if anak == arrReservasiEko[i].nama[h] {
				cari=arrReservasiEko[i].kodpas[h]
				var temu bool
				temu=true
				for g:=0; g<arrReservasiEko[i].jumlah; g++ {
					if arrReservasiEko[i].kodpas[g]==cari&&temu {
						fmt.Println("Orang tua dari anak tersebut adalah ",arrReservasiEko[i].nama[g])
						temu=false
					}
				}
			}
		}
	}
}

func showVegetarian () {
	fmt.Println("DAFTAR NAMA PENUMPANG KELAS BISNIS YANG VEGETARIAN")
	for i:=0; i<bookedBis; i++ {
		for j :=0; j<arrReservasiBis[i].jumlah; j++ {
			if arrReservasiBis[i].dietary[j]==true {
				fmt.Println(arrReservasiBis[i].nama[j])
			}
		}			
	}
	fmt.Println()
	fmt.Println("DAFTAR NAMA PENUMPANG KELAS EKONOMI YANG VEGETARIAN")
	for i:=0; i<bookedEko; i++ {
		for j :=0; j<arrReservasiEko[i].jumlah; j++ {
			if arrReservasiEko[i].dietary[j]==true {
				fmt.Println(arrReservasiEko[i].nama[j])
			}
		}			
	}
}

func showKidsandParents () {
	var arrUmur[30]int
	var arrNama[30]string
	var indeks[30]int
	var idx int
	for i:=0; i<bookedBis; i++ {
		for j :=0; j<arrReservasiBis[i].jumlah; j++ {
			if arrReservasiBis[i].usia[j]>2&&arrReservasiBis[i].usia[j]<18 {
				arrUmur[idx]=arrReservasiBis[i].usia[j]
				arrNama[idx]=arrReservasiBis[i].nama[j]
				indeks[idx]=idx
				idx++
			}
		}		
	}
	for i:=0; i<bookedEko; i++ {
		for j :=0; j<arrReservasiEko[i].jumlah; j++ {
			if arrReservasiEko[i].usia[j]>2&&arrReservasiEko[i].usia[j]<18 {
				arrUmur[idx]=arrReservasiEko[i].usia[j]
				arrNama[idx]=arrReservasiEko[i].nama[j]
				indeks[idx]=idx
				idx++
			}
		}		
	}
	for i:=0; i<idx; i++ {
	}
	i:=idx-1
	for i>=0 {
		maks:=0
		j:=1
		for j<=i {
			if arrUmur[j]>arrUmur[maks] {
				maks=j
			}
			j++
		}
		t:=indeks[maks]
		indeks[maks]=indeks[i]
		indeks[i]=t
		i--
	}
	for i:=0; i<idx; i++ {
		fmt.Println(arrNama[indeks[i]], "dengan umur", arrUmur[indeks[i]])
	}
}

func showKursiKosong () {
	fmt.Println("KURSI KOSONG BISNIS")
	for i := 0; i < 3; i ++{
		for j :=0; j<4; j++ {
			fmt.Print(kursi_Bisnis[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("KURSI KOSONG EKONOMI")

	for i:=0; i<7; i++ {
		for j:=0; j<6; j++ {
			fmt.Print(kursi_Ekonomi[i][j], " ")
		}
		fmt.Println()
	}
}

func catatData(nama,kodpas string, umur int) {
	f, err := os.OpenFile("APLIKASI_PENERBANGAN.txt", os.O_APPEND|os.O_CREATE, 0600) //if FOUND = append file. if !FOUND create file.

	if err != nil {
		fmt.Printf("error creating file: %v", err)
	}
	defer f.Close()
	for i := 0; i < 1; i++ { // Generating...
		_, err = f.WriteString(fmt.Sprintf("Nama = %s\nKode paspor = %s\nUmur = %d\n\n", nama, kodpas, umur)) // writing...
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}
}