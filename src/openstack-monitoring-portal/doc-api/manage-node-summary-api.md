# Server API
* [manageNode Summary .GET       /v1/openstack/manageNode/summary](#10)
* [RabbitMq Summary .GET         /v1/openstack/manageNode/rabbitMqSummary](#11)
* [TopProcess By Cpu .GET        /v1/openstack/manageNode/:hostname/topProcessCpu](#12)
* [TopProcess By Memory .GET     /v1/openstack/manageNode/:hostname/topProcessMem](#13)

## Server API

<div id='10'/>

### **manage node summary** :  ``GET /v1/openstack/manageNode/summary``

#### Manage Node 자원상태 요약정보
    
   Openstack의 Manage Node별 자원사용량 요약정보를 조회한다.
   
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
        	<td><span>hostname</span></td>
           	<td><span>Manage Node HostName</span></td>
	    </tr>
            <tr class="">
        	<td><span>cpuUsage</span></td>
           	<td><span>Manage Node CPU 사용률</span></td>
	    </tr>
	        <tr class="">
            <td><span>memoryUsage</span></td>
            <td><span>Manage Node Memory 사용률</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>memoryMbMax</span></td>
            <td><span>Manage Node Memory 총 용량(MB)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>memoryUsedMb</span></td>
            <td><span>Manage Node의 사용중인 Memory 용량(MB)</span></td>
        </tr>        
        </tr>
            <tr class="">
            <td><span>diskGbMax</span></td>
            <td><span>Manage Node의 Disk 총용량(GB)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>diskGbUsed</span></td>
            <td><span>Manage Node의 사용중인 Disk 용량(GB)</span></td>
        </tr>        
            <tr class="">
            <td><span>agentStatus</span></td>
            <td><span>Compute Node에 설치된 Metric Agent의 상태</span></td>
        </tr>
	</tbody>
</table>


```
[request]
Manage Node 자원상태 요약정보 (parameter없음)
curl -X GET -G 'http://localhost:8080/v1/openstack/manageNode/summary'
```

##### response
```
[  
   {  
      "hostname":"block1",
      "cpuUsage":11.9,
      "memoryUsage":64.54,
      "memoryMbMax":128911,
      "memoryUsedMb":83193,
      "diskGbMax":104.7314453125,
      "diskGbUsed":4.3837890625,
      "agentStatus":"OK"
   },
   {  
      "hostname":"controller",
      "cpuUsage":1.2,
      "memoryUsage":63.94,
      "memoryMbMax":128751,
      "memoryUsedMb":82325,
      "diskGbMax":331.505859375,
      "diskGbUsed":59.4697265625,
      "agentStatus":"OK"
   }
]
```


<div id='11'/>

### **rabbitMqSummarysummary** :  ``GET /v1/openstack/manageNode/rabbitMqSummary``

#### Manage Node 자원상태 요약정보

   Openstack rabbitMq의 자원사용량 요약정보를 조회한다.

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
        	<td><span>connections</span></td>
           	<td><span>Connection수</span></td>
	    </tr>
            <tr class="">
        	<td><span>channels</span></td>
           	<td><span>channels수</span></td>
	    </tr>
	        <tr class="">
            <td><span>queues</span></td>
            <td><span>queues 수</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>exchanges</span></td>
            <td><span>exchanges 수</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>diskMbLimitFree</span></td>
            <td><span>disk Limit Free 용량(MB)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>diskMbFree</span></td>
            <td><span>diskMbFree 용량(MB)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>memoryMbLimit</span></td>
            <td><span>Memory 총 용량(MB)</span></td>
        </tr>
            <tr class="">
            <td><span>memoryMbUsed</span></td>
            <td><span>사용중인 Memory 용량(MB)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>fileDescriptorTotal</span></td>
            <td><span>file Descriptor 총수</span></td>
        </tr>
            <tr class="">
            <td><span>fileDescriptorUsed</span></td>
            <td><span>사용중인 file Descriptor수</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>erlangProcessLimit</span></td>
            <td><span>사용가능한 Process수</span></td>
        </tr>
            <tr class="">
            <td><span>erlangProcessUsed</span></td>
            <td><span>사용중인 Process수</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>socketsLimit</span></td>
            <td><span>Socket 이 사용가능한 File Descriptor 수</span></td>
        </tr>
            <tr class="">
            <td><span>socketsUsed</span></td>
            <td><span>Socket 이 사용중인 File Descriptor 수</span></td>
        </tr>
	</tbody>
</table>


```
[request]
Manage Node 자원상태 요약정보 (parameter없음)
curl -X GET -G 'http://localhost:8080/v1/openstack/manageNode/rabbitMqSummary'
```

##### response
```
{
   "connections":316,
   "channels":316,
   "queues":196,
   "consumers":272,
   "exchanges":118,
   "NodeResources":{
      "diskMbFree":261639.84375,
      "diskMbLimitFree":47.6837158203125,
      "memoryMbUsed":792.1396942138672,
      "memoryMbLimit":51500.54999923706,
      "fileDescriptorUsed":342,
      "fileDescriptorTotal":1024,
      "erlangProcessUsed":4079,
      "erlangProcessLimit":1048576,
      "socketsUsed":317,
      "socketsLimit":829
   }
}
```

<div id='12'/>

### **Top Process By Cpu** :  ``GET /v1/openstack/manageNode/:hostname/topProcessCpu``

#### Node Top Process By Cpu

   Openstack Node에서 실행중인 Top Process 정보를 조회한다.

<h4>Request Parameters</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
            	<th>Optional</th>
            	<th>Example Values</th>
	</tr>
            <tbody>
                <tr class="">
                    <td><span>hostName</span></td>
                    <td><span>Node Host Name</span></td>
                    <td><span>require</span></td>
                    <td><span>ex)controller</span></td>
                </tr>
            </tbody>
    </thead>
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
        	<td><span>index</span></td>
           	<td><span>Top Process Order</span></td>
	    </tr>
            <tr class="">
        	<td><span>processName</span></td>
           	<td><span>Process 명</span></td>
	    </tr>
	        <tr class="">
            <td><span>useage</span></td>
            <td><span>사용률</span></td>
        </tr>
	</tbody>
</table>


```
[request]
Node에서 실행중인 Top Process 정보
curl -X GET -G 'http://localhost:8080/v1/openstack/manageNode/controller/topProcessCpu'
```

##### response
```
[
    {
        "index": 1,
        "processName": "rabbitmq",
        "usage": 11.333333333333334
    },
    {
        "index": 2,
        "processName": "nova-conductor",
        "usage": 9.171428571428569
    },
    ...
]

```


<div id='13'/>

### **Top Process By Memory** :  ``GET /v1/openstack/manageNode/:hostname/topProcessMem``

#### Node Top Process By Memory

   Openstack Node에서 실행중인 Top Process 정보를 조회한다.

<h4>Request Parameters</h4>
<table>
    <thead>
     	<tr>
        	<th>Name</th>
           	<th>Description</th>
            	<th>Optional</th>
            	<th>Example Values</th>
	</tr>
            <tbody>
                <tr class="">
                    <td><span>hostName</span></td>
                    <td><span>Node Host Name</span></td>
                    <td><span>require</span></td>
                    <td><span>ex)controller</span></td>
                </tr>
            </tbody>
    </thead>
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
        	<td><span>index</span></td>
           	<td><span>Top Process Order</span></td>
	    </tr>
            <tr class="">
        	<td><span>processName</span></td>
           	<td><span>Process 명</span></td>
	    </tr>
	        <tr class="">
            <td><span>useage</span></td>
            <td><span>사용률(MB)</span></td>
        </tr>
	</tbody>
</table>


```
[request]
Node에서 실행중인 Top Process 정보
curl -X GET -G 'http://localhost:8080/v1/openstack/manageNode/controller/topProcessMem'
```

##### response
```
[
    {
        "index": 1,
        "processName": "nova-api",
        "usage": 9870
    },
    {
        "index": 2,
        "processName": "cinder-api",
        "usage": 3778
    },
    ...
]
```

