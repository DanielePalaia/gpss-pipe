# gpss-pipe
gpss integration with pipes

This component is listening forever to a pipe (when EOF is reached it ask to gpss to write the info). </br>

1) a gpss server needs to be initialized and working </br>

2) Let's create a named pipe in Linux: </br></br>
    **mkfifo mypipe </br>**

3) Let's create the destination table in greenplum (whatever table is fine if it is coherent with the input fields) ex. </br>

CREATE TABLE people(id int, name varchar(1000), surname varchar(1000), email varchar(1000), gender varchar(10)); </br>

4) Configure the program properties file that needs to be in the path you are running the software (./bin/linux/properties.ini), where specify the path of the pipe created and the delim set as input field separator (; in case of csv) </br>

```
   GpssAddress=10.91.51.23:50007
   GreenplumAddress=10.91.51.23
   GreenplumPort=5533
   GreenplumUser=gpadmin
   GreenplumPasswd=
   Database=test
   SchemaName=public
   TableName=people
   PipePath=./mypipe
   Delim=;
   Batch=100
```

5) Run the software (./bin/macosx/pipegpss or ./bin/linux/pipegpss) </br>

```
    Danieles-MacBook-Pro:bin dpalaia$ ./pipegpss
    2019/03/14 15:58:11 Starting the connector and reading properties in the properties.ini file
    2019/03/14 15:58:11 Properties read: Connecting to the Grpc server specified
    2019/03/14 15:58:11 Connected to the grpc server
    2019/03/14 15:58:11 delegating to pipe client
    2019/03/14 15:58:11 Opening named pipe: ./mypipe for reading
    2019/03/14 15:58:11 waiting for someone to write something in the pipe

6) submit the example csv file provided in the pipe (./bin/macosx/data.csv): it contains 1000 elements </br>

    cat data.csv >> mypipe </br>

7) you should see some logs in the pipegpss screen and the table populated with 1K elements </br>

```
    test=# select count(*) from people;
     count 
    -------
      1000
    (1 row)
