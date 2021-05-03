# README ServiceGlue

## How to add new config

**At the moment just .yml config files are support**

1. Set custom config path or use default config path which is `./config/"env"`
1.1 set env with -env flag by start, default env is "local"
2. Define your struct  
2.1 struct name must be the same config file name  
2.2 struct fields must be the same in config file
3. Just give the `config.ReadConfig` method your struct and except interface{} which is your struct.
4. You can easily extract your struct:

```golang
i := config.ReadConfig(yourpkg.Websocket{})
cfg := i.(yourpkg.Websocket)
```