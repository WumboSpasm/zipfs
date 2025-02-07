package zipfs_test

import (
	"net/http"

	"github.com/FlashpointProject/zipfs"
)

func Example() error {
	fs, err := zipfs.New("testdata/testdata.zip")
	if err != nil {
		return err
	}

	extensions := []string{"html", "htm"}
	return http.ListenAndServe(":8080", zipfs.FileServer(fs, "test/base/api/", "", true, extensions, nil))
}
