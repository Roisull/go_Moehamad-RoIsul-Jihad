# summary System Design

# A. ERD
    -> Entitas (Entity)
    Representasi dari objek atau konsep dalam dunia nyata yang dapat disimpan dalam basis data. Misalnya, "Pengguna", "Produk", "Pesanan", dll.

    -> Atribut (Attribute)
    Informasi atau properti yang terkait dengan entitas. Misalnya, untuk entitas "Pengguna", atribut dapat mencakup nama, alamat, email, dan sebagainya.

    -> Hubungan (Relationship)
    Keterkaitan atau koneksi antara dua atau lebih entitas. Ini menunjukkan bagaimana entitas berinteraksi satu sama lain. Contohnya adalah hubungan "Mempunyai" antara "Pengguna" dan "Pesanan".

    -> Kardinalitas
    Menunjukkan jumlah entitas terkait yang dapat terjadi dalam suatu hubungan. Misalnya, hubungan "Satu-ke-Banyak" antara "Pengguna" dan "Pesanan" menunjukkan bahwa satu pengguna dapat memiliki banyak pesanan.

    -> Primary Key (Kunci Utama)
    Atribut atau kumpulan atribut yang secara unik mengidentifikasi setiap entitas dalam suatu tabel atau koleksi data. Kunci utama digunakan untuk memastikan integritas data.

    -> Foreign Key (Foreign Key)
    Atribut atau kumpulan atribut yang menghubungkan dua tabel atau entitas dalam basis data. Kunci asing menghubungkan entitas dalam suatu hubungan.

    -> Normalisasi
    Proses desain basis data untuk mengorganisir struktur data sehingga meminimalkan redundansi dan memungkinkan pengelolaan data yang efisien.

    -> Notasi ERD
    Ada beberapa notasi yang digunakan untuk menggambar ERD, termasuk notasi Chen, Crow's Foot, dan Barker.

    -> Tipe-Tipe Hubungan
    1. Satu-ke-Satu (One-to-One): Satu entitas pada sisi pertama hanya dapat terhubung dengan satu entitas pada sisi kedua, dan sebaliknya.
    2. Satu-ke-Banyak (One-to-Many): Satu entitas pada sisi pertama dapat terhubung dengan banyak entitas pada sisi kedua, tetapi tidak sebaliknya.
    3. Banyak-ke-Banyak (Many-to-Many): Banyak entitas pada sisi pertama dapat terhubung dengan banyak entitas pada sisi kedua, dan sebaliknya.

# B. Use Case Diagram
    -> Aktor(actor)
    Pemangku kepentingan eksternal yang berinteraksi dengan sistem. Aktor dapat berupa pengguna manusia, sistem eksternal, atau entitas lain yang terlibat dalam interaksi dengan sistem.

    -> Use Case
    Representasi dari fungsionalitas atau tindakan yang dapat dilakukan oleh sistem. Ini menggambarkan aktivitas atau tugas tertentu yang dilakukan oleh satu atau lebih aktor.

    -> Hubungan Antara Aktor dan Use Case
    Garis yang menghubungkan aktor dengan use case menunjukkan keterlibatan aktor dalam skenario tertentu. Aktor berpartisipasi dalam use case jika mereka terhubung dengannya.

    -> Penggunaan Kasus Alternatif
    Use case juga dapat memiliki skenario alternatif atau ekstensi yang menunjukkan jalur yang berbeda dari jalur utama. Hal ini memungkinkan untuk menangani situasi-situasi yang tidak terduga atau penanganan kesalahan.

    -> Inklusi dan Ekstensi (Inclusion and Extension)
    Inklusi adalah relasi di antara use case yang menunjukkan bahwa satu use case dapat memasukkan fungsionalitas dari use case lainnya. Ekstensi adalah relasi yang menunjukkan bahwa satu use case dapat memperpanjang fungsionalitas dari use case lainnya.

    -> Penggunaan Kasus Abstrak (Abstract Use Case)
    Use case abstrak adalah use case yang memiliki fungsionalitas umum dan dapat digunakan sebagai kerangka untuk use case yang lebih spesifik.

    -> Sistem (System Boundary)
    Batas fisik atau logis yang membatasi sistem yang diwakili dalam use case diagram.

    -> Deskripsi Use Case
    Informasi tambahan atau deskripsi singkat yang dapat disertakan untuk menjelaskan lebih lanjut tentang fungsi use case.

    -> Tujuan Use Case Diagram
    Memberikan pandangan tingkat tinggi tentang fungsionalitas sistem dan bagaimana aktor berinteraksi dengannya.
    Memfasilitasi komunikasi antara tim pengembang dan pemangku kepentingan untuk memahami kebutuhan dan skenario pengguna.

# C. Redis
    Redis adalah sebuah sistem manajemen basis data (Database Management System) berkinerja tinggi, berbasis memori (in-memory), dan sumber daya terdistribusi. Redis dikenal sebagai "store data di dalam memori" atau "data struktur dalam memori" karena data disimpan dalam RAM untuk memungkinkan akses data yang sangat cepat.

# D. Neo4j
    Neo4j adalah sebuah sistem manajemen basis data grafik (Graph Database Management System) yang dirancang khusus untuk menyimpan, mengelola, dan mengakses data yang memiliki struktur grafik. Dalam Neo4j, data diwakili sebagai simpul (nodes) dan hubungan (relationships) yang memungkinkan representasi yang efisien dari koneksi kompleks antara entitas.

# E. Cassandra
    Apache Cassandra adalah sebuah sistem manajemen basis data distribusi terdistribusi dan tingkat tinggi (distributed and high availability database management system) yang dirancang untuk menangani volume data besar dan skala yang besar. Cassandra termasuk dalam kategori basis data NoSQL, yang berarti tidak mengikuti model relasional tradisional dan lebih terfokus pada skalabilitas horizontal dan kinerja tinggi.