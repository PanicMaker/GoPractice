package geecache

import (
	"fmt"
	"io"
	"log"
	"net/http/httptest"
	"testing"
)

func TestHTTPPool_ServeHTTP(t *testing.T) {
	var db = map[string]string{
		"Tom":  "630",
		"Jack": "589",
		"Sam":  "567",
	}

	NewGroup("scores", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	req := httptest.NewRequest("GET", "http://localhost:9999/_geecache/scores/Tom", nil)
	w := httptest.NewRecorder()

	httpPool := NewHTTPPool(":9999")
	httpPool.ServeHTTP(w, req)
	bytes, _ := io.ReadAll(w.Result().Body)

	if string(bytes) != "630" {
		t.Fatal("630, but got", string(bytes))
	}
}
