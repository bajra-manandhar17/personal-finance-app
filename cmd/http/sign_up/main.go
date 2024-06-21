package sign_up

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/bajra-manandhar17/personal-finance-app/internal/business/user"
	"github.com/bajra-manandhar17/personal-finance-app/internal/config"
	"github.com/bajra-manandhar17/personal-finance-app/internal/db/query"
	"github.com/bajra-manandhar17/personal-finance-app/internal/helper/httphelper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	postgresDbProvider *gorm.DB
	userService        user.UserService
)

func init() {
	ctx := context.Background()
	localPostgresDbProvider, err := config.NewPostgresDbProvider(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	postgresDbProvider = localPostgresDbProvider
	query := query.Use(postgresDbProvider)

	userService = user.NewUserService(query)
}

func handler(request *http.Request) (httphelper.HttpResponse, error) {
	ctx := request.Context()

	var reqData user.RegisterNewUserReq
	if err := httphelper.MapAndValidateBody(request.Body, &reqData); err != nil {
		return httphelper.MapErrorToApiResponse(err)
	}

	if err := userService.RegisterNewUser(ctx, reqData); err != nil {
		return httphelper.MapErrorToApiResponse(err)
	}

	return httphelper.PrepareApiResponse(nil)
}

func RegisterNewUserHandler(c *gin.Context) {
	httpResp, _ := handler(c.Request)

	// Unmarshal the JSON response body
	var responseBody map[string]interface{}
	if err := json.Unmarshal([]byte(httpResp.Body), &responseBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error unmarshalling response body"})
		return
	}

	c.JSON(httpResp.StatusCode, responseBody)
}
