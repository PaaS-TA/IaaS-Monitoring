# Node Detail View API
* [Openstack Node Cpu Usage .GET                 /v1/openstack/node/:hostname/cpuUsage](#10)
* [Openstack Node Cpu Load Average .GET          /v1/openstack/node/:hostname/cpuLoad](#11)
* [Openstack Node Memory Swap Usage .GET         /v1/openstack/node/:hostname/swapUsage](#12)
* [Openstack Node Memory memory Usage .GET       /v1/openstack/node/:hostname/memUsage](#13)
* [Openstack Node Memory disk Usage by mountPoint .GET        /v1/openstack/node/:hostname/diskUsage](#14)
* [Openstack Node Memory disk Read Kbyte by mountPoint .GET   /v1/openstack/node/:hostname/diskRead](#15)
* [Openstack Node Memory disk Write Kbyte by mountPoint .GET  /v1/openstack/node/:hostname/diskWrite](#16)
* [Openstack Node Memory Network Kbyte .GET                   /v1/openstack/node/:hostname/networkIo](#17)
* [Openstack Node Memory Network Error .GET                   /v1/openstack/node/:hostname/networKError](#18)
* [Openstack Node Memory Network Drop Packet .GET             /v1/openstack/node/:hostname/networKError](#19)

## Node Detail View API

<div id='10'/>

### **openstack Node Cup Usage** :  ``GET /v1/openstack/node/:hostname/cpuUsage``

#### openstack Node의 CPU 사용률
    
   openstack Node의 CPU 사용률 요약정보를 조회한다.
   
<h4>Request Parameters</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
            	<th>Optional</th>
            	<th>Example Values</th>
	</tr>
    </thead>
            <tbody>
                <tr class="">
                    <td><span>hostname</span></td>
                    <td><span>조회하려는 host명</span></td>
                    <td><span>required</span></td>
                    <td><span>ex) compute1</span></td>
                </tr>
                 <tr class="">
                    <td><span>defaultTimeRange</span></td>
                        <td><span>현재 시간기준으로 조회할 시간 범위</span></td>
                        <td><span>option(timeRangeFrom/To 없을때 필수)</span></td>
                        <td><span>ex) 10m</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeFrom</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 from</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 2h</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeTo</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 To</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 1h</span></td>
                </tr>
                    <tr class="">
                    <td><span>groupBy</span></td>
                        <td><span>조회된 Data의 평균치를 구할 시간 단위</span></td>
                        <td><span>required</span></td>
                        <td><span>ex)15s</span></td>
                </tr>
            </tbody>
</table>


<h4>Response</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
	</tr>
    </thead>
    	<tbody>
            <tr class="">
                <td><span>time</span></td>
                <td><span>시간(unix time)</span></td>
            </tr>
            <tr class="">
                <td><span>usage</span></td>
                <td><span>CPU 사용률</span></td>
            </tr>
	</tbody>
</table>


```
[request]
openstack Node들의 CPU사용률 조회
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/cpuUsage?defaultTimeRange=10m&groupBy=1m'
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/cpuUsage?timeRangeFrom=10m&timeRangeTo=30m&groupBy=1m'
```

##### response
```
[
   {
      "metric":[
         {
            "time":1502330400,
            "usage":20.7
         },
         {
            "time":1502330460,
            "usage":21.95
         },
         ...
      ],
      "name":"cpuUsage"
   }
]
```



<div id='11'/>

### **openstack Node Cup Load Average** :  ``GET /v1/openstack/node/:hostname/cpuLoad``

#### openstack Node의 CPU Load Average

   openstack Node의 CPU Load Average정보를 조회한다.(1/5/15 Minute)

<h4>Request Parameters</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
            	<th>Optional</th>
            	<th>Example Values</th>
	</tr>
    </thead>
            <tbody>
                <tr class="">
                    <td><span>hostname</span></td>
                    <td><span>조회하려는 host명</span></td>
                    <td><span>required</span></td>
                    <td><span>ex) compute1</span></td>
                </tr>
                 <tr class="">
                    <td><span>defaultTimeRange</span></td>
                        <td><span>현재 시간기준으로 조회할 시간 범위</span></td>
                        <td><span>option(timeRangeFrom/To 없을때 필수)</span></td>
                        <td><span>ex) 10m</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeFrom</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 from</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 2h</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeTo</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 To</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 1h</span></td>
                </tr>
                    <tr class="">
                    <td><span>groupBy</span></td>
                        <td><span>조회된 Data의 평균치를 구할 시간 단위</span></td>
                        <td><span>required</span></td>
                        <td><span>ex)15s</span></td>
                </tr>
            </tbody>
</table>


<h4>Response</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
	</tr>
    </thead>
    	<tbody>
        <tr class="">
        	<td><span>time</span></td>
           	<td><span>시간(unix time)</span></td>
	    </tr>
        <tr class="">
        	<td><span>usage</span></td>
           	<td><span>CPU Load Time</span></td>
	    </tr>
	</tbody>
</table>


```
[request]
openstack Node들의 CPU Load Time정보 조회
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/cpuLoad?defaultTimeRange=10m&groupBy=1m'
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/cpuLoad?timeRangeFrom=10m&timeRangeTo=30m&groupBy=1m'
```

##### response
```
[
   {
      "metric":[
         {
            "time":1502330400,
            "usage":20.7
         },
         {
            "time":1502330460,
            "usage":21.95
         },
         ...
      ],
      "name":"cpuLoad"
   }
]
```

<div id='12'/>

### **openstack Node Swap Memory Usage** :  ``GET /v1/openstack/node/:hostname/swapUsage``

#### openstack Node의 Swap Memory Usage

   openstack Node의 Swap Memory Usage 정보를 조회한다.

<h4>Request Parameters</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
            	<th>Optional</th>
            	<th>Example Values</th>
	</tr>
    </thead>
            <tbody>
                <tr class="">
                    <td><span>hostname</span></td>
                    <td><span>조회하려는 host명</span></td>
                    <td><span>required</span></td>
                    <td><span>ex) compute1</span></td>
                </tr>
                 <tr class="">
                    <td><span>defaultTimeRange</span></td>
                        <td><span>현재 시간기준으로 조회할 시간 범위</span></td>
                        <td><span>option(timeRangeFrom/To 없을때 필수)</span></td>
                        <td><span>ex) 10m</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeFrom</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 from</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 2h</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeTo</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 To</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 1h</span></td>
                </tr>
                    <tr class="">
                    <td><span>groupBy</span></td>
                        <td><span>조회된 Data의 평균치를 구할 시간 단위</span></td>
                        <td><span>required</span></td>
                        <td><span>ex)15s</span></td>
                </tr>
            </tbody>
</table>


<h4>Response</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
	</tr>
    </thead>
    	<tbody>
        <tr class="">
        	<td><span>time</span></td>
           	<td><span>시간(unix time)</span></td>
	    </tr>
            <tr class="">
        	<td><span>usage</span></td>
           	<td><span>Swap 사용률</span></td>
	    </tr>
	</tbody>
</table>


```
[request]
openstack Node들의 Memory Swap 사용률 정보 조회
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/swapUsage?defaultTimeRange=10m&groupBy=1m'
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/swapUsage?timeRangeFrom=10m&timeRangeTo=30m&groupBy=1m'
```

##### response
```
[
   {
      "metric":[
         {
            "time":1502330400,
            "usage": 0
         },
         {
            "time":1502330460,
            "usage": 0
         },
         ...
      ],
      "name":""memorySwap""
   }
]
```


<div id='13'/>

### **openstack Node Memory Usage** :  ``GET /v1/openstack/node/:hostname/mempUsage``

#### openstack Node의 Memory 사용률

   openstack Node의 Memory 사용률 정보를 조회한다.

<h4>Request Parameters</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
            	<th>Optional</th>
            	<th>Example Values</th>
	</tr>
    </thead>
            <tbody>
                <tr class="">
                    <td><span>hostname</span></td>
                    <td><span>조회하려는 host명</span></td>
                    <td><span>required</span></td>
                    <td><span>ex) compute1</span></td>
                </tr>
                 <tr class="">
                    <td><span>defaultTimeRange</span></td>
                        <td><span>현재 시간기준으로 조회할 시간 범위</span></td>
                        <td><span>option(timeRangeFrom/To 없을때 필수)</span></td>
                        <td><span>ex) 10m</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeFrom</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 from</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 2h</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeTo</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 To</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 1h</span></td>
                </tr>
                    <tr class="">
                    <td><span>groupBy</span></td>
                        <td><span>조회된 Data의 평균치를 구할 시간 단위</span></td>
                        <td><span>required</span></td>
                        <td><span>ex)15s</span></td>
                </tr>
            </tbody>
