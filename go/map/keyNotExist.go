package main

type data struct {
	name string
}

func main() {
	m := map[string]*data{"x": {"one"}}
	m["y"].name = "what?" //???  SIGSEGV(segmentation violation) 段错误
}
