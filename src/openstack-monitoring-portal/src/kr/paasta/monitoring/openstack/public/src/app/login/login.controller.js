(function() {
  'use strict';

  angular
    .module('monitoring')
    .controller('LoginController', LoginController);

  /** @ngInject */
  function LoginController($scope, $timeout, $location, $http, $exceptionHandler, apiUris, cache, constants, loginService) {
    var vm = this;
    vm.scope = $scope;
    vm.Math = Math;

    angular.element('.loginUiView').css('height','100%');

    vm.credentials = {
      username: '',
      password: '',
      rememberMe: false
    };
    /*vm.login = function() {
      vm.scope.loading = true;
      if(vm.email == 'admin' && vm.password == 'admin') {
        cache.setUser({name: 'Admin', email: $scope.email}, constants.expire);
        $location.path('/');
      } else {
        var message = '이메일 또는 비밀번호가 일치하지 않습니다.';
        $timeout(function() { $exceptionHandler(message, {code: 403, message: message}); }, 500);
      }
    };*/
    vm.login = function() {
      vm.scope.loading = true;

      var credentials = {
        username: vm.email,
        password: vm.password
      };

      loginService.ping().then(
        function() {

          loginService.authenticate(credentials).then(
            function (response) {
              if (!response) {
                vm.scope.loading = false;
                $scope.login.password = '';
                var vMsg = 'Username or password is incorrect. Please check back again.';
                $timeout(function() { $exceptionHandler(vMsg, {code: 401, message: vMsg}); }, 500);
                $location.path('/login');
              } else {
                var data = response.data;
                cache.setUser({name: data.username}, constants.expire);
                $location.path('/');
              }
            },
            function(reason) {
              vm.scope.loading = false;
              $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
            }
          );
        },
        function(reason) {
          vm.scope.loading = false;
          $timeout(function() { $exceptionHandler(reason.data.message, {code: reason.data.HttpStatus, message: reason.data.message}); }, 500);
        }
      );
    };

  }
})();
