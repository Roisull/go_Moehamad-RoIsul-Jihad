# summary materi docker

perngertian: sebuah platform perangkat lunak yang dibuat untuk memungkinkan developer untuk membungkus, mendistribusikan dan menjalankan aplikasi dalam lingkungan yang diisolasi disebuat juga container.

# point kunci tentang docker

1. Containerization
   mengemas aplikasi dan semua dependensinya ke dalam unit terisolasi yang disebut container. Ini memastikan bahwa aplikasi dapat dijalankan dengan konsisten di berbagai lingkungan.

2. Portabilitas
   Container Docker dapat dijalankan di berbagai sistem operasi dan lingkungan tanpa perubahan, menyediakan portabilitas yang tinggi. Ini memudahkan pengembangan, pengujian, dan distribusi aplikasi.

3. Lingkungan Terisolasi
   Setiap container menjalankan aplikasi dan dependensinya di lingkungan terisolasi. Ini memungkinkan aplikasi berjalan secara independen tanpa konflik dengan aplikasi atau dependensi lainnya di sistem host.

4. Docker Image
   Docker menggunakan konsep gambar (image) sebagai template untuk membuat container. Image berisi sistem file dan konfigurasi yang diperlukan untuk menjalankan aplikasi tertentu.

5. DockerFile
   Dockerfile adalah skrip yang digunakan untuk membangun Docker image. Ini berisi instruksi untuk mengatur lingkungan, menginstal dependensi, dan mengkonfigurasi aplikasi.

6. Docker Hub
   Docker Hub adalah registry yang menyediakan repositori publik dan pribadi untuk Docker images. Pengembang dapat berbagi dan mendistribusikan image melalui Docker Hub.

7. Kontainer vs Mesin Virtual
   Container Docker lebih ringan dibandingkan mesin virtual karena mereka menggunakan kernel sistem operasi host dan berbagi sumber daya dengan sistem host. Mesin virtual, di sisi lain, memerlukan hypervisor untuk menyediakan lingkungan virtual penuh.

8. Orkestrasi
   Docker mendukung alat-orangkastasi seperti Docker Compose dan Kubernetes untuk mengelola dan mendeploy aplikasi yang terdiri dari beberapa container.

9. Skalabilitas
   Docker memungkinkan pengguna untuk dengan mudah menambah atau mengurangi jumlah instance container sesuai kebutuhan, mendukung skala horizontal dengan mudah.

10. Kompatibilitas
    Docker kompatibel dengan banyak bahasa pemrograman dan teknologi, membuatnya cocok untuk berbagai jenis aplikasi dan mikroservis arsitektur.
