# TELEPHONE RECHARGE MACHINE HTTP SERVER API

## 1. Database Initialization Process

```bash
https://telephone.http.vsensetech.in/root/database/init
```

### HTTP Method → GET

### HTTP Responses

json message with http status

json message format

```json
{"message":"__server_status_message__"}
```

HTTP status codes

400 → Bad Request (User side error)

500 → Internal server error (Server side error)

200 → ok (Operation successful without any error)

## 2. Creating Admin Account

```bash
https://telephone.http.vsensetech.in/root/create/admin
```

### HTTP Method → POST

HTTP Request format

```json
{
    "admin_name":"vsense",
    "password":"Vsense@2024"
}
```

## 3. Admin Login

```bash
https://telephone.http.vsensetech.in/login/admin
```

### HTTP Method → POST

HTTP Request format

```json
{
    "admin_name":"vsense",
    "password":"Vsense@2024"
}
```

HTTP Response format

```json
{
    "admin_id": "4b8de3d3-e726-44db-924d-720ee555dbb1"
}
```

## 4. Delete Admin Account

```bash
https://telephone.http.vsensetech.in/root/delete/admin/__ADMIN_ID__
```

### HTTP Method → GET

## 5. Create User Account

```bash
https://telephone.http.vsensetech.in/admin/create/user
```

### HTTP Method → POST

HTTP Request format

```json
{
    "user_name":"vsense",
    "password":"Vsense@2024"
}
```

## 6. Delete User Account

```bash
https://telephone.http.vsensetech.in/admin/delete/user/__USER_ID__
```

### HTTP Method → GET

## 7. Create Machine

```bash
https://telephone.http.vsensetech.in/admin/create/machine/__ADMIN_ID__
```

### HTTP Method → POST

HTTP Request format

```json
{
    "machine_id" :"vs24rm01",
    "label":"hostel 1 machine"
}
```

## 8. Get Machines

```bash
https://telephone.http.vsensetech.in/admin/machines/__ADMIN_ID__
```

### HTTP Method → GET

HTTP Response format

```json
{
    "machines": [
        {
            "machine_id": "vs24rm01",
            "label": "hostel 1 machine",
            "balance": 0,
            "update_timestamp": "2024-11-15 15:44:18.115733728 +0000 UTC m=+42.059345349"
        }
    ]
}
```

## 9. Delete Machine

```bash
https://telephone.http.vsensetech.in/admin/delete/machine/__MACHINE_ID__
```

### HTTP Method → GET

## 10. Recharge Machine

```bash
https://telephone.http.vsensetech.in/admin/recharge/machine/__MACHINE_ID__
```

### HTTP Method → POST

HTTP Request format

```json
{
    "amount":1000
}
```

## 11. Get Recharge Machine History

```bash
https://telephone.http.vsensetech.in/admin/recharge/history/__MACHINE_ID__
```

### HTTP Method → GET

HTTP Response format

```json
{
    "recharge_history": [
        {
            "amount": 9000,
            "timestamp": "2024-11-15T16:01:53.235967Z"
        },
        {
            "amount": 8000,
            "timestamp": "2024-11-15T16:04:04.887999Z"
        },
        {
            "amount": 800,
            "timestamp": "2024-11-15T16:04:15.053672Z"
        }
    ]
}
```

## 12. User Login

```bash
https://telephone.http.vsensetech.in/login/user
```

### HTTP Method → POST

HTTP Request format

```json
{
    "user_name":"vsense",
    "password":"Vsense@2024"
}
```

HTTP Response format

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjI3OTExNDAsInVzZXJfaWQiOiJlZDc3ODk4NC02NmZjLTRiYzQtYTk3ZS03OTAyYjFmNGIyNzkiLCJ1c2VyX25hbWUiOiJ2c2Vuc2UifQ.Y0azTbt39oRcHcbyiq6n2D9-sSEGIEGIiuw_KS_6ewY"
}
```

## 13. Get Machine Balance

```bash
https://telephone.http.vsensetech.in/user/machine/balance/__MACHINE_ID__
```

### HTTP Method → GET

HTTP Response format

```json
{
    "balance": 17800
}
```

## 14. Deduct Machine Balance

```bash
https://telephone.http.vsensetech.in/user/deduct/machine/balance/__MACHINE_ID__
```

### HTTP Method → POST

HTTP Request format

```json
{
    "amount":95
}
```

## 15. Get Expense History [User]

```bash
https://telephone.http.vsensetech.in/user/expense/history/__MACHINE_ID__
```

### HTTP Method → GET

HTTP Response format

```json
{
    "expense_history": [
        {
            "machine_id": "",
            "amount": 95,
            "timestamp": "2024-11-15T16:17:53.285817Z"
        },
        {
            "machine_id": "",
            "amount": 95,
            "timestamp": "2024-11-15T16:23:27.089603Z"
        }
    ]
}
```

## 16. Get Expense History [Admin]

```bash
https://telephone.http.vsensetech.in/admin/expense/history/__MACHINE_ID__
```

### HTTP Method → GET

HTTP Response format

```json
{
    "expense_history": [
        {
            "machine_id": "",
            "amount": 95,
            "timestamp": "2024-11-15T16:17:53.285817Z"
        },
        {
            "machine_id": "",
            "amount": 95,
            "timestamp": "2024-11-15T16:23:27.089603Z"
        }
    ]
}
```
