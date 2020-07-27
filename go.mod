module wawandco/milo

go 1.14

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/stretchr/testify v1.5.1
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/PuerkitoBio/goquery => github.com/wawandco/goquery v1.5.2-0.20200727152614-d4cd420698c5

replace golang.org/x/net => github.com/wawandco/net v0.0.0-20200727155331-94802980ad79
