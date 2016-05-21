define(['./module', 'moment-timezone'], function (directives, moment) {
    'use strict';
    directives.directive('uiTimezone', ['$timezones', function ($timezones) {
        
        return {
            restrict: 'A',
            scope: {
                ngModel: '=',
                country: '='
            },
            template: '<select ng-model="ngModel" class="form-control">'+
                      '    <option ng-repeat="(val, zone) in timezones" value="{{val}}" ng-selected="val == ngModel">{{zone}}</option>'+
                      '</select>',
            link: function(scope, element) {
                scope.$watch('country', function(val){
                    if(typeof($timezones.countries[val]) == 'undefined'){
                        scope.timezones = $timezones.timezones;
                    } else {
                        scope.timezones = $timezones.countries[val];
                    }
                });
            }
        }
    }]);
});
