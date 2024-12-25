#!/bin/bash

# Nama file .envrc
ENVRC_FILE=".envrc"

# Cek apakah file .envrc sudah ada
if [ -f "$ENVRC_FILE" ]; then
    echo "$ENVRC_FILE sudah ada. Menghapus file yang lama."
    rm "$ENVRC_FILE"
fi

# Buat file .envrc baru dan tambahkan variabel lingkungan
echo "Membuat file $ENVRC_FILE dengan variabel lingkungan."

cat <<EOL > "$ENVRC_FILE"
export ADDR=":3000"
# Tambahkan variabel lain di bawah ini
# export ANOTHER_VAR="value"
EOL

echo "$ENVRC_FILE telah dibuat dengan variabel ADDR."
