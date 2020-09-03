package gee

import (
  "net/http"
  "strings"
)

//静态路由到动态路由 使用前缀树来实现*匹配
type router struct {
  roots    map[string]*node
  handlers map[string]HandlerFunc
}


func (r *router) handle(c *Context) {
 n, params := r.getRoute(c.Method, c.Path)
 if n != nil {
   key := c.Method + "-" + n.pattern
   c.Params = params
   c.handlers = append(c.handlers, r.handlers[key])
 } else {
   c.handlers = append(c.handlers, func(c *Context) {
     c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
   })
 }
 c.Next()
}


func newRouter() *router {
return &router{
roots:    make(map[string]*node),
handlers: make(map[string]HandlerFunc),
}
}


func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
  parts := parsePattern(pattern)

  key := method + "-" + pattern
  _, ok := r.roots[method]
  if !ok {
    r.roots[method] = &node{}
  }
  r.roots[method].insert(pattern, parts, 0)
  r.handlers[key] = handler
}

// Only one * is allowed
func parsePattern(pattern string) []string {
  vs := strings.Split(pattern, "/")

  parts := make([]string, 0)
  for _, item := range vs {
    if item != "" {
      parts = append(parts, item)
      if item[0] == '*' {
        break
      }
    }
  }
  return parts
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
  searchParts := parsePattern(path)
  params := make(map[string]string)
  root, ok := r.roots[method]

  if !ok {
    return nil, nil
  }

  n := root.search(searchParts, 0)

  if n != nil {
    parts := parsePattern(n.pattern)
    for index, part := range parts {
      if part[0] == ':' {
        params[part[1:]] = searchParts[index]
      }
      if part[0] == '*' && len(part) > 1 {
        params[part[1:]] = strings.Join(searchParts[index:], "/")
        break
      }
    }
    return n, params
  }

  return nil, nil
}



func (n *node) insert(pattern string, parts []string, height int) {
  if len(parts) == height {
    n.pattern = pattern
    return
  }

  part := parts[height]
  child := n.matchChild(part)
  if child == nil {
    child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
    n.children = append(n.children, child)
  }
  child.insert(pattern, parts, height+1)
}
func (n *node) search(parts []string, height int) *node {
  if len(parts) == height || strings.HasPrefix(n.part, "*") {
    if n.pattern == "" {
      return nil
    }
    return n
  }

  part := parts[height]
  children := n.matchChildren(part)

  for _, child := range children {
    result := child.search(parts, height+1)
    if result != nil {
      return result
    }
  }

  return nil
}

