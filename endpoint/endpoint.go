package endpoint

import (
	"case2/model"
	"case2/usecase"
	"net/http"
	"strconv"

	b64 "encoding/base64"

	"github.com/labstack/echo/v4"
	zlog "github.com/rs/zerolog/log"
)

func MakeCreateProductEndpoint(u usecase.ProductUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			createRequest model.CreateRequest
			err           error
			resp          model.CreateResponse
		)

		if err = c.Bind(&createRequest); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		if err = c.Validate(createRequest); err != nil {
			zlog.Info().Interface("error", err).Msg("Validate Param Create")
			return err
		}

		if resp, err = u.Create(c.Request().Context(), createRequest); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func MakeDetailProductEndpoint(u usecase.ProductUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			detailRequest model.DetailRequest
			err           error
			resp          model.DetailResponse
			sDec          []byte
			productID     int
		)

		base64ID := c.Param("id")
		if sDec, err = b64.URLEncoding.DecodeString(base64ID); sDec == nil || err != nil {
			zlog.Info().Interface("error", err).Msg("Validate Param Detail")
			return c.String(http.StatusBadRequest, "invalid Parameter")
		}

		strDec := string(sDec)
		if productID, err = strconv.Atoi(strDec); err != nil {
			zlog.Info().Interface("error", err).Msg("Failed conv str to int")
			return c.String(http.StatusBadRequest, "invalid Parameter")
		}

		detailRequest.ProductID = int64(productID)
		if resp, err = u.Detail(c.Request().Context(), detailRequest); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func MakeListProductEndpoint(u usecase.ProductUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			listRequest model.ListRequest
			err         error
			resp        *model.ListResponse
		)

		listRequest.Search = c.QueryParam("search")
		page, _ := strconv.Atoi(c.QueryParam("page"))
		listRequest.Page = int32(page)

		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		listRequest.Limit = int32(limit)

		if resp, err = u.List(c.Request().Context(), &listRequest); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func MakeUpdateProductEndpoint(u usecase.ProductUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			updateRequest model.UpdateRequest
			err           error
			resp          model.UpdateResponse
			sDec          []byte
			productID     int
		)

		if err = c.Bind(&updateRequest); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		if err = c.Validate(updateRequest); err != nil {
			zlog.Info().Interface("error", err).Msg("Validate Param Create")
			return err
		}

		base64ID := c.Param("id")
		if sDec, err = b64.URLEncoding.DecodeString(base64ID); sDec == nil || err != nil {
			zlog.Info().Interface("error", err).Msg("Validate Param Detail")
			return c.String(http.StatusBadRequest, "invalid Parameter")
		}

		strDec := string(sDec)
		if productID, err = strconv.Atoi(strDec); err != nil {
			zlog.Info().Interface("error", err).Msg("Failed conv str to int")
			return c.String(http.StatusBadRequest, "invalid Parameter")
		}

		updateRequest.ProductID = int64(productID)
		if resp, err = u.Update(c.Request().Context(), updateRequest); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func MakeDeleteProductEndpoint(u usecase.ProductUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			deleteRequest model.DeleteRequest
			err           error
			resp          model.DeleteResponse
			sDec          []byte
			productID     int
		)

		base64ID := c.Param("id")
		if sDec, err = b64.URLEncoding.DecodeString(base64ID); sDec == nil || err != nil {
			zlog.Info().Interface("error", err).Msg("Validate Param Detail")
			return c.JSON(http.StatusBadRequest, "invalid Parameter")
		}

		strDec := string(sDec)
		if productID, err = strconv.Atoi(strDec); err != nil {
			zlog.Info().Interface("error", err).Msg("Failed conv str to int")
			return c.JSON(http.StatusBadRequest, "invalid Parameter")
		}

		deleteRequest.ProductID = int64(productID)
		if resp, err = u.Delete(c.Request().Context(), deleteRequest); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}
