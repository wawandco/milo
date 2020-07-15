module wawandco/milo

go 1.14

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/stretchr/testify v1.5.1
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/PuerkitoBio/goquery => github.com/wawandco/goquery v1.5.2-0.20200714195702-e514b22bcb73

replace golang.org/x/net => github.com/wawandco/net v0.0.0-20200714194128-a771946c73f6
