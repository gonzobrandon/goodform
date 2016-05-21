define(['./module'], function (directives) {
    'use strict';
    directives.directive('loading', [ function () {
        
        return {
            restrict: 'A',
            transclude: true,
            scope: {
                loading: '='
            },
            template: '<i class="fa fa-spin fa-sun-o"></i>',
            link: function(scope, element) {
                element.css('text-align', 'center');
                scope.$watch('loading', function(){
                    if(!scope.loading)
                        element.addClass('hidden');
                    else
                        element.removeClass('hidden')
                });
            }
        }
    }]);
});
