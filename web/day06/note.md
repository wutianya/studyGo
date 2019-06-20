# Go web服务

## HTTP 方法

- POST
- GET
- PUT
- DELETE
  
## PUT 与 POST的区别

使用PUT时需要准确的知道哪一项资源将会被替换, 而使用POST只会创建出一项新资源以及一个新的url. 简而言之, POST用于创建一项全新的资源,  
PUT则用于替换一项已经存在的资源.

## JOSN方法

- Unmarshal  
  将切边存储的数据解封至结构中 json.Unmarshal(body, &post)
- MarshalIndent  
  将结构体数据解析为切片 json.MarshalIndent(&post, "", "\t\t")
