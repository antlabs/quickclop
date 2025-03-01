package main

import (
	"fmt"
	"net/url"
)

type Options struct {
	Name    string   `clop:"-n,--name" usage:"名称"`
	Age     int      `clop:"-a,--age" usage:"年龄"`
	Website url.URL  `clop:"-w,--website" usage:"网站" default:"https://github.com/antlabs/quickclop"`
	ApiURL  *url.URL `clop:"--api" usage:"API地址" env:"API_URL"`
}

func main() {
	opts := &Options{}
	if err := opts.Parse(nil); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Name: %s, Age: %d\n", opts.Name, opts.Age)
	fmt.Printf("Website: %s\n", opts.Website.String())
	if opts.ApiURL != nil {
		fmt.Printf("API URL: %s\n", opts.ApiURL.String())
	}
}
