# ecommerce

1) Start the mysql 

2) Clone the repository
   ```
    git clone https://github.com/shivshankarsingh98/redis-monitor.git
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

   already default values are set to:
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

8) Open the below address in browser and click button "Stream Redis Metrics"
   ```
   http://127.0.0.1:8080/
   ```


   ```



