# arvantest-monolith
Arvan test inital monolith architecture

```
localhost:8080 wallet system
localhost:8081 discount system
```

```
localhost:8080/get-balance gives balance of user with given request body:

{
	"phone_number":"9126116"
}
```
with result:
```
{
    "balance": 2110
}
```
```
localhost:8080/update-balance updates balance of user that submitted voucher code(if valid).
```
input:

```
{
	"phone_number":"9126116",
	"amount":10
}
```

result :
```
{
    "result": "Ok"
}
```
mean all done and it's fine.
Ps. Actuallt it should've been grpc internal communication from voucher(discount) service to wallet. Due to building microservice separate from monolith and learning lots of new stuff, I had to skip implementing it.

localhost:8081/enable-voucher-code

I assumed that we have limited voucher codes and each one has specific amount of gift credit. This  api gets the code and the credit we want and enables it in database.

input:

```
{
	"voucher_code":"98CB-7558-JF9U",
	"amount": 1000000
}

output:
{
    "result": "OK"
}
```

```
localhost:8081/submit-voucher-code submits specific voucher code for specific user. 
```
input:
```
{
	"phone_number":"9126116",
	"voucher_code":"98CB-7558-JF9U"
}
output:
{
    "result": "OK"
}
```

This API also calls update-balance internally.
```
localhost:8081/get-voucher-code-status/{voucher-code}
```

returns all voucher entities that are active and the users that used it.

result:
```
{
    "used_voucher_code_users": [
        {
            "PhoneNumber": "9126113",
            "VoucherCode": "98CB-7558-JF9U",
            "IsUsed": true
        },
        {
            "PhoneNumber": "9126116",
            "VoucherCode": "98CB-7558-JF9U",
            "IsUsed": true
        },
        {
            "PhoneNumber": "9126117",
            "VoucherCode": "98CB-7558-JF9U",
            "IsUsed": true
        }
    ]
}
```


