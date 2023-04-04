# simple-socket (full-duplex)

## Cara menggunakan 
- buka 2 terminal dengan posisi endpoint pada server dan client
- jalankan server dahulu dengan cara `go run server.go`
- jalankan client dengan cara `go run client.go`

setelah client dan server berjalan, message dari client akan muncul pada server dan message server akan muncul pada client. ini dinamakan `full-duplex` yang mana komunikasi 2 arah bisa berjalan secara bersamaan.

komunikasi ini sering kita lihat pada fitur chat aplikasi