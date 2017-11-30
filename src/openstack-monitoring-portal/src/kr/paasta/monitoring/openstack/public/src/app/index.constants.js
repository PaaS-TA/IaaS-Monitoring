/* global malarkey:false, moment:false */
(function() {
  'use strict';

  var apiHost = "http://localhost:8080";
  var apiUris = {
    summary:                      apiHost + "/v1/openstack/summary",
    computeNodeSummary:           apiHost + "/v1/openstack/computeNode/summary",
    manageNodeSummary:            apiHost + "/v1/openstack/manageNode/summary",
    projectSummary:               apiHost + "/v1/openstack/projects/summary",

    computeTopProcessByCpu:       apiHost + "/v1/openstack/manageNode/:hostname/topProcessCpu",
    computeTopProcessByMemory:    apiHost + "/v1/openstack/manageNode/:hostname/topProcessMem",
    computeRabbitMqSummary:       apiHost + "/v1/openstack/manageNode/rabbitMqSummary",

    manageTopProcessByCpu:        apiHost + "/v1/openstack/manageNode/:hostname/topProcessCpu",
    manageTopProcessByMemory:     apiHost + "/v1/openstack/manageNode/:hostname/topProcessMem",
    manageRabbitMqSummary:        apiHost + "/v1/openstack/manageNode/rabbitMqSummary",

    projectInstanceList:          apiHost + "/v1/openstack/projects/:projectId/instances",

    nodeList:                     apiHost + "/v1/openstack/node",
    nodeCpuUsageList:             apiHost + "/v1/openstack/node/:hostname/cpuUsage",
    nodeCpuLoad1mList:            apiHost + "/v1/openstack/node/:hostname/cpuLoad",
    nodeMemorySwapList:           apiHost + "/v1/openstack/node/:hostname/swapUsage",
    nodeMemoryUsageList:          apiHost + "/v1/openstack/node/:hostname/memUsage",
    nodeDiskUsageList:            apiHost + "/v1/openstack/node/:hostname/diskUsage",
    nodeDiskIOReadList:           apiHost + "/v1/openstack/node/:hostname/diskRead",
    nodeDiskIOWriteList:          apiHost + "/v1/openstack/node/:hostname/diskWrite",
    nodeNetworkIOKByteList:       apiHost + "/v1/openstack/node/:hostname/networkIo",
    nodeNetworkErrorList:         apiHost + "/v1/openstack/node/:hostname/networkError",
    nodeNetworkDroppedPacketList: apiHost + "/v1/openstack/node/:hostname/networkDropPacket",
    nodeRabbitMQList:             apiHost + "/v1/openstack/node/:hostname/rabbitMqSummary",

    instanceCpuUsageList:         apiHost + "/v1/openstack/projects/:instanceId/cpuUsage",
    instanceMemoryUsageList:      apiHost + "/v1/openstack/projects/:instanceId/memUsage",
    instanceDiskIOReadList:       apiHost + "/v1/openstack/projects/:instanceId/diskRead",
    instanceDiskIOWriteList:      apiHost + "/v1/openstack/projects/:instanceId/diskWrite",
    instanceNetworkIOKByteList:   apiHost + "/v1/openstack/projects/:instanceId/networkIo",
    instanceNetworkPacketList:    apiHost + "/v1/openstack/projects/:instanceId/networkPackets",

    defaultRecentLogs:            apiHost + "/v1/openstack/log/recent",
    specificTimeRangeLogs:        apiHost + "/v1/openstack/log/specific",

    alarmNotification:            apiHost + "/v1/alarm/notification",
    alarmNotificationId:          apiHost + "/v1/alarm/notification/:id",
    alarmDefinition:              apiHost + "/v1/alarm/definition",
    alarmDefinitionId:            apiHost + "/v1/alarm/definition/:id",
    alarmStatus:                  apiHost + "/v1/alarm/status",
    alarmStatusCount:             apiHost + "/v1/alarm/status/count",
    alarmStatusId:                apiHost + "/v1/alarm/:alarmId/status",
    alarmStatusHistory:           apiHost + "/v1/alarm/:alarmId/history",
    alarmActionList:              apiHost + "/v1/alarm/:alarmId/action",
    alarmAction:                  apiHost + "/v1/alarm/action",
    alarmActionId:                apiHost + "/v1/alarm/action/:id",

    ping:                         apiHost + "/v1/ping",
    login:                        apiHost + "/v1/login",
    logout:                       apiHost + "/v1/logout"

  };

  var nodeChartConfig = [
    {id: 1, name: 'CPU Usage', func: 'nodeCpuUsageList', type: 'lineChart', percent: true, axisLabel: '%'},
    {id: 2, name: 'CPU Load Average', func: 'nodeCpuLoad1mList', type: 'lineChart', percent: false, axisLabel: 'Count per 1 minute'},
    {id: 3, name: 'Swap Usage', func: 'nodeMemorySwapList', type: 'lineChart', percent: true, axisLabel: '%'},
    {id: 4, name: 'Memory Usage', func: 'nodeMemoryUsageList', type: 'lineChart', percent: true, axisLabel: '%'},
    {id: 5, name: 'Disk Usage', func: 'nodeDiskUsageList', type: 'lineChart', percent: true, axisLabel: '%'},
    {id: 6, name: 'Disk IO Read', func: 'nodeDiskIOReadList', type: 'lineChart', percent: false, axisLabel: 'KB'},
    {id: 7, name: 'Disk IO Write', func: 'nodeDiskIOWriteList', type: 'lineChart', percent: false, axisLabel: 'KB'},
    {id: 8, name: 'Network IO KByte', func: 'nodeNetworkIOKByteList', type: 'lineChart', percent: false, axisLabel: 'KB'},
    {id: 9, name: 'Network Error', func: 'nodeNetworkErrorList', type: 'lineChart', percent: false, axisLabel: 'Count'},
    {id: 10, name: 'Network Dropped Packet', func: 'nodeNetworkDroppedPacketList', type: 'lineChart', percent: false, axisLabel: 'Count'}
  ];

  var projectChartConfig = [
    {id: 1, name: 'CPU Usage', func: 'instanceCpuUsageList', type: 'lineChart', percent: true, axisLabel: '%'},
    {id: 2, name: 'Memory Usage', func: 'instanceMemoryUsageList', type: 'lineChart', percent: true, axisLabel: '%'},
    {id: 3, name: 'Disk IO Read', func: 'instanceDiskIOReadList', type: 'lineChart', percent: false, axisLabel: 'KB'},
    {id: 4, name: 'Disk IO Write', func: 'instanceDiskIOWriteList', type: 'lineChart', percent: false, axisLabel: 'KB'},
    {id: 5, name: 'Network IO KByte', func: 'instanceNetworkIOKByteList', type: 'lineChart', percent: false, axisLabel: 'KB'},
    {id: 6, name: 'Network IO Packet', func: 'instanceNetworkPacketList', type: 'lineChart', percent: false, axisLabel: 'Count'}
  ];

  var constants = {
    version: '0.0.1',
    expire: '0.0.1'
  };

  angular
    .module('monitoring')
    .constant('malarkey', malarkey)
    .constant('moment', moment)
    .constant('apiUris', apiUris)
    .constant('nodeChartConfig', nodeChartConfig)
    .constant('projectChartConfig', projectChartConfig)
    .constant('constants', constants);

})();
