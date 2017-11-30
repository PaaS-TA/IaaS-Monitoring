(function() {
  'use strict';

  angular
    .module('monitoring')
    .factory('mainService', MainService);

  /** @ngInject */
  function MainService($http, apiUris) {
    var service = {};

    service.openStackSummary = function(){
      return $http.get(apiUris.summary);
    };

    return service;
  }
})();
