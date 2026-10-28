1. Apa fungsi utama dari struktur perulangan for dalam program Go? Sebutkan dan
jelaskan tiga komponen pada perulangan for (inisialisasi, kondisi, dan post statement)!

2. Bahasa Go tidak memiliki perulangan while dan do-while. Jelaskan perbedaan antara
for dengan tiga komponen dan for yang hanya menggunakan kondisi! Berikan contoh
kode masing-masing.

Jawab :
1. Untuk menjalankan suatu instruksi hingga mencapai kondisi yang diinginkan.
    - inisialisasi: dijalankan sekali dan dipakai untuk mengisi kepada variabel penghitung. contoh : "i := 1"
    - kondisi: di evaluasi pada awal iterasi, untuk memeriksa apakah nilai nya TRUE atau FALSE. Jika TRUE maka akan tetap berjalan, jika FALSE maka berhenti
    - post statement: dijalankan saat akhir iterasi, biasanya digunakan untuk menambah atau mengurangi nilai pada variabel. contoh : "i++"

2. - for dengan tiga komponen: dipakai ketika kita tahu sampai kapan perulangan itu dijalankan. contoh :
    "for i := 0; i < 5; i++ {
        fmt.Println("Nilai i:", i)
    }"


    - for dengan kondisi saja : digunakan ketika kita menginginkan kode nya tetap berjalan pada kondisi tertentu . contoh :
    "i := 0
    for i < 5 {
        fmt.Println("Nilai i:", i)
        i++
    }"