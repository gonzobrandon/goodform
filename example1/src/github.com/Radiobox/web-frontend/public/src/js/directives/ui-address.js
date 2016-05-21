define(['./module'], function (directives) {
    'use strict';
    directives.directive('address', ['$interval', '$window', '$timezones', function ($interval, $window, $timezones) {
        
        return {
            restrict: 'A',
            transclude: true,
            scope: {
                address: '@',
                ngModel: '='
            },
            templateUrl: '/partials/directives/address.tpl.html',
            compile: function(element, attrs) {
                if(typeof(attrs.required) != 'undefined')
                    element.find('input').attr('required', 'true');
                return {
                    post: function(scope, element, attrs) {
                        scope.countries = $timezones.countries;
                        var form = element.parentForm().attr('name');
                        if(form && scope.$parent[form])
                            scope.form = scope.$parent[form];
                        
                        scope.$watch(function(){ return JSON.stringify(scope.ngModel); }, function(){
                            if(!scope.ngModel)
                                scope.ngModel = {
                                    city: '',
                                    state: '',
                                    country: ''
                                }
                            if(typeof(scope.ngModel.country) == 'undefined' || typeof(scope.countries[scope.ngModel.country]) == 'undefined')
                                scope.ngModel.country = 'United States';
                        });
                        if(scope.address)
                            scope.type = scope.address;
                    }
                }
            }
        }
    }]);
});
