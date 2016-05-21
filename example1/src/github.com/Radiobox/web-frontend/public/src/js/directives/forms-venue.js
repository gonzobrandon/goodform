/*
    DIRECTIVE FOR VENUE FORMS
    
    Example uses
    
    New Venue
    
    <div class="col-xs-12" form-venue on-cancel="addNew = !addNew" on-create="pushNew($response)"></div>
    
    Edit Venue
    
    <div class="col-xs-12" form-venue model-id="venue.id" on-cancel="toggleEditor($index)" ng-model="venues[$index]" on-update="toggleEditor($index)"></div>
*/

define(['./module'], function (directives) {
    'use strict';
    var venueTemplate = {
        name: '',
        email: '',
        address: {},
        url: '',
        venue_blurb: '',
        timezone: ''
    }
    directives.directive('formVenue', function ($http, $filter, userService, $location, $parse, $timeout) {
        return {
            restrict: 'A',
            templateUrl: '/partials/forms/venue.tpl.html',
            scope: {
                modelLink: '=', //broken, links are too inconsistent
                modelId: '=',
                ngModel: '='
            },
            link: function (scope, element, attrs) {
                if(attrs.onCancel)
                    scope.onCancel = function(){ return $parse(attrs.onCancel)(scope.$parent); };
                
                scope.loading = false;
                scope.showErrors = false;
                scope.user = userService;
                scope.errors = {};
                
                        
                
                
                var whenLogged = function() {
                    scope.venue = angular.copy(venueTemplate);
                    
                    if(scope.modelId){
                        var modelLink = '/api/venues/' + scope.modelId;
                    } else if(scope.modelLink){
                        var modelLink = scope.modelLink;
                    }
                    
                    if(modelLink){
                        scope.preloading = true;
                        $http.get(modelLink).success(function(data){
                            angular.forEach(scope.venue, function(val, key){
                                scope.venue[key] = data.response[key];
                            });
                            scope.preloading = false;
                        });
                    }
                    scope.save = function() {
                        scope.showErrors = true;
                        if(modelLink){
                            $http.patch( modelLink, scope.venue).success(function(dat){
                                if(attrs.onUpdate)
                                    $parse(attrs.onUpdate)(scope.$parent, {$response:dat});
                                angular.extend(scope.ngModel, dat.response);
                            }).error(function(dat){
                                scope.errors = dat.notifications.input;
                            });
                        } else {
                            $http.post( '/api/venues/', scope.venue).success(function(dat){
                                if(attrs.onCreate)
                                    $parse(attrs.onCreate)(scope.$parent, {$response:dat});
                                scope.ngModel = dat.response;
                                scope.venue = angular.copy(venueTemplate);
                            }).error(function(dat){
                                scope.errors = dat.notifications.input;
                            });
                        }
                    }
                };
                userService.onLogin(whenLogged, true, scope);
    
            }
        }
    });
    

    //Looks the same as the form now, but will be a very different selection directive once venues are loaded so...
    //DO NOT REFACTOR WITH ABOVE
    directives.directive('selectVenue', function ($http, $filter, userService, $location, $parse, $timeout) {
        return {
            restrict: 'A',
            template:   '<div class="form-group">' + 
                        '    <label ng-click="save()">Venue Name</label>' +
                        '    <span class="warn" slide-down="errors.name">{{errors.name}}</span>' + 
                        '    <input class="form-control" name="name" placeholder="" type="text" ng-model="venue.name"  />' + 
                        '</div>' + 
                        '<span class="warn" slide-down="errors.address">{{errors.address}}</span>' + 
                        '<div address ng-model="venue.address" errors="errors" ></div>' + 
                        '<div class="form-group">' + 
                        '    <label>Timezone</label>' + 
                        '    <span class="warn" slide-down="errors.timezone">{{errors.timezone}}</span>' + 
                        '    <div ui-timezone ng-model="venue.timezone" country="venue.address.country"></div>' + 
                        '</div>',
            scope: {
                modelLink: '=', //broken, links are too inconsistent
                modelId: '=',
/*                 ngModel: '=', */
                saveTrigger: '=',
                loadingBool: '=',
                metaTimezone: '=',
                metaId: '=',
                onSave: '&'
            },
            link: function (scope, element, attrs) {
                scope.loading = scope.loadingBool;
                if(attrs.onCancel)
                    scope.onCancel = function(){ return $parse(attrs.onCancel)(scope.$parent); };
                scope.showErrors = false;
                scope.user = userService;
                scope.errors = {};
                var modelLink = '';
                
                var whenLogged = function() {
                    scope.venue = angular.copy(venueTemplate);
                    scope.ngModel = scope.venue;
                    scope.$watch('saveTrigger', function(val){
                        if(val == 0){
                            return false;
                        }
                        scope.save();
                    });
                    scope.$watch('modelId', function(val){
                        if(scope.modelId){
                            modelLink = '/api/venues/' + scope.modelId;
                        } else {
                            scope.venue = angular.copy(venueTemplate);
                            modelLink = false;
                        }
                        
                        if(modelLink){
                            scope.preloading = true;
                            $http.get(modelLink).success(function(data){
                                angular.forEach(scope.venue, function(val, key){
                                    scope.venue[key] = data.response[key];
                                });
                                scope.preloading = false;
                            });
                        }
                    });
                    scope.$watch('venue.timezone', function(val){
                        scope.metaTimezone = val;
                    });
                    scope.save = function() {
                        scope.showErrors = true;
                        if(modelLink){
                            $http.patch( modelLink, scope.venue).success(function(dat){
                                $timeout(function(){
                                    if(attrs.onSave)
                                        scope.onSave()
                                });
                                angular.extend(scope.ngModel, dat.response);
                                scope.errors = {};
                            }).error(function(dat){
                                scope.loadingBool = false;
                                scope.errors = dat.notifications.input;
                            });
                        } else {
                            $http.post( '/api/venues/', scope.venue).success(function(dat){
                                scope.errors = {};
                                scope.ngModel = dat.response;
                                scope.metaTimezone = scope.ngModel.timezone;
                                scope.modelId = dat.response.id;
                                modelLink = dat.meta.location;
                                $timeout(function(){
                                    if(attrs.onSave)
                                        scope.onSave()
                                });
                            }).error(function(dat){
                                scope.loadingBool = false;
                                scope.errors = dat.notifications.input;
                            });
                        }
                    }
                };
                userService.onLogin(whenLogged, true, scope);
    
            }
        }
    });
});
