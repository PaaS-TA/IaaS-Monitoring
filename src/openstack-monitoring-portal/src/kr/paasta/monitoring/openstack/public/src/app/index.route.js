(function() {
  'use strict';

  angular
    .module('monitoring')
    .config(routerConfig);

  /** @ngInject */
  function routerConfig($stateProvider, $urlRouterProvider) {
    var navbar = {
      template: '<div><acme-navbar></acme-navbar></div>'
    };
    $stateProvider
      .state('login', {
        url: '/login',
        views: {
          'login': {
            templateUrl: 'app/login/login.html',
            controller: 'LoginController',
            controllerAs: 'login'
          }
        }
      });

    $stateProvider
      .state('main', {
        url: '/',
        views: {
          'navbar': navbar,
          'body': {
            templateUrl: 'app/main/main.html',
            controller: 'MainController',
            controllerAs: 'main'
          }
        }
      });

    $stateProvider
      .state('compute_node', {
        url: '/compute_node',
        views: {
          'navbar': navbar,
          'body': {
            templateUrl: 'app/node/compute_node.html',
            controller: 'ComputeNodeController',
            controllerAs: 'cnd'
          }
        }
      });

    $stateProvider
      .state('compute_node_detail', {
        url: '/compute_node/:hostname',
        views: {
          'navbar': navbar,
          'body': {
            templateUrl: 'app/node/node_detail.html',
            controller: 'NodeDetailController'
          }
        }
      });

    $stateProvider
      .state('manage_node', {
        url: '/manage_node',
        views: {
          'navbar': navbar,
          'body': {
            templateUrl: 'app/node/manage_node.html',
            controller: 'ManageNodeController',
            controllerAs: 'mnd'
          }
        }
      });

    $stateProvider
      .state('manage_node_detail', {
        url: '/manage_node/:hostname',
        views: {
          'navbar': navbar,
          'body': {
            templateUrl: 'app/node/node_detail.html',
            controller: 'NodeDetailController'
          }
        }
      });

    $stateProvider
      .state('project', {
        url: '/project',
        views: {
          'navbar': navbar,
          'body': {
            templateUrl: 'app/project/project.html',
            controller: 'ProjectController',
            controllerAs: 'pjt'
          }
        }
      });

    $stateProvider
      .state('project_detail', {
        url: '/project/:instanceId',
        views: {
          'navbar': navbar,
          'body': {
            templateUrl: 'app/project/project_detail.html',
            controller: 'ProjectDetailController'
          }
        }
      });

    $stateProvider
      .state('alarm_notification', {
        url: '/alarm_notification',
        views: {
          'navbar': navbar,
          'body': {
            templateUrl: 'app/alarm/alarm_notification.html',
            controller: 'AlarmNotificationController',
            controllerAs: 'aln'
          }
        }
      });

    $stateProvider
      .state('alarm_definition', {
        url: '/alarm_definition',
        views: {
          'navbar': navbar,
          'body': {
            templateUrl: 'app/alarm/alarm_definition.html',
            controller: 'AlarmDefinitionController',
            controllerAs: 'ald'
          }
        }
      });

    $stateProvider
      .state('alarm_definition_detail', {
        url: '/alarm_definition/:id',
        views: {
          'navbar': navbar,
          'body': {
            templateUrl: 'app/alarm/alarm_definition_detail.html',
            controller: 'AlarmDefinitionDetailController',
            controllerAs: 'ald'
          }
        }
      });

    $stateProvider
      .state('alarm_status', {
        url: '/alarm_status',
        views: {
          'navbar': navbar,
          'body': {
            templateUrl: 'app/alarm/alarm_status.html',
            controller: 'AlarmStatusController',
            controllerAs: 'ast'
          }
        }
      });

    $stateProvider
      .state('alarm_status_detail', {
        url: '/alarm_status/:id',
        views: {
          'navbar': navbar,
          'body': {
            templateUrl: 'app/alarm/alarm_status_detail.html',
            controller: 'AlarmStatusDetailController',
            controllerAs: 'ast'
          }
        }
      });

    $urlRouterProvider.otherwise('/');
  }

})();
