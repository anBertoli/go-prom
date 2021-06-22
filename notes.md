### Latency

##### Cumulative latency histogram
Absolute number of requests to "/" divided by latency windows.  

`latency_microseconds_bucket{path="/"}`

##### Rate latency histogram
Per-second number of requests to "/" divided by latency windows (average on 1 min).

`rate(latency_microseconds_bucket{path="/"}[1m])`

##### Absolute number of latency measures 
Absolute number of latency measures for "/" endpoint.

`latency_microseconds_count{path="/"}`

##### Rate of latency measures
Per-second number of latency measures for "/" (average on 1 min).

`rate(latency_microseconds_count{path="/"}[1m])`

##### Percentage of latency windows  
Per-second percentage of latency windows for "/".

`100 * rate(latency_microseconds_bucket{path="/"}[1m]) / ignoring(le) group_left rate(latency_microseconds_count{path="/"}[1m])`

##### Per-second average of latencies
Per second average latency for each endpoint (average on 1 min).

`rate(latency_microseconds_sum[1m]) / rate(latency_microseconds_count[1m])`

##### Global of latencies
Global average latency for each endpoint (average on 1 min).

`latency_microseconds_sum / latency_microseconds_count`