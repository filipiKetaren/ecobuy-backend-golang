package product

import (
	"ecobuy/services/product"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	ProductService product.ProductServiceInterface
}

func NewProductController(ps product.ProductServiceInterface) *ProductController {
	return &ProductController{
		ProductService: ps,
	}
}

func (pc *ProductController) GetProductsController(c echo.Context) error {
	category := c.QueryParam("category")

	// Ambil query parameter untuk pagination
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Set default nilai pagination jika tidak diisi
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	products, err := pc.ProductService.GetProducts(category, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Gagal mengambil data produk",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil mendapatkan daftar produk",
		"data":    products,
	})
}

func (pc *ProductController) GetProductDetailController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "ID produk tidak valid",
		})
	}

	product, err := pc.ProductService.GetProductByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Produk tidak ditemukan",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil mendapatkan detail produk",
		"data":    product,
	})
}
