./grpc_cli  ls 192.168.1.183:50051 --l

```


./grpc_cli  ls 192.168.1.183:50051 --l
filename: grpc_reflection/v1alpha/reflection.proto
package: grpc.reflection.v1alpha;
service ServerReflection {
  rpc ServerReflectionInfo(stream grpc.reflection.v1alpha.ServerReflectionRequest) returns (stream grpc.reflection.v1alpha.ServerReflectionResponse) {}
}

filename: routeguide.proto
package: routeguide;
service RouteGuide {
  rpc GetFeature(routeguide.Point) returns (routeguide.Feature) {}
  rpc ListFeatures(routeguide.Rectangle) returns (stream routeguide.Feature) {}
  rpc RecordRoute(stream routeguide.Point) returns (routeguide.RouteSummary) {}
  rpc RouteChat(stream routeguide.RouteNote) returns (stream routeguide.RouteNote) {}
}



./grpc_cli  ls 192.168.1.183:50051 routeguide.RouteGuide.GetFeature --l
  rpc GetFeature(routeguide.Point) returns (routeguide.Feature) {}
```


```
./grpc_cli  call  192.168.1.183:50051 routeguide.RouteGuide.GetFeature  "latitude:409146138,longitude:-746188906"
connecting to 192.168.1.183:50051
name: "Berkshire Valley Management Area Trail, Jefferson, NJ, USA"
location {
  latitude: 409146138
  longitude: -746188906
}

Rpc succeeded with OK status
```