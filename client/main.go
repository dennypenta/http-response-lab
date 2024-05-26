package main

import (
	model "client/proto"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime/pprof"

	"google.golang.org/protobuf/proto"
)

func decode(r io.Reader) (model.AutoList, error) {
	var m model.AutoList
	data, err := io.ReadAll(r)
	if err != nil {
		return model.AutoList{}, err
	}
	if err := proto.Unmarshal(data, &m); err != nil {
		return model.AutoList{}, err
	}

	return m, nil
}

func main() {
	res, err := http.DefaultClient.Get("http://localhost:8000")
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	models, err := decode(res.Body)
	if err != nil {
		panic(err.Error())
	}

	if len(models.Autos) == 0 {
		panic("models must not be empty")
	}
	fmt.Println("DONE")
	fmt.Println(len(models.Autos))

	f, _ := os.Create("mem.pprof")
	pprof.WriteHeapProfile(f)
	f.Close()
}