</table>


<h4>Response</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
	</tr>
    </thead>
    	<tbody>
        <tr class="">
        	<td><span>time</span></td>
           	<td><span>시간(unix time)</span></td>
	    </tr>
            <tr class="">
        	<td><span>usage</span></td>
           	<td><span>Memory 사용률</span></td>
	    </tr>
	</tbody>
</table>


```
[request]
openstack Node들의 Memory 사용률 정보 조회
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/memUsage?defaultTimeRange=10m&groupBy=1m'
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/memUsage?timeRangeFrom=10m&timeRangeTo=30m&groupBy=1m'
```

##### response
```
[
   {
      "metric":[
         {
            "time":1502330400,
            "usage": 60.5
         },
         {
            "time":1502330460,
            "usage": 60.8
         },
         ...
      ],
      "name":"""memoryUsage"""
   }
]
```

<div id='14'/>

### **openstack Node Disk Usage By MountPoint** :  ``GET /v1/openstack/node/:hostname/diskUsage``

#### openstack Node의 Disk 사용률 정보

   openstack Node의 Mount Point별 Disk 사용률 정보를 조회한다.

<h4>Request Parameters</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
            	<th>Optional</th>
            	<th>Example Values</th>
	</tr>
    </thead>
            <tbody>
                <tr class="">
                    <td><span>hostname</span></td>
                    <td><span>조회하려는 host명</span></td>
                    <td><span>required</span></td>
                    <td><span>ex) compute1</span></td>
                </tr>
                 <tr class="">
                    <td><span>defaultTimeRange</span></td>
                        <td><span>현재 시간기준으로 조회할 시간 범위</span></td>
                        <td><span>option(timeRangeFrom/To 없을때 필수)</span></td>
                        <td><span>ex) 10m</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeFrom</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 from</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 2h</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeTo</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 To</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 1h</span></td>
                </tr>
                    <tr class="">
                    <td><span>groupBy</span></td>
                        <td><span>조회된 Data의 평균치를 구할 시간 단위</span></td>
                        <td><span>required</span></td>
                        <td><span>ex)15s</span></td>
                </tr>
            </tbody>
