package middlewares

import (
	"net/http"
	"os"

	"test-api/index/advanced/constants"
	"test-api/index/advanced/utils"

	"github.com/gin-gonic/gin"
)

// SimpleAuthMiddleware adalah contoh middleware untuk ngecek token dari header
// Cara kerjanya: Sebelum request masuk ke Controller utama, request akan difilter disini.
func SimpleAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mengambil Token dari HTTP Header "Authorization"
		token := c.GetHeader("Authorization")

		// Mengecek token buatan. Real case biasanya memparsing & memvalidasi JWT.
		// Untuk contoh kita hanya mengecek apakah nilainya cocok dengan SECRET_KEY di .env
		secret := os.Getenv("SECRET_KEY")
		if token != secret {
			// Batalkan request dengan standar format error dan msg dari constants
			utils.ResponseError(c, http.StatusUnauthorized, constants.RCAuthError, constants.MsgAksesDitolak)
			c.Abort()
			return
		}

		// Jika valid, teruskan request ke Controller selanjutnya (misal GetAllUsers)
		c.Next()
	}
}
