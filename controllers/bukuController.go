package controllers

import (
	"context"
	"net/http"
	"quizz3-buku/config"

	"github.com/gin-gonic/gin"
)

type Buku struct {
	ID          int     `json:"id_buku"`
	Judul       string  `json:"judul_buku"`
	Penulis     string  `json:"penulis"`
	Penerbit    string  `json:"penerbit"`
	TahunTerbit int     `json:"tahun_terbit"`
	Harga       float64 `json:"harga"`
	Stok        int     `json:"stok"`
	IDKategori  int     `json:"id_kategori"`
}

// GetBuku menangani Get All dan Get by ID
func GetBuku(c *gin.Context) {
	id := c.Param("id")
	
	if id != "" {
		// Logic Get By ID
		var b Buku
		err := config.DB.QueryRow(context.Background(), 
			"SELECT id_buku, judul_buku, penulis, penerbit, tahun_terbit, harga, stok, id_kategori FROM buku WHERE id_buku=$1", id).
    		Scan(&b.ID, &b.Judul, &b.Penulis, &b.Penerbit, &b.TahunTerbit, &b.Harga, &b.Stok, &b.IDKategori)
		
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
			return
		}
		c.JSON(http.StatusOK, b)
		return
	}

	// Logic Get All
	rows, _ := config.DB.Query(context.Background(), "SELECT id_buku, judul_buku, penulis, penerbit, tahun_terbit, harga, stok, id_kategori FROM buku")
	defer rows.Close()

	var bukuList []Buku
	for rows.Next() {
		var b Buku
		rows.Scan(&b.ID, &b.Judul, &b.Penulis, &b.Penerbit, &b.TahunTerbit, &b.Harga, &b.Stok, &b.IDKategori)
		bukuList = append(bukuList, b)
	}
	c.JSON(http.StatusOK, bukuList)
}

func CreateBuku(c *gin.Context) {
	var b Buku
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec(context.Background(),
		`INSERT INTO buku (judul_buku, penulis, penerbit, tahun_terbit, harga, stok, id_kategori) 
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		b.Judul, b.Penulis, b.Penerbit, b.TahunTerbit, b.Harga, b.Stok, b.IDKategori)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal simpan data: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Buku berhasil ditambahkan"})
}

func UpdateBuku(c *gin.Context) {
	id := c.Param("id")
	var b Buku
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec(context.Background(),
		`UPDATE buku SET 
			judul_buku=$1, penulis=$2, penerbit=$3, tahun_terbit=$4, 
			harga=$5, stok=$6, id_kategori=$7 
		 WHERE id_buku=$8`,
		b.Judul, b.Penulis, b.Penerbit, b.TahunTerbit, b.Harga, b.Stok, b.IDKategori, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil diupdate"})
}

func DeleteBuku(c *gin.Context) {
	id := c.Param("id")
	_, err := config.DB.Exec(context.Background(), "DELETE FROM buku WHERE id_buku=$1", id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil dihapus"})
}