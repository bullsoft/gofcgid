## 1. Build

go get ...

go install ...

## 2. Start the server

```
Listen on:  0.0.0.0 : 8888
```


## 3. client connect

```
➜  ~  telnet 127.0.0.1 8888
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
```

## 4. Server log

```
Accept from:  127.0.0.1:61149

```

## 5. Issue post in client and response form server

```
hello=world&foo=bar
X-Powered-By: PHP/5.4.24
Content-type: text/html

array(2) {
  ["hello"]=>
  string(5) "world"
  ["foo"]=>
  string(3) "bar"
}
array(2) {
  ["foo"]=>
  string(3) "bar"
  ["hello"]=>
  string(5) "world"
}
hello, world1
```

## 6. Server log

```
We get  21  bytes:  hello=world&foo=bar
```


