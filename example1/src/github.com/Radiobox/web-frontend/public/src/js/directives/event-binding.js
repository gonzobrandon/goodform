define(['./module'], function (directives) {
    'use strict';
    directives.directive('ngClicktapstart', [ function () {
        return {
            restrict: 'A',
            scope: {
                ngClicktapstart: '&'
            },
            link: function(scope, element) {
                element.bind('mousedown touchstart', function(e){ scope.ngClicktapstart({$event: e}) } );
            }
        }
    }]);
    directives.directive('ngEnter', [ function () {
        return {
            restrict: 'A',
            scope: {
                ngEnter: '&'
            },
            link: function(scope, element) {
                element.bind('keyup', function(e){ 
                    if (e.keyCode == 13) {
                        e.preventDefault();
                        scope.$apply(function(){
                            return scope.ngEnter({$event: e})
                        });
                        return false;
                    }
                });
            }
        }
    }]);
});
