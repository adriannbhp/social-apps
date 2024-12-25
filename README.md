
Dari informasi yang Anda berikan, tampaknya ada beberapa hal yang perlu diperiksa untuk memastikan bahwa aplikasi Anda menggunakan port yang benar sesuai dengan variabel lingkungan yang telah Anda atur. Mari kita telusuri langkah-langkah berikut:

1. Periksa File .envrc
   Pastikan file .envrc Anda berisi:

bash

export ADDR=":3030"
2. Izinkan direnv untuk Memuat .envrc
   Setelah Anda membuat atau mengubah file .envrc, Anda perlu memberi izin kepada direnv untuk memuat file tersebut. Jalankan perintah berikut di terminal:

bash

direnv allow
3. Cek Variabel Lingkungan
   Setelah menjalankan direnv allow, periksa apakah variabel lingkungan ADDR telah diatur dengan benar dengan menjalankan:

bash

echo $ADDR
Outputnya harus :3030. Jika outputnya tidak sesuai, maka ada masalah dengan pemuatan file .envrc.

4. Debugging di Kode Go
   Pastikan Anda menambahkan log di dalam fungsi main untuk melihat nilai yang diambil dari variabel lingkungan:

go

log.Printf("ADDR: %s", cfg.addr) // Tambahkan log ini untuk debugging
5. Periksa Proses Build
   Pastikan bahwa Anda membangun aplikasi Anda dengan benar. Anda dapat mencoba membangun aplikasi secara manual dengan perintah berikut:

bash

go build -o ./bin/main ./cmd/api
Setelah itu, pastikan file main.exe ada di dalam folder bin.

6. Restart air
   Setelah melakukan perubahan, pastikan untuk menghentikan dan memulai ulang air. Ini penting karena air mungkin tidak mendeteksi perubahan pada file .envrc jika tidak di-restart.

Hentikan air dengan menekan Ctrl + C.
Jalankan kembali air:
bash

air
7. Periksa Output Log
   Setelah menjalankan aplikasi, periksa output log untuk memastikan bahwa server dimulai pada alamat yang benar. Anda harus melihat log yang menunjukkan alamat yang digunakan, misalnya:


ADDR: :3030
server has started at :3030
8. Cek Versi air
   Pastikan Anda menggunakan versi terbaru dari air. Anda dapat memperbarui air dengan menjalankan:

bash

go get -u github.com/cosmtrek/air
9. Cek Izin File
   Pastikan file .envrc memiliki izin yang benar. Anda dapat memeriksa izin file dengan perintah berikut:

bash

ls -l .envrc
10. Cek Konfigurasi Shell
    Pastikan bahwa hook direnv telah ditambahkan ke file konfigurasi shell Anda (seperti ~/.bashrc atau ~/.zshrc). Contoh untuk Bash:

bash

eval "$(direnv hook bash)"
Setelah menambahkan baris ini, muat ulang file konfigurasi shell Anda:

bash

source ~/.bashrc  # untuk Bash
# atau
source ~/.zshrc   # untuk Zsh
Kesimpulan