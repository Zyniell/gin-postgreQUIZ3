package controllers

import (
	"quizz3-buku/config" // Ganti dengan path modul yang benar jika perlu
	"net/http"
	"github.com/gin-gonic/gin"
)

// Struct Buku disesuaikan dengan skema tabel buku
type Buku struct {
	ID          int     `json:"id_buku"`
	Judul       string  `json:"judul_buku"`
	Penulis     string  `json:"penulis"`
	Penerbit    string  `json:"penerbit"`
	TahunTerbit int     `json:"tahun_terbit"`
	Harga       float64 `json:"harga"`
	Stok        int     `json:"stok"`
	IDKategori  *int    `json:"id_kategori"` // Gunakan pointer jika bisa bernilai null
}

// --- CREATE ---
func CreateBuku(c *gin.Context) {
	var newBuku Buku

	if err := c.ShouldBindJSON(&newBuku); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid"})
		return
	}

	// Validasi minimal: Judul wajib diisi
	if newBuku.Judul == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Judul buku wajib diisi"})
		return
	}

	query := `INSERT INTO buku (judul_buku, penulis, penerbit, tahun_terbit, harga, stok, id_kategori) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id_buku`
	
	err := config.DB.QueryRow(query, 
		newBuku.Judul, 
		newBuku.Penulis, 
		newBuku.Penerbit, 
		newBuku.TahunTerbit, 
		newBuku.Harga, 
		newBuku.Stok, 
		newBuku.IDKategori).Scan(&newBuku.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Buku berhasil ditambahkan", "data": newBuku})
}

// --- READ ALL ---
func GetAllBuku(c *gin.Context) {
	var results []Buku

	sqlStatement := `SELECT id_buku, judul_buku, penulis, penerbit, tahun_terbit, harga, stok, id_kategori FROM buku`

	rows, err := config.DB.Query(sqlStatement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data buku"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var b Buku
		err = rows.Scan(&b.ID, &b.Judul, &b.Penulis, &b.Penerbit, &b.TahunTerbit, &b.Harga, &b.Stok, &b.IDKategori)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal scanning data"})
			return
		}
		results = append(results, b)
	}

	if results == nil {
		results = []Buku{}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil semua data buku",
		"data":    results,
	})
}

// --- READ BY ID ---
func GetBukuByID(c *gin.Context) {
	id := c.Param("id")
	var b Buku

	query := `SELECT id_buku, judul_buku, penulis, penerbit, tahun_terbit, harga, stok, id_kategori 
              FROM buku WHERE id_buku = $1`
	
	err := config.DB.QueryRow(query, id).
		Scan(&b.ID, &b.Judul, &b.Penulis, &b.Penerbit, &b.TahunTerbit, &b.Harga, &b.Stok, &b.IDKategori)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, b)
}

// --- UPDATE ---
func UpdateBuku(c *gin.Context) {
	id := c.Param("id")
	var b Buku

	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah"})
		return
	}

	if b.Judul == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Judul buku tidak boleh kosong"})
		return
	}

	query := `UPDATE buku SET judul_buku=$1, penulis=$2, penerbit=$3, tahun_terbit=$4, harga=$5, stok=$6, id_kategori=$7 
              WHERE id_buku=$8`
	
	res, err := config.DB.Exec(query, b.Judul, b.Penulis, b.Penerbit, b.TahunTerbit, b.Harga, b.Stok, b.IDKategori, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update data"})
		return
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data buku berhasil diperbarui"})
}

// --- DELETE ---
func DeleteBuku(c *gin.Context) {
	id := c.Param("id")

	res, err := config.DB.Exec("DELETE FROM buku WHERE id_buku = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal hapus data"})
		return
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil dihapus"})
}

// Struct untuk Kategori
type Kategori struct {
	IDKategori   int    `json:"id_kategori"`
	NamaKategori string `json:"nama_kategori"`
}

// Tambahkan fungsi ini agar tidak undefined
func CreateKategori(c *gin.Context) {
	var newKategori Kategori
	if err := c.ShouldBindJSON(&newKategori); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	query := `INSERT INTO kategori (nama_kategori) VALUES ($1) RETURNING id_kategori`
	err := config.DB.QueryRow(query, newKategori.NamaKategori).Scan(&newKategori.IDKategori)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newKategori})
}