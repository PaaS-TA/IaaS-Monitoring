(function() {
  'use strict';

  angular
    .module('monitoring')
    .factory('manageNodeService', ManageNodeService);

  /** @ngInject */
  function ManageNodeService($http, apiUris) {
    var service = {};

    service.manageNodeSummary = function(hostname){
      var config = {
        params: {'hostname': hostname},
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.manageNodeSummary, config);
    };

    service.manageTopProcessByCpu = function(condition){
      return $http.get(apiUris.manageTopProcessByCpu.replace(":hostname", condition.hostname));
    };

    service.manageTopProcessByMemory = function(condition){
      return $http.get(apiUris.manageTopProcessByMemory.replace(":hostname", condition.hostname));
    };

    service.manageRabbitMqSummary = function(){
      return $http.get(apiUris.manageRabbitMqSummary);
    };

    return service;
  }

  angular
    .module('monitoring')
    .factory('computeNodeService', ComputeNodeService);

  /** @ngInject */
  function ComputeNodeService($http, apiUris, common) {
    var service = {};

    service.computeNodeSummary = function(hostname){
      var config = {
        params: {'hostname': hostname},
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.computeNodeSummary, config);
    };

    service.nodeCpuUsageList = function(condition) {
      var config = {
        params: common.setDtvParam(condition),
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.nodeCpuUsageList.replace(":hostname", condition.hostname), config);
    };

    service.nodeCpuLoad1mList = function(condition) {
      var config = {
        params: common.setDtvParam(condition),
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.nodeCpuLoad1mList.replace(":hostname", condition.hostname), config);
    };

    service.nodeMemorySwapList = function(condition) {
      var config = {
        params: common.setDtvParam(condition),
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.nodeMemorySwapList.replace(":hostname", condition.hostname), config);
    };

    service.nodeMemoryUsageList = function(condition) {
      var config = {
        params: common.setDtvParam(condition),
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.nodeMemoryUsageList.replace(":hostname", condition.hostname), config);
    };

    service.nodeDiskUsageList = function(condition) {
      var config = {
        params: common.setDtvParam(condition),
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.nodeDiskUsageList.replace(":hostname", condition.hostname), config);
    };

    service.nodeDiskIOReadList = function(condition) {
      var config = {
        params: common.setDtvParam(condition),
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.nodeDiskIOReadList.replace(":hostname", condition.hostname), config);
    };

    service.nodeDiskIOWriteList = function(condition) {
      var config = {
        params: common.setDtvParam(condition),
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.nodeDiskIOWriteList.replace(":hostname", condition.hostname), config);
    };

    service.nodeNetworkIOKByteList = function(condition) {
      var config = {
        params: common.setDtvParam(condition),
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.nodeNetworkIOKByteList.replace(":hostname", condition.hostname), config);
    };

    service.nodeNetworkErrorList = function(condition) {
      var config = {
        params: common.setDtvParam(condition),
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.nodeNetworkErrorList.replace(":hostname", condition.hostname), config);
    };

    service.nodeNetworkDroppedPacketList = function(condition) {
      var config = {
        params: common.setDtvParam(condition),
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.nodeNetworkDroppedPacketList.replace(":hostname", condition.hostname), config);
    };

    service.nodeRabbitMQList = function(condition) {
      var config = {
        params: common.setDtvParam(condition),
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.nodeRabbitMQList.replace(":hostname", condition.hostname), config);
    };

    return service;
  }
})();
