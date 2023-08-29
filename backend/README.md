run:
`go run main.go`


leslie, for live reload, you are using air, not gin.

air still seems to have a hiccup sometimes due to what seems to be a race condition.

users claim to have fixed this hiccup with a rebuild delay in .air.toml, but that did not work for you.

run with live reload:

`air`



take breaks drink water