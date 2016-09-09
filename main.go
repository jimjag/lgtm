package main

import (
	"github.com/AusDTO/lgtm/router"
	"github.com/AusDTO/lgtm/router/middleware"
	"net/http"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/ianschenck/envflag"
	_ "github.com/joho/godotenv/autoload"
)

var (
	env_port = envflag.Int("PORT", 8000, "port envvar")
	cert     = envflag.String("SERVER_CERT", "", "")
	key      = envflag.String("SERVER_KEY", "", "")

	debug = envflag.Bool("DEBUG", false, "")
)

func main() {
	envflag.Parse()

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

	handler := router.Load(
		ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true),
		middleware.Version,
		middleware.Store(),
		middleware.Remote(),
		middleware.Cache(),
	)

	if *cert != "" {
		logrus.Fatal(
			http.ListenAndServeTLS(":"+strconv.Itoa(*env_port), *cert, *key, handler),
		)
	} else {
		logrus.Fatal(
			http.ListenAndServe(":"+strconv.Itoa(*env_port), handler),
		)
	}
}
