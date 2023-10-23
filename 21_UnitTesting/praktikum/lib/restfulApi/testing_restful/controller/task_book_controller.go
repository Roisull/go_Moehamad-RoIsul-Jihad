package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/user/go_MoehamadRoisulJihad/21_UnitTesting/praktikum/lib/restfulApi/testing_restful/config"
	"github.com/user/go_MoehamadRoisulJihad/21_UnitTesting/praktikum/lib/restfulApi/testing_restful/model"
)

// mengambil semua data buku
func GetBooksController(c echo.Context) error {
	var books []model.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{

		"message": "Success get all books",
		"books": books,
	})
}

// mengambil data buku berdasarkan ID
func GetBookByIdController(c echo.Context) error {

	// mendapatkan id dari parameter url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID Format",
		})
	}

	// temukan pengguna dengan ID yang sesuai
	var book model.Book
	result := config.DB.First(&book, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message" : "Books Not Found",
		})
	}

	return c.JSON(http.StatusOK, book)
}

// membuat data buku baru
func CreateBookController(c echo.Context) error {
	book := model.Book{}

	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create new book",
		"book": book,
	})
}

// mengahapus data buku berdasarkan ID
func DeleteBookController(c echo.Context) error {

	// mendapatkan id buku dari parameter url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID Format",
		})
	}

	// temukan buku dengan ID yang sesuai
	var book model.Book
	result := config.DB.First(&book, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message" : "Books Not Found",
		})
	}

	// hapus data buku dari database
	config.DB.Delete(&book)

	// kirim response json dengan pesan sukses hapus data buku
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Book deleted successfully",
	})
}

// update & mengubah data buku berdasarkan ID
func UpdateBookController(c echo.Context) error {

	// mendapatkan id buku dari parameter url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID Format",
		})
	}

	// temukan buku dengan ID yang sesuai
	var book model.Book
	result := config.DB.First(&book, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message" : "Books Not Found",
		})
	}

	// binding data baru dari permintaan
	newBookData := new(model.Book)
	if err := c.Bind(newBookData); err != nil {
		return err
	}

	// memperbarui data buku
	book.Judul = newBookData.Judul
	book.Penulis = newBookData.Penulis
	book.Penerbit = newBookData.Penerbit

	// simpan perubahan ke database
	config.DB.Save(&book)

	// kirim respons json dengan pesan sukses dan data buku berhasil di perbarui
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Book Updated Successfully",
		"user": book,
	})
}
