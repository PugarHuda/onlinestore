##Ecommerce With Golang Project

# You can start the project with below commands
```json
docker-compose up -d
go run main.go
```
SIGNUP FUNCTION API CALL (POST REQUEST)
http://localhost:8000/users/signup

```json
{
    "first_name":"huda",
    "last_name" :"aduh",
    "email":"musangpandan@gamil.com",
    "password":"musangpandan",
    "phone":"372908428"
}
```
Response :"Successfully Signed Up!!"

LOGIN FUNCTION API CALL (POST REQUEST)

http://localhost:8000/users/login

```json
{
    "email":"musangpandan@gamil.com",
    "password":"musangpandan"
}
```
response will be like this

```json
{
    "_id": "65ca1933d54f68947d56e0ee",
    "first_name": "huda",
    "last_name": "aduh",
    "password": "$2a$14$XpLap230vfrnJB9KPMp2MuV1DcRhT9D6VWrLz657c/LaJbGZyH2MS",
    "email": "musangpandan@gamil.com",
    "phone": "372908428",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im11c2FuZ3BhbmRhbkBnYW1pbC5jb20iLCJGaXJzdF9OYW1lIjoiaHVkYSIsIkxhc3RfTmFtZSI6ImFkdWgiLCJVaWQiOiI2NWNhMTkzM2Q1NGY2ODk0N2Q1NmUwZWUiLCJleHAiOjE3MDc4Mjk5Mzl9.PlqkZnSgsitFjU-VA_NZ2QV8MB5p9o3mxJxez6dE_uA",
    "Refresh_Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X05hbWUiOiIiLCJMYXN0X05hbWUiOiIiLCJVaWQiOiIiLCJleHAiOjE3MDgzNDgzMzl9.17-e2V44v7u9zvMFbwkV_220wNfBFqAewmmnze3Ly3M",
    "created_at": "2024-02-12T13:12:19Z",
    "updtaed_at": "2024-02-12T13:12:19Z",
    "user_id": "65ca1933d54f68947d56e0ee",
    "usercart": [],
    "address": [],
    "orders": []
}
```
Admin add Product Function (POST REQUEST)

http://localhost:8000/admin/addproduct

```json
{
  "product_name": "Macbook",
  "price": 34536,
  "rating": 9,
  "image": "yuhuuu.jpg"
}
```
Response : "Successfully added our Product Admin!!"

View all the Products in db GET REQUEST

http://localhost:8000/users/productview

Response

```json
[
    {
        "Product_ID": "65ca201bd54f68947d56e0f0",
        "product_name": "Apple Vision Pro",
        "price": 3599,
        "rating": 9,
        "image": "VisionPro.jpg"
    },
    {
        "Product_ID": "65ca2094d54f68947d56e0f2",
        "product_name": "Lenovo",
        "price": 2342,
        "rating": 9,
        "image": "yaaa.jpg"
    },
    {
        "Product_ID": "65ca20abd54f68947d56e0f4",
        "product_name": "Macbook",
        "price": 34536,
        "rating": 9,
        "image": "yuhuuu.jpg"
    }
]
```
Search Product by regex function (GET REQUEST)
defines the word search sorting http://localhost:8000/users/search?name=Le

response:

```json
[
    {
        "Product_ID": "65ca2094d54f68947d56e0f2",
        "product_name": "Lenovo",
        "price": 2342,
        "rating": 9,
        "image": "yaaa.jpg"
    }
]
```
Adding the Products to the Cart (GET REQUEST)

http://localhost:8000/addtocart?id=xxxproduct_idxxx&userID=xxxxxxuser_idxxxxxx

Corresponding MongoDB query

Removing Item From the Cart (GET REQUEST)

http://localhost:8000/addtocart?id=xxxxxxx&userID=xxxxxxxxxxxx

Listing the item in the users' cart (GET REQUEST) and total price

http://localhost:8000/listcart?id=xxxxxxuser_idxxxxxxxxxx

Adding the Address (POST REQUEST)

http://localhost:8000/addadress?id=user_id**\*\***\***\*\***

The Address array is limited to two values home and work address more than two address is not acceptable

```json
{
  "house_name": "rumah hantu",
  "street_name": "ihtakut",
  "city_name": "sukamars",
  "pin_code": "777777"
}
```
Editing the Home Address(PUT REQUEST)

http://localhost:8000/edithomeaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx

Editing the Work Address(PUT REQUEST)

http://localhost:8000/editworkaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx

Delete Addresses(GET REQUEST)

http://localhost:8000/deleteaddresses?id=xxxxxxxxxuser_idxxxxxxxxxxxxx

delete both addresses

Cart Checkout Function and placing the order(GET REQUEST)

After placing the order the items have to be deleted from cart functonality added

http://localhost:8000?id=xxuser_idxxx

Instantly Buying the Products(GET REQUEST) http://localhost:8000?userid=xxuser_idxxx&pid=xxxxproduct_idxxxx
