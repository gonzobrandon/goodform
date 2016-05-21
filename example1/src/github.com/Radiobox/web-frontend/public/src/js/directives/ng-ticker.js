define(['./module'], function (directives) {
    'use strict';
    directives.directive('ngTicker', ['$interval', function ($interval) {
        
        return {
            restrict: 'C',
            transclude: true,
            scope: {
                ngTicker: '='
            },
            template: '<span ng-mouseover="stop()" ng-mouseout="start()">'
                        +'<span class="ticker-block" ng-transclude></span>'
                        +'<span class="ticker-block" ng-transclude></span>'
                        +'<span class="ticker-block" ng-transclude></span>'
                        +'<span class="ticker-block" ng-transclude></span>'
                    +'</span>',
            link: function(scope, element) {
                var span = element.find('span')[0];
                var $span = angular.element(span);
                var maxDist = 0;
                var movement = 0;
                var speed = 1;
                function move(){
                    maxDist = span.offsetWidth/2;
                    movement += speed;
                    if(movement > maxDist)
                        movement -= maxDist;
                    $span.css('margin-left','-' + movement + 'px');
                }
                var interval = $interval(move, 20);
                scope.stop = function(){
                    $interval.cancel(interval);
                }
                scope.start = function(){
                    interval = $interval(move, 20);
                }
            }
        }
    }]);
});
