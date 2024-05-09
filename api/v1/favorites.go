package v1

import (
	"Mall/consts"
	"Mall/pkg/utils/ctl"
	"Mall/pkg/utils/log"
	"Mall/service"
	"Mall/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateFavoriteHandler 创建收藏
func CreateFavoriteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.FavoriteCreateReq
		if err := c.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		l := service.GetFavoriteSrv()
		resp, err := l.FavoriteCreate(c.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// ListFavoritesHandler 收藏夹详情接口
func ListFavoritesHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.FavoritesServiceReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		if req.PageSize == 0 {
			req.PageSize = consts.BasePageSize
		}

		l := service.GetFavoriteSrv()
		resp, err := l.FavoriteList(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// DeleteFavoriteHandler 删除收藏夹
func DeleteFavoriteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.FavoriteDeleteReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetFavoriteSrv()
		resp, err := l.FavoriteDelete(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}
