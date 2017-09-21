# Server API
* [openstack_Summary .GET       /v1/openstack/summary](#10)


## Server API

<div id='10'/>
### **openstack Summary** :  ``GET /v1/openstack/summary``

#### openstack Hypervisor 자원상태 요약정보
    
   Openstack의 Compute Node들의 자원사용량 요약정보를 조회한다.
   
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
        	<td><span>totalVms</span></td>
           	<td><span>Openstack의 전체 Instance 수</span></td>
	    </tr>
            <tr class="">
        	<td><span>runningVms</span></td>
           	<td><span>Openstack Instance중 Running중인 Instance수</span></td>
	    </tr>
            <tr class="">
        	<td><span>totalVcpus</span></td>
           	<td><span>Hypervisoer 전체 VCPU수</span></td>
	    </tr>
	        <tr class="">
            <td><span>usedVcpus</span></td>
            <td><span>Hypervisoer에서 사용중인 VCPU수</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>totalMemoryMb</span></td>
            <td><span>Hypervisoer에서 전체 Memory 용량(MB)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>memoryUsedMb</span></td>
            <td><span>Hypervisoer에서 사용중인 Memory 용량(MB)</span></td>
        </tr>        
        </tr>
            <tr class="">
            <td><span>freeMemoryMb</span></td>
            <td><span>Hypervisoer에서 사용가능한 Memory 용량(MB)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>totalDiskGb</span></td>
            <td><span>Hypervisoer에서 전체 Disk 용량(GB)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>usedDiskGb</span></td>
            <td><span>Hypervisoer에서 사용중인 Disk 용량(GB)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>freeDiskGb</span></td>
            <td><span>Hypervisoer에서 사용가능안 Disk용량(GB)</span></td>
        </tr>
        </tr>
            <tr class="">
            <td><span>diskAvailableLeastGb</span></td>
            <td><span>Hypervisoer에서 실제로 사용가능안 Disk용량(GB)</span></td>
        </tr>
	</tbody>
</table>


```
[request]
openstack Hypervisor 자원상태 요약정보 조회 (parameter없음)
curl -X GET -G 'http://localhost:8080/v1/openstack/summary'
```

##### response
```
{  
   "totalVms":44,
   "runningVms":43,
   "totalVcpus":64,
   "usedVcpus":68,
   "totalMemoryMb":257822,
   "memoryUsedMb":151040,
   "freeMemoryMb":106782,
   "totalDiskGb":1576,
   "usedDiskGb":1171,
   "freeDiskGb":405,
   "diskAvailableLeastGb":307
}
```

