# Time Complexity

-> Time Complexity, bagaimana kita mengitung sebuah algoritma atau sebuah sistem akan berjalan dalam hal waktu.
-> bagaimana cara kita menghitung waktu tersebut?
    1. kita harus menentukan, seberapa banyak operasi dominan yang dilakukan
        operasi dominan adalah operasi yang primitif, misalnya penambahan, perkalian, assigment.

-> ada beberapa jenis dari time compelexity
    1. Constant Time o(1) = ketika kita menginputkan berapapun, maka operasi yang ada dalam algoritma didalamnya hanya dilakukan satu kali.
    2. Linear Time o(n) = efisiensi di mana waktu eksekusi algoritma tumbuh secara linier seiring dengan pertambahan ukuran input. Dengan kata lain, semakin besar ukuran input yang diberikan kepada algoritma, semakin banyak langkah yang diperlukan untuk menjalankan algoritma akan bertambah secara proporsional.
    3. Logarithmic Time o(log n) = waktu eksekusi algoritma tumbuh secara logaritmik seiring dengan pertambahan ukuran input. Dengan kata lain, semakin besar ukuran input yang diberikan kepada algoritma, waktu eksekusi akan bertambah, tetapi pertambahan ini akan berlangsung dengan pertumbuhan logaritmik.
    4. Quadratic Time o(n^2) = algoritma tumbuh secara kuadratik seiring dengan pertambahan ukuran input. Dengan kata lain, semakin besar ukuran input yang diberikan kepada algoritma, waktu eksekusi akan bertambah secara kuadratik.

-> Space Complexity = adalah konsep dalam analisis algoritma yang berkaitan dengan penggunaan sumber daya memori oleh suatu algoritma selama eksekusi. Dalam istilah sederhana, space complexity mengukur berapa banyak memori yang diperlukan oleh suatu algoritma untuk menjalankan operasinya sesuai dengan ukuran input.