 
 version=$(cat version.config)
 echo "Publishing version ${version}"
 echo go list -m github.com/akrck02/valhalla-core-sdk@${version} 

 GOPROXY=proxy.golang.org go list -m github.com/akrck02/valhalla-core-sdk@${version}