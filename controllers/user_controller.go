package controllers

import (
	"net/http"

	"advanced/config"
	"advanced/constants"
	"advanced/structs"
	"advanced/utils"

	"github.com/gin-gonic/gin"
)

// CreateUser adalah endpoint untuk register atau memasukkan data baru
func CreateUser(c *gin.Context) {
	var input structs.User

	// BindJSON otomatis mengecek validasi berdasar struct tag `binding:"required,email"` dll
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, constants.RCInvalidInput, constants.MsgInvalidInput+": "+err.Error())
		return
	}

	// Simpan ke database menggunakan GORM
	if err := config.DB.Create(&input).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, constants.RCDBError, constants.MsgGagalSimpanUser+": "+err.Error())
		return
	}

	// Jangan kembalikan password di response JSON demi keamanan
	input.Password = "KNTL"

	utils.ResponseSuccess(c, http.StatusCreated, constants.MsgSuksesBikinData, input)
}

// GetAllUsers adalah endpoint untuk mendapatkan semua user
func GetAllUsers(c *gin.Context) {
	var users []structs.User

	// Cek error jika GORM gagal query
	if err := config.DB.Find(&users).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, constants.RCUserError, constants.MsgGagalAmbilUser)
		return
	}

	// Filter atau sembunyikan password sebelum dikirim ke response
	for i := range users {
		users[i].Password = "KNTL"
	}

	utils.ResponseSuccess(c, http.StatusOK, constants.MsgSuksesAmbilData, users)
}
