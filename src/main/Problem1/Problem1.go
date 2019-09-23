package main

import "fmt"

/*Alur Pikir Verical bar 
Vertical bar memiliki tinggi maksimal samadengan nilai tertinggi dari nilai array yang ada.
Sedangkan lebarnya samadengan jumlah bilangan pada array.
Maka dapat disimpulkan untuk pembukaannya kita cari dahulu nilai tertinggi dari nilai array yang ada.

Vertical bar hanya terdiri dari simbol "|" dan " ". Disini permasalahannya adalah bagaimana menentukan kapan simbol "|" muncul dan kapan simbol " " muncul.
Untuk menentukannya kita perlu melihat pola hubungan antara angka, simbol "|" dan " " tiap kolomnya, dimana simbol "|" berada pada baris sebelum angka, dan berurutan ke atas sampai simbol "|" berjumlah angka yang bersangutan.
Dari situ dapat kita lihat bahwa simbol " " mengisi bagian sisa dari tinggi maksimal terhadap simbol "|".
Oleh karena itu kita perlu membuat sebuah variabel yang menyimpan selisih dari nilai maksimal terhadap nilai bilangan masing" nya. 
Setelah itu tampilkan simbol "|" pada baris yang memiliki nilai lebih / samadengan selisih itu tadi. Dan ketika baris kurang dari selisih tersebut, tampilkan simbol " "
*/

/*Alur Pikir Insertion Sort
Insertion sort dilakukan dengan pembandingan antara nilai sekarang dengan nilai sebelumnya. Oleh sebab itu, perulangan dimulai dari bilangan ke 2.
Alurnya dimulai dengan pengecekan apakah bilangan sekarang lebih dari / kurang dari bilangan sebelumnya ? jika iya, nilai akan ditukar.
Kemudian pengecekan akan dilakukan terhadap nilai berikutnya terhadap nilai sekarang - sebelumnya.
Dan begitu terus pengecekan dilakukan sampai nilai terakhir dicek.
*/ 

func main() {
    length := 0
    fmt.Printf("Masukkan jumlah bilangan dalam array : ")
    fmt.Scanln(&length)
    fmt.Println("Masukkan bilangan satu per-satu : ")
    numbers := make([]int, length)
    for i := 0; i < length; i++ {
        fmt.Scanln(&numbers[i])
    }

	fmt.Printf("Pilih tindakan :\n 1.Visualisasi\n 2.Pengurutan (Ascending)\n 3.Pengurutan (Ascending Reverse)\n 4.Pengurutan (Descending)\n\nInputkan pilihan anda :" )
	var no string
	fmt.Scanln(&no)
	if (no == "1") {
		fmt.Println("Visualisasi\n")
    	visual(length, numbers)
	}else if (no == "2"){
		fmt.Println("Pengurutan (Ascending)\n")
		insertionSort(length, numbers)
	}else if (no == "3"){
		fmt.Println("Pengurutan (Ascending Reverse)\n")
		insertionSortReverse(length, numbers)
	}else if (no == "4"){
		fmt.Println("Pengurutan (Descending)\n")
		insertionSortDescending(length, numbers)
	}else{
		fmt.Println("Inputan salah, kembali keawal\n")
	}
}
//Fungsi untuk menampilkan visualisasi vertical bar tiap array
func visual(length int, numbers []int) {
	//pertama" cari nilai terbesar pada array dan simpan pada variabel bernama temp. Kemudian samadengankan dengan nilai dari array pertama / ke 0
    temp := numbers[0]
    //buat perulangan sebanyak panjang array
    for i:=0; i<length; i++{
    	//buat kondisi yang menentukan apakah nilai dari array ke [i] lebih dari temp, jika iya. Samadengankan temp dengan nilai dari array ke [i] tadi
        if (numbers[i]>temp) {
            temp = numbers[i]
        }
    }
	//buat variabel temporari lagi dengan nama x untuk menyimpan hasil pengurangan antara nilai terbesar dengan nilai yang ada pada array
    //lakukan perulangan sebanyak 2x
    //perulangan pertama untuk menentukan banyak baris yang akan dibuat, dan perulangan kedua untuk menentukan banyak kolom yang akan dibuat
    x := 0
    for i:=0; i<temp;i++{
        for j:=0; j<length;j++{
            x = temp - numbers[j]
            //buat kondisi yang menentukan apakah i (baris yang sedang aktif) memiliki nilai lebih dari hasil pengurangan tadi.
            //Jika iya, print simbol " |", sedangkan jika tidak maka print "  "           
            if (i>=x){
                fmt.Printf(" |")
            }else{
                fmt.Printf("  ")
            }
        }
        //setelah pengerjaan satu baris selesai, beri enter untuk memisahkan baris sekarang dengan setelahnya
        fmt.Println("\n")
    }
    //print array semua nilai tepat setelah simbol" dibuat
    fmt.Println(numbers, "\n")
}
//Fungsi untuk pengurutan menggunakan insertion sort
func insertionSort(length int, numbers []int) {
	//panggil fungsi visual yang sebelumnya telah dibuat untuk menampilkan visualisasi awal sebelum disorting
	visual(length, numbers)	
	//buat variabel temp untuk menyimpan nilai terkecil dan variabel j untuk menyimpan nilai posisi perulangan setelah posisi sekarang
	var temp, j int
	//buat perulangan mundur agar pengecekan dilakukan kanan ke kiri
	for i:=length-2; i>=0;i--{
		//sama dengankan dulu variabel temp dengan nilai array pertama
		temp=numbers[i]
		//seperti dijelaskan sebelumnya, samadengankan j dengan posisi perulangan sekarang + 1
		j=i+1
		//buat perulangan dengan syarat j tidak lebih dari (panjang array - 1) sekaligus nilai array ke [j] harus kurang dari temp
		//selama syarat terpenuhi, nilai array ke [j-1] / array di bagian kanan akan disamadengankan dengan nilai array ke [j]
		//kemudian j ditambahkan untuk melakukan pengecekan pada seluruh kolom dari baris yang sama
		for (j<=length-1) && (numbers[j]<temp){
			numbers[j-1]=numbers[j]
			j++
		}
		//jika kondisi tidak memenuhi / perulangan telah dilakukan, samadengankan nilai array ke [j-1] dengan nilai temp yang merupakan nilai dari array ke [i] 
		numbers[j-1]=temp
		//tampilkan visualisasi dari masing" percobaan pengurutan
		visual(length, numbers)
	}
}
//Fungsi InsertionSort versi dari kiri ke kanan prosesnya hampir sama hanya berbeda di perulangannya
//disini perulangannya dimulai dari angka kecil ke besar (maju)
func insertionSortReverse(length int, numbers []int) {
	visual(length, numbers)	
	var temp, j int
	for i:=1; i<length;i++{
		temp=numbers[i]
		j=i-1
		for (j>=0) && (numbers[j]>temp){
			numbers[j+1]=numbers[j]
			j--
		}
		numbers[j+1]=temp
		visual(length, numbers)
	}
}
//Fungsi InsertionSort versi descending dan dilakukan dari kiri ke kanan
//Alurnya hampir sama hanya beda di bagian nilai yang ditukar saja
func insertionSortDescending(length int, numbers []int) {
	visual(length, numbers)
	var temp, j int
	for i:=1; i<length;i++{
		temp=numbers[i]
		j=i
		for (j>0) && (numbers[j-1]<temp){
			numbers[j]=numbers[j-1]
			j--
		}
		numbers[j]=temp
		visual(length, numbers)
	}
}
