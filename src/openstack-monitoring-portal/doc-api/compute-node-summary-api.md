# Server API
* [compute_node_summary .GET       /v1/openstack/compute_node/summary](#10)


## Server API

<div id='10'/>
### **compute node summary** :  ``GET /v1/openstack/compute_node/summary``

#### Compute Node 자원상태 요약정보
    
   Openstack의 Compute Node별 자원사용량 요약정보를 조회한다.
   
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
        	<td><span>nodeId</span></td>
           	<td><span>Compute Node ID</span></td>
	    </tr>
            <tr class="">
        	<td><span>hostname</span></td>
           	<td><span>Compute Node HostName</span></td>
	    </tr>
            <tr class="">
        	<td><span>hostIp</span></td>
           	<td><span>Compute Node Host IP</span></td>
	    </tr>
	        <tr class="">
            <td><span>type</span></td>
            <td><span>Hypervisoer Type</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>vcpusMax</span></td>
            <td><span>Hypervisor의 VCPU 총 개수</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>vcpusUsed</span></td>
            <td><span>Hypervisor에서 사용중인 VCPU 개수</span></td>
        </tr>        
        </tr>
            <tr class="">
            <td><span>memoryMbMax</span></td>
            <td><span>Hypervisor의 총 Memory용량(MB)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>memoryMbUsed</span></td>
            <td><span>Hypervisor의 사용중인 Memory용량(MB)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>memoryMbFree</span></td>
            <td><span>Hypervisor의 사용가능한 Memory용량(MB)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>diskGbMax</span></td>
            <td><span>Hypervisor의 Disk 총 용량(GB)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>diskGbUsed</span></td>
            <td><span>Hypervisor의 사용중인 Disk 용량(GB)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>diskGbFree</span></td>
            <td><span>Hypervisor의 사용가능한 Disk 용량(GB)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>diskAbilableLeast</span></td>
            <td><span>Hypervisor의 실제로 사용가능한 Disk 용량(GB)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>state</span></td>
            <td><span>Hypervisor의 상태(up/down)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>status</span></td>
            <td><span>Hypervisor의 사용가능 상태(enabled/diabled)</span></td>
        </tr>
            <tr class="">
            <td><span>totalVms</span></td>
            <td><span>Hypervisor 전체 Instance 수</span></td>
        </tr>
            <tr class="">
            <td><span>runningVms</span></td>
            <td><span>Hypervisor의 Running중인 Instance수</span></td>
        </tr>
            <tr class="">
            <td><span>cpuUsage</span></td>
            <td><span>Compute Node의 Cpu사용률</span></td>
        </tr>
            <tr class="">
            <td><span>memUsage</span></td>
            <td><span>Compute Node의 Memroy사용률</span></td>
        </tr>
            <tr class="">
            <td><span>agentStatus</span></td>
            <td><span>Compute Node에 설치된 Metric Agent의 상태</span></td>
        </tr>
	</tbody>
</table>


```
[request]
Compute Node 자원상태 요약정보 조회 (parameter없음)
curl -X GET -G 'http://localhost:8080/v1/openstack/compute_node/summary'
```

##### response
```
[  
   {  
      "nodeId":1,
      "hostname":"compute2",
      "hostIp":"10.10.1.173",
      "type":"QEMU",
      "vcpusMax":32,
      "vcpusUsed":31,
      "memoryMbMax":128911,
      "memoryMbUsed":72704,
      "memoryMbFree":56207,
      "diskGbMax":788,
      "diskGbUsed":571,
      "diskGbFree":217,
      "diskAbilableLeast":168,
      "state":"up",
      "status":"enabled",
      "totalVms":21,
      "runningVms":21,
      "cpuUsage":28.9,
      "memUsage":48.06,
      "agentStatus":"OK"
   },
   {  
      "nodeId":2,
      "hostname":"compute1",
      "hostIp":"10.10.1.174",
      "type":"QEMU",
      "vcpusMax":32,
      "vcpusUsed":37,
      "memoryMbMax":128911,
      "memoryMbUsed":78336,
      "memoryMbFree":50575,
      "diskGbMax":788,
      "diskGbUsed":600,
      "diskGbFree":188,
      "diskAbilableLeast":139,
      "state":"up",
      "status":"enabled",
      "totalVms":23,
      "runningVms":22,
      "cpuUsage":21.7,
      "memUsage":67.46,
      "agentStatus":"OK"
   }
]
```

