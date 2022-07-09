// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// )

// type Coffee struct{
// 	Title string `json:"title"`
// 	Description string `json:"description"`
// }
// type LogTransport struct {
//     Transport http.RoundTripper
// }

// func (lt *LogTransport) transport() http.RoundTripper {
//     if lt.Transport == nil {
//         return http.DefaultTransport
//     }
//     return lt.Transport
// }

// func (lt *LogTransport) CancelRequest(req *http.Request) {
//     type canceler interface {
//         CancelRequest(*http.Request)
//     }
//     if cr, ok := lt.transport().(canceler); ok {
//         cr.CancelRequest(req)
//     }
// }

// func (lt *LogTransport) RoundTrip(req *http.Request) (*http.Response, error) {
//     res, err := lt.transport().RoundTrip(req)

//     log.Printf("%d\t%s\t%s\n", res.StatusCode, req.Method, req.URL.String())

//     return res, err
// }

// func main(){
// 	client := &http.Client{
//         Transport: &LogTransport{},
//     }

// 	res, err := client.Get("https://api.sampleapis.com/coffee/hot")
// 	if err != nil {
// 		// handle error 今回は実装しない
// 	}
// 	// コネクションをcloseする
// 	defer res.Body.Close()

// 	// httpリクエストした返り値の型は *http.Response
// 	fmt.Printf("resの型は、%T\n", res)

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		// handle error 今回は実装しない
// 	}
// 	// io.ReadAllの帰り値はbyte型の配列（json）
// 	fmt.Printf("bodyの型は、%T\n", body)

// 	// 扱いやすいようにjsonを構造体に変換する
// 	var coffee []Coffee
// 	if err := json.Unmarshal(body, &coffee); err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Printf("coffeeの型は、%T\n", coffee)

// 	// 配列をループで処理する
// 	for _, value := range coffee{
// 		fmt.Printf("%+v\n", value)
// 	}
// }