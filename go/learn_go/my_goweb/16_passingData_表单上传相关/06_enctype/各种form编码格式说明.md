```

application/x-www-form-urlencoded方式是Jquery的Ajax请求默认方式，这种方式的好处就是浏览器都支持，
在请求发送过程中会对数据进行序列化处理，以键值对形式？key1=value1&key2=value2的方式发送到服务器，如果用Jquery，
它内部已经进行了处理，如果自己写原生的Ajax请求，就需要自己对数据进行序列化。



application/json，随着json规范的越来越流行，并且浏览器支持程度原来越好，
许多开发人员易application/json作为请求content-type，告诉服务器请求的主题内容是json格式的字符串，服务器端会对json字符串进行解析，这种方式的好处就是前端人员不需要关心数据结构的复杂度，
只要是标准的json格式就能提交成功，application/json数据格式越来越得到开发人员的青睐。


multipart/form-data
multipart/form-data 。 指定传输数据为二进制类型，比如图片、mp3、文件。


enctype="text/plain"
text/plain是纯文本传输的意思，在发送邮件时要设置这种编码类型，否则会出现接收时编码混乱的问题，网络上经常拿text/plain和 text/html做比较，其实这两个很好区分，前者用来传输纯文本文件，后者则是传递html代码的编码类型，
```