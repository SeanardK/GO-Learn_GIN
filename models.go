package main

import (
	"errors"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{
		ID:      1,
		Title:   "Pengenalan Golang",
		Content: "Go, juga dikenal sebagai Golang, adalah bahasa pemrograman yang dikembangkan di Google oleh Robert Griesemer, Rob Pike, dan Ken Thompson. Go dirilis pada tahun 2009 dan sejak itu telah mendapatkan popularitas yang signifikan karena kemudahan penggunaannya dan kinerja yang tinggi.",
	},
	{
		ID:      2,
		Title:   "Penggunaan Interface dalam Go",
		Content: "Interface adalah salah satu fitur kunci dalam bahasa pemrograman Go. Interface memungkinkan Anda untuk mendefinisikan kontrak yang harus dipatuhi oleh tipe data lain. Hal ini memungkinkan Anda untuk membuat kode yang lebih fleksibel dan mudah digunakan.",
	},
	{
		ID:      3,
		Title:   "Penggunaan Goroutines dalam Go",
		Content: "Goroutines adalah cara Go mengelola konkurensi. Mereka memungkinkan Anda untuk menjalankan beberapa fungsi secara bersamaan tanpa harus membuat thread secara eksplisit. Ini membuat pengelolaan konkurensi menjadi lebih mudah dan aman.",
	},
	{
		ID:      4,
		Title:   "Penggunaan Channel dalam Go",
		Content: "Channel adalah cara komunikasi antara Goroutines dalam Go. Mereka memungkinkan Goroutines untuk mengirim dan menerima data satu sama lain. Ini sangat berguna dalam menghindari race conditions dan memfasilitasi komunikasi yang aman antara Goroutines.",
	},
	{
		ID:      5,
		Title:   "Penggunaan Slice dalam Go",
		Content: "Slice adalah tipe data yang sangat fleksibel dalam Go. Mereka memungkinkan Anda untuk membuat dan memanipulasi koleksi data dengan mudah. Slice sering digunakan dalam pemrograman Go untuk operasi dengan array yang dinamis.",
	},
	{
		ID:      6,
		Title:   "Pemrograman Fungsional dalam Go",
		Content: "Meskipun Go bukanlah bahasa pemrograman fungsional murni, namun ia mendukung beberapa konsep pemrograman fungsional seperti fungsi sebagai tipe data dan closure. Hal ini memungkinkan Anda untuk menulis kode yang lebih deklaratif dan ekspresif.",
	},
	{
		ID:      7,
		Title:   "Penggunaan Struct dalam Go",
		Content: "Struct adalah tipe data yang memungkinkan Anda untuk menggabungkan beberapa tipe data menjadi satu kesatuan logis. Ini sangat berguna untuk merepresentasikan data yang kompleks dalam program Go.",
	},
	{
		ID:      8,
		Title:   "Manajemen Error dalam Go",
		Content: "Go memiliki pendekatan yang unik dalam manajemen error. Biasanya, fungsi yang mungkin menghasilkan kesalahan akan mengembalikan dua nilai: nilai aktual dan objek error. Ini memungkinkan Anda untuk dengan mudah menangani kesalahan dalam program Anda.",
	},
	{
		ID:      9,
		Title:   "Penggunaan Pointer dalam Go",
		Content: "Pointer adalah fitur penting dalam Go yang memungkinkan Anda untuk bekerja dengan referensi langsung ke lokasi memori. Hal ini memungkinkan Anda untuk mengakses dan memanipulasi data secara langsung, tanpa harus membuat salinan data yang besar.",
	},
	{
		ID:      10,
		Title:   "Penggunaan Map dalam Go",
		Content: "Map adalah struktur data kunci-nilai dalam Go yang memungkinkan Anda untuk menyimpan dan mengakses data dengan cepat. Map sangat berguna ketika Anda perlu mengaitkan nilai dengan kunci tertentu, dan sering digunakan dalam berbagai aplikasi Go.",
	},
}

func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
