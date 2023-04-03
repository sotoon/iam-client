r=$(go test -bench=.  ./... -run=^$ -benchmem)
echo "$r"
echo "$r" | grep ns/op
