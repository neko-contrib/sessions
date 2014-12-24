#sessions

Session middleware for [Neko](https://github.com/rocwong/neko)

## Usage

~~~ go
package main
import (
  "github.com/rocwong/neko"
  "github.com/neko-contrib/sessions"
)

func main() {
  m := neko.New()

  m.Use(neko.Sessions("sess_neko", sessions.NewCookieStore([]byte("secret123"))))

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

