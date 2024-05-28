package router

import (
	"bytes"
	"easy-images/config"
	"easy-images/internal/common"
	"easy-images/internal/handler"
	"easy-images/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func Bind(port int) {
	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()

	route.Use(cros)
	route.Use(panicError)
	route.POST("/upload", upload)
	route.Static("/images", "images")
	addr := ":" + strconv.Itoa(port)
	logger.Infof("server start by http://0.0.0.0%s", addr)
	if err := route.Run(addr); err != nil {
		logger.Fatal(err)
	}
}

func cros(context *gin.Context) {
	method := context.Request.Method
	context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
	context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
	context.Header("Access-Control-Allow-Headers", "*")
	context.Header("Access-Control-Expose-Headers", "*")
	context.Header("Access-Control-Max-Age", "172800")
	context.Header("Access-Control-Allow-Credentials", "false")
	context.Set("content-type", "application/json")

	if method == "OPTIONS" {
		context.Status(http.StatusOK)
		return
	}

	uid := uuid.NewString()
	// 请求打印
	data, err := httputil.DumpRequest(context.Request, false)
	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("\n------ START REQUEST %s ---------\n%s", uid, data)
	}

	//处理请求
	context.Next()

	// 结束处理
	logger.Infof("\n------ END REQUEST %s ---------", uid)
}

func panicError(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("response error: %v", r)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": map[string]string{
					"message": fmt.Sprintf("%v", r),
				},
			})
		}
	}()

	//处理请求
	ctx.Next()
}

func upload(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	uploadToken := config.Config.GetString("TOKEN")
	if token != "Bearer "+uploadToken {
		common.Error(ctx, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := ctx.Request.ParseMultipartForm(10 << 20); err != nil {
		common.Error(ctx, -1, err)
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		common.Error(ctx, -1, err)
		return
	}

	file, err := image.Open()
	if err != nil {
		common.Error(ctx, -1, err)
		return
	}

	buffer := new(bytes.Buffer)
	if _, err := buffer.ReadFrom(file); err != nil {
		common.Error(ctx, -1, err)
		return
	}

	path, err := handler.Handler(buffer.Bytes(), image.Filename)
	if err != nil {
		common.Error(ctx, -1, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"URL": config.Config.GetString("DOMAIN") + "/" + path,
	})
}
