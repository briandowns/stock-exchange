# Stock-Exchange Engine

The engine is responsible for blah blah blah

## Building Docker Image
Run the following command from the root of the `engine` package to build your image `engine`:

`docker build -t engine .`

## Running in Docker Container
Once you have successfully built your image, use the following command to run your `engine` image:

`docker run -d --name engine -p 7777:7777 engine`

## Test 
Test that the engine is running in the container using the following CURL command:

`curl $DOCKER_HOST:7777/api/v1/stats`

This should result in a response similar to following:

```
{
   "pid": 5,
   "uptime": "18.579360806s",
   "uptime_sec": 18.579360806,
   "time": "2016-06-22 15:27:23.360069418 +0000 UTC",
   "unixtime": 1466609243,
   "status_code_count": {},
   "total_status_code_count": {},
   "count": 0,
   "total_count": 0,
   "total_response_time": "0",
   "total_response_time_sec": 0,
   "average_response_time": "0",
   "average_response_time_sec": 0
}
```

