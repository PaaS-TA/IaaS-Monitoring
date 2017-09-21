(function() {
  'use strict';

  angular
    .module('monitoring')
    .controller('ProjectController', ProjectController);

  /** @ngInject */
  function ProjectController($scope, $timeout, $location, $sce, $exceptionHandler, projectService) {
    var vm = this;
    vm.scope = $scope;
    vm.Math = Math;

    var hash = $location.hash();
    $location.hash('');
    if(hash) {
      vm.searchCondition = hash;
    }

    (vm.getProjectSummary = function() {
      vm.scope.loading = true;
      vm.selectedProject = false;
      vm.projectInstanceList = null;
      projectService.projectSummary(vm.searchCondition).then(
        function(result) {
          vm.projectSummary = result.data;
          if(hash) {
            angular.forEach(vm.projectSummary, function(projectSummary) {
              if(projectSummary.name == hash) {
                vm.selectProject(projectSummary);
                hash = undefined;
              }
            });
          } else {
            if(vm.projectSummary) {
              vm.selectProject(vm.projectSummary[0]);
            } else {
              vm.selectedProjectName = null;
              vm.searchInstanceName = '';
              vm.projectInstanceList = null;
              vm.projectInstanceTotalCount = 0;
            }
          }
          // vm.scope.loading = false;
        },
        function(reason) {
          vm.scope.loading = false;
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    })();

    var oldObj = {};
    vm.selectProject = function(obj) {
      oldObj['select'] = '';
      obj['select'] = 'active';
      vm.selectedProject = obj;
      vm.selectedProjectName = $sce.trustAsHtml('[ <span style="color:#337ab7;">' + obj.name + '</span> ]');
      vm.searchInstanceName = '';
      vm.projectInstanceList = null;
      vm.projectInstanceTotalCount = 0;
      marker = '';
      vm.getInstanceList();
      oldObj = obj;
    };

    var limit = 10;
    var marker = '';
    vm.searchInstance = function() {
      if(vm.selectedProject == false) {
        var message = 'Project가 선택되지 않았습니다.';
        $exceptionHandler(message, {code: null, message: message});
        return;
      }
      marker = '';
      vm.projectInstanceList = null;
      vm.projectInstanceTotalCount = 0;
      vm.getInstanceList();
    };
    vm.getInstanceList = function() {
      if(vm.selectedProject == false) {
        var message = 'Project가 선택되지 않았습니다.';
        $exceptionHandler(message, {code: null, message: message});
        return;
      }
      vm.scope.loading = true;

      var params = {
        'hostname': vm.searchInstanceName,
        'limit': limit,
        'marker': marker
      };
      projectService.projectInstanceList(vm.selectedProject.id, params).then(
        function(result) {
          if(vm.selectedProject.id == result.data.projectId) {
            vm.projectInstanceList = vm.projectInstanceList == null ? [] : vm.projectInstanceList;
            vm.projectInstanceList = vm.projectInstanceList.concat(result.data.metric);
            vm.projectInstanceTotalCount = result.data.totalCnt;
            vm.moreButton = '<strong>더 보 기</strong> (총 ' + vm.projectInstanceTotalCount + ' 건)';
            if(vm.projectInstanceList) {
              marker = vm.projectInstanceList[(vm.projectInstanceList.length-1)].instance_id;
            }
            if(vm.projectInstanceList.length >= vm.projectInstanceTotalCount) {
              vm.moreButton = '(총 ' + vm.projectInstanceTotalCount + '건)';
            }
          }
          vm.scope.loading = false;
        },
        function(reason) {
          vm.scope.loading = false;
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    };

    /********** reload **********/
    vm.scope.$on('broadcast:reload', function() {
      vm.getProjectSummary();
    });
  }


  angular
    .module('monitoring')
    .controller('ProjectDetailController', ProjectDetailController);

  /** @ngInject */
  function ProjectDetailController($scope, $log, $stateParams, $interval, $timeout, $rootScope,
                                   projectService, projectChartConfig, common, cookies, nvd3Generator, $exceptionHandler) {
    var vm = this;
    vm.scope = $scope;
    vm.Math = Math;

    vm.scope.loading = true;

    /********** Chart **********/
    vm.scope.gridsterOpts = {
      margins: [20, 20],
      columns: 3,
      rowHeight: 250,
      mobileModeEnabled: false,
      swapping: true,
      draggable: {
        handle: 'h4',
        stop: function(event, $element, widget) {
          $timeout(function(){
            $log.log(event+','+$element+','+widget);
          },400)
        }
      },
      resizable: {
        enabled: true,
        handles: ['n', 'e', 's', 'w', 'ne', 'se', 'sw', 'nw'],
        minWidth: 200,
        layoutChanged: function() {
        },

        // optional callback fired when resize is started
        start: function(event, $element, widget) {
          $timeout(function(){
            $log.log(event+','+$element+','+widget);
          },400)
        },

        // optional callback fired when item is resized,
        resize: function(event, $element, widget) {
          if (widget.chart.api) widget.chart.api.update();
        },

        // optional callback fired when item is finished resizing
        stop: function(event, $element, widget) {
          $timeout(function(){
            if (widget.chart.api) widget.chart.api.update();
          },400)
        }
      }
    };

    vm.scope.dashboard = {
      widgets: []
    };

    vm.scope.events = {
      resize: function(e, scope){
        $timeout(function(){
          if (scope.api && scope.api.update) scope.api.update();
        },200)
      }
    };

    vm.scope.config = { visible: false };

    $timeout(function(){
      vm.scope.config.visible = true;
    }, 200);

    angular.element(window).on('resize', function(){
      vm.scope.$broadcast('resize');
    });

    var charts = projectChartConfig;

    // 조회조건 설정
    var instanceId = $stateParams.instanceId;
    var condition = {
      instanceId: instanceId,
      groupBy: cookies.getGroupBy()
    };
    if(cookies.getDefaultTimeRange() == 'custom') {
      condition['timeRangeFrom'] = common.timeDifference(new Date().getTime(), cookies.getTimeRangeFrom());
      condition['timeRangeTo'] = common.timeDifference(new Date().getTime(), cookies.getTimeRangeTo());
    } else {
      condition['defaultTimeRange'] = cookies.getDefaultTimeRange();
    }

    var count = 0;
    for(var i in charts) {
      var chartOpt = charts[i];
      if(chartOpt.func) {
        (function(opt, cnt) {
          projectService[opt.func](condition).then(
            function (result) {
              // if(opt.func == 'nodeNetworkDroppedPacketList') console.info(JSON.stringify(result.data));
              vm.scope.setWidget(cnt, opt, result.data);
            },
            function (reason) {
              vm.scope.setWidget(cnt, opt);
              $log.error(reason);
              $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
            }
          );
        })(chartOpt, count);
        count++;
      }
    }
    var stop = $interval(function() {
      if(count == (vm.scope.dashboard.widgets).length) {
        $interval.cancel(stop);
        vm.scope.dashboard.widgets.sort(common.CompareForSort);
        vm.scope.loading = false;
      }
    }, 500);

    vm.scope.setWidget = function(index, dtvOpt, jsonArr) {
      var col = dtvOpt.col==undefined?Math.floor(index%3):dtvOpt.col;
      var row = dtvOpt.row==undefined?Math.floor(index/3):dtvOpt.row;
      var sizeX = dtvOpt.sizeX==undefined?1:dtvOpt.sizeX;
      var sizeY = dtvOpt.sizeY==undefined?1:dtvOpt.sizeY;

      var widget = {
        col: col, row: row, sizeX: sizeX, sizeY: sizeY, name: dtvOpt.name, id: dtvOpt.id, type: dtvOpt.type,
        chart: {
          options: nvd3Generator[dtvOpt.type].options(),
          api: {}
        },
        func: dtvOpt.func
      };
      if(dtvOpt.percent && jsonArr) {
        widget.chart.options.chart.forceY = [0, 100];
        widget.chart.options.chart.yAxis.tickFormat = function (d) { return d3.format('.0%')(d/100); };
        widget.percent = dtvOpt.percent;
      }
      if(dtvOpt.axisLabel) {
        widget.chart.options.chart.yAxis.axisLabel = dtvOpt.axisLabel;
        widget.chart.options.chart.yAxis.axisLabelDistance = -5;
        widget.chart.options.chart.margin.left = 55;
        widget.axisLabel = dtvOpt.axisLabel;
      }
      if(jsonArr) {
        var value, arr = [];
        var checkData = 0;
        for(var i in jsonArr) {
          value = jsonArr[i].metric==null?[{time:0,usage:0}]:jsonArr[i].metric;
          arr.push({values: value, key: jsonArr[i].name});
          for (var j in value) {
            if(value[j] != null) {
              if (checkData < value[j].usage) {
                checkData = value[j].usage;
              }
            }
          }
        }
        widget.chart.data = arr;

        if(dtvOpt.type != 'list') {
          if(checkData < 5 && widget.chart.options.chart.forceY == undefined) {
            widget.chart.options.chart.forceY = [0, 5];
          }
          if(checkData > 10000) {
            widget.chart.options.chart.yAxis.axisLabelDistance = 20;
            widget.chart.options.chart.margin.left = 80;
          }
        }
      } else {
        widget.chart.options.chart.forceY = false;
      }
      vm.scope.dashboard.widgets.push(widget);
    };

    // 조회주기 및 GroupBy 설정
    var savedCustom = false;
    vm.scope.saveTimeRange = function () {
      var instanceId = $stateParams.instanceId;
      var condition = {
        instanceId: instanceId,
        groupBy: cookies.getGroupBy()
      };
      if(vm.scope.selTimeRange == 'custom') {
        condition['timeRangeFrom'] = common.timeDifference(new Date().getTime(), vm.scope.timeRangeFrom);
        condition['timeRangeTo'] = common.timeDifference(new Date().getTime(), vm.scope.timeRangeTo);
        savedCustom = true;
      } else {
        condition['defaultTimeRange'] = vm.scope.selTimeRange;

        condition.defaultTimeRange = vm.scope.selTimeRange;
        condition.groupBy = vm.scope.selGroupBy;
        savedCustom = false;
      }
      angular.forEach(vm.scope.dashboard.widgets, function(widget, index) {
        (function(opt, idx) {
          projectService[opt.func](condition).then(
            function (result) {
              if(result) {
                var jsonArr = result.data;
                var value, arr = [];
                for(var i in jsonArr) {
                  value = jsonArr[i].metric==null?[{time:0,usage:0}]:jsonArr[i].metric;
                  arr.push({values: value, key: jsonArr[i].name});
                }
                vm.scope.dashboard.widgets[idx].chart.data = arr;
                vm.scope.dashboard.widgets[idx].loading = false;
              }
            },
            function (reason, status) {
              $timeout(function() { $exceptionHandler(reason.Message, {code: status, message: reason.Message}); }, 500);
            }
          );
        })(widget, index);
      });
    };

    // time range 선택 시 그에 해당하는 group by 선택
    vm.scope.selectGroupBy = function() {
      vm.scope.selGroupBy = common.getGroupingByTimeRange(vm.scope.selTimeRange, vm.scope.timeRangeFrom, vm.scope.timeRangeTo);
    };

    // 팝업에서 save 하지 않은 경우 원래 값을 유지
    angular.element('#timeRange').on('hidden.bs.modal', function () {
      if(savedCustom != true) {
        vm.scope.selTimeRange = cookies.getDefaultTimeRange();
        vm.scope.selGroupBy = cookies.getGroupBy();
      }
    });

    /********** reload **********/
    vm.scope.$on('broadcast:reload', function() {
      if(vm.scope.selectedTab == 'logs') {
        vm.scope.logSearch();
      } else {
        angular.forEach(vm.scope.dashboard.widgets, function(widget) {
          widget.loading = true;
        });
        vm.scope.saveTimeRange();
      }
    });
    vm.scope.$on('broadcast:saveTimeRange', function (event, data) {
      vm.scope.selTimeRange = data.selTimeRange;
      vm.scope.selGroupBy = data.selGroupBy;
      vm.scope.selRefreshTime = data.selRefreshTime;
      vm.scope.timeRangeFrom = data.timeRangeFrom;
      vm.scope.timeRangeTo = data.timeRangeTo;
      vm.scope.saveTimeRange();
    });
  }
})();
