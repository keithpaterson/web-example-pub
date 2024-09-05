package request

import (
	"net/http"
	"strconv"
	"strings"
)

var (
	methodsWithId = [...]string{http.MethodGet, http.MethodGet, http.MethodDelete, http.MethodPut, http.MethodPatch}
)

// try to read the ID value from the request.
//
//	If it is present return (id, true)
//	Otherwise return (0, false)
func GetIdValue(r *http.Request, name string) (int, bool) {
	for _, m := range methodsWithId {
		if r.Method == m {
			return ExtractId(r, name)
		}
	}
	return 0, false
}

func ExtractId(r *http.Request, name string) (int, bool) {
	uriSegments := strings.Split(r.URL.Path, "/")

	// Expect the path to be soemthing like
	//   'http://any/path/prefix/resource_name/id/....'
	// we will find the resource-name, assume the next element is an id
	index := -1
	for i, seg := range uriSegments {
		if seg == name {
			index = i + 1
			break
		}
	}

	if index < 0 || index >= len(uriSegments) || uriSegments[index] == "" {
		// No ID segment
		return 0, false
	}

	// else convert to int
	if id, err := strconv.Atoi(uriSegments[index]); err == nil {
		return id, true
	}

	// string conversion error, just return false here
	// we either found:
	//   a non-int id (currently invalid), or
	//   some resource property was specifed that's not really an ID (shouldn't happen).
	// either way we don't have an ID to return
	//
	// TODO: add an error to the log (maybe a warning actually)
	return 0, false
}
