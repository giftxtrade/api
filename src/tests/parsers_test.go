package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/giftxtrade/api/src/controllers"
	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
)

func TestGetBearerToken(t *testing.T) {
    // Test properly formatted authorization field
    {
        token1, err := utils.GetBearerToken("Bearer my_token")
        if token1 != "my_token" && err == nil {
            t.Errorf("did not parse bearer token properly")
        }
    }

    // Test with empty bearer token
    {
        _, err := utils.GetBearerToken("Bearer")
        if err == nil {
            t.Errorf("did not parse bearer token properly")
        }
    }

    // Test with no value
    {
        _, err := utils.GetBearerToken("")
        if err == nil {
            t.Errorf("did not parse bearer token properly")
        }
    }
}

func TestGetJwtClaims(t *testing.T) {
    // Test with correct key and claims
    {
        key := "abcd123"
        jwt := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbWFpbCI6ImV4YW1wbGVAZW1haWwuY29tIiwidXNlcm5hbWUiOiJleGFtcGxlIn0.OTAHN2L8ag5n20Xbfa-T4PRpLvsnp4pNFb8D3zHluyg"

        claims_map, err := utils.GetJwtClaims(jwt, key)
        username := claims_map["username"].(string)
        email := claims_map["email"].(string)
        if err != nil || email != "example@email.com" || username != "example" {
            fmt.Println(err)
            t.Fail()
        }
    }

    // Test with correct jwt and incorrect key
    {
        key := "incorrect key"
        jwt := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2NDIzOTc5OTksImV4cCI6MTY3MzkzMzk5OSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsImVtYWlsIjoiZXhhbXBsZUBlbWFpbC5jb20iLCJ1c2VybmFtZSI6ImV4YW1wbGUifQ.fBJbtYyIJuHA6Ip8OlQuVmDrHlIhtSAlx7S3lUBK_qM"

        if _, err := utils.GetJwtClaims(jwt, key); err == nil {
            t.Fail()
        }
    }

    // Test with correct key and incorrect jwt
    {
        key := "q34859t8jsvdh1"
        jwt := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2NDIzOTc5OTksImV4cCI6MTY3MzkzMzk5OSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsImBtYWlsIjoiZXhhbXBsZUBlbWFpbC5jb20iLCJ1c2VybmFtZSI6ImV4YW1wbGUifQ.Ohw7jfG65CzgiTB-DZMVoKl67APTeJrwrmHd3Ex9KX0"

        if _, err := utils.GetJwtClaims(jwt, key); err == nil {
            t.Fail()
        }
    }
}

func TestGenerateTokens(t *testing.T) {
    {
        user := database.User{
            ID: 1,
            Email: "johndoe@example.com",
            Name: "John Doe",
        }
        user_service := services.UserService{}
        jwt1, err1 := user_service.GenerateJWT("123", &user)
        jwt2, err2 := user_service.GenerateJWT("1234", &user)

        if err1 != nil || err2 != nil || jwt1 == jwt2 {
            t.Fail()
        }
    }
}

func TestParseAuthContext(t *testing.T) {
    {
        ctx := context.Background()
        user := types.User{
            ID: 2,
            Email: "johndoe2@example.com",
            Name: "John Doe 2",
        }
        token := "my random token"
        ctx = context.WithValue(ctx, controllers.AUTH_KEY, types.Auth{Token: token, User: user})
        parsed_auth := controllers.ParseAuthContext(ctx)

        if parsed_auth.User != user {
            t.Fail()
        }

        if parsed_auth.Token != token {
            t.Fail()
        }
    }
}