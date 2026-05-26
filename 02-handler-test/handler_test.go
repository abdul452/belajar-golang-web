package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic handler
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})

	mux.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "About Page")
	})

	mux.HandleFunc("/contact", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Contact Page")
	})

	// 📌 CATATAN 1: Rute "/images/" (Menggunakan garis miring di akhir)
	// Sifat rute ini adalah CATCH-ALL (Menangkap semua sub-URL di bawahnya).
	// Jika ada user mengakses "/images/", "/images/abdul", "/images/123", atau "/images/apapun/itu",
	// selama tidak ada rute lain yang lebih spesifik, maka semuanya akan ditangkap oleh handler ini.
	//
	// Contoh: Akses "http://localhost:8080/images/abdul" ➡️ Tetap memunculkan "Images Page" (Bukan 404 Not Found).
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Images Page")
	})

	// 📌 CATATAN 2: Rute "/images/thumbnails/" (Rute yang Lebih Spesifik)
	// Di Go, ServeMux itu pintar. Jika ada beberapa rute yang mirip (sama-sama berawalan /images/),
	// Go akan selalu memprioritaskan rute yang polanya PALING PANJANG dan PALING SPESIFIK.
	//
	// Contoh: Akses "http://localhost:8080/images/thumbnails/" ➡️ Otomatis masuk ke sini dan memunculkan "Thumbnail Images Page",
	// karena sistem mendeteksi rute ini lebih cocok dan presisi daripada rute "/images/" yang terlalu umum di atas.
	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Thumbnail Images Page")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Method)     // isinya contoh GET, POST, PUT, DELETE, dll
		fmt.Fprintln(writer, request.RequestURI) // isinya contoh /, /about, /contact, dll
		fmt.Fprintln(writer, request.Proto)      // isinya contoh HTTP/1.1
		fmt.Fprintln(writer, request.Header)     // isinya contoh User-Agent, Accept, Content-Type, dll
		fmt.Fprintln(writer, request.URL)        // isinya contoh http://localhost:8080/, http://localhost:8080/about, dll
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
