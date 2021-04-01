# SSH client by golang

## encapsulation the library from 
> github.com/helloyi/go-sshclient

- include a simple cli
- multiple setting for choose

## How to use

1. init config

```
go run main.go

initial the new config // if you see this, you were initial the config success
```

2. write the config

- write your machine information to the setting.json

    ```json
    {
      "hosts": [
        {
          "name": "thisIsOnlyForComment",
          "domain": "vmhost:22",
          "user": "root",
          "password": "123456"
        }
      ]
    }
    ```

3. running again

```
go run main.go
==============Current Store Host============
Order: 0, {Name:thisIsOnlyForComment Domain:vmhost:22 User:root Password:123456}
What is your command? input Order to connect or exit to leave
```

4. input the Order 

- just enjoy it

