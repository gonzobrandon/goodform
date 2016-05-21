/*
    DIRECTIVE FOR ALBUM FORMS
    
    Example uses
    
    New Album
    
    <div class="col-xs-12" form-album on-cancel="addNew = !addNew" on-create="pushNew($response)"></div>
    
    Edit Album
    
    <div class="col-xs-12" form-album model-id="album.id" on-cancel="toggleEditor($index)" ng-model="albums[$index]" on-update="toggleEditor($index)"></div>
*/

define(['./module'], function (directives) {
    'use strict';
    var albumTemplate = {
        title: '',
        pic_square: ''
    }
    directives.directive('formAlbum', function ($http, $filter, userService, $location, $parse, $timeout) {
        return {
            restrict: 'A',
            templateUrl: '/partials/forms/album.tpl.html',
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
                scope.errors = {};
                
                
                
                var whenLogged = function() {
                    scope.album = angular.copy(albumTemplate);       

                    
                    if(scope.modelId){
                        var modelLink = '/api/albums/' + scope.modelId;
                    } else if(scope.modelLink){
                        var modelLink = scope.modelLink;
                    }
                    
                    if(modelLink){
                        scope.preloading = true;
                        $http.get(modelLink).success(function(data){
                            angular.forEach(scope.album, function(val, key){
                                if(data.response[key] !== null)
                                    scope.album[key] = data.response[key];
                            });
                            scope.preloading = false;
                        });
                    }
                    scope.save = function() {
                        scope.showErrors = true;
                        if(modelLink){
                            $http.patch( modelLink, scope.album).success(function(dat){
                                if(attrs.onUpdate)
                                    $parse(attrs.onUpdate)(scope.$parent, {$response:dat});
                                angular.extend(scope.ngModel, dat.response);
                            }).error(function(dat){
                                scope.errors = dat.notifications.input;
                            });
                        } else {
                            scope.album.album_artist = scope.artistId;
                            $http.post( '/api/albums/', scope.album).success(function(dat){
                                if(attrs.onCreate)
                                    $parse(attrs.onCreate)(scope.$parent, {$response:dat});
                                scope.ngModel = dat.response;
                                scope.album = angular.copy(albumTemplate);
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
