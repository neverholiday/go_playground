package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	googleoauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type Handler struct {
	Config *oauth2.Config
}

func newHandler(config *oauth2.Config) *Handler {
	return &Handler{Config: config}
}

func (h *Handler) login(c echo.Context) error {
	authUrl := h.Config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, authUrl)
}

func (h *Handler) callback(c echo.Context) error {

	ctx := c.Request().Context()
	code := c.FormValue("code")
	token, err := h.Config.Exchange(c.Request().Context(), code)
	if err != nil {
		return c.String(http.StatusInternalServerError, "unable exchange token")
	}

	client := h.Config.Client(ctx, token)

	srv, err := googleoauth2.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}
	userInfo, err := srv.Userinfo.Get().Fields("email").Context(ctx).Do()
	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}

	// save token file
	jsonByte, err := json.Marshal(token)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}

	slog.Info("write file token.json")
	os.WriteFile("token.json", jsonByte, 0644)

	return c.String(http.StatusOK, fmt.Sprintf("login with %v successfully", userInfo.Email))
}

func main() {

	json, err := os.ReadFile("credentials.json")
	if err != nil {
		panic(err)
	}

	config, err := google.ConfigFromJSON(json, googleoauth2.UserinfoEmailScope)
	if err != nil {
		panic(err)
	}

	authUrl := config.AuthCodeURL("test-state")
	fmt.Println(authUrl)

	h := newHandler(config)

	e := echo.New()
	e.GET("/login", h.login)
	e.GET("/callback", h.callback)
	e.Logger.Fatal(e.Start(":8080"))
}
