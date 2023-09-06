# resume sigkat ConcurrentProgramming

Concurrent programming dalam Go (Golang) adalah kemampuan untuk menjalankan beberapa tugas (goroutine) secara bersamaan, sehingga memungkinkan eksekusi konkuren yang efisien. Berikut adalah beberapa poin kunci tentang concurrent programming di Go:

a. Goroutines: Go menggunakan goroutine sebagai unit utama untuk konkurensi. Goroutines adalah fungsi-fungsi ringan yang dapat dijalankan secara asinkron. Anda dapat membuat dan menjalankan banyak goroutine dengan mudah, dan mereka diatur oleh runtime Go, bukan oleh sistem operasi.

b. Channels: Channels adalah mekanisme komunikasi yang digunakan untuk berbagi data antara goroutine. Mereka memungkinkan goroutine untuk berkomunikasi dengan aman dan sinkron, mencegah race condition dan deadlock. Dengan channels, Anda dapat mengirim dan menerima data antara goroutine.

c. Select Statement: select adalah pernyataan khusus di Go yang digunakan untuk mengatur komunikasi antara goroutine melalui channels. Ini memungkinkan Anda untuk memilih operasi yang siap untuk dijalankan pada channels yang berbeda, menghadle kasus konkurensi yang lebih kompleks.

d. Concurrency Patterns: Go menyediakan berbagai pola konkurensi yang sudah ada dan mudah digunakan. Ini termasuk pola seperti fan-out, fan-in, worker pools, dan banyak lagi. Pola-pola ini memungkinkan Anda untuk merancang dan mengimplementasikan konkurensi dengan lebih mudah dan aman.

e. Synchronization: Go memiliki fitur built-in seperti sync package yang digunakan untuk mengkoordinasikan goroutine. Ini termasuk sync.WaitGroup untuk menunggu selesainya goroutine, sync.Mutex untuk menghindari race condition, dan banyak lagi.

f. Concurrency Safety: Bahasa Go dirancang dengan keselamatan konkurensi yang tinggi. Ini berarti bahwa bahasa memiliki desain yang meminimalkan kesalahan umum dalam konkurensi, sehingga Anda dapat menulis kode konkurensi yang lebih aman dan mudah dipahami.