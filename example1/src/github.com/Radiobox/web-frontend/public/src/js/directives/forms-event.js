define(['./module', 'moment-timezone'], function (directives, moment) {
    'use strict';
    directives.directive('formEvent', function ($http, $filter, userService, $location, $parse, $timezones) {
        return {
            restrict: 'A',
            templateUrl: '/partials/forms/event.tpl.html',
            scope: {
                modelLink: '=', //broken, links are too inconsistent
                modelId: '=',
                ngModel: '=',
                artistId: '='
            },
            link: function (scope, element, attrs) {
                if(attrs.onCancel)
                    scope.onCancel = function(){ return $parse(attrs.onCancel)(scope.$parent); };
                
                scope.loading = false;
                scope.showErrors = false;
                scope.user = userService;
                scope.meta = {
                    timezone: false
                };
                scope.triggerVenueSave = 0;
                
                
                var whenLogged = function() {
                    var defaultDate = new Date();
                    defaultDate.setHours(20);
                    defaultDate.setMinutes(0);
                    defaultDate.setSeconds(0);
                    defaultDate.setMilliseconds(0);
                    var eventTemplate = {
                        title: '',
/*                         phone: '', */
/*                         url: '', */
                        scheduled_start: angular.copy(defaultDate),
                        scheduled_end: angular.copy(defaultDate),
                        artist: false,
                        venue: false
                    }
                    var localeFormat = 'YYYY-MM-DDTHH:mm:ss';
                    scope.startTimeOpen = false;
                    scope.endTimeOpen = false;
                    scope.event = angular.copy(eventTemplate);
                    
                    if(scope.modelId){
                        var modelLink = '/api/events/' + scope.modelId;
                    } else if(scope.modelLink){
                        var modelLink = scope.modelLink;
                    }
                    
                    if(modelLink){
                        scope.preloading = true;
                        $http.get(modelLink).success(function(data){
                            angular.forEach(scope.event, function(val, key){
                                if(data.response[key] !== null)
                                    scope.event[key] = data.response[key];
                                
                                //convert user objects to id for patch
                                if(typeof(scope.event[key].id) != 'undefined')
                                    scope.event[key] = scope.event[key].id;
                                
                                if(key == 'scheduled_start' || key == 'scheduled_end') {
                                    scope.event[key] = moment(moment(scope.event[key]).tz(data.response.venue.timezone).format(localeFormat)).format();
                                }
                            });
                            scope.preloading = false;
                        });
                    }
                    
                    scope.saveEvent = function(e) {
                    
                        scope.showErrors = true;
                        var input = angular.copy(scope.event);
                        if(scope.artistId)
                            input.artist = scope.artistId;
                        input.scheduled_start = moment.tz(moment(input.scheduled_start).format(localeFormat), scope.meta.timezone).format();
                        input.scheduled_end = moment.tz(moment(input.scheduled_end).format(localeFormat), scope.meta.timezone).format();
                        
                        
                        scope.showErrors = true;
                        if(modelLink){
                            $http.patch( modelLink, input).success(function(dat){
                                if(attrs.onUpdate)
                                    $parse(attrs.onUpdate)(scope.$parent, {$response:dat});
                                angular.extend(scope.ngModel, dat.response);
                            }).error(function(dat){
                                scope.loading = false;
                                scope.errors = dat.notifications.input;
                            });
                        } else {
                            scope.showErrors = true;
                            $http.post( '/api/events/', input).success(function(dat){
                                if(attrs.onCreate)
                                    $parse(attrs.onCreate)(scope.$parent, {$response:dat});
                                scope.event = angular.copy(eventTemplate);
                                scope.loading = false;
                                scope.errors = {};
                            }).error(function(dat){
                                scope.loading = false;
                                scope.errors = dat.notifications.input;
                            });
                        }
                        
                    }
                    
                    scope.save = function() {
                        var input = scope.event;
                        scope.triggerVenueSave++;
                        scope.loading = true;
                    }
                    scope.$watch('event.scheduled_start', function(val){
                        if(scope.event.scheduled_start.valueOf() > scope.event.scheduled_end.valueOf())
                            scope.event.scheduled_end = angular.copy(scope.event.scheduled_start);
                    })
                };
                userService.onLogin(whenLogged, true, scope);
    
            }
        }
    });
});
