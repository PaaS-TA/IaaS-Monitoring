(function() {
  'use strict';

  angular
    .module('monitoring')
    .controller('AlarmNotificationController', AlarmNotificationController);

  /** @ngInject */
  function AlarmNotificationController($scope, $timeout, $window, $sce, $exceptionHandler, alarmNotificationService) {
    var vm = this;
    vm.scope = $scope;
    vm.Math = Math;

    vm.scope.loading = true;

    vm.checkedCnt = 0;

    var limit = 10;
    var offset = 0;
    vm.alarmNotificationList = [];
    (vm.getAlarmNotificationList = function() {
      var params = {
        'offset': offset,
        'limit': limit
      };
      alarmNotificationService.alarmNotificationList(params).then(
        function(result) {
          if(result.data.data) vm.alarmNotificationList = vm.alarmNotificationList.concat(result.data.data);
          vm.totalCount = result.data.totalCnt;
          vm.moreButton = '<strong>더 보 기</strong> (총 ' + vm.totalCount + ' 건)';
          if(vm.alarmNotificationList) {
            offset = vm.alarmNotificationList.length;
          }
          if(vm.alarmNotificationList.length >= vm.totalCount) {
            vm.moreButton = '(총 ' + vm.totalCount + '건)';
          }
          vm.scope.loading = false;
        },
        function(reason) {
          vm.scope.loading = false;
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    })();

    vm.checkAllNotification = function() {
      angular.forEach(vm.alarmNotificationList, function(notification) {
        if(vm.selectAll) {
          notification.select = true;
          vm.checkedCnt++;
        } else {
          notification.select = false;
          vm.checkedCnt--;
        }
      });
    };

    vm.checkNotification = function(obj) {
      if(obj.select) {
        vm.checkedCnt++;
      } else {
        vm.checkedCnt--;
      }
    };

    vm.getAlarmNotification = function(obj) {
      vm.detail = angular.copy(obj);
    };

    vm.saveAlarmNotification = function() {
      vm.scope.loading = true;

      var data = {
        name: vm.detail.name,
        address: vm.detail.email
      };
      var func = null;
      if(vm.detail.id) {
        data.id = vm.detail.id;
        data.period = vm.detail.period;
        func = alarmNotificationService.updateAlarmNotification(data);
      } else {
        func = alarmNotificationService.insertAlarmNotification(data);
      }
      func.then(
        function() {
          vm.getAlarmNotificationList();
          vm.scope.loading = false;
        },
        function(reason) {
          vm.scope.loading = false;
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    };

    vm.deleteAlarmNotification = function() {
      var deleteCnt = 0;
      var deletedCnt = 0;
      angular.forEach(vm.alarmNotificationList, function(notification, index) {
        if(notification.select) {
          deleteCnt++;
          alarmNotificationService.deleteAlarmNotification(notification.id).then(
            function() {
              vm.alarmNotificationList.splice(index, 1);
              vm.checkedCnt--;
              deletedCnt++;
              if(deleteCnt == deletedCnt) {
                vm.getAlarmNotificationList();
              }
            },
            function(reason) {
              vm.scope.loading = false;
              $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
            }
          );
        }
      });
    };

    /********** reload **********/
    vm.scope.$on('broadcast:reload', function() {
      vm.scope.loading = true;
      vm.checkedCnt = 0;
      limit = 10;
      offset = 0;
      vm.alarmNotificationList = [];
      vm.getAlarmNotificationList();
    });
  }


  angular
    .module('monitoring')
    .controller('AlarmDefinitionController', AlarmDefinitionController);

  /** @ngInject */
  function AlarmDefinitionController($scope, $timeout, $exceptionHandler, alarmDefinitionService) {
    var vm = this;
    vm.scope = $scope;
    vm.Math = Math;

    vm.scope.loading = true;

    vm.checkedCnt = 0;

    var limit = 10;
    var offset = 0;
    vm.alarmDefinitionList = [];
    vm.searchAlarmDefinition = function() {
      offset = 0;
      vm.alarmDefinitionList = [];
      vm.totalCount = 0;
      vm.getAlarmDefinitionList();
    };
    (vm.getAlarmDefinitionList = function() {
      var params = {
        'offset': offset,
        'limit': limit
      };
      if(vm.searchCondition) {
        params['name'] = vm.searchCondition;
      }
      if(vm.selectedSeverity) {
        params['severity'] = vm.selectedSeverity;
      }
      alarmDefinitionService.alarmDefinitionList(params).then(
        function(result) {
          if(result.data.data) vm.alarmDefinitionList = vm.alarmDefinitionList.concat(result.data.data);
          vm.totalCount = result.data.totalCnt;
          vm.moreButton = '<strong>더 보 기</strong> (총 ' + vm.totalCount + ' 건)';
          if(vm.alarmDefinitionList) {
            offset = vm.alarmDefinitionList.length;
          }
          if(vm.alarmDefinitionList.length >= vm.totalCount) {
            vm.moreButton = '(총 ' + vm.totalCount + '건)';
          }
          vm.scope.loading = false;
        },
        function(reason) {
          vm.scope.loading = false;
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    })();

    vm.checkAllDefinition = function() {
      angular.forEach(vm.alarmDefinitionList, function(definition) {
        if(vm.selectAll) {
          definition.select = true;
          vm.checkedCnt++;
        } else {
          definition.select = false;
          vm.checkedCnt--;
        }
      });
    };

    vm.checkDefinition = function(obj) {
      if(obj.select) {
        vm.checkedCnt++;
      } else {
        vm.checkedCnt--;
      }
    };

    vm.deleteAlarmDefinition = function() {
      var deleteCnt = 0;
      var deletedCnt = 0;
      angular.forEach(vm.alarmDefinitionList, function(definition) {
        if(definition.select) {
          deleteCnt++;
          alarmDefinitionService.deleteAlarmDefinition(definition.id).then(
            function() {
              vm.checkedCnt--;
              deletedCnt++;
              if(deleteCnt == deletedCnt) {
                vm.getAlarmDefinitionList();
                vm.scope.loading = false;
              }
            },
            function(reason) {
              vm.scope.loading = false;
              $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
            }
          );
        }
      });
    };

    /********** reload **********/
    vm.scope.$on('broadcast:reload', function() {
      vm.scope.loading = true;
      vm.searchAlarmDefinition();
    });

  }


  angular
    .module('monitoring')
    .controller('AlarmDefinitionDetailController', AlarmDefinitionDetailController);

  /** @ngInject */
  function AlarmDefinitionDetailController($scope, $stateParams, $timeout, $location, $exceptionHandler, alarmDefinitionService, projectService, alarmNotificationService) {
    var vm = this;
    vm.scope = $scope;
    vm.Math = Math;

    vm.scope.loading = true;

    vm.pageTitle = $stateParams.id == 'new' ? null : 'Alarm Definition Update';
    vm.alarmDefinition = {};
    vm.alarmDefinition['severity'] = 'HIGH';
    vm.alarmDefinition['matchBy'] = 'hostname';

    /***** get alarm definition *****/
    vm.getAlarmDefinition = function() {
      alarmDefinitionService.alarmDefinition($stateParams.id).then(
        function(result) {
          vm.alarmDefinition = result.data;

          // expression
          var obj_expression = vm.alarmDefinition.expression;
          var gate = obj_expression.indexOf(' and ') >= 0 ? ' and ' : ' or ';
          var gate2 = gate == ' and ' ? ' or ' : ' and ';
          var tmp_expression = obj_expression.split(gate);
          for(var i in tmp_expression) {
            if(tmp_expression[i].indexOf(gate2) >= 0) {
              var or_expression = tmp_expression[i].split(gate2);
              tmp_expression.splice(i, 1);
              for(var j in or_expression) {
                tmp_expression.push(or_expression[j]);
              }
            }
          }
          var arr_expression = [];
          angular.forEach(tmp_expression, function(expression) {
            expression = expression.replace(/(\s*)/g, '');
            var json = {};
            json['func'] = expression.substr(0, expression.indexOf('('));
            var metric = expression.slice(expression.indexOf('(')+1, expression.indexOf(')'));
            json['metric'] = metric;
            if(metric.indexOf('{') >= 0) {
              json['metric'] = metric.substr(0, metric.indexOf('{'));
              json['dimension'] = metric.slice(metric.indexOf('{')+1, metric.indexOf('}'));
            }
            var tmp = expression.substr(expression.indexOf(')')+1);
            var len = tmp.indexOf('=') < 0 ? 1 : 2;
            json['operation'] = tmp.substr(0, len);
            json['value'] = parseInt(tmp.substr(len));
            arr_expression.push(json);
          });
          vm.alarmDefinition.arrExpression = arr_expression;

          vm.scope.loading = false;
        },
        function(reason) {
          vm.scope.loading = false;
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    };

    /***** dimension setup modal *****/
    vm.measurementType = '';
    vm.scope.showDimensionModal = false;
    vm.selDimension = 'hostname';
    var modalDimensionIndex = 0;
    vm.setDimensionModal = function(obj, index) {
      vm.scope.loadingModal = true;

      var modalDimension = obj.dimension;
      modalDimensionIndex = index;
      vm.scope.dimensionTitle = modalDimension == undefined ? 'Set Dimension' : modalDimension;
      if(modalDimension) {
        vm.scope.dimensionTitle = modalDimension;
        vm.modalDimension = modalDimension.substr(0, modalDimension.indexOf('='));
        vm.modalDimensionValue = modalDimension.substr(modalDimension.indexOf('=')+1);
      } else {
        vm.modalDimension = 'hostname';
        vm.scope.dimensionTitle = 'Set Dimension';
      }
      vm.scope.showDimensionModal = !vm.scope.showDimensionModal;

      switch (obj.metric){
        case 'cpu.percent':
        case 'mem.usable_perc':
        case 'disk.space_used_perc':
          vm.measurementType = 'node';
          break;
        case 'cpu.utilization_norm_perc':
        case 'mem.free_perc':
          vm.measurementType = 'vm';
          break;
        default:
          vm.measurementType = 'node';
          break;
      }
      if(vm.measurementType == 'node') {
        alarmDefinitionService.nodeList().then(
          function(result) {
            vm.nodeList = result.data;
            vm.selDimensionValue1 = vm.nodeList[0];
            vm.scope.loadingModal = false;
          },
          function(reason) {
            vm.scope.loadingModal = false;
            $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
          }
        );
      } else {
        getProjectList();
      }
    };
    vm.setDimension = function() {
      var dimensionStr = vm.measurementType == 'node' ? 'hostname=' + vm.selDimensionValue1 : 'hostname=' + vm.selDimensionValue2.name;
      vm.alarmDefinition.arrExpression[modalDimensionIndex].dimension = dimensionStr;
    };

    /***** alarm receiver setup modal *****/
    vm.scope.showAlarmActionModal = false;
    vm.scope.alarmActionTitle = 'Alarm Receiver';
    vm.setAlarmActionModal = function() {
      vm.scope.loadingModal = true;
      vm.scope.showAlarmActionModal = !vm.scope.showAlarmActionModal;
      rLimit = 10;
      rOffset = 0;
      vm.rAlarmNotificationList = [];
      vm.rGetAlarmNotificationList();
    };
    // Show selected rows
    vm.selectedAlarmAction = [];
    vm.selectAlarmAction = function(obj) {
      if(vm.selectedAlarmAction.indexOf(obj) >= 0) {
        obj['select'] = '';
        vm.selectedAlarmAction.splice(vm.selectedAlarmAction.indexOf(obj), 1);
      } else {
        obj['select'] = 'active';
        vm.selectedAlarmAction.push(obj);
      }
    };
    vm.setAlarmAction = function() {
      if(vm.selectedAlarmAction) {
        if(vm.alarmDefinition.alarmAction) {
          vm.alarmDefinition.alarmAction = vm.alarmDefinition.alarmAction.concat(vm.selectedAlarmAction);
        } else {
          vm.alarmDefinition['alarmAction'] = [];
          vm.alarmDefinition.alarmAction = vm.selectedAlarmAction;
        }
        vm.selectedAlarmAction = [];
      }
    };
    vm.deleteAlarmAction = function(index) {
      vm.alarmDefinition.alarmAction.splice(index, 1);
    };

    /***** init *****/
    vm.addAlarmDefinition = function() {
      vm.alarmDefinition.arrExpression.push(
        {
          func: 'max',
          metric: 'cpu.percent',
          operation: '>'
        }
      );
      vm.scope.loading = false;
    };
    if($stateParams.id == 'new') {
      vm.alarmDefinition['arrExpression'] = [];
      vm.addAlarmDefinition();
    } else {
      vm.getAlarmDefinition();
    }

    /***** save alarm definition *****/
    vm.saveAlarmDefinition = function() {
      vm.scope.loadingModal = true;

      // make expression
      var expression = '';
      var arrExpression = vm.alarmDefinition.arrExpression;
      for(var i in arrExpression) {
        expression += arrExpression[i].func + '(' + arrExpression[i].metric;
        if(arrExpression[i].dimension) {
          expression += '{' + arrExpression[i].dimension + '}';
        }
        expression += ') ' + arrExpression[i].operation + ' ' + arrExpression[i].value;
        if((i+1) < arrExpression.length) {
          expression += ' ' + arrExpression[i].gate + ' ';
        }
      }

      var alarm_actions = [];
      angular.forEach(vm.alarmDefinition.alarmAction, function(alarmAction) {
        alarm_actions.push(alarmAction.id);
      });

      var body = {
        name: vm.alarmDefinition.name,
        severity: vm.alarmDefinition.severity,
        expression: expression,
        alarm_actions: alarm_actions,
        description: vm.alarmDefinition.description
      };

      var func = '';
      if($stateParams.id == 'new') {
        body['match_by'] = [vm.alarmDefinition.matchBy];
        func = 'insertAlarmDefinition';
      } else {
        body['id'] = $stateParams.id;
        body['match_by'] = vm.alarmDefinition.matchBy;
        func = 'updateAlarmDefinition';
      }
      alarmDefinitionService[func](body).then(
        function() {
          vm.scope.loadingModal = false;
          $location.path('/alarm_definition');
        },
        function(reason) {
          vm.scope.loadingModal = false;
          $timeout(function() { $exceptionHandler(reason.data, {code: reason.status, message: reason.data}); }, 500);
        }
      );
      vm.scope.loadingModal = false;
    };

    /***** get project & instance list => for dimension setup modal *****/
    vm.getProjectList = function() {
      projectService.projectSummary().then(
        function(result) {
          vm.projectList = result.data;
          vm.selDimensionValue1 = vm.projectList[0];
          getInstanceList(vm.selDimensionValue1.id);
        },
        function(reason) {
          vm.scope.loadingModal = false;
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    };
    vm.getInstanceList = function(id) {
      vm.scope.loadingModal = true;
      var params = {
        'hostname': '',
        'limit': 100,
        'marker': ''
      };
      projectService.projectInstanceList(id, params).then(
        function(result) {
          vm.instanceList = result.data.metric;
          vm.selDimensionValue2 = vm.instanceList[0];
          vm.scope.loadingModal = false;
        },
        function(reason) {
          vm.scope.loadingModal = false;
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    };
    function getProjectList() {
      vm.getProjectList();
    }
    function getInstanceList(id) {
      vm.getInstanceList(id);
    }

    /***** get project & instance list => for alarm receiver setup modal *****/
    var rLimit = 10;
    var rOffset = 0;
    vm.rAlarmNotificationList = [];
    vm.rGetAlarmNotificationList = function() {
      var params = {
        'offset': rOffset,
        'limit': rLimit
      };
      alarmNotificationService.alarmNotificationList(params).then(
        function(result) {
          if(result.data.data) vm.rAlarmNotificationList = vm.rAlarmNotificationList.concat(result.data.data);
          vm.rTotalCount = result.data.totalCnt;
          vm.rMoreButton = '<strong>더 보 기</strong> (총 ' + vm.rTotalCount + ' 건)';
          if(vm.rAlarmNotificationList) {
            rOffset = vm.rAlarmNotificationList.length;
            // Disable already existing action from list
            if(vm.alarmDefinition.alarmAction) {
              angular.forEach(vm.alarmDefinition.alarmAction, function(alarmAction) {
                angular.forEach(vm.rAlarmNotificationList, function(notification) {
                  if(alarmAction.id == notification.id) {
                    notification['disabled'] = true;
                  }
                });
              });
            }
          }
          if(vm.rAlarmNotificationList.length >= vm.rTotalCount) {
            vm.rMoreButton = '(총 ' + vm.rTotalCount + '건)';
          }
          vm.scope.loadingModal = false;
        },
        function(reason) {
          vm.scope.loadingModal = false;
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    };

  }


  angular
    .module('monitoring')
    .controller('AlarmStatusController', AlarmStatusController);

  /** @ngInject */
  function AlarmStatusController($scope, $timeout, $exceptionHandler, alarmStatusService) {
    var vm = this;
    vm.scope = $scope;
    vm.Math = Math;

    vm.scope.loading = true;

    var limit = 10;
    var offset = 0;
    vm.alarmStatusList = [];
    vm.searchAlarmStatus = function() {
      offset = 0;
      vm.alarmStatusList = [];
      vm.totalCount = 0;
      vm.getAlarmStatusList();
    };
    vm.selectedState = 'ALARM';
    (vm.getAlarmStatusList = function() {
      var params = {
        'offset': offset,
        'limit': limit,
        'state': vm.selectedState
      };
      if(vm.selectedSeverity) {
        params['severity'] = vm.selectedSeverity;
      }
      alarmStatusService.alarmStatusList(params).then(
        function(result) {
          if(result.data.data) vm.alarmStatusList = vm.alarmStatusList.concat(result.data.data);
          vm.totalCount = result.data.totalCnt;
          vm.moreButton = '<strong>더 보 기</strong> (총 ' + vm.totalCount + ' 건)';
          if(vm.alarmStatusList) {
            offset = vm.alarmStatusList.length;
          }
          if(vm.alarmStatusList.length >= vm.totalCount) {
            vm.moreButton = '(총 ' + vm.totalCount + '건)';
          }
          vm.scope.loading = false;
        },
        function(reason) {
          vm.scope.loading = false;
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    })();

    /********** reload **********/
    vm.scope.$on('broadcast:reload', function() {
      vm.scope.loading = true;
      vm.searchAlarmStatus();
    });
  }


  angular
    .module('monitoring')
    .controller('AlarmStatusDetailController', AlarmStatusDetailController);

  /** @ngInject */
  function AlarmStatusDetailController($scope, $stateParams, $window, $timeout, $exceptionHandler, alarmStatusService) {
    var vm = this;
    vm.scope = $scope;
    vm.Math = Math;

    vm.scope.loading = true;

    var alarmId = $stateParams.id;
    (vm.getAlarmStatus = function() {
      alarmStatusService.alarmStatus(alarmId).then(
        function(result) {
          vm.detail = result.data;
          vm.getAlarmStatusHistory('1d');
          vm.getAlarmActionList();
        },
        function(reason) {
          vm.scope.loading = false;
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    })();

    // set panel height of alarm history
    angular.element('.alarmHistory').height(angular.element('.alarmDetail').height());
    angular.element($window).on('resize', function () {
      angular.element('.alarmHistory').height(angular.element('.alarmDetail').height());
    });

    vm.getAlarmStatusHistory = function(timeRange) {
      vm.timeRange = timeRange;
      var params = {
        timeRange: timeRange
      };
      alarmStatusService.alarmStatusHistory(alarmId, params).then(
        function(result) {
          vm.alarmStatusHistoryList = result.data;
          vm.scope.loading = false;
        },
        function(reason) {
          vm.scope.loading = false;
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    };

    // 조치 이력 조회
    vm.getAlarmActionList = function() {
      alarmStatusService.alarmActionList(alarmId).then(
        function(result) {
          vm.alarmActionList = result.data;
        },
        function(reason) {
          vm.scope.loading = false;
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    };
    // 조치 이력 등록
    vm.insertAlarmAction = function() {
      vm.scope.loading = true;
      if(vm.alarmActionDesc == '' || vm.alarmActionDesc == null || vm.alarmActionDesc == undefined) {
        vm.scope.loading = false;
        $exceptionHandler('', {code: null, message: '조치내용이 입력되지 않았습니다.'});
        return;
      }
      var body = {
        alarmId: alarmId,
        alarmActionDesc: vm.alarmActionDesc
      };
      alarmStatusService.insertAlarmAction(body).then(
        function() {
          vm.getAlarmActionList();
          vm.alarmActionDesc = '';
          vm.scope.loading = false;
        },
        function(reason) {
          vm.scope.loading = false;
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    };

    // 조치내용 수정 필드 활성화
    vm.setModifying = function(index) {
      vm.modifying = index;
    };
    // 조치내용 수정
    vm.updateAction = function(obj) {
      vm.scope.loading = true;
      vm.modifying = -1;
      var body = {
        alarmActionDesc: obj.alarmActionDesc
      };
      alarmStatusService.updateAlarmAction(obj.id, body).then(
        function() {
          vm.getAlarmActionList();
          vm.scope.loading = false;
        },
        function(reason) {
          vm.scope.loading = false;
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    };

    // 조치내용 삭제
    vm.deleteAction = function(id) {
      vm.scope.loading = true;
      alarmStatusService.deleteAlarmAction(id).then(
        function() {
          vm.getAlarmActionList();
          vm.scope.loading = false;
        },
        function(reason) {
          vm.scope.loading = false;
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    };
  }

})();
