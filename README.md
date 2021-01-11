# ecommerce

1) Start the mysql server 

2) Clone the repository
   ```
    git clone https://github.com/shivshankarsingh98/ecommerce.git
   ```

3) Set GOPATH
   ```
   dir=$(pwd)
   export GOPATH=$dir/ecommerce/
   ```
   
4) Move to src where main.go is located
   ```
   cd $dir/ecommerce/src
   ```

5) Set this variable in main.go

   ```
   var (
	   mysqlUser = "root"
	   mysqlPass = "Pass1234"
	   databaseName = "ecommerce"
   )
   ```

6) Download and install third-party Go packages
   ```
   go get -v
   ```


7) Run the app 
   ```
   go run main.go
   ```
   
8) Example product catalouge
![alt text](https://github.com/shivshankarsingh98/ecommerce/blob/main/example-product-catalogue.jpeg?raw=true)