</table>


<h4>Response</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
	</tr>
    </thead>
    	<tbody>
    	<tr class="">
                	<td><span>name</span></td>
                   	<td><span>mountPoint Name</span></td>
        </tr>
            <tr class="">
        	<td><span>time</span></td>
           	<td><span>시간(unix time)</span></td>
	    </tr>
            <tr class="">
        	<td><span>usage</span></td>
           	<td><span>Disk Memory 사용률</span></td>
	    </tr>
	</tbody>
</table>


```
[request]
openstack Node들의 Memory Swap 사용률 정보 조회
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/diskUsage?defaultTimeRange=10m&groupBy=1m'
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/diskUsage?timeRangeFrom=10m&timeRangeTo=30m&groupBy=1m'
```

##### response
```
[
   {
      "metric":[
         {
            "time":1502335680,
            "usage":22.5
         },
         {
            "time":1502335740,
            "usage":22.5
         },
        ...
      ],
      "name":"/"
   },
   {
      "metric":[
         {
            "time":1502335680,
            "usage":0.7
         },
         {
            "time":1502335740,
            "usage":0.7
         },
         ...
      ],
      "name":"/boot/efi"
   }
]
```


<div id='15'/>

### **openstack Node Disk Read Kbyte By MountPoint** :  ``GET /v1/openstack/node/:hostname/diskRead``

