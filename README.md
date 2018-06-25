# go.orm
## 一个go的orm RestfulAPI demo

成功启动后

使用浏览器或者postman

**get : http://localhost:3000/user/query/{id}**
根据id查询用户信息

**post： http://localhost:3000/user/register**  
注册一个用户，id自增
需要一串inbody的json作为参数，例如:
{
&emsp; "name":"xpppp",
&emsp; "phone":"111111",
&emsp; "address":"china"
}

**delete： http://localhost:3000/user/delete/{id}**
删除此id的用户

**put： http://localhost:3000/user/update**
根据id修改用户信息
需要一串inbody的json作为参数，例如:
{
&emsp;	"id": 2,
&emsp;	"name":"xpppp",
&emsp;	"phone":"22222",
&emsp;	"address":"米国"
}