define(['./module', 'facebook'], function (directives) {
    'use strict';
    directives.directive('formArtist', function ($http, $filter, userService, $location, $parse, $timeout) {
        return {
            restrict: 'A',
            templateUrl: '/partials/forms/artist.tpl.html',
            scope: {
                modelLink: '=', //broken, links are too inconsistent
                modelId: '=',
                ngModel: '='
            },
            link: function (scope, element, attrs) {
                var whenLogged = function() {
                    if(attrs.onCancel)
                        scope.onCancel = function(){ return $parse(attrs.onCancel)(scope.$parent); };
                    scope.loading = false;
                    scope.showErrors = false;
                    scope.user = userService;
                    var artistTemplate = {
                        username: '',
                        contact_email: userService.userObj.email,
                        hometown_address: {
                            city: '',
                            state_province: '',
                            country: '',
                            zip_postal: ''
                        },
                        keywords: [],
                        pic_cover: '',
                        pic_square: '',
                        band_users: {},
                        booking_user: userService.userId,
                        manager_user: userService.userId,
                        description: ''
                    }
                    scope.artist = angular.copy(artistTemplate);
                    scope.temp = {
                        tag: ''
                    };
                    scope.meta = {
                        facebook: ''
                    }
                    scope.tag = '';
                    
                    if(scope.modelId){
                        var modelLink = '/api/artists/' + scope.modelId;
                    } else if(scope.modelLink){
                        var modelLink = scope.modelLink;
                    }
                    
                    if(modelLink){
                        scope.preloading = true;
                        $http.get(modelLink).success(function(data){
                            angular.forEach(scope.artist, function(val, key){
                                if(data.response[key] !== null)
                                    scope.artist[key] = data.response[key];
                                
                                //convert user objects to id for patch
                                if(typeof(scope.artist[key].id) != 'undefined')
                                    scope.artist[key] = scope.artist[key].id;
                            });
                            scope.preloading = false;
                        });
                    }
                    
                    scope.addTag = function(tag) {
                        if(typeof(tag) == 'string')
                            scope.temp.tag = tag;
                        if(scope.artist.keywords.indexOf(scope.temp.tag) == -1 && scope.temp.tag != '')
                            scope.artist.keywords.push(scope.temp.tag);
                        scope.temp.tag = '';
                        
                        return false;
                    };
                    
                    scope.removeTag = function(tag) {
                        scope.artist.keywords.splice(scope.artist.keywords.indexOf(tag),1);
                    };
                    
                    scope.fbImport = function(){
                        FB.api(
                            '/' + scope.meta.facebook,
                            function (response) {
                                scope.$apply(function(){
                                    if (response && !response.error) {
                                        scope.artist.username = response.name || scope.artist.username;
                                        scope.artist.description = response.about || scope.artist.description;
                                        scope.addTag(response.genre);
                                        if(response.cover)
                                            scope.artist.pic_cover = response.cover;
                                        FB.api('/' + response.username + '/picture?width=720&height=720', function(response){
                                            if(response.data){
                                                if(response.data.url && !response.data.is_silhouette){
                                                    scope.artist.pic_square = { source: response.data.url };
                                                }
                                            }
                                        });
                                    }
                                });
                            }
                        );
                    }
                    
                    scope.save = function(e) {
                        
                        
                        scope.showErrors = true;
                        if(modelLink){
                            $http.patch( modelLink, scope.artist).success(function(dat){
                                if(attrs.onUpdate)
                                    $parse(attrs.onUpdate)(scope.$parent, {$response:dat});
                                angular.extend(scope.ngModel, dat.response);
                            }).error(function(dat){
                                scope.errors = dat.notifications.input;
                            });
                        } else {
                            scope.showErrors = true;
                            $http.post( '/api/artists/', scope.artist).success(function(dat){
                                if(attrs.onCreate)
                                    $parse(attrs.onCreate)(scope.$parent, {$response:dat});
                                scope.artist = angular.copy(artistTemplate);
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
});
