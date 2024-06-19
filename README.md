# Synapsis Test

Hello mas/mba/bapak/ibu,

First of all, I would like to apologize profusely for being unable to complete the synapsis test due to illness.

I had previously informed the Synapsis HR team, but the Synapsis HR team gave me a second chance, but it turned out that I had not completely recovered. Thanks for second opportunity as well.

I really tried to at least submit this test to respect the Hiring Manager / Human Resource team from Synapsis, I hope sir/madam/madam/madam will forgive me for the problems I experienced.

I have accepted that by doing this I will be deemed not to have passed to the next stage, because I have not been able to complete all the requirements given.

However, let me explain some of the concerns or todo lists in my mind for this test question, for my self to learn:

### list of todo
**required:**
1. create flow Customer can add product to shopping cart
2. create flow Customers can see a list of products that have been added to the shopping cart
3. create flow Customer can delete product list in shopping cart
4. create flow Customers can checkout and make payment transactions
5. create swagger with open api standar implementation
6. implement CI/CD with docker as containerization

**for my self:**
- buffer algorithm for order flow to minimize error
- use payment gateway xendit development environment for payment transaction 
- design roles database
- implement self registry with private S3 and using portainer as webhook


> "I hope you can find the best candidate." - Ananda Affan Fattahila

## List CURL API 

### Register

**body request:**
- username: string
- password: string, minimum 6 char

```
curl --location 'https://synapsis.fanzru.dev/accounts/register' \
--header 'Content-Type: application/json' \
--data '{
    "username": "fanzru",
    "password": "fanzru123"
}'
```

### Login
**body request:**
- username: string
- password: string, minimum 6 char

```
curl --location 'https://synapsis.fanzru.dev/accounts/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "fanzru",
    "password": "fanzru123"
}'
```
### Find Product
query parameters:
- categoryId : integer, in this case you can use 1,2,3,4,5
```
curl --location 'https://synapsis.fanzru.dev/product/?categoryId=1'
```
### Find Product per Category
query parameters:
- categoryId : integer, in this case you can use 1,2,3,4,5
```
curl --location 'https://synapsis.fanzru.dev/categories/?categoryId=1'
```


Stacks : Golang, MySQL, Nginx, Ubuntu...

  

# **Documentation**

![Entity Relationship Diagram](https://github.com/fanzru/synapsis-test/blob/main/erd.png)
  

Using Domain Drive Design with Golang

  

## Migration Database

  

> If you have change in database please create sql syntax to rollback, because infra need sleep.

  

Please Install `golang-migrate` to migrate database, to install `golang-migrate` you can read this documentation:

  

**Windows :**

```

https://verinumbenator.medium.com/installing-golang-migrate-on-windows-b4b3df9b97b2

```

**Mac :**

```

brew install golang-migrate

```

**Ubuntu :**

```

https://www.geeksforgeeks.org/how-to-install-golang-migrate-on-ubuntu/

```

## Migration

if have done to install `golang-migrate`, please prepare your database and create new database, for example `kerjago_db`.

**Migrate Create**

Create a migration file. You can find the file at `migration` folder

```

make migrate-create NAME=namefile

```

  
  
  

**Migrate Up**

To migrate all your migration file

```

make migrate-up

```

  

**Migrate Down**

To delete all your schema with migration

```

make migrate-down

```

  

**Migrate Rollback**

to run migration down only `N` step(s) behind

```

make migrate-rollback N=yournumberrunmigrationdown

```

  

**Fixing your Migration**

What happend if your database is dirty?

You can fix your migration first and then using foce command with the version you want.

If you're happend to get `error: Dirty database version 16. Fix and force version.`

Then you want to run:

```

make migrate-force VERSION=15

```

Reference: https://github.com/golang-migrate/migrate/issues/282#issuecomment-530743258# synapsis-test
