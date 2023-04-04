<b>Endpoint</b><br>
Endpoint adalah URL atau titik akhir dari sebuah API. Setiap API memiliki beberapa endpoint yang berbeda, yang masing-masing memberikan akses ke fungsionalitas atau data tertentu.

Endpoint biasanya diatur dalam struktur hierarkis, yang memungkinkan pengguna API untuk mengakses berbagai jenis informasi atau fungsionalitas dengan cara yang terstruktur dan terorganisir. Contohnya, sebuah API untuk aplikasi e-commerce mungkin memiliki endpoint yang berbeda untuk menampilkan daftar produk, menambahkan produk ke keranjang belanja, melakukan pembayaran, dan mengirimkan konfirmasi pesanan.

Endpoint biasanya diidentifikasi dengan URL yang spesifik dan terdiri dari beberapa bagian, seperti nama domain, path, dan parameter

<br><br><br>
<img src="https://docs.bmc.com/docs/bhom/221/files/1060822088/1109335498/1/1631692675194/API_20_08.PNG" alt="endpoint" />
<br><br><br>

Penjelasan:

- METHOD -> Dalam konteks API, "method" mengacu pada tipe permintaan HTTP yang digunakan untuk berinteraksi dengan endpoint API tertentu. Ada beberapa jenis method HTTP yang dapat digunakan dalam API, dan masing-masing memiliki fungsinya sendiri-sendiri. contoh method GET, POST, DELETE, PUT , etc.

- SCHEME -> bagian pertama dari URL yang menentukan protokol atau cara akses yang digunakan untuk mengakses sebuah sumber daya yang diidentifikasi oleh URL tersebut. Scheme biasanya dituliskan sebelum tanda titik dua  :  dan diikuti oleh "//". biasanya METHOD menggunakan HTTPS dan HTTP.
    - HTTP: digunakan untuk mengakses sumber daya melalui protokol HTTP (Hypertext Transfer Protocol), yang merupakan protokol standar untuk bertukar data di web. biasanya bersifat private atau hanya beberapa orang saja yang mengetahui aksesnya.
    
    - HTTPS: digunakan untuk mengakses sumber daya melalui protokol HTTPS (Hypertext Transfer Protocol Secure), yang menggunakan enkripsi SSL/TLS untuk memberikan keamanan dan privasi pada pertukaran data di web. biasa bersifat global dan siapapun bisa mengakses.

- Portal Url / Host -> nama atau alamat dari server atau sistem komputer yang menyediakan akses ke sumber daya yang diminta melalui URL. Host biasanya ditempatkan setelah scheme dan tanda titik dua  :  dan diikuti oleh path atau endpoint tertentu yang diminta.

- Base Path -> bagian dari URL yang merupakan awalan dari endpoint atau path lainnya pada API. Base path ini sering kali digunakan pada API dengan banyak endpoint dan sumber daya, di mana semua endpoint tersebut terletak pada alamat yang sama, sehingga dapat disederhanakan dengan menggunakan base path.

- Endpoint -> bagian dari URL yang menunjukkan alamat dari sebuah endpoint pada API. Endpoint path ini berbeda-beda pada setiap API tergantung dari struktur dan fungsionalitas API tersebut.