package main

func main() {
	server := NewServer("0.0.0.0", 123)
	server.Start()
}
