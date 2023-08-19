# Simple Golang in-memory cache with TTL

This package let you save get and delete values from in-memory cache. You can also add a TTL for each value

Installation:

```go get -u github.com/NIRumiantsev/golang-cache-ttl```

Usage:
```
func main() {
	c := New()
	c.Set("UserId", 1, time.Second)

	userId, err := c.Get("UserId")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userId) // Works correctly

	time.Sleep(time.Second * 2)

	_, err = c.Get("UserId")

	if err != nil {
		log.Fatal(err) // Error occurs
	}
}
```