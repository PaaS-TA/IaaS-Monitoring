# Server API
* [project_summary .GET         /v1/openstack/projects/summaries](#10)
* [project_Instances .GET       /v1/openstack/projects/:projectId/instances"](#11)
* [Instance CpuUsage .GET       /v1/openstack/projects/:instanceId/cpuUsage](#12)
* [Instance MemoryUsage .GET    /v1/openstack/projects/:instanceId/memUsage](#13)
* [Instance DiskRead KByte .GET  /v1/openstack/projects/:instanceId/diskRead](#14)
* [Instance DiskWrite KByte .GET  /v1/openstack/projects/:instanceId/diskWrite](#15)
* [Instance Network IO KByte .GET  /v1/openstack/projects/:instanceId/networkIo](#16)
* [Instance Network IO Packet .GET  /v1/openstack/projects/:instanceId/networkPackets](#17)


## Server API

<div id='10'/>

### **project Summary** :  ``GET /v1/openstack/projects/summary``

#### Project 목록 및 요약정보를 조회한다.
    
   Openstack의 Project 목록 및 요약정보를 조회한다.
   
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
        	<td><span>id</span></td>
           	<td><span>project Id</span></td>
	    </tr>
            <tr class="">
        	<td><span>name</span></td>
           	<td><span>project Name</span></td>
	    </tr>
	        <tr class="">
            <td><span>description</span></td>
            <td><span>Project Description</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>enabled</span></td>
            <td><span>Project Enabled(true/false)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>instancesLimit</span></td>
            <td><span>Project에서 사용가능한 Instance 수(Quota)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>instancesUsed</span></td>
            <td><span>Project에서 사용중인 Instance 수</span></td>
        </tr>        
        </tr>
            <tr class="">
            <td><span>vcpusLimit</span></td>
            <td><span>Project에서 사용가능한 vcpu 수(Quota)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>vcpusUsed</span></td>
            <td><span>Project에서 사용중인 vcpu 수</span></td>
        </tr>
        </tr>        
            <tr class="">
            <td><span>memoryMbLimit</span></td>
            <td><span>Project에서 사용가능한 Memory양(Quota)</span></td>
        </tr>
        </tr>
            <tr class="">
                <td><span>memoryMbUsed</span></td>
                <td><span>Project에서 사용중인 Memory양(MB)</span></td>
        </tr>
        </tr>
             <tr class="">
                <td><span>floatingIpsLimit</span></td>
                <td><span>Project에서 사용가능한 FloatingIP 수(Quota)</span></td>
        </tr>
        </tr>
             <tr class="">
                <td><span>floatingIpsUsed</span></td>
                <td><span>Project에서 사용중인 FloatingIP 수</span></td>
        </tr>
        </tr>
             <tr class="">
                <td><span>securityGroupsLimit</span></td>
                <td><span>Project에서 사용가능한 securityGroups 수(Quota)</span></td>
        </tr>
        </tr>
             <tr class="">
                <td><span>securityGroupsUsed</span></td>
                <td><span>Project에서 사용중인 securityGroups 수</span></td>
        </tr>
        </tr>
             <tr class="">
                <td><span>volumeStorageLimit</span></td>
                <td><span>Project에서 사용가능한 volumeStorage 수(Quota)</span></td>
        </tr>
        </tr>
             <tr class="">
                <td><span>volumeStorageUused</span></td>
                <td><span>Project에서 사용중인 volumeStorage 수</span></td>
        </tr>
        </tr>
             <tr class="">
                <td><span>volumeStorageLimitGb</span></td>
                <td><span>Project에서 사용가능한 VolumnStorage 양(GB) </span></td>
        </tr>
        </tr>
             <tr class="">
                <td><span>volumeStorageUsedGb</span></td>
                <td><span>Project에서 사용중인 volumeStorage양(GB) </span></td>
        </tr>
	</tbody>
</table>


```
[request]
Openstack의 Project 목록 및 요약정보를 조회한다. (parameter없음)
curl -X GET -G 'http://localhost:8080/v1/openstack/projects/summary'
```

##### response
```
[
    {
        "description": "",
        "enabled": true,
        "floatingIpsLimit": 100,
        "floatingIpsUsed": 10,
        "id": "9c1a27e20412473b843dbf32bdec2390",
        "instancesLimit": 150,
        "instancesUsed": 42,
        "memoryMbLimit": 182400,
        "memoryMbUsed": 141312,
        "name": "admin",
        "securityGroupsLimit": 30,
        "securityGroupsUsed": 3,
        "vcpusLimit": 150,
        "vcpusUsed": 63,
        "volumeStorageLimit": 30,
        "volumeStorageLimitGb": 1000,
        "volumeStorageUsed": 16,
        "volumeStorageUsedGb": 253
    },
    {
        "description": "",
        "enabled": true,
        "floatingIpsLimit": 2,
        "floatingIpsUsed": 0,
        "id": "da9572dd8c0a4857aa12d9c4cbb76306",
        "instancesLimit": 15,
        "instancesUsed": 1,
        "memoryMbLimit": 10240,
        "memoryMbUsed": 512,
        "name": "demo",
        "securityGroupsLimit": 5,
        "securityGroupsUsed": 2,
        "vcpusLimit": 50,
        "vcpusUsed": 1,
        "volumeStorageLimit": 10,
        "volumeStorageLimitGb": 100,
        "volumeStorageUsed": 0,
        "volumeStorageUsedGb": 0
    }
]
```


<div id='11'/>

### **project Instances** :  ``GET /v1/openstack/projects/:projectId/instances``

#### project에 속한 Instance 목록 및 자원상태 요약정보 조회

   Openstack project에 속한 Instance 목록 및 자원상태 요약정보 조회

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
                    <td><span>projectId</span></td>
                    <td><span>project Id</span></td>
                    <td><span>require</span></td>
                    <td><span>ex)c5e1012f-9455-47ff-bbb7-1cff5e1f9894</span></td>
                </tr>
                <tr class="">
                    <td><span>limit</span></td>
                    <td><span>page Size</span></td>
                    <td><span>require</span></td>
                    <td><span>ex)10</span></td>
                </tr>
                <tr class="">
                    <td><span>marker</span></td>
                    <td><span>paging 처리시 다음 Page요청시 현재 Page의 마지막 instanceId</span></td>
                    <td><span>option</span></td>
                    <td><span>ex)9c1a27e20412473b843dbf32bdec2390</span></td>
                </tr>
                <tr class="">
                    <td><span>hostname</span></td>
                    <td><span>조회하려는 instance 명</span></td>
                    <td><span>option</span></td>
                    <td><span>ex)monasca</span></td>
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
        	<td><span>instanceId</span></td>
           	<td><span>instance(vm) Id</span></td>
	    </tr>
            <tr class="">
        	<td><span>name</span></td>
           	<td><span>instance(vm) Name</span></td>
	    </tr>
	        <tr class="">
            <td><span>cpuUsage</span></td>
            <td><span>cpu 사용률</span></td>
        </tr>
            <tr class="">
            <td><span>memoryUsage</span></td>
            <td><span>memory 사용률</span></td>
        </tr>
            <tr class="">
            <td><span>vcpus</span></td>
            <td><span>할당된 vcpu 수</span></td>
        </tr>
        <tr class="">
            <td><span>memoryMb</span></td>
            <td><span>할당된 Memory(MB)</span></td>
        </tr>
        <tr class="">
            <td><span>disk_gb</span></td>
            <td><span>할당된 Disk(GB)</span></td>
        </tr>
        <tr class="">
            <td><span>flavor</span></td>
            <td><span>flavorType</span></td>
        </tr>
        <tr class="">
            <td><span>tenant_id</span></td>
            <td><span>projectid</span></td>
        </tr>
        <tr class="">
            <td><span>zone</span></td>
            <td><span>zone name</span></td>
        </tr>
        <tr class="">
            <td><span>state</span></td>
            <td><span>상태</span></td>
        </tr>
        <tr class="">
            <td><span>upTime</span></td>
            <td><span>Instance start Time</span></td>
        </tr>
        <tr class="">
            <td><span>startedAt</span></td>
            <td><span>Instance 생성일시</span></td>
        </tr>
        <tr class="">
            <td><span>projectId</span></td>
            <td><span>Project ID</span></td>
        </tr>
        <tr class="">
            <td><span>totoalCnt</span></td>
            <td><span>instance count in a project</span></td>
        </tr>
	</tbody>
</table>


```
[request]
Project내 Instance 정보 조회
curl -X GET -G 'http://localhost:8080/v1/openstack/projects/da9572dd8c0a4857aa12d9c4cbb76306/instances?limit=2'
curl -X GET -G 'http://localhost:8080/v1/openstack/projects/da9572dd8c0a4857aa12d9c4cbb76306/instances?limit=2&hostname=monasca'
```

##### response
```
{
    "metric": [
        {
            "address": [
                "115.68.151.163"
            ],
            "cpuUsage": 3,
            "disk_gb": 1,
            "ended_at": "",
            "flavor": "m1.tiny",
            "instance_id": "c5e1012f-9455-47ff-bbb7-1cff5e1f9894",
            "memoryUsage": 0,
            "memory_mb": 512,
            "name": "demo-test",
            "started_at": "2017-07-06T04:27:23",
            "state": "ACTIVE",
            "tenant_id": "da9572dd8c0a4857aa12d9c4cbb76306",
            "uptime": 4226153.0,
            "vcpus": 1,
            "zone": "nova"
        }
    ],
    "projectId": "da9572dd8c0a4857aa12d9c4cbb76306",
    "totalCnt": 1
    }
```


<div id='12'/>

### **Instance Cpu Usage** :  ``GET /v1/openstack/projects/:instanceId/cpuUsage``

#### Instance CPU Usage

   Instance Cpu Usage 정보를 조회한다.

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
                    <td><span>instanceId</span></td>
                    <td><span>instanceId</span></td>
                    <td><span>required</span></td>
                    <td><span>ex)5e1012f-9455-47ff-bbb7-1cff5e1f9894</span></td>
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
           	<td><span>Cpu 사용률</span></td>
	    </tr>
	</tbody>
</table>


```
[request]
Instance의 CPU 사용률 정보 조회
curl -X GET -G 'http://localhost:8080/v1/openstack/projects/929c08d7-b9e1-48c5-a49b-7ab9b10e9f3c/cpuUsage?defaultTimeRange=10m&groupBy=1m'
curl -X GET -G 'http://localhost:8080/v1/openstack/projects/929c08d7-b9e1-48c5-a49b-7ab9b10e9f3c/cpuUsage?timeRangeFrom=10m&timeRangeTo=30m&groupBy=1m'
```

##### response
```
[
   {
      "metric":[
         {
            "time":1502330400,
            "usage": 12
         },
         {
            "time":1502330460,
            "usage": 12
         },
         ...
      ],
      "name":"cpu"
   }
]
```


<div id='13'/>

### **Instance Memory Usage** :  ``GET /v1/openstack/projects/:instanceId/memUsage``

#### Instance Memory Usage

   Instance Memory Usage 정보를 조회한다.

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
                    <td><span>instanceId</span></td>
                    <td><span>instanceId</span></td>
                    <td><span>required</span></td>
                    <td><span>ex)5e1012f-9455-47ff-bbb7-1cff5e1f9894</span></td>
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
Instance의 Memory 사용률 정보 조회
curl -X GET -G 'http://localhost:8080/v1/openstack/projects/929c08d7-b9e1-48c5-a49b-7ab9b10e9f3c/memUsage?defaultTimeRange=10m&groupBy=1m'
curl -X GET -G 'http://localhost:8080/v1/openstack/projects/929c08d7-b9e1-48c5-a49b-7ab9b10e9f3c/memUsage?timeRangeFrom=10m&timeRangeTo=30m&groupBy=1m'
```

##### response
```
[
   {
      "metric":[
         {
            "time":1502330400,
            "usage": 12
         },
         {
            "time":1502330460,
            "usage": 12
         },
         ...
      ],
      "name":"cpu"
   }
]
```

<div id='14'/>

### **Instance Disk Read Kbyte** :  ``GET /v1/openstack/projects/:instanceId/diskRead``

#### Instance Disk Read(Kbyte)

   Instance Disk Read(Kbyte)정보를 조회한다.

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
                    <td><span>instanceId</span></td>
                    <td><span>instance Id</span></td>
                    <td><span>required</span></td>
                    <td><span>ex)5e1012f-9455-47ff-bbb7-1cff5e1f9894</span></td>
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
           	<td><span>Disk Read Kbyte</span></td>
	    </tr>
	</tbody>
</table>


```
[request]
Instance의 Disk Read Kbyte 정보 조회
curl -X GET -G 'http://localhost:8080/v1/openstack/projects/929c08d7-b9e1-48c5-a49b-7ab9b10e9f3c/diskRead?defaultTimeRange=10m&groupBy=1m'
curl -X GET -G 'http://localhost:8080/v1/openstack/projects/929c08d7-b9e1-48c5-a49b-7ab9b10e9f3c/diskRead?timeRangeFrom=10m&timeRangeTo=30m&groupBy=1m'
```

##### response
```
[
   {
      "metric":[
         {
            "time":1502330400,
            "usage": 12
         },
         {
            "time":1502330460,
            "usage": 12
         },
         ...
      ],
      "name":"cpu"
   }
]
```


<div id='15'/>

### **Instance Disk Write Kbyte** :  ``GET /v1/openstack/projects/:instanceId/diskWrite``

#### Instance Disk Read(Kbyte)

   Instance Disk Read(Kbyte)정보를 조회한다.

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
                    <td><span>instanceId</span></td>
                    <td><span>instance Id</span></td>
                    <td><span>required</span></td>
                    <td><span>ex)5e1012f-9455-47ff-bbb7-1cff5e1f9894</span></td>
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
           	<td><span>Disk Read Kbyte</span></td>
	    </tr>
	</tbody>
</table>


```
[request]
Instance의 Disk Write Kbyte 정보 조회
curl -X GET -G 'http://localhost:8080/v1/openstack/projects/929c08d7-b9e1-48c5-a49b-7ab9b10e9f3c/diskWrite?defaultTimeRange=10m&groupBy=1m'
curl -X GET -G 'http://localhost:8080/v1/openstack/projects/929c08d7-b9e1-48c5-a49b-7ab9b10e9f3c/diskWrite?timeRangeFrom=10m&timeRangeTo=30m&groupBy=1m'
```

##### response
```
[
   {
      "metric":[
         {
            "time":1502330400,
            "usage": 12
         },
         {
            "time":1502330460,
            "usage": 12
         },
         ...
      ],
      "name":"cpu"
   }
]
```


<div id='16'/>

### **Instance Network In/Out Kbyte** :  ``GET /v1/openstack/projects/:instanceId/networkIo``

#### Instance Disk Read(Kbyte)

   Instance Network In/Out Kbyte 정보를 조회한다.

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
                    <td><span>instanceId</span></td>
                    <td><span>instance Id</span></td>
                    <td><span>required</span></td>
                    <td><span>ex)5e1012f-9455-47ff-bbb7-1cff5e1f9894</span></td>
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
           	<td><span>Network In/Out Kbyte</span></td>
	    </tr>
	</tbody>
</table>


```
[request]
Instance의 Network IO Kbyte 정보 조회
curl -X GET -G 'http://localhost:8080/v1/openstack/projects/929c08d7-b9e1-48c5-a49b-7ab9b10e9f3c/networkIo?defaultTimeRange=10m&groupBy=1m'
curl -X GET -G 'http://localhost:8080/v1/openstack/projects/929c08d7-b9e1-48c5-a49b-7ab9b10e9f3c/networkIo?timeRangeFrom=10m&timeRangeTo=30m&groupBy=1m'
```

##### response
```
[
    {
        "metric": [
            {
                "time": 1503553020,
                "usage": 9.61
            },
            ...
        ],
        "name": "in"
    },
    {
        "metric": [
            {
                "time": 1503553020,
                "usage": 24.93
            },
            ...
        ],
        "name": "out"
    }
]

```




<div id='17'/>

### **Instance Network In/Out Packet** :  ``GET /v1/openstack/projects/:instanceId/networkPackets``

#### Instance Disk Read(Kbyte)

   Instance Network In/Out Packets 정보를 조회한다.

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
                    <td><span>instanceId</span></td>
                    <td><span>instance Id</span></td>
                    <td><span>required</span></td>
                    <td><span>ex)5e1012f-9455-47ff-bbb7-1cff5e1f9894</span></td>
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
           	<td><span>Network In/Out packets</span></td>
	    </tr>
	</tbody>
</table>


```
[request]
Instance의 Network IO Packet 정보 조회
curl -X GET -G 'http://localhost:8080/v1/openstack/projects/929c08d7-b9e1-48c5-a49b-7ab9b10e9f3c/networkPackets?defaultTimeRange=10m&groupBy=1m'
curl -X GET -G 'http://localhost:8080/v1/openstack/projects/929c08d7-b9e1-48c5-a49b-7ab9b10e9f3c/networkPackets?timeRangeFrom=10m&timeRangeTo=30m&groupBy=1m'
```

##### response
```
[
    {
        "metric": [
            {
                "time": 1503553020,
                "usage": 9.61
            },
            ...
        ],
        "name": "in"
    },
    {
        "metric": [
            {
                "time": 1503553020,
                "usage": 24.93
            },
            ...
        ],
        "name": "out"
    }
]

```