#### openstack Node의 Disk Read Kbyte

   openstack Node의 Disk Read(KByte) 정보를 조회한다.

<h4>Request Parameters</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
            	<th>Optional</th>
            	<th>Example Values</th>
	</tr>
    </thead>
            <tbody>
                <tr class="">
                    <td><span>hostname</span></td>
                    <td><span>조회하려는 host명</span></td>
                    <td><span>required</span></td>
                    <td><span>ex) compute1</span></td>
                </tr>
                 <tr class="">
                    <td><span>defaultTimeRange</span></td>
                        <td><span>현재 시간기준으로 조회할 시간 범위</span></td>
                        <td><span>option(timeRangeFrom/To 없을때 필수)</span></td>
                        <td><span>ex) 10m</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeFrom</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 from</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 2h</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeTo</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 To</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 1h</span></td>
                </tr>
                    <tr class="">
                    <td><span>groupBy</span></td>
                        <td><span>조회된 Data의 평균치를 구할 시간 단위</span></td>
                        <td><span>required</span></td>
                        <td><span>ex)15s</span></td>
                </tr>
            </tbody>
</table>


<h4>Response</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
	</tr>
    </thead>
    	<tbody>
    	<tr class="">
            <td><span>name</span></td>
            <td><span>mountPoint Name</span></td>
        </tr>
            <tr class="">
        	<td><span>time</span></td>
           	<td><span>시간(unix time)</span></td>
	    </tr>
            <tr class="">
        	<td><span>usage</span></td>
           	<td><span>Disk Read KByte</span></td>
	    </tr>
	</tbody>
</table>


```
[request]
openstack Node들의 Disk Read 정보 조회
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/diskRead?defaultTimeRange=10m&groupBy=1m'
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/diskRead?timeRangeFrom=10m&timeRangeTo=30m&groupBy=1m'
```

##### response
```
[
   {
      "metric":[
         {
            "time":1502335680,
            "usage":22.5
         },
         {
            "time":1502335740,
            "usage":22.5
         },
        ...
      ],
      "name":"/"
   },
   {
      "metric":[
         {
            "time":1502335680,
            "usage":0.7
         },
         {
            "time":1502335740,
            "usage":0.7
         },
         ...
      ],
      "name":"/boot/efi"
   }
]
```


<div id='16'/>

### **openstack Node Disk Write Kbyte By MountPoint** :  ``GET /v1/openstack/node/:hostname/diskWrite``

#### openstack Node의 Disk Write Kbyte

   openstack Node의 Disk Write(KByte) 정보를 조회한다.

<h4>Request Parameters</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
            	<th>Optional</th>
            	<th>Example Values</th>
	</tr>
    </thead>
            <tbody>
                <tr class="">
                    <td><span>hostname</span></td>
                    <td><span>조회하려는 host명</span></td>
                    <td><span>required</span></td>
                    <td><span>ex) compute1</span></td>
                </tr>
                 <tr class="">
                    <td><span>defaultTimeRange</span></td>
                        <td><span>현재 시간기준으로 조회할 시간 범위</span></td>
                        <td><span>option(timeRangeFrom/To 없을때 필수)</span></td>
                        <td><span>ex) 10m</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeFrom</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 from</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 2h</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeTo</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 To</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 1h</span></td>
                </tr>
                    <tr class="">
                    <td><span>groupBy</span></td>
                        <td><span>조회된 Data의 평균치를 구할 시간 단위</span></td>
                        <td><span>required</span></td>
                        <td><span>ex)15s</span></td>
                </tr>
            </tbody>
</table>


<h4>Response</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
	</tr>
    </thead>
    	<tbody>
    	<tr class="">
            <td><span>name</span></td>
            <td><span>mountPoint Name</span></td>
        </tr>
            <tr class="">
        	<td><span>time</span></td>
           	<td><span>시간(unix time)</span></td>
	    </tr>
            <tr class="">
        	<td><span>usage</span></td>
           	<td><span>Disk write Kbyte</span></td>
	    </tr>
	</tbody>
</table>


