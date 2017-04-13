# GoJsoner
Loading JSON file in Golang and discarding comments.

## Installation

Install GoJsoner using the "go get" command:

```
go get github.com/sipt/GoJsoner
```

## Examples

```Go
result, err := Discard(`
		{//test comment1
			"name": "测试",
			/**
			test comment2
			1
			2
			3
			end
			*/
			"age":26 //test comment3
			/*****/
		}
	`)
```

output

```
{"name":"测试","age":26}
```


