package main

import (
  "fmt"
  "gee"
  "net/http"
  "time"
)
type student struct {
  Name string
  Age  int8
}
func formatAsDate(t time.Time) string {
  year, month, day := t.Date()
  return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

//封装一个context 简化操作接口调用 context承载当前信息 百宝箱一样供调用
func main() {
  r := gee.New()
  //处理由此导致的越界问题
  r.GET("/panic", func(c *gee.Context) {
    names := []string{"geektutu"}
    c.String(http.StatusOK, names[0])
  })

  r.GET("/", func(c *gee.Context) {
    c.String(http.StatusOK, "Hello Geektutu\n")
  })

  r.Run(":9999")

  //实现了映射表 提供了用户注册静态路由
  //r := gee.New()
  ////实现路由分组
  ////实现用户自定义功能
  //r.Use(gee.Logger()) // global midlleware
  //
  ////实现静态分离
  //r.SetFuncMap(template.FuncMap{
  //  "formatAsDate": formatAsDate,
  //})
  //r.LoadHTMLGlob("templates/*")
  //r.Static("/assets", "./static")
  //
  //stu1 := &student{Name: "Geektutu", Age: 20}
  //stu2 := &student{Name: "Jack", Age: 22}
  //r.GET("/", func(c *gee.Context) {
  //  c.HTML(http.StatusOK, "css.tmpl", nil)
  //})
  //
  //r.GET("/students", func(c *gee.Context) {
  //  c.HTML(http.StatusOK, "arr.tmpl", gee.H{
  //    "title":  "gee",
  //    "stuArr": [2]*student{stu1, stu2},
  //  })
  //})
  //
  //r.GET("/date", func(c *gee.Context) {
  //  c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
  //    "title": "gee",
  //    "now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
  //  })
  //})
  //
  //r.Run(":9999")


  //r.GET("/", func(c *gee.Context) {
  //  c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
  //})

  //v1 := r.Group("/v1")
  //{
  //  v1.GET("/", func(c *gee.Context) {
  //    c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
  //  })
  //
  //  v1.GET("/hello", func(c *gee.Context) {
  //    // expect /hello?name=geektutu
  //    c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
  //  })
  //}
  //v2 := r.Group("/v2")
  //{
  //  v2.GET("/hello/:name", func(c *gee.Context) {
  //    // expect /hello/geektutu
  //    c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
  //  })
  //  v2.POST("/login", func(c *gee.Context) {
  //    c.JSON(http.StatusOK, gee.H{
  //      "username": c.PostForm("username"),
  //      "password": c.PostForm("password"),
  //    })
  //  })

  //}

  //前后端分离所产生的问题资源利用，并发，数据库
  //但前后分离的一大问题在于，页面是在客户端渲染的，比如浏览器，这对于爬虫并不友好。Google 爬虫已经能够爬取渲染后的网页，但是短期内爬取服务端直接渲染的 HTML 页面仍是主流。
  //r.Run(":9999")
}
