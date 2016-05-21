define(['./module'], function (directives) {
    'use strict';
    directives.directive('uiDatetime', ['$interval', '$window', '$timestamp', function ($interval, $window, $timestamp) {
        
        return {
            restrict: 'A',
            transclude: true,
            scope: {
                ngModel: '='
            },
            templateUrl: '/partials/directives/ui-datetime.tpl.html',
            link: function(scope, element) {
                
                scope.dateOpened = false;
                scope.openDate = function($event) {
                    $event.preventDefault();
                    $event.stopPropagation();
                
                    scope.dateOpened = true;
                };
                scope.changed = function(){
                    
                };
            }
        }
    }]);
});
