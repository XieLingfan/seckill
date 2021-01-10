package product

import (
	"context"
	"fmt"
	"net/http"
	product_srv "product_srv/proto/product"
	"strconv"
	"time"
	"web/utils"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-grpc"
)

func GetProductList(ctx *gin.Context) {
	currentPage := ctx.DefaultQuery("currentPage", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")

	// 和srv通信获取products数据
	service := grpc.NewService()

	product_service := product_srv.NewProductsService("xlf.product.srv.product_srv", service.Client())

	rep, err := product_service.ProductList(context.TODO(), &product_srv.ProductsRequest{
		CurrentPage: utils.StrToInt(currentPage),
		Pagesize:    utils.StrToInt(pageSize),
	})

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":         rep.Code,
		"msg":          rep.Msg,
		"products":     rep.Products,
		"total":        rep.Total,
		"current_page": rep.Current,
		"page_size":    rep.PageSize,
	})

}

func ProductAdd(ctx *gin.Context) {

	name := ctx.PostForm("name")
	price := ctx.PostForm("price")
	num := ctx.PostForm("num")
	unit := ctx.PostForm("unit")
	desc := ctx.PostForm("desc")

	file, err := ctx.FormFile("pic")

	// 和srv通信获取front_users数据
	service := grpc.NewService()

	product_service := product_srv.NewProductsService("xlf.product.srv.product_srv", service.Client())
	fmt.Println(product_service.ProductList)

	fmt.Println("+++++++++++++++++")
	fmt.Println(utils.StrToFloat32(price))
	if err != nil {
		// 发生错误，不保存file
		fmt.Println(err)

		rep, _ := product_service.ProductAdd(context.TODO(), &product_srv.ProductAddRequest{
			Name:  name,
			Num:   utils.StrToInt(num),
			Price: utils.StrToFloat32(price),
			Unit:  unit,
			Desc:  desc,
		})

		ctx.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})

	}

	unix_int64 := time.Now().Unix()
	unix_str := strconv.FormatInt(unix_int64, 10)
	file_path := "upload/" + unix_str + file.Filename

	ctx.SaveUploadedFile(file, file_path)

	rep, err := product_service.ProductAdd(context.TODO(), &product_srv.ProductAddRequest{
		Name:  name,
		Num:   utils.StrToInt(num),
		Price: utils.StrToFloat32(price),
		Unit:  unit,
		Desc:  desc,
		Pic:   file_path,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})

}

func ProductDel(ctx *gin.Context) {
	id := ctx.PostForm("id")
	fmt.Println(id)

	// 和srv通信获取products数据
	service := grpc.NewService()

	product_service := product_srv.NewProductsService("xlf.product.srv.product_srv", service.Client())

	rep, _ := product_service.ProductDel(context.TODO(), &product_srv.ProductDelRequest{
		Id: utils.StrToInt(id),
	})

	ctx.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})

}

func ProductToEdit(ctx *gin.Context) {

	id := ctx.Query("id")
	fmt.Println("=====================")
	fmt.Println(id)

	// 和srv通信获取products数据
	service := grpc.NewService()

	product_service := product_srv.NewProductsService("xlf.product.srv.product_srv", service.Client())
	rep, err := product_service.ProductToEdit(context.TODO(), &product_srv.ProductToEditRequest{
		Id: utils.StrToInt(id),
	})

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "没有查询到数据",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":       rep.Code,
		"msg":        rep.Msg,
		"product":    rep.Product,
		"img_base64": utils.Img2Base64(rep.Product.Pic),
	})

}

func ProductDoEdit(ctx *gin.Context) {

	id := ctx.PostForm("id")
	name := ctx.PostForm("name")
	price := ctx.PostForm("price")
	num := ctx.PostForm("num")
	unit := ctx.PostForm("unit")
	desc := ctx.PostForm("desc")

	file, err := ctx.FormFile("pic")

	// 和srv通信获取front_users数据
	service := grpc.NewService()

	product_service := product_srv.NewProductsService("xlf.product.srv.product_srv", service.Client())
	fmt.Println(product_service.ProductList)

	fmt.Println("+++++++++++++++++")
	fmt.Println(utils.StrToFloat32(price))
	if err != nil {
		// 发生错误，不保存file
		fmt.Println(err)

		rep, _ := product_service.ProductDoEdit(context.TODO(), &product_srv.ProductEditRequest{
			Id:    utils.StrToInt(id),
			Name:  name,
			Num:   utils.StrToInt(num),
			Price: utils.StrToFloat32(price),
			Unit:  unit,
			Desc:  desc,
		})

		ctx.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})

	}

	unix_int64 := time.Now().Unix()
	unix_str := strconv.FormatInt(unix_int64, 10)
	file_path := "upload/" + unix_str + file.Filename

	ctx.SaveUploadedFile(file, file_path)

	rep, err := product_service.ProductDoEdit(context.TODO(), &product_srv.ProductEditRequest{
		Id:    utils.StrToInt(id),
		Name:  name,
		Num:   utils.StrToInt(num),
		Price: utils.StrToFloat32(price),
		Unit:  unit,
		Desc:  desc,
		Pic:   file_path,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})

}
