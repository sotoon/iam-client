r=$(go test -bench=.  ./... -run=^$ -benchmem)
echo "$r"
echo
echo =============[ Final Result ]=============
echo "$r" | grep ns/op
