package main

import (
	"web"
	"time"
	"strconv"
	)

func postForId(id int64) BlogPost {
    post, _ := Db.GetPost(id)
    return post
}

func postsForDay(date *time.Time) []BlogPost {
	posts, _ := Db.GetPostsForDate(*date)
	return posts
}

func postsForMonth(date *time.Time) []BlogPost {
    posts, _ := Db.GetPostsForMonth(*date)
    return posts
}

func index(ctx *web.Context) string {
   css := ctx.Params["css"]
	 if len(css) > 0 {
	 	SetCSS(ctx, css)
	 }
		
		posts := postsForMonth(time.LocalTime())//Db.GetLastNPosts(10)
    s := RenderHeader(ctx)
    s += RenderPosts(&posts)
    s += RenderFooter()
	return s
}

func post(ctx *web.Context) string {
    id_s := ctx.Params["id"]
    id, _ := strconv.Atoi64(id_s)
    
    post := postForId(id)
        
    s := RenderHeader(ctx)
    s += "<ul>"
    s += RenderPost(&post, true)
    s += "</ul>"
    s += RenderFooter()

    return s
}
