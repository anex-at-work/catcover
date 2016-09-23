# Just concatenation tool for coverage files

## Before...

I know, you can use a bash like this:
```
echo 'mode: atomic' > coverage.txt && go list ./... | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.txt' && rm coverage.tmp
```

(from [here](https://github.com/pierrre/gotestcover)).

## But:

If you use [ginkgo](https://onsi.github.io/ginkgo/), then you can use this utility do to the next steps:
* run `ginkgo -cover -r`
* concatenate all *\*.coverprofile* into *all_coverage.coverprofile*
* remove all excess lines with `mode: atomic`
* make *coverage.html* file with `go tool cover` utility

### Just compile and run `./coverage` in the right place

And, in the end, you can open *coverage.html* in your browser and reach 100% coverage 8)
