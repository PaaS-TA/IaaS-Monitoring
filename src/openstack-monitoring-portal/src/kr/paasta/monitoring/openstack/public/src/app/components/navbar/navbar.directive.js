(function() {
  'use strict';

  angular
    .module('monitoring')
    .directive('acmeNavbar', acmeNavbar);

  /** @ngInject */
  function acmeNavbar() {
    var directive = {
      restrict: 'E',
      templateUrl: 'app/components/navbar/navbar.html',
      scope: {
        creationDate: '=',
        eventHandler: '&ngClick'
      },
      controller: NavbarController,
      controllerAs: 'vm',
      bindToController: true/*,
      link: function(scope, elem, attrs) {
        // click 1
        elem.bind('click', function(e) {
          scope.$apply(function () {
            console.log(angular.element(e.target).parents('.dropdown-menu').is('ul'));
            if(angular.element(e.target).parents('.dropdown-menu').is('ul')) {
              // angular.element(e.target).parent().children("active").remove();
              angular.element(e.target).parent().addClass('active');
            }
          });
        });

        // click 2
        scope.eventHandler = function(e) {
          angular.element(e.target).parent().addClass('active');
        };
      }*/
    };

    return directive;

    /** @ngInject */
    function NavbarController($scope, $rootScope, $stateParams, $location, $interval, $timeout,
                              moment, common, cookies, cache, $exceptionHandler, alarmStatusService, loginService) {

      switch($location.path()){
        case '/manage_node' :
        case '/manage_node/'+$stateParams.hostname :
          $scope.selected = 'mnd';
          break;
        case '/compute_node' :
        case '/compute_node/'+$stateParams.hostname :
          $scope.selected = 'cnd';
          break;
        case '/project' :
          $scope.selected = 'prj';
          break;
        case '/alarm_notification' :
          $scope.selected = 'aln';
          break;
        case '/alarm_definition' :
        case '/alarm_definition/'+$stateParams.id :
          $scope.selected = 'ald';
          break;
        case '/alarm_status' :
        case '/alarm_status/'+$stateParams.id :
          $scope.selected = 'ast';
          break;
        default:
          $scope.selected = '';
      }

      $scope.alarms = 0;
      ($scope.getAlarms = function() {
        var params = {
          state: 'ALARM'
        };
        alarmStatusService.alarmStatusCount(params).then(
          function(result) {
            $scope.alarms = result.data.totalCnt;
          },
          function(reason) {
            $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
          }
        );
      })();

      ($scope.spinning = function() {
        $scope.spin = true;
        var stop = $interval(function() {
          if(angular.element('body').find('.fa-spinner').is(':visible') == false) {
            $interval.cancel(stop);
            $scope.spin = false;
          }
        }, 500);
      })();

      $scope.username = cache.getUser().name;
      $scope.logout = function() {
        loginService.logout().then(
          function() {
            cache.clear();
            $location.path('/login');
          },
          function(reason) {
            $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
            cache.clear();
            $location.path('/login');
          }
        );
      };

      $scope.reload = function() {
        $scope.getAlarms();
        $scope.spinning();
        $rootScope.$broadcast('broadcast:reload');
      };

      /********** TimeRange & GroupBy **********/
      $scope.selTimeRange = cookies.getDefaultTimeRange();
      $scope.selGroupBy = cookies.getGroupBy();
      $scope.selRefreshTime = cookies.getRefreshTime();
      // Time Range 수동설정 달력
      $scope.timeRangeFrom = moment();
      $scope.timeRangeTo = moment();
      $scope.timeRangeTo.startOf('day').fromNow();
      $scope.optionsFrom = {format: 'YYYY.MM.DD HH:mm'};
      $scope.optionsTo = {format: 'YYYY.MM.DD HH:mm'};
      $scope.updateTimeRange = function (dateFrom, dateTo) {
        $scope.optionsFrom.maxDate = dateTo;
        $scope.optionsTo.minDate = dateFrom;
        $scope.optionsFromDate = $scope.optionsFrom.maxDate._d;
        $scope.optionsToDate = $scope.optionsTo.minDate._d;
        if($scope.selTimeRange == 'custom') $scope.selGroupBy = common.selectGroupingByCustomTimeRange(dateTo, dateFrom);
      };
      $timeout(function() {
        $scope.updateTimeRange($scope.timeRangeTo, $scope.timeRangeFrom);
      });

      ($scope.getTimeRangeString = function() {
        $timeout(function() {
          if($scope.selTimeRange == 'custom') {
            var toMonth = (new Date($scope.optionsToDate).getMonth()+1).toString().length === 1 ? '0'+(new Date($scope.optionsToDate).getMonth()+1).toString() : (new Date($scope.optionsToDate).getMonth()+1).toString();
            var toDate = new Date($scope.optionsToDate).getDate().toString().length === 1 ? '0'+new Date($scope.optionsToDate).getDate().toString() : new Date($scope.optionsToDate).getDate().toString();
            var toHours = new Date($scope.optionsToDate).getHours().toString().length === 1 ? '0'+new Date($scope.optionsToDate).getHours().toString() : new Date($scope.optionsToDate).getHours().toString();
            var toMinutes = new Date($scope.optionsToDate).getMinutes().toString().length === 1 ? '0'+new Date($scope.optionsToDate).getMinutes().toString() : new Date($scope.optionsToDate).getMinutes().toString();
            var toSeconds = new Date($scope.optionsToDate).getSeconds().toString().length === 1 ? '0'+new Date($scope.optionsToDate).getSeconds().toString() : new Date($scope.optionsToDate).getSeconds().toString();

            var fromMonth = (new Date($scope.optionsFromDate).getMonth()+1).toString().length === 1 ? '0'+(new Date($scope.optionsFromDate).getMonth()+1).toString() : (new Date($scope.optionsFromDate).getMonth()+1).toString();
            var fromDate = new Date($scope.optionsFromDate).getDate().toString().length === 1 ? '0'+new Date($scope.optionsFromDate).getDate().toString() : new Date($scope.optionsFromDate).getDate().toString();
            var fromHours = new Date($scope.optionsFromDate).getHours().toString().length === 1 ? '0'+new Date($scope.optionsFromDate).getHours().toString() : new Date($scope.optionsFromDate).getHours().toString();
            var fromMinutes = new Date($scope.optionsFromDate).getMinutes().toString().length === 1 ? '0'+new Date($scope.optionsFromDate).getMinutes().toString() : new Date($scope.optionsFromDate).getMinutes().toString();
            var fromSeconds = new Date($scope.optionsFromDate).getSeconds().toString().length === 1 ? '0'+new Date($scope.optionsFromDate).getSeconds().toString() : new Date($scope.optionsFromDate).getSeconds().toString();

            var to = new Date($scope.optionsToDate).getFullYear()+
              '.' +toMonth+
              '.' +toDate+
              ' ' +toHours+
              ':' +toMinutes+
              ':' +toSeconds;
            var from = new Date($scope.optionsFromDate).getFullYear()+
              '.'+fromMonth+
              '.'+fromDate+
              ' '+fromHours+
              ':'+fromMinutes+
              ':'+fromSeconds;
            $scope.timeRangeString = (to)+' to '+(from);
          } else {
            $scope.timeRangeString = angular.element("input[name='radioTimeRange']:checked").parent().text();
          }
        });
      })();

      // 조회주기 및 GroupBy 설정
      $scope.saveTimeRange = function () {
        if($scope.selTimeRange == 'custom') {
          cookies.putDefaultTimeRange($scope.selTimeRange);
          cookies.putTimeRangeFrom(Number($scope.timeRangeFrom));
          cookies.putTimeRangeTo(Number($scope.timeRangeTo));
        } else {
          cookies.putDefaultTimeRange($scope.selTimeRange);
          cookies.putGroupBy($scope.selGroupBy);
        }
        cookies.putRefreshTime($scope.selRefreshTime);
        var datas = {
          selTimeRange: $scope.selTimeRange,
          selGroupBy: $scope.selGroupBy,
          selRefreshTime: $scope.selRefreshTime,
          timeRangeFrom: $scope.timeRangeFrom,
          timeRangeTo: $scope.timeRangeTo
        };
        angular.element('body').find('.modal-backdrop').hide();
        $rootScope.$broadcast('broadcast:saveTimeRange', datas);
        $scope.getTimeRangeString();
      };

      // time range 선택 시 그에 해당하는 group by 선택
      $scope.selectGroupBy = function() {
        $scope.selGroupBy = common.getGroupingByTimeRange($scope.selTimeRange, $scope.timeRangeFrom, $scope.timeRangeTo);
      };

      /********** Auto Refresh **********/
      var refreshInterval;
      ($scope.runRefreshInterval = function() {
        if(cookies.getRefreshTime() !== 'off' && angular.isUndefined(cookies.getRefreshTime())) {
          var refreshTime = common.getMillisecondsRefreshTime(cookies.getRefreshTime());
          refreshInterval = $interval(function() {
            $rootScope.$broadcast('broadcast:reload');
          }, refreshTime);
        }
      })();
      $scope.$on('broadcast:reload', function(){
        $interval.cancel(refreshInterval);
        $scope.runRefreshInterval();
      });
      $scope.$on('$stateChangeStart', function(){
        $interval.cancel(refreshInterval);
      });

      /********** modal **********/
      $scope.timeRangeTop = function($event){
        var offsetTop = angular.element($event.target).prop('offsetTop');
        angular.element('.time-range').css('top',(offsetTop+180)-angular.element(window).scrollTop());
      };

    }
  }

})();
