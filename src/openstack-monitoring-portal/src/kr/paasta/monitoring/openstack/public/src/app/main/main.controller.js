(function() {
  'use strict';

  angular
    .module('monitoring')
    .controller('MainController', MainController);

  /** @ngInject */
  function MainController($scope, $timeout, $exceptionHandler, mainService, computeNodeService, manageNodeService, projectService) {
    var vm = this;
    vm.scope = $scope;
    vm.Math = Math;

    vm.scope.loading = true;

    (vm.getOpenStackSummary = function() {
      mainService.openStackSummary().then(
        function (result) {
          vm.summary = result.data;
          vm.usageVcpu = Math.round((vm.summary.vcpuUsed / vm.summary.vcpuTotal) * 100);
          vm.usageMemory = Math.round((vm.summary.memoryMbUsed / vm.summary.memoryMbTotal) * 100);
          vm.usageDisk = Math.round((vm.summary.diskGbUsed / vm.summary.diskGbTotal) * 100);
          vm.usageVms = Math.round((vm.summary.vmRunning / vm.summary.vmTotal) * 100);
          vm.init(vm.usageVcpu, vm.usageMemory, vm.usageDisk, vm.usageVms);
          vm.scope.loading = false;
        },
        function(reason) {
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    })();

    vm.init = function(cpu, mem, disk, inst) {
      vm.scope.cpuPercent = cpu;
      vm.scope.memPercent = mem;
      vm.scope.diskPercent = disk;
      vm.scope.instance = inst;
      /*if(!$scope.$$phase) { // Error: $digest already in progress
        vm.scope.$apply(function() {
          vm.scope.cpuPercent = cpu;
          vm.scope.memPercent = mem;
          vm.scope.diskPercent = disk;
          vm.scope.instance = inst;
        });
      }*/
    };

    (vm.getComputeNodeSummary = function() {
      computeNodeService.computeNodeSummary().then(
        function (result) {
          vm.computeNodeSummary = result.data;
        },
        function(reason) {
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      )
    })();

    (vm.getManageNodeSummary = function() {
      manageNodeService.manageNodeSummary().then(
        function (result) {
          vm.manageNodeSummary = result.data;
        },
        function(reason) {
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      )
    })();

    (vm.getProjectSummary = function() {
      projectService.projectSummary().then(
        function (result) {
          // $log.log(JSON.stringify(result.data));
          vm.projectSummary = result.data;
        },
        function(reason) {
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      )
    })();

    /********** reload **********/
    vm.scope.$on('broadcast:reload', function() {
      vm.scope.loading = true;
      vm.getOpenStackSummary();
      vm.getComputeNodeSummary();
      vm.getManageNodeSummary();
      vm.getProjectSummary();
    });
  }
})();
