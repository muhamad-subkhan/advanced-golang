<b>Headers</b><br>
Headers atau header fields adalah bagian dari sebuah request atau response pada HTTP (Hypertext Transfer Protocol) yang digunakan untuk menyampaikan informasi tambahan tentang request atau response tersebut. Header terdiri dari pasangan nama dan nilai (key-value pairs) yang disampaikan dalam bentuk teks ASCII.

Pada request, headers digunakan untuk memberikan informasi tambahan tentang request, seperti jenis konten yang diminta (content type), bahasa yang digunakan (accept-language), dan data autentikasi yang dibutuhkan (authorization). Sedangkan pada response, headers digunakan untuk memberikan informasi tambahan tentang response, seperti jenis konten yang disampaikan (content type), panjang data (content length), dan status response (status code).
<br><br><br>

<img src='https://ithelp.ithome.com.tw/upload/images/20211021/20135414kiBgWoQlxB.png' alt='image' />
<br><br><br>

Beberapa contoh headers yang umum digunakan pada HTTP request dan response adalah sebagai berikut:

- Accept: header ini digunakan pada request untuk menunjukkan jenis konten yang diinginkan oleh client, seperti text/plain atau application/json.

- Content-Type: header ini digunakan pada request atau response untuk menunjukkan jenis konten yang disampaikan, seperti text/html atau image/png.

- Authorization: header ini digunakan pada request untuk menyampaikan data autentikasi, seperti token atau username dan password.

- User-Agent: header ini digunakan pada request untuk memberikan informasi tentang client yang melakukan request, seperti jenis browser atau aplikasi.

- Cache-Control: header ini digunakan pada request atau response untuk memberikan instruksi tentang caching, seperti apakah response boleh di-cache atau tidak.

Headers sangat penting dalam pengiriman data melalui protokol HTTP karena dapat mempengaruhi bagaimana request dan response tersebut diterima dan diproses oleh server dan client. Headers juga dapat memberikan informasi tambahan yang berguna bagi pengembangan dan pemeliharaan aplikasi web atau API.