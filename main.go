package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Post struct {
	Id     int    `json:"id"`
	Userid int    `json:"userId"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body"`
}

func readJsonFile() io.Reader {
	f, err := os.Open("posts.json")

	if err != nil {
		log.Fatal(err)
	}

	return f
}

func main() {
	// post := Post{
	// 	Id:     333,
	// 	Userid: 35,
	// 	Title:  "",
	// 	Body:   "shshshshshhshshhshshshhshshshhshshshshs",
	// }

	// res, err := json.Marshal(post)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(res))

	//////////////////////

	// jsonData := `{
	// 	"userId": 1,
	// 	"id": 1,
	// 	"title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
	// 	"body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
	//   }`

	// var data map[string]interface{}

	// err := json.Unmarshal([]byte(jsonData), &data)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("map: %#v\n", data)

	// var post Post

	// err := json.Unmarshal([]byte(jsonData), &post)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%#v\n", post)

	// var out bytes.Buffer

	// json.Indent(&out, []byte(jsonData), "", " ")

	// out.WriteTo(os.Stdout)

	//////////////////////

	dec := json.NewDecoder(readJsonFile())
	enc := json.NewEncoder(os.Stdout)

	for {
		var post Post

		if err := dec.Decode(&post); err != nil {
			log.Fatal(err)
			return
		}

		if err := enc.Encode(&post); err != nil {
			log.Fatal(err)
			return
		}
	}

}
