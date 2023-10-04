# ORM dan Code Struktur-MVC

# 1 ORM (Object relational Mapping)
Object Relational Mapping (ORM) adalah teknik dalam pengembangan perangkat lunak yang memungkinkan pengembang untuk bekerja dengan basis data menggunakan objek dan kelas seperti cara mereka bekerja dengan struktur data biasa di dalam bahasa pemrograman.

Di Golang, terdapat beberapa pustaka ORM yang memungkinkan Anda untuk melakukan operasi basis data dengan menggunakan struktur dan tipe data seperti dalam bahasa pemrograman pada umumnya.

    1 -> GORM : GORM adalah salah satu pustaka ORM yang paling populer di Golang. Ia menyediakan API yang sederhana namun kuat untuk berinteraksi dengan basis data. GORM mendukung berbagai jenis basis data termasuk MySQL, PostgreSQL, SQLite, dan lainnya.

    2 -> XORM : XORM adalah pustaka ORM lain yang cukup terkenal di Golang. Ia memiliki fitur yang kuat dan dukungan untuk berbagai jenis basis data. XORM menyediakan API yang mudah digunakan untuk mengakses dan memanipulasi data.

    3 -> Ent : Ent adalah pustaka ORM baru yang dirancang khusus untuk Golang. Ia menawarkan pendekatan yang modern dan kuat untuk mengelola struktur data dan basis data. Ent memiliki fitur yang kuat untuk menghasilkan dan memanipulasi data.

    4 -> Qbs : Qbs adalah ORM ringan untuk Golang yang dirancang dengan fokus pada kinerja tinggi. Ia menawarkan API yang sederhana dan mudah digunakan untuk melakukan operasi basis data.

    5 -> Gorp : Gorp adalah pustaka ORM yang lebih sederhana dan ringan dibandingkan dengan beberapa ORM lainnya. Ia menyediakan fungsionalitas dasar untuk berinteraksi dengan basis data.

# 2 Code Struktur-MVC
Model-View-Controller (MVC) adalah pola desain perangkat lunak yang memisahkan aplikasi menjadi tiga komponen utama: Model, View, dan Controller. Setiap komponen memiliki tanggung jawab spesifik dalam aplikasi.

    1 -> Model (M)
        - Merupakan representasi dari data atau objek bisnis dalam aplikasi.
        - Bertanggung jawab untuk mengakses dan memanipulasi data, serta menangani logika bisnis.
        - Tidak harus mengetahui atau berkomunikasi langsung dengan tampilan (View) atau pengontrol (Controller).

    2 -> View (V)
        - Bertanggung jawab untuk menangani tampilan atau presentasi data kepada pengguna.
        - Tidak memiliki pengetahuan tentang logika bisnis atau cara mengakses data.
        - Mengambil informasi dari Model dan mempresentasikannya dengan cara yang sesuai.

    3 -> Controller (C)
        - Berfungsi sebagai perantara antara Model dan View.
        - Menanggapi tindakan atau permintaan pengguna dan mengambil tindakan yang sesuai, seperti memperbarui Model atau memperbarui tampilan.
        - Bertanggung jawab untuk mengarahkan alur kontrol dan memutuskan tampilan mana yang harus ditampilkan.