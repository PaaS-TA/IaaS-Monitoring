(function() {
  'use strict';

  angular
    .module('monitoring')
    .factory('alarmNotificationService', alarmNotificationService);

  /** @ngInject */
  function alarmNotificationService($http, apiUris) {
    var service = {};

    service.alarmNotificationList = function(params){
      var config = {
        params: params,
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.alarmNotification, config);
    };

    service.insertAlarmNotification = function(data){
      var config = {
        headers : {'Content-Type': 'application/json', 'Accept': 'application/json'}
      };
      return $http.post(apiUris.alarmNotification, data, config);
    };

    service.updateAlarmNotification = function(data){
      var config = {
        headers : {'Content-Type': 'application/json', 'Accept': 'application/json'}
      };
      return $http.put(apiUris.alarmNotificationId.replace(":id", data.id), data, config);
    };

    service.deleteAlarmNotification = function(id){
      var config = {
        headers : {'Accept' : 'application/json'}
      };
      return $http.delete(apiUris.alarmNotificationId.replace(":id", id), config);
    };

    return service;
  }


  angular
    .module('monitoring')
    .factory('alarmDefinitionService', alarmDefinitionService);

  /** @ngInject */
  function alarmDefinitionService($http, apiUris) {
    var service = {};

    service.alarmDefinitionList = function(params){
      var config = {
        params: params,
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.alarmDefinition, config);
    };

    service.alarmDefinition = function(id){
      var config = {
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.alarmDefinitionId.replace(":id", id), config);
    };

    service.insertAlarmDefinition = function(data){
      var config = {
        headers : {'Accept' : 'application/json'}
      };
      return $http.post(apiUris.alarmDefinition, data, config);
    };

    service.updateAlarmDefinition = function(data){
      var config = {
        headers : {'Content-Type': 'application/json', 'Accept': 'application/json'}
      };
      return $http.patch(apiUris.alarmDefinitionId.replace(":id", data.id), data, config);
    };

    service.deleteAlarmDefinition = function(id){
      var config = {
        headers : {'Accept' : 'application/json'}
      };
      return $http.delete(apiUris.alarmDefinitionId.replace(":id", id), config);
    };

    service.nodeList = function(){
      var config = {
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.nodeList, config);
    };

    return service;
  }


  angular
    .module('monitoring')
    .factory('alarmStatusService', alarmStatusService);

  /** @ngInject */
  function alarmStatusService($http, apiUris) {
    var service = {};

    service.alarmStatusList = function(params){
      var config = {
        params: params,
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.alarmStatus, config);
    };

    service.alarmStatusCount = function(params){
      var config = {
        params: params,
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.alarmStatusCount, config);
    };

    service.alarmStatus = function(alarmId){
      var config = {
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.alarmStatusId.replace(":alarmId", alarmId), config);
    };

    service.alarmStatusHistory = function(alarmId, params){
      var config = {
        params: params,
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.alarmStatusHistory.replace(":alarmId", alarmId), config);
    };

    service.alarmActionList = function(alarmId){
      var config = {
        headers : {'Accept' : 'application/json'}
      };
      return $http.get(apiUris.alarmActionList.replace(":alarmId", alarmId), config);
    };

    service.insertAlarmAction = function(data){
      var config = {
        headers : {'Accept' : 'application/json'}
      };
      return $http.post(apiUris.alarmAction, data, config);
    };

    service.updateAlarmAction = function(id, data){
      var config = {
        headers : {'Content-Type': 'application/json', 'Accept': 'application/json'}
      };
      return $http.put(apiUris.alarmActionId.replace(":id", id), data, config);
    };

    service.deleteAlarmAction = function(id){
      var config = {
        headers : {'Accept' : 'application/json'}
      };
      return $http.delete(apiUris.alarmActionId.replace(":id", id), config);
    };

    return service;
  }
})();
