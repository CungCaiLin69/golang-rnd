package middleware

import (
	"fmt"
	"golang-rnd/initializers"
	"golang-rnd/lib"
	"golang-rnd/schema"
	"net/http"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type RoleActionEnum string

const (
	Create   RoleActionEnum = "C"
	Read     RoleActionEnum = "R"
	Update   RoleActionEnum = "U"
	Delete   RoleActionEnum = "D"
	Upload   RoleActionEnum = "A"
	Download RoleActionEnum = "B"
	Archive  RoleActionEnum = "AV"
)

func (r RoleActionEnum) String() string {
	return string(r)
}

func (r RoleActionEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(r))
}

func (r *RoleActionEnum) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	switch str {
	case "C", "R", "U", "D", "A", "B", "AV":
		*r = RoleActionEnum(str)
	default:
		return fmt.Errorf("invalid RoleActionEnum value: %s", str)
	}
	return nil
}

func ValidateProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			apiErr := lib.NewApiError("JWT Invalid", []string{"Authorization token required"}, false, nil)
			c.JSON(http.StatusBadRequest, apiErr)
			c.Abort()
			return
		}

		claims, err := jwtParse(token)
		if err != nil {
			apiErr := lib.NewApiError("JWT Invalid", []string{err.Error()}, false, nil)
			c.JSON(http.StatusUnauthorized, apiErr)
			c.Abort()
			return
		}

		c.Set("UserAccount", claims)
		c.Next()
	}
}

func jwtParse(tokenString string) (*schema.JwtCommunicator, error) {
	claims := &schema.JwtCommunicator{}
	config := initializers.LoadConfig()
	secretKey := []byte(config.JwtSecret)

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			apiErr := lib.NewApiError("Unexpected signing method", nil, false, nil)
			return nil, apiErr
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		apiErr := lib.NewApiError("Cannot validate token", []string{"Invalid token"}, false, nil)
		return nil, apiErr
	}

	return claims, nil
}

func ValidateMatrix(role RoleActionEnum) gin.HandlerFunc {
	if role == "" {
		role = Read
	}

	return func(c *gin.Context) {
		uac, exists := c.Get("UserAccount")
		if !exists {
			apiErr := lib.NewApiError("Matrix Failed", []string{"Unauthorized"}, false, nil)
			c.JSON(http.StatusUnauthorized, apiErr)
			c.Abort()
			return
		}

		userAccount, ok := uac.(*schema.JwtCommunicator)
		if !ok {
			apiErr := lib.NewApiError("Matrix Failed", []string{"Unauthorized"}, false, nil)
			c.JSON(http.StatusUnauthorized, apiErr)
			c.Abort()
			return
		}

		if userAccount.UserMatrix == nil {
			apiErr := lib.NewApiError("Matrix Failed", []string{"Unauthorized"}, false, nil)
			c.JSON(http.StatusUnauthorized, apiErr)
			c.Abort()
			return
		}

		if !hasPermission(userAccount.UserMatrix, role) {
			apiErr := lib.NewApiError("Matrix Failed", []string{"Permission Denied"}, false, nil)
			c.JSON(http.StatusForbidden, apiErr)
			c.Abort()
			return
		}

		c.Set("UserMatrix", userAccount.UserMatrix)
		c.Next()
	}
}

func hasPermission(matrix *schema.UserMatrix, role RoleActionEnum) bool {
	switch role {
	case Create:
		return matrix.IsCreate != nil && *matrix.IsCreate
	case Read:
		return matrix.IsRead != nil && *matrix.IsRead
	case Update:
		return matrix.IsUpdate != nil && *matrix.IsUpdate
	case Delete:
		return matrix.IsDelete != nil && *matrix.IsDelete
	case Upload:
		return matrix.IsUpload != nil && *matrix.IsUpload
	case Download:
		return matrix.IsDownload != nil && *matrix.IsDownload
	default:
		return false
	}
}
