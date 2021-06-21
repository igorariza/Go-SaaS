package main

import (
	"log"
	"net/http"

	"github.com/igorariza/Go-SaaS/tree/main/controllers"
)

func main() {
	api := &controllers.API{}
	if err := http.ListenAndServe(":8080", api); err != nil {
		log.Println(err)
	}
}

// import (
// 	"net/http"
// )
// func shiftPath(p string) (head, tail string) {
// 	p = path.Clean("/" + p)
// 	i := strings.Index(p[1:], "/") + 1
// 	if i <= 0 {
// 	return p[1:], "/"
// 	}
// 	return p[1:i], p[i:]
//    }
//    func (h *App) log(next http.Handler) http.Handler {
//    return http.HandlerFunc(func(w http.ResponseWriter, r
//    *http.Request) {
// 	start := time.Now()
// 	next.ServeHTTP(w, r)
// 	log.Println(time.Since(start), r.Method, r.URL.Path)
// 	})
//    }
//    type key int
//    const (
// 	ctxTestKey key = 1
// 	ctxUserID = 2
//    )
//    func (h *App) test(next http.Handler) http.Handler {
//    return http.HandlerFunc(func(w http.ResponseWriter, r
//    *http.Request) {
// 	ctx := r.Context()
// 	ctx = context.WithValue(ctx, ctxTestKey, "hello world")
// 	next.ServeHTTP(w, r.WithContext(ctx))
// 	})
//    }

// func (u User) Detail(w http.ResponseWriter, r *http.Request) {
// 		ctx := r.Context()
// 		v := ctx.Value(ctxTestKey)
// 		id := ctx.Value(ctxUserID)
// 		w.Write([]byte(fmt.Sprintf("value of context is %s for user id %d", v, id)))
// }

// type User struct{}
// func (u User) ServeHTTP(w http.ResponseWriter, r *http.Request){
// 		var head string
// 		head, r.URL.Path = shiftPath(r.URL.Path)
// 		if head == "profile" {
// 		u.Profile(w, r)
// 		return
// 		} else if head == "detail" {
// 		head, _ := shiftPath(r.URL.Path)
// 		i, err := strconv.Atoi(head)
// 		if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 		}
// 		ctx := r.Context()
// 		ctx = context.WithValue(ctx, ctxUserID, i)
// 		u.Detail(w, r.WithContext(ctx))
// 		return
// 		}
// 		http.Error(w, "not found", http.StatusNotFound)
// }

// var i interface{} = User{}
// 	next = &Route{
// 	Logger: true,
// 	Tester: true,
// 	Handler: i.(http.Handler),
// 	}
// func (h *App) ServeHTTP(w http.ResponseWriter,
// 						r *http.Request) {
//  var next *Route
//  var head string
//  head, r.URL.Path = shiftPath(r.URL.Path)
// 	if len(head) == 0 {
// 	next = &Route{
// 		Logger: true,
// 		Handler: http.HandlerFunc(func(w http.ResponseWriter,
// 									   r *http.Request) {
// 									   w.Write([]byte("home page"))
// 								}),
// 		}
// 	} else if head == "user" {
// 		var i interface{} = User{}
// 			next = &Route{
// 			Logger: true,
// 			Tester: true,
// 			Handler: i.(http.Handler),
// 		}
// 	} else {
//  http.Error(w, "not found", http.StatusNotFound)
//  return
//  }
//  if next.Logger {
//  next.Handler = h.log(next.Handler)
//  }
//  if next.Tester {
//  next.Handler = h.test(next.Handler)
//  }
//  next.Handler.ServeHTTP(w, r)
// }
// type Route struct {
// 	Logger bool
// 	Tester bool
// 	Handler http.Handler
//    }

// type App struct {
// 	User *Route
// }
// func main() {
// 	app := &App{
// 		User: &Route{
// 			Logger: true,
// 			Tester: true,
// 		},
// 		Billing: &Route{
// 			Logger: true,
// 			Tester: false,
// 		},
// 	}
// 	http.ListenAndServe(":8080", app)
// }

// // func main() {
// // 	http.HandleFunc("/time", getTime)
// // 	http.Handle("/iseven", isEven(http.HandlerFunc(getTime)))
// // }

// // func getTime(w http.ResponseWriter, r *http.Request) {
// // 	if r.Method != "GET" {
// //    http.Error(w, "only GET method are allowed",
// //    http.StatusMethodNotAllowed)
// // 	return
// // 	}

// // 	w.Write([]byte(time.Now().String()))
// // }

// // func isEven(h http.Handler) http.Handler {
// // return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// //  if time.Now().Second()%2 == 0 {
// //  		h.ServeHTTP(w, r)
// //  	}
// // 	http.Error(w, "current time second is odd, cannot serve the response",
// // 			   http.StatusInternalServerError)
// //  })
// // }
