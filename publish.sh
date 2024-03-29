 
 version=$(cat version.txt)
 echo "Publishing version ${version}"
 GOPROXY=proxy.golang.org go list -m github.com/akrck02/valhalla-core-sdk@${version}