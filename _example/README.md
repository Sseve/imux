### example

- 接口测试
1. curl http://localhost:9990/ping  
   {"code":200,"message":"pong"}
2. curl 'http://localhost:9990/pong?foo=FOO&bar=BAR'  
   {"bar":"BAR","code":200,"foo":"FOO"}
---
1. curl http://localhost:9990/v1/foo/666 -H 'Authorization:abc'  
   {"code":200,"message":"Get foo id: 666"}
2. curl http://localhost:9990/v1/foo -H 'Authorization:abc' -d '{"foo": "FOO", "bar": "BAR"}'  
   {"code":200,"message":"test foo bind success"}
---
1. curl -X GET http://localhost:9990/v2/hello  
   {"code":200,"message":"Hello Get"}
2. curl -X POST http://localhost:9990/v2/hello  
   {"code":200,"message":"Hello Post"}
3. curl -X DELETE http://localhost:9990/v2/hello  
   {"code":200,"message":"Hello Delete"}
4. curl -X PUT http://localhost:9990/v2/hello  
   {"code":200,"message":"Hello Put"}
