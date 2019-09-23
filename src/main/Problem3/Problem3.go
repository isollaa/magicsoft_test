package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "log"
    "path/filepath"
)

func main() {
    var (
        root, temp string
        files, srcS, srcT []string
        err   error
        source, target []string
        sourceContent, targetContent []string
    )
    //Alamat file yang mulai di cek
    root = "src/main/Problem3"
    files, err = filePathWalkDir(root)
    if err != nil {
        panic(err)
    }
    //perulangan untuk membaca file sebanyak file yang ada
    for _, file := range files {
        //simbol "\" saya replace dengan "/" soalnya filepath terkadang mainnya pake simbol "\" sedangkan string tidak bisa nerima simbol tersebut / simbol dianggap punya perintah
        temp = strings.ReplaceAll(string(file), `\`, "/")
        //pengecekan nama file antara yang mengandung kata "Source" dengan yang mengandung kata "Target" dan dipisahkan variabel" tersendiri.
        if strings.Contains(temp, "Source"){
            //variabel untuk menyimpan alamat full (mulai dari src/main/Problem3/..) dari masing" file pada folder Source
            srcS = append(srcS, temp)
            //variabel untuk menyimpan isi konten dari masing" file pada folder Source
            sourceContent = append(sourceContent, string(getContent(temp)))
            //variabel untuk menyimpan alamat yang sudah terpotong (src/main/Problem3/Source)nya dari masing" file pada folder Source. Saya potong agar outputnya kelihatan sama dengan dicontoh saja sih.
            source =  append(source, after(temp, "Source"))
        //Untuk yang Target urang lebih sama dengan yang Source
        }else if strings.Contains(temp, "Target"){
            srcT = append(srcT, temp)
            targetContent = append(targetContent, string(getContent(temp)))
            target =  append(target, after(temp, "Target"))
        }
    }   
    fmt.Printf("Pilih tindakan :\n 1.Task 1\n 2.Task 2\n\nInputkan pilihan anda :" )
    var no string
    fmt.Scanln(&no)
    if (no == "1") {
        fmt.Println("Task 1\n")
        task1(source, target);
    }else if (no == "2"){
        fmt.Println("Task 2\n")
        //pemanggilan method yang berfungsi untuk mendapatkan posisi sumber dari file yang sama antara Source dengan Target pada array
        getSimilarSourcePosition(source, target);
        //variabel" ini dikosongkan untuk menghapus data yang sebelumnya ada didalamnya (sumber data/ konten data)
        sourceContent = nil;
        targetContent = nil;
        source = nil;
        target = nil;
        //perulangan untuk mengisi variabel tadi lagi dengan sumber dan konten yang baru berdasarkan alamat srcS yang ke tempS[i](tempS didapatkan saat method getSimilarSourcePosition dijalankan)
        for i:=0; i<len(tempS); i++{
            sourceContent = append(sourceContent, string(getContent(srcS[tempS[i]])))
            source = append(source, after(srcS[tempS[i]], "Source"))
        }
        for i:=0; i<len(tempT); i++{
            targetContent = append(targetContent,string(getContent(srcT[tempT[i]])))
            target = append(target, after(srcT[tempT[i]], "Target"))
        }
        //pemanggilan method untuk menampilkan file yang termodifikasi
        showModifiedFile(source, sourceContent, targetContent)
    }else{
        fmt.Println("Inputan salah, kembali keawal\n")
    }
}
//fungsi untuk mengtrim dan hanya menyimpan data setelah spesifik kata yang diinputkan
func after(value string, a string) string {
    // Get substring after a string.
    pos := strings.LastIndex(value, a)
    if pos == -1 {
        return ""
    }
    adjustedPos := pos + len(a)
    if adjustedPos >= len(value) {
        return ""
    }
    return value[adjustedPos:len(value)]
}
//fungsi untuk membaca seluruh file pada direktori yang diinginkan
func filePathWalkDir(root string) ([]string, error) {
    var files []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            files = append(files, path)
        }
        return nil
    })
    return files, err
}
//
func task1(source, target []string) {
    //variabel untuk alat bantu perhitungan, counter 1 untuk yang source, counter 2 untuk target.
    //kemudian dicek, apabila nama file pada source sama dengan nama pada target, maka counter akan bertambah
    var counter, counter2 int
	for i:=0; i<len(source); i++{
    	for j:=0; j<len(target); j++{
    		if source[i] == target[j]{
    			counter++
    		}
    		if source[j] == target[i]{
    			counter2++
    		}
    	}
        //jika counter tidak bertambah, maka nama file pada perulangan ini akan ditampilkan
    	if counter == 0{
			fmt.Println(source[i]+"  NEW")
    	}
    	if counter2 == 0{
    		fmt.Println(target[i]+"  DELETED")
    	}	
        //di akhir perulangan selalu di 0 kan kembali counternya
    	counter = 0
    	counter2 = 0
    }
}
//fungsi untuk mendapatkan isi dari file dan di returnkan dalam bentuk string
func getContent(temp string) string{
    file, err := os.Open(temp)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var content string
    for scanner.Scan() {           
        content = scanner.Text()
    }
    return content
}
//variabel untuk menyimpan posisi" dari file yang sama antara source dengan target
var tempS,tempT []int;
func getSimilarSourcePosition(source, target []string) {
    var counter, counter2  int
	for i:=0; i<len(source); i++{
    	for j:=0; j<len(target); j++{
    		if source[i] == target[j]{
    			counter++
    		}
    		if source[j] == target[i]{
    			counter2++
    		}
    	}
    	if counter == 0{
			fmt.Println(source[i]+"  NEW")
    	}//jika counter lebih dari 0, maka tempS diisi dengan i
        else{
			tempS = append(tempS, i)
    	}
    	if counter2 == 0{
    		fmt.Println(target[i]+"  DELETED")
    	}else{
			tempT = append(tempT, i)
    	}
    	counter = 0
    	counter2 = 0
    }
}
//fungsi untuk menampilkan file dengan nama yang sama dengan isi berbeda / termodifikasi
func showModifiedFile(source, sourceContent, targetContent []string) {
    for i:=0; i<len(sourceContent); i++{
    	if sourceContent[i]!=targetContent[i]{
		fmt.Println(source[i]+"  MODIFIED")
		}
    }
}

