package main

import "fmt"

type Vertex struct {
	Lat, long float64
}

func main() {

	m := make(map[string]Vertex)

	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m)
	/* map[Bell Labs:{40.68433 -74.39967}] */


	m2 := map[string]Vertex{
		"Bell Labs": Vertex{40.68433, -74.39967},
		"Google": Vertex{37.42202, -122.08408},
	}
	fmt.Println(m2)
	/* map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}] */

	delete(m2, "Bell Labs")
	fmt.Println(m2)
	/* map[Google:{37.42202 -122.08408}] */

	/* _ は宣言するけど後で使わない変数。m2["Bell Labs"]は第1戻り値に{0,0},第2戻り値にfalseを返す */
	_, ok := m2["Bell Labs"]
	fmt.Println("m2 has key 'Bell Labs':", ok)
	/* m2 has key 'Bell Labs': false */
}