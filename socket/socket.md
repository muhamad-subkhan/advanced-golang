# Socket (Web Socket)<br>
Websocket adalah standar baru untuk komunikasi realtime pada web dan aplikasi mobile. Websocket dirancang untuk diterapkan di browser dan server web. Namun, dapat digunakan oleh aplikasi client atau server. 

Web socket merupakan protocol penyedia saluran komunikasi full-duplex (komunikasi dua arah secara bersamaan) melalui TCP tunggal.


### Manfaat Web Socket
- memungkinkan server mendorong data kepada klien yang terhubung.
- mengurangi lalu lintas jaringan yang tidak perlu dan latecy menggunakan full-duplex melalui koneksi tunggal.
- Streaming melalui proxy dan firewall, mendukung komunikasi simultan hulu ke hilir.
- kompetibel denga pre-websocket dunia dengan cara beralih dari koneksi HTTP ke websockets.

<br>
<img src="https://anotherorion.com/wp-content/uploads/2015/02/Simplex-half-full-Duplex.png" alt="image" />
<br><br><br>

### Simplex<br>
Simplex adalah komunikasi yang tidak memungkinkan penerima dan pengirim saling bertukar informasi. Pada komunikasi ini sinyal-sinyal dikirim hanya satu arah saja dalam waktu yang bersamaan. karena melalui satu arah saja, komunikasi ini tidak terjadi secara interaktif, informasi hanya disampaikan melalui satu titik saja.

<img src="https://www.pintarkomputer.com/wp-content/uploads/2014/09/simplex.png" alt="image" />
<br><br><br><br><br><br>

### Half Duplex<br>
half duplex adalah metode yang memungkinkan komunikasi antara dua belah pihak yaitu pengirim dan penerima dapat saling berbagi informasi dan berkomunikasi secara interaktif, tetapi tidak dalam waktu yang bersamaan.

<img src="https://www.pintarkomputer.com/wp-content/uploads/2014/09/half-duplex.png" alt="image" />
<br><br><br><br><br><br>

### Full Duplex atau Duplex<br>
Full duplex adalah metode yang memungkinkan komunikasi antara dua belah pihak yaitu pengirim dan penerima dapat saling berbagi informasi dan berkomunikasi secara interaktif dan dalam waktu yang bersamaan.

<img src="https://www.pintarkomputer.com/wp-content/uploads/2014/09/full-duplex.png" alt="image" />
<br><br><br><br><br><br>