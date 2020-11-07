package main

type data struct {
	name string
}

func main() {
	m := map[string]*data{"x": {"one"}}

	// ???  SIGSEGV(segmentation violation) 段错误
	m["y"].name = "what?"
}
