(function() {
  'use strict';

  angular
    .module('monitoring')
    .factory('projectService', ProjectService);

  /** @ngInject */
  function ProjectService($http, apiUris, common) {
    var service = {};

    service.projectSummary = function(projectName){
      var config = {
        params: {'projectName': projectName},
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.projectSummary, config);
    };

    service.projectInstanceList = function(id, params){
      var config = {
        params: params,
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.projectInstanceList.replace(":projectId", id), config);
    };

    service.instanceCpuUsageList = function(condition) {
      var config = {
        params: common.setDtvParam(condition),
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.instanceCpuUsageList.replace(":instanceId", condition.instanceId), config);
    };

    service.instanceMemoryUsageList = function(condition) {
      var config = {
        params: common.setDtvParam(condition),
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.instanceMemoryUsageList.replace(":instanceId", condition.instanceId), config);
    };

    service.instanceDiskIOReadList = function(condition) {
      var config = {
        params: common.setDtvParam(condition),
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.instanceDiskIOReadList.replace(":instanceId", condition.instanceId), config);
    };

    service.instanceDiskIOWriteList = function(condition) {
      var config = {
        params: common.setDtvParam(condition),
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.instanceDiskIOWriteList.replace(":instanceId", condition.instanceId), config);
    };

    service.instanceNetworkIOKByteList = function(condition) {
      var config = {
        params: common.setDtvParam(condition),
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.instanceNetworkIOKByteList.replace(":instanceId", condition.instanceId), config);
    };

    service.instanceNetworkPacketList = function(condition) {
      var config = {
        params: common.setDtvParam(condition),
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.instanceNetworkPacketList.replace(":instanceId", condition.instanceId), config);
    };

    return service;
  }
})();
