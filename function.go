// function.go
package function

import (
	"bufio"
	"net/http"
	"strconv"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	scanner := bufio.NewScanner(r.Body)
	fuel := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		if mass, err := strconv.ParseInt(line, 10, 64); err != nil {
			http.Error(w, "Invalid input:"+line, http.StatusNotAcceptable)
			return
		} else {
			fuel += calcFuel(mass)
		}
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(strconv.FormatInt(fuel, 10)))
}

func calcFuel(mass int64) int64 {
	fuel := mass / 3 - 2
	fuel2 := fuel / 3 - 2
	for fuel2 > 0 {
		fuel += fuel2
		fuel2 = fuel2 / 3 - 2
	}
	return fuel
}