# online-shopping-cart

Requirement for starting project
<ul>
  <li>Docker</li>
  <li>Docker-compose</li>
</ul>
<br>

Start project
```shell
docker-compose up
```
<br>

Reset project database
```shell
docker-compose down --volumes
```

Accessing swagger
```shell
http://localhost:8080/docs/index.html
```
<br>
Sample account for login

```shell
ADMIN
{
  "email": "admin@example.com",
  "password": "password"
}

CUSTOMER
{
  "email": "customer1@example.com",
  "password": "password"
}
```
