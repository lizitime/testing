package main

import (
	"fmt"
	"os"
	"encoding/json"

	//"flag"
	//"log"
	//"net/url"
	//"os/signal"

	//"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("Hello, world. 你好，世界！")

	var nowUnix int64 = 1504856025
	now := time.Unix(nowUnix, 0)
	fmt.Println(now)
	fmt.Println()
	
	now = time.Now()
	fmt.Println(now)

	var i int = 0
	if i > 0 {
		i = 2
		fmt.Println("1234")
	} else {
		i = 3
		fmt.Println("5678")
	}

	//testJson()

	//  fmt.Println(now.Unix())
	//  fmt.Println(now.UnixNano())
}

func testJson() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	// b, err := json.Marshal(group)
	b, err := json.MarshalIndent(group, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	fmt.Println("")
}

//var addr = flag.String("addr", "localhost:8080", "http service address")
//
//func main() {
//	fmt.Println("addr is :", *addr)
//	flag.Parse()
//	log.SetFlags(0)
//
//	interrupt := make(chan os.Signal, 1)
//	signal.Notify(interrupt, os.Interrupt)
//
//	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
//	log.Printf("connecting to %s", u.String())
//
//	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
//	if err != nil {
//		log.Fatal("dial:", err)
//	}
//	defer c.Close()
//
//	done := make(chan struct{})
//
//	go func() {
//		defer c.Close()
//		defer close(done)
//		for {
//			_, message, err := c.ReadMessage()
//			if err != nil {
//				log.Println("read:", err)
//				return
//			}
//			log.Printf("recv: %s", message)
//		}
//	}()
//
//	ticker := time.NewTicker(time.Second)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case t := <-ticker.C:
//			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
//			if err != nil {
//				log.Println("write:", err)
//				return
//			}
//		case <-interrupt:
//			log.Println("interrupt")
//			// To cleanly close a connection, a client should send a close
//			// frame and wait for the server to close the connection.
//			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
//			if err != nil {
//				log.Println("write close:", err)
//				return
//			}
//			select {
//			case <-done:
//			case <-time.After(time.Second):
//			}
//			c.Close()
//			return
//		}
//	}
//}
