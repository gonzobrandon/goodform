define(['./module',], function (controllers) {
    'use strict';

    controllers.controller('alertHeaderCtrl', ['$scope', 'rootService', function ($scope, rootService) {

        $scope.alerts = rootService.alerts;

        $scope.addAlert = rootService.addAlert;

        $scope.closeAlert = rootService.closeAlert;

    }]);


    controllers.controller('headerCtrl', ['$scope', '$modal', '$log', '$http', '$filter', '$cookieStore', 'userService', 'rootService', 'playerService', '$window',   function ($scope, $modal, $log, $http, $filter, $cookieStore, userService, rootService, playerService, $window) {

        $scope.player = playerService;

        $scope.userService = userService;
        $scope.mobileVisible = false;
        $scope.mobileShow = function () {
            $scope.mobileVisible = !$scope.mobileVisible;
        };

        $scope.openSignup = userService.openSignup;
        var openSignup = $scope.openSignup;
        $scope.openLogin = userService.openLogin;

        $scope.logout = function(){
            userService.doLogout();
        }
        $scope.window = $window;


    }]);
});
