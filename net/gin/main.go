package main

import (
	"encoding/json"
	"fmt"
	//"io"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Req struct {
	Id int64 `json:"id"`
}

type Res struct {
	Message string `json:"message"`
}

func main() {
	// 设置release或bug模式
	//gin.SetMode(gin.ReleaseMode)
	//gin.DisableConsoleColor()

	//logf, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(logf)
	//gin.DefaultErrorWriter = io.MultiWriter(logf)

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	/*engine.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp,
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage)
	}))*/

	setupRouter(engine)

	//err := engine.Run(":8080")
	s := &http.Server{
		Addr:           ":8080",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 3 seconds.")
	}
	log.Println("Server exiting")
}

var db = make(map[string]string)

func setupRouter(engine *gin.Engine) {
	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := engine.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar   [Basic Zm9vOmJhcg==]
		"manu": "123", // user:manu password:123  [Basic bWFudToxMjM=]
	}))
	{
		authorized.POST("admin", func(c *gin.Context) {
			user := c.MustGet(gin.AuthUserKey).(string)

			// Parse JSON
			var json struct {
				Value string `json:"value" binding:"required"`
				//curl -XPOST 'http://127.0.0.1:8080/admin' -H 'Authorization: Basic Zm9vOmJhcg==' -H 'Content-type: application/json' -d '{"value":"xioamisd"}'
			}

			if c.Bind(&json) == nil {
				db[user] = json.Value
				c.JSON(http.StatusOK, gin.H{"status": "ok"})
			} else {
				c.String(http.StatusBadRequest, "err msg\n")
			}
		})
	}
	// Get user value
	engine.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	v1 := engine.Group("/home")
	{
		v1.GET("/user/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "hello %s\n", name)
		})
		v1.GET("/user/:name/*action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			c.String(http.StatusOK, "%s is %s\n", name, action[1:])
		})
		v1.POST("/msg", func(c *gin.Context) {
			msg := c.PostForm("msg")
			c.JSON(http.StatusOK, gin.H{
				"code":  http.StatusOK,
				"error": msg,
			})
		})
		v1.POST("/data", func(c *gin.Context) {
			data, err := c.GetRawData()
			if err != nil {
				c.JSON(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
				return
			}

			req := Req{}
			err = json.Unmarshal(data, &req)
			if err != nil {
				panic(fmt.Sprintf("err: %s", err.Error()))
				//c.AbortWithError(http.StatusBadRequest, err)
				//return
			}

			c.JSON(http.StatusOK, &Res{
				Message: fmt.Sprintf("ok, receive [%d]!", req.Id),
			})
		})
		v1.GET("/redirect", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
		})
		v1.GET("/cookie", func(c *gin.Context) {
			cookie, err := c.Cookie("gin_cookie")
			if err != nil {
				cookie = "NotSet"
				c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
			}
			fmt.Printf("Cookie: %s\n", cookie)
			//curl -XGET 'http://127.0.0.1:8080/home/cookie' -H 'Cookie: gin_cookie=x'
		})
		v1.StaticFS("/dir", http.Dir("."))
		v1.Static("/File", ".")
		v1.StaticFile("/image", "./test.jpeg")
	}
}
