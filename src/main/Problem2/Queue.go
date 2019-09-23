package main

import (
	"fmt"
)

/*Untuk bagian ini saya masih kurang paham sih sama pemakaian interface, berhubung ini dihubungkan dengan QueueTest.go,
saya mikirnya kelas ini bakal diagregasi sama kelas itu, dan ternyata beda.
Untuk pengujian kelas ini, masih saya lakukan secara individu(run kelas ini sendiri) tidak melalui QueueTest.go
Jadi disini saya hanya menjelaskan fungsi-fungsi dasar dalam queue saja..*/

func main() {
	size := 7
	q := New(size)
	fmt.Println("", q.Len())
	q.Push(6)
	q.Push(5)
	q.Push(4)
	q.Push(3)
	q.Push(1)
	q.Push(2)
	q.Push(2)
	fmt.Println("", q.Len())
	fmt.Println("", q.Keys())
}

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

type UniqueQueue struct{
	keys []interface{}
}
//fungsi untuk mengetahui banyak data dalam array yang tidak samadengan null.
//data akan dicek sebanyak jumlah data dalam array, dan apabila tidak samadengan null. Maka len akan mereturn nilai tersebut.
func (q *UniqueQueue) Len() int{
	count:=0
	for i:=0; i<len(q.keys); i++{
		if q.keys[i]!=nil{
			count++
		}
	}
	return count
}
//fungsi untuk memasukkan data kedalam array.
//Karena ini adalah unique queue.. Maka sebelum penginputan data, program harus melakukan pengecekan dahulu.
//Apakah data yang diinputkan sama dengan data yang ada dalam array. Jika data sama, maka penginputan akan diabaikan, 
//bila sebaliknya, data akan diinputkan seperti biasa. 
//Sehubungan dengan fixed size, maka sebelum penginputan, jumlah data non-null akan di cek tehadap panjang array.
//Jika panjangnya sudah sama, maka pengepop an akan dilakukan terhadap nilai paling kanan / nilai yang pertama dimasukkan
func (q *UniqueQueue) Push(key interface{}){
	if q.Contains(key)==false{
		if q.Len() == len(q.keys){
			q.Pop()
		}
		if q.Len()==0{
			q.keys[len(q.keys)-1] = key		
		}else if q.Len()==len(q.keys){
			q.keys[len(q.keys)-(q.Len())] = key
		}else{
			q.keys[len(q.keys)-(q.Len()+1)] = key
		}
		q.Len()
	}
}
//Fungsi pop dilakukan dengan perulangan yang dilakukan sebanyak jumlah data non-null dikurang 1 karena peritungan array dilakukan mulai dari 0 bukan 1.
//Dalam setiap perulangannya, nilai akan digeser 1 step ke kanan (nilai yang paling awal diinputkan) dan nilai paling awal tersebut akan dikeluarkan dari array.
func (q *UniqueQueue) Pop()interface{}{
	for i:=q.Len()-1; i>0; i--{
		q.keys[i] = q.keys[i-1]
	}
	return q.keys
}
//Fungsi ini digunakan untuk mengecek apakah data yang diinputkan sudah ada dalam array / belum.
//Jika data sudah ada, maka fungsi ini akan mereturn true, jika belum akan mereturn false
func (q *UniqueQueue) Contains(key interface{}) bool{
	var isContains bool = false
	for i:=0; i<len(q.keys); i++{
		if key == q.keys[i]{
			isContains = true
		}
	}
	return isContains 
}
func (q *UniqueQueue) Keys()[]interface{}{
	return q.keys
}
func New(size int) Queue {
	var q Queue = &UniqueQueue{make([]interface{},size)}
	return q
}