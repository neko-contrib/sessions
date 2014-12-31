#sessions
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/neko-contrib/sessions)
[![GoCover](http://gocover.io/_badge/github.com/neko-contrib/sessions)](http://gocover.io/github.com/neko-contrib/sessions)

[Neko](https://github.com/rocwong/neko) middleware/handler that provides a Session service.

## Usage

~~~ go
package main
import (
  "github.com/rocwong/neko"
  "github.com/neko-contrib/sessions"
)

func main() {
  m := neko.New()

  m.Use(sessions.Sessions("sess_neko", sessions.NewCookieStore([]byte("secret123"))))

  m.GET("/", func (ctx *neko.Context) {
    ctx.Session.Set("myvalue", "Session Save")
    ctx.Text("Session Save")
  })
  m.GET("/get", func (ctx *neko.Context) {
    ctx.Text(ctx.Session.Get("myvalue"))
  })

  m.GET("/flash", func (ctx *neko.Context) {
    ctx.Session.AddFlash("Flash Session")
    ctx.Text("Flash Save")
  })
  m.GET("/getflash", func (ctx *neko.Context) {
    ctx.Text(ctx.Session.Flashes())
  })

  m.Run(":3000")
}

~~~

