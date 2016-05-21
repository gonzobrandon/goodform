define(['./module'], function (directives) {
    'use strict';
    directives.directive('linebreaks', [ function () {
        
        return {
            restrict: 'C',
            link: function(scope, element) {
                scope.$watch(function(){ return element[0].innerHTML; }, function(val){
                    element[0].innerHTML = element[0].innerHTML.replace(/\n/g, '<br />');
                })
            }
        }
    }]);
});
