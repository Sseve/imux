### example

- 接口测试
1. curl http://localhost:9990/ping  
   {"code":200,"message":"pong"}
2. curl 'http://localhost:9990/pong?foo=FOO&bar=BAR'  
   {"bar":"BAR","code":200,"foo":"FOO"}
---
1. curl http://localhost:9990/api/foo/666 -H 'Authorization:abc'  
   {"code":200,"message":"Get foo id: 666"}
2. curl http://localhost:9990/api/foo -H 'Authorization:abc' -d '{"foo": "FOO", "bar": "BAR"}'  
   {"code":200,"message":"test foo bind success"}
---
1. curl -XGET http://localhost:9990/v1/hello  
   {"code":200,"message":"Hello Get"}
2. curl -XPOST http://localhost:9990/v1/hello  
   {"code":200,"message":"Hello Post"}
3. curl -XDELETE http://localhost:9990/v1/hello  
   {"code":200,"message":"Hello Delete"}
4. curl -XPUT http://localhost:9990/v1/hello  
   {"code":200,"message":"Hello Put"}
