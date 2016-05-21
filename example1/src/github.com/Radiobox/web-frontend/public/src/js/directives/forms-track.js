/*
    DIRECTIVE FOR TRACK FORMS
    
    Example uses
    
    New Track
    
    <div class="col-xs-12" form-track on-cancel="addNew = !addNew" on-create="pushNew($response)"></div>
    
    Edit Track
    
    <div class="col-xs-12" form-track model-id="track.id" on-cancel="toggleEditor($index)" ng-model="tracks[$index]" on-update="toggleEditor($index)"></div>
*/

define(['./module'], function (directives) {
    'use strict';
    var trackTemplate = {
        title: '',
        artist: '',
        album: '',
        media: '',
        preview_media: ''
    };
    var allowedFileTypes = [
        'audio/mp3',
        'audio/flac',
        'audio/aac'
    ]
    directives.directive('formTrack', function ($http, $filter, userService, $location, $parse, $timeout, $upload) {
        return {
            restrict: 'A',
            templateUrl: '/partials/forms/track.tpl.html',
            scope: {
                modelLink: '=', //broken, links are too inconsistent
                modelId: '=',
                ngModel: '=',
                albumId: '=',
                artistId: '='
            },
            link: function (scope, element, attrs) {
                if(attrs.onCancel)
                    scope.onCancel = function(){ return $parse(attrs.onCancel)(scope.$parent); };
                
                scope.loading = false;
                scope.showErrors = false;
                scope.user = userService;
                scope.errors = {};
                
                scope.dropSupported = false;
                
                var whenLogged = function() {
                    
                    scope.track = angular.copy(trackTemplate);       

                    
                    if(scope.modelId){
                        var modelLink = '/api/tracks/' + scope.modelId;
                    } else if(scope.modelLink){
                        var modelLink = scope.modelLink;
                    }
                    
                    if(modelLink){
                        scope.preloading = true;
                        $http.get(modelLink).success(function(data){
                            angular.forEach(scope.track, function(val, key){
                                scope.track[key] = data.response[key];
                            });
                            scope.preloading = false;
                        });
                    }
                    
                    var fileInput = element.find('input');
                    angular.forEach(fileInput, function(val){
                        if(angular.element(val).hasClass('fileUpload'))
                            fileInput = val;
                    });
                    
                    scope.chooseFile = function(){
                        fileInput.click();
                    };
                    
                    scope.$watch('track.media', function(val){
                        if(val == ''){
                            scope.fileName = false;
                        }
                    })
                    
                    scope.save = function() {
                        scope.showErrors = true;
                        if(modelLink){
                            $http.patch( modelLink, scope.track).success(function(dat){
                                if(attrs.onUpdate)
                                    $parse(attrs.onUpdate)(scope.$parent, {$response:dat});
                                angular.extend(scope.ngModel, dat.response);
                            }).error(function(dat){
                                
                            });
                        } else {
                            scope.track.artist = scope.artistId;
                            scope.track.album = scope.albumId;
                            $http.post( '/api/tracks/', scope.track).success(function(dat){
                                if(attrs.onCreate)
                                    $parse(attrs.onCreate)(scope.$parent, {$response:dat});
                                scope.ngModel = dat.response;
                                scope.track = angular.copy(trackTemplate);
                            }).error(function(dat){
                                scope.errors = dat.notifications.input;
                            });
                        }
                    }
                    scope.onFileSelect = function($files) {
                        var headers = $http.defaults.headers.common;
                        
                        var file = $files[0];
                        scope.fileName = file.name;
                        
                        if(allowedFileTypes.indexOf(file.type) == -1){
                            scope.errors.media = 'The file selected is not supported. Please use mp3, flac, or aac file types.';
                            return false;
                        }
                        
                        scope.imageLoaded = false;
                        scope.uploading = 1;
                        scope.upload = $upload.upload({
                            url: '/api/media/',
                            method: 'POST',
                            headers: headers,
                            file: file
                        }).progress(function(evt) {
                            if(scope.uploading){
                                scope.uploading = parseInt(98 * evt.loaded / evt.total)+1;
                                scope.$apply();
                            }
                        }).success(function(data, status, headers, config) {
                            // file is uploaded successfully
                            scope.uploading = false;
                            scope.track.media = data.response.id;
                            scope.track.preview_media = data.response.id;
                        }).error(function() {
                            scope.uploading = false;
                            alert('something went wrong with your upload, please try again');
                        });
                    }
                };
                scope.$watch(function(){ return element.find('input').length; }, whenLogged);
                userService.onLogin(false, true, scope);
    
            }
        }
    });
    
});
