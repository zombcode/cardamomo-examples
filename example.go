package main

import (
  "../../cardamomo"
	"fmt"
	"time"
)

type Box struct {
  Size BoxSize
  Color  string
  Open   bool
}

type BoxSize struct {
  Width  int
  Height int
}

func main() {

	// HTTP

  c := cardamomo.Instance("8000")
  c.SetDevDebugMode(true)

  c.Get("/", func(req cardamomo.Request, res cardamomo.Response) {
		res.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    res.Writer.Header().Set("Access-Control-Allow-Headers", "Accept,Authorization,Cache-Control,Content-Type,DNT,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Mx-ReqToken,X-Requested-With");

    res.Send("Hello world!");
  })

  c.Get("/routeget1", func(req cardamomo.Request, res cardamomo.Response) {
    res.Send("Hello route get 1!");
  })

	c.Get("/routeget2/:param1/and/:param2", func(req cardamomo.Request, res cardamomo.Response) {
    res.Send("Hello route get 1 with param1 = " + req.GetParam("param1", "") + " and param2 = " + req.GetParam("param2", "") + "!");
  })

  c.Get("/routeget3/:param/{{a([a-zA-Z0-9]+)b$}}", func(req cardamomo.Request, res cardamomo.Response) {
    res.Send("Hello! This route uses REGEX! Only URL that use parameters between 'a' and 'b'");
  })

  c.Get("/rendertest", func(req cardamomo.Request, res cardamomo.Response) {
    res.Render("socket.html", cardamomo.JSONC{
      "data": cardamomo.JSONC{
        "title": "Hello world!",
        "desc": "Lorem ipsum!",
      },
    });
  })

  c.Get("/routejson", func(req cardamomo.Request, res cardamomo.Response) {
    boxsize := BoxSize {
      Width:  10,
    	Height: 20,
    }

    box := Box {
    	Size: boxsize,
    	Color:  "blue",
    	Open:   false,
    }

    res.SendJSON(box)
  })

  c.Post("/routepost1", func(req cardamomo.Request, res cardamomo.Response) {
    res.Send("Hello route post 1!");
  })

	c.Post("/fileupload", func(req cardamomo.Request, res cardamomo.Response) {
		req.MoveUploadedFile("foo", "./tmp/filetest.jpg");

		res.Send("Hello file upload!");
  })

  c.Base("/base1", func(router *cardamomo.Router) {
    router.Get("/routeget1", func(req cardamomo.Request, res cardamomo.Response) {
      res.Send("Hello route base1/routeget1!");
    })
    router.Post("/routepost1", func(req cardamomo.Request, res cardamomo.Response) {
      res.Send("Hello route base1/routepost1!");
    })
		router.Base("/base2", func(router *cardamomo.Router) {
			router.Get("/routeget1", func(req cardamomo.Request, res cardamomo.Response) {
	      res.Send("Hello route base1/base2/routeget1!");
	    })
		});
  })

	// Cookies

	c.Get("/setcookie/:key/:value", func(req cardamomo.Request, res cardamomo.Response) {
		key := req.GetParam("key", "")
		value := req.GetParam("value", "")

		expire := time.Now().AddDate(0, 0, 1) // Expires in one day!
		req.SetCookie(key, value, "/", "localhost", expire, 86400, false, false) // key, value, path, domain, expiration, max-age, httponly, secure

    res.Send("Added cookie \"" + key + "\"=\"" + value + "\"");
  })

	c.Get("/getcookie/:key", func(req cardamomo.Request, res cardamomo.Response) {
		key := req.GetParam("key", "")

		cookie := req.GetCookie(key, "empty cookie!"); // key, defaultValue

    res.Send("The value for cookie \"" + key + "\" is \"" + cookie + "\"");
  })

	c.Get("/deletecookie/:key", func(req cardamomo.Request, res cardamomo.Response) {
		key := req.GetParam("key", "")

		req.DeleteCookie(key, "/", "localhost"); // key, defaultValue

    res.Send("Deleted cookie \"" + key + "\"");
  })

	// Error handler

	c.SetErrorHandler(func (code string, req cardamomo.Request, res cardamomo.Response) {
		fmt.Printf("\nError: %s\n", code)

		if( code == "404" ) {
			res.Send("Error 404!");
		}
	})

	// Sockets

	socket := c.OpenSocket()
  socket.Cluster(cardamomo.SocketClusterParams{ // You can use this lines for cluster testing
    Hosts: []cardamomo.SocketClusterHost{ // Write a list of ALL servers included in the cluster
      cardamomo.SocketClusterHost{
        Host: "192.168.0.214", // Use the server 1 IP
        Port: "8000", // Use the server 1 PORT
        Master: true, // Only ONE server can be MASTER
      },
      cardamomo.SocketClusterHost{
        Host: "192.168.0.214", // Use the server 2 IP
        Port: "8001", // Use the server 2 PORT
        Master: false, // Only ONE server can be MASTER
      },
    },
    Password: "examplepass",
  })
  socket.SendClient("as12df34gh56", "testing", cardamomo.JSONC{"foo":"bar"}); // This is for testing communication between sockets, you can use a real client ID for that

	socket.OnSocketBase("/base1", func(client *cardamomo.SocketClient) {
		fmt.Printf("\n\nBase 1 new client!\n\n")

		client.OnSocketAction("onDisconnect", func(sparams map[string]interface{}) {
			fmt.Printf("\n\nDisconnect!\n\n")
		})

		client.OnSocketAction("action1", func(sparams map[string]interface{}) {
			fmt.Printf("\n\nAction 1!\n\n")

			fmt.Printf("\n\nParam: %s\n\n", sparams["param_1"])
			fmt.Printf("\n\nParam: %s\n\n", sparams["param_2"].(map[string]interface{})["inner_1"])
			fmt.Printf("\n\nParam: %d\n\n", int(sparams["param_2"].(map[string]interface{})["inner_2"].([]interface{})[1].(float64)))

			client.Send("action1", sparams)
		})
	})

	c.Get("/socket/broadcast/:message", func(req cardamomo.Request, res cardamomo.Response) {
		params := make(map[string]interface{})
		params["message"] = req.GetParam("message", "Default message")

		socket.Send("broadcast", params);
		socket.SendBase("/base1", "broadcast", params);
    res.Send("Broadcast \"" + req.GetParam("message", "Default message") + "\" sended!");
  })

	c.Get("/socket/chat/:id/:message", func(req cardamomo.Request, res cardamomo.Response) {
		params := make(map[string]interface{})
		params["message"] = req.GetParam("message", "Default message")

		socket.SendClient(req.GetParam("id", ""), "chat", params);

    res.Send("Message \"" + req.GetParam("message", "Default message") + "\" sended to client \"" + req.GetParam("id", "Empty client!") + "\"!");
  })

  // Scripts
  scripts := cardamomo.NewScripts() // You need this line in order to initialize scripts objects

  scripts.AddScript(func() {
    fmt.Printf("\n10 seconds! \n");
  }, 10)

  scriptDatetime, err := time.Parse("3 04 00 PM", "9 46 00 PM") // This execute task at 9:46 every days
  fmt.Printf("\n\nTime: %s\n\n", scriptDatetime);

  if err == nil {
    scripts.AddScriptAtTime(func() {
      fmt.Printf("\nConcrete time! \n");
    }, scriptDatetime)
  }

  c.Run()
}
