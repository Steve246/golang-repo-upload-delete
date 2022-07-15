package controller

import (
	"errors"
	"golang-upload-download/delivery/api"
	"golang-upload-download/model"
	"golang-upload-download/usecase"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	router        *gin.Engine
	ucProduct     usecase.CreateProductUseCase
	ucProductList usecase.ListProductUseCase
	ucFindProduct usecase.FindProductUseCase
	api.BaseApi
}

//https://github.com/gin-gonic/gin#multiparturlencoded-form
//https://github.com/gin-gonic/gin#upload-files
func (p *ProductController) createNewProduct(c *gin.Context) {

	//kita buat form inputnya multiport//form data
	productId := c.PostForm("productId")
	productName := c.PostForm("productName")
	file, fileHeader, err := c.Request.FormFile("productImg")
	if err != nil {
		p.Failed(c, errors.New("failed get file"))
	}
	log.Println(fileHeader.Filename) //gambarbagus.jpg

	fileName := strings.Split(fileHeader.Filename, ".")
	if len(fileName) != 2 {
		p.Failed(c, errors.New("unrecognized file extension"))
	} //ini kita pisahin jadi gambarbagus dan jpg, terus cek extensionnya sama atau ngak 
	newProduct := model.Product{
		ProductId:   productId,
		ProductName: productName,
	}

	//fileName[1] ambil jpg
	err = p.ucProduct.CreateProduct(&newProduct, file, fileName[1])
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, newProduct)
}

func (p *ProductController) findAllProduct(c *gin.Context) {
	products, err := p.ucProductList.List()
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, products)
}
func (p *ProductController) findImageById(c *gin.Context) {
	productId := c.Param("id")
	product, err := p.ucFindProduct.ById(productId)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.SuccessDownload(c, product.ImgPath)
}

func NewProductController(
	router *gin.Engine,
	ucProduct usecase.CreateProductUseCase,
	ucProductList usecase.ListProductUseCase,
	ucFindProduct usecase.FindProductUseCase) *ProductController {
	// Disini akan terdapat kumpulan semua request method yang dibutuhkan
	controller := ProductController{
		router:        router,
		ucProduct:     ucProduct,
		ucProductList: ucProductList,
		ucFindProduct: ucFindProduct,
	}

	// ini method-method nya
	router.POST("/product", controller.createNewProduct)
	router.GET("/product", controller.findAllProduct)
	router.GET("/product/image/:id", controller.findImageById)
	return &controller
}