```
[request]
openstack Node들의 Disk Write 정보 조회
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/diskWrite?defaultTimeRange=10m&groupBy=1m'
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/diskWrite?timeRangeFrom=10m&timeRangeTo=30m&groupBy=1m'
```

##### response
```
[
   {
      "metric":[
         {
            "time":1502335680,
            "usage":22.5
         },
         {
            "time":1502335740,
            "usage":22.5
         },
        ...
      ],
      "name":"/"
   },
   {
      "metric":[
         {
            "time":1502335680,
            "usage":0.7
         },
         {
            "time":1502335740,
            "usage":0.7
         },
         ...
      ],
      "name":"/boot/efi"
   }
]
```


<div id='17'/>

### **openstack Node Network IO Kbyte ** :  ``GET /v1/openstack/node/:hostname/networKIo``

#### openstack Node의 Network IO Kbyte

   openstack Node의 Network IO Kbyte 정보를 조회한다.

<h4>Request Parameters</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
            	<th>Optional</th>
            	<th>Example Values</th>
	</tr>
    </thead>
            <tbody>
                <tr class="">
                    <td><span>hostname</span></td>
                    <td><span>조회하려는 host명</span></td>
                    <td><span>required</span></td>
                    <td><span>ex) compute1</span></td>
                </tr>
                 <tr class="">
                    <td><span>defaultTimeRange</span></td>
                        <td><span>현재 시간기준으로 조회할 시간 범위</span></td>
                        <td><span>option(timeRangeFrom/To 없을때 필수)</span></td>
                        <td><span>ex) 10m</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeFrom</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 from</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 2h</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeTo</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 To</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 1h</span></td>
                </tr>
                    <tr class="">
                    <td><span>groupBy</span></td>
                        <td><span>조회된 Data의 평균치를 구할 시간 단위</span></td>
                        <td><span>required</span></td>
                        <td><span>ex)15s</span></td>
                </tr>
            </tbody>
</table>


<h4>Response</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
	</tr>
    </thead>
    	<tbody>
    	<tr class="">
            <td><span>name</span></td>
            <td><span>network In/Out Error</span></td>
        </tr>
            <tr class="">
        	<td><span>time</span></td>
           	<td><span>시간(unix time)</span></td>
	    </tr>
            <tr class="">
        	<td><span>usage</span></td>
           	<td><span>Network In/Out Kbyte</span></td>
	    </tr>
	</tbody>
</table>


```
[request]
openstack Node들의 Network IO(Kybte) 정보 조회
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/networKIo?defaultTimeRange=10m&groupBy=30s'
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/networKIo?timeRangeFrom=10m&timeRangeTo=30m&groupBy=30s'
```

##### response
```
[
   {
      "metric":[
         {
            "time":1502335680,
            "usage":22.5
         },
         {
            "time":1502335740,
            "usage":22.5
         },
        ...
      ],
      "name":"networkInKbyte"
   },
   {
      "metric":[
         {
            "time":1502335680,
            "usage":2779.24
         },
         {
            "time":1502335740,
            "usage":2779.24
         },
         ...
      ],
      "name":"networkOutKbyte"
   }
]
```

<div id='18'/>

### **openstack Node Network IO Error ** :  ``GET /v1/openstack/node/:hostname/networKError``

#### openstack Node의 Network Error

   openstack Node의 Network Error 정보를 조회한다.

<h4>Request Parameters</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
            	<th>Optional</th>
            	<th>Example Values</th>
	</tr>
    </thead>
            <tbody>
                <tr class="">
                    <td><span>hostname</span></td>
                    <td><span>조회하려는 host명</span></td>
                    <td><span>required</span></td>
                    <td><span>ex) compute1</span></td>
                </tr>
                 <tr class="">
                    <td><span>defaultTimeRange</span></td>
                        <td><span>현재 시간기준으로 조회할 시간 범위</span></td>
                        <td><span>option(timeRangeFrom/To 없을때 필수)</span></td>
                        <td><span>ex) 10m</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeFrom</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 from</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 2h</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeTo</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 To</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 1h</span></td>
                </tr>
                    <tr class="">
                    <td><span>groupBy</span></td>
                        <td><span>조회된 Data의 평균치를 구할 시간 단위</span></td>
                        <td><span>required</span></td>
                        <td><span>ex)15s</span></td>
                </tr>
            </tbody>
