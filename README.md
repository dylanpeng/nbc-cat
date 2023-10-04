# nbc-cat
## gen model

- install

```
go install gorm.io/gen/tools/gentool@latest
```

- usega

```
gentool -c etc/gen.yml
gentool -c etc/gen.yml --tables "users"

or

make gen-model
```