</table>


<h4>Response</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
	</tr>
    </thead>
    	<tbody>
    	<tr class="">
            <td><span>name</span></td>
            <td><span>network In/Out Error</span></td>
        </tr>
            <tr class="">
        	<td><span>time</span></td>
           	<td><span>시간(unix time)</span></td>
	    </tr>
            <tr class="">
        	<td><span>usage</span></td>
           	<td><span>Network IO Error</span></td>
	    </tr>
	</tbody>
</table>


```
[request]
openstack Node들의 Network IO(Kybte) 정보 조회
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/networKError?defaultTimeRange=10m&groupBy=30s'
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/networKError?timeRangeFrom=10m&timeRangeTo=30m&groupBy=30s'
```

##### response
```
[
   {
      "metric":[
         {
            "time":1502335680,
            "usage":22.5
         },
         {
            "time":1502335740,
            "usage":22.5
         },
        ...
      ],
      "name":"networkInError"
   },
   {
      "metric":[
         {
            "time":1502335680,
            "usage":2779.24
         },
         {
            "time":1502335740,
            "usage":2779.24
         },
         ...
      ],
      "name":"networkOutError"
   }
]
```



<div id='19'/>

### **openstack Node Network Drop Packet ** :  ``GET /v1/openstack/node/:hostname/networkDropPacket``

#### openstack Node의 Network networKDropPacket

   openstack Node의 Network networKDropPacket 정보를 조회한다.

<h4>Request Parameters</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
            	<th>Optional</th>
            	<th>Example Values</th>
	</tr>
    </thead>
            <tbody>
                <tr class="">
                    <td><span>hostname</span></td>
                    <td><span>조회하려는 host명</span></td>
                    <td><span>required</span></td>
                    <td><span>ex) compute1</span></td>
                </tr>
                 <tr class="">
                    <td><span>defaultTimeRange</span></td>
                        <td><span>현재 시간기준으로 조회할 시간 범위</span></td>
                        <td><span>option(timeRangeFrom/To 없을때 필수)</span></td>
                        <td><span>ex) 10m</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeFrom</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 from</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 2h</span></td>
                </tr>
                    <tr class="">
                    <td><span>timeRangeTo</span></td>
                        <td><span>구간 조회시 사용하는 parameter로 조회 시간 To</span></td>
                        <td><span>option (defaultTimeRange없을때 필수)</span></td>
                        <td><span>ex) 1h</span></td>
                </tr>
                    <tr class="">
                    <td><span>groupBy</span></td>
                        <td><span>조회된 Data의 평균치를 구할 시간 단위</span></td>
                        <td><span>required</span></td>
                        <td><span>ex)15s</span></td>
                </tr>
            </tbody>
</table>


<h4>Response</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
	</tr>
    </thead>
    	<tbody>
    	<tr class="">
            <td><span>name</span></td>
            <td><span>network In/Out Error</span></td>
        </tr>
            <tr class="">
        	<td><span>time</span></td>
           	<td><span>시간(unix time)</span></td>
	    </tr>
            <tr class="">
        	<td><span>usage</span></td>
           	<td><span>Network Drop Packet수</span></td>
	    </tr>
	</tbody>
</table>


```
[request]
openstack Node들의 Network IO(Kybte) 정보 조회
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/networkDropPacket?defaultTimeRange=10m&groupBy=30s'
curl -X GET -G 'http://localhost:8080/v1/openstack/node/compute1/networkDropPacket?timeRangeFrom=10m&timeRangeTo=30m&groupBy=30s'
```

##### response
```
[
   {
      "metric":[
         {
            "time":1502335680,
            "usage":0
         },
         {
            "time":1502335740,
            "usage":0
         },
        ...
      ],
      "name":"networkInDroppedPacket"
   },
   {
      "metric":[
         {
            "time":1502335680,
            "usage":0
         },
         {
            "time":1502335740,
            "usage":0
         },
         ...
      ],
      "name":"networkOutDroppedPacket"
   }
]
```