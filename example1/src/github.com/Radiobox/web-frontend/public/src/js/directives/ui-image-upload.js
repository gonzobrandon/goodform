define(['./module'], function (directives) {
    'use strict';
    directives.directive('imageWiz', ['$upload', '$http', function ($upload, $http) {
        return {
            restrict: 'A',
            scope: {
                aspect: '@aspect' || false,
                ngModel: '=',
                origWidth: '@origWidth',
            },
            templateUrl: '/partials/directives/image_upload.tpl.html',
            link: function(scope, element) {
                var cropLoaded = false;
                scope.cropApplied = false;
                if(!scope.ngModel)
                    scope.ngModel = '';
                
                var truncateNum = function(val){ return Math.round(val*10)/10 };
                var cropToString = function(){ return truncateNum(scope.crop.left) + ':' + truncateNum(scope.crop.right) + ':' + truncateNum(scope.crop.top) + ':' + truncateNum(scope.crop.bottom)};
                
                scope.crop = {
                    left: 0,
                    right: 0,
                    top: 0,
                    bottom: 0
                }
                
                var loadedCropCheck = false;
                var loadedCrop = angular.copy(scope.crop);
                
                var handleUpload = function(data){
                    autoCrop = true;
                    cropLoaded = true;
                    scope.uploading = false;
                    scope.preview = data.response.image_path;
                    scope.ngModel = data.response.image_id;
                    angular.element(scope.dropZone).css('max-width', '100%').css('height', 'auto');
                }
                
                scope.$watch('ngModel', function(val){
                    if(scope.ngModel && scope.ngModel != ''){
                        if(typeof(val) == 'object'){
                            $http.post('/api/images', {url: val.source}).success(function(data){
                                
                                if(val.offset_x)
                                    loadedCrop.left = val.offset_x * 100/851;
                                if(val.offset_y)
                                    loadedCrop.top = val.offset_y * 100/315;
                                handleUpload(data);
                                
                            });
                            return false;
                        }
                        if(scope.ngModel.indexOf('/') > -1){
                            scope.preview = scope.ngModel;
                            scope.ngModel = scope.ngModel.split('/');
                            scope.ngModel = scope.ngModel[scope.ngModel.length-1].split('.')[0];
                        }
                        scope.preview = '/api/images/' + scope.ngModel + '.jpg?crop-top=0';
                        if(!cropLoaded){
                            $http.get('/api/images/' + scope.ngModel).success(function(data){
                                scope.crop = data.response.options['crop-percent'];
                                loadedCrop = data.response.options['crop-percent'];
                                angular.forEach(scope.crop, function(val, key){
                                    scope.crop[key] = val*100;
                                });
                                loadedCropCheck = cropToString();
                                scope.cropApplied = true;
                            })
                        }
                        cropLoaded = true;
                    }
                });
                
                scope.dWidth = '100%';
                scope.chooseFile = function(){
                    element.find('input')[0].click();
                };
                
                var autoCrop = false;
                scope.$watch('preview', function(){
                    if(scope.preview){
                        var img = document.createElement('img');
                        img.src = scope.preview;
                        angular.element(img).bind('load',function(){
                            scope.crop = loadedCrop;
                            scope.imageLoaded = true;
                            scope.$apply();
                            if(autoCrop){
                                scope.applyCrop();
                                autoCrop = false;
                            }
                        });
                    }
                });
                
                
                
                var result = false;
                angular.forEach(element.find('div'), function(el){
                    var $el = angular.element(el);
                    if($el.hasClass('drop-zone')){
                        scope.dropZone = el;
                    }
                    if($el.hasClass('image-crop')){
                        scope.cropZone = el;
                    }
                });
                
                scope.dropSupported = ('draggable' in scope.dropZone) || ('ondragstart' in scope.dropZone && 'ondrop' in scope.dropZone);
                
                var $el = angular.element(scope.dropZone);
                
                                
                var dHeight = scope.dropZone.offsetHeight;
                
                var setPadding = function(val){
                    if(scope.aspect  && scope.ngModel == ''  && !scope.preview){
                        if(scope.origWidth){
                            $el.css('max-width', scope.origWidth+'px');
                        }
                        dHeight = scope.dropZone.offsetWidth/parseFloat(scope.aspect);
                        $el.css('height', dHeight+'px');
                        var divs = $el.find('div');
                        angular.element(divs[0]).css('padding', '0 15px')
                        angular.element(divs[0]).css('padding', (dHeight-divs[0].offsetHeight)/2+'px 15px');
                        angular.element(divs[1]).css('display','block');
                        angular.element(divs[1]).css('padding', '0 15px')
                        angular.element(divs[1]).css('padding', (dHeight-divs[1].offsetHeight)/2+'px 15px');
                        angular.element(divs[1]).css('display','');
                    } else {
                        $el.css('max-width', '100%').css('height', 'auto');
                    }
                }
                
                scope.$watch(function(){
                    if(scope.aspect  && scope.ngModel == ''  && !scope.preview){
                        return scope.dropZone.offsetWidth/parseFloat(scope.aspect);
                        
                    } else {
                        return false;
                    }
                }, setPadding);
                
                
                

                
                scope.onFileSelect = function($files) {
                    var headers = $http.defaults.headers.common;
                    
                    var file = $files[0];
                    scope.imageLoaded = false;
                    scope.uploading = 1;
                    scope.upload = $upload.upload({
                        url: '/api/images/',
                        method: 'POST',
                        headers: headers,
                        file: file
                    }).progress(function(evt) {
                        if(scope.uploading){
                            scope.uploading = parseInt(98 * evt.loaded / evt.total)+1;
                            scope.$apply();
                        }
                    }).success(function(data, status, headers, config) {
                        handleUpload(data);
                    }).error(function() {
                        scope.uploading = false;
                        alert('something went wrong with your upload, please try again');
                    });
                }
                
                
                scope.$watch(function(){
                    return cropToString();
                }, function(val){
                    if(loadedCropCheck && loadedCropCheck == val)
                        scope.cropApplied = true;
                    else
                        scope.cropApplied = false;
                });
                scope.applyCrop = function() {
                    $http.patch('/api/images/' + scope.ngModel, {
                        'crop-left': scope.crop.left,
                        'crop-right': scope.crop.right,
                        'crop-top': scope.crop.top,
                        'crop-bottom': scope.crop.bottom
                    }).success(function(response){
                        loadedCropCheck = cropToString();
                        scope.cropApplied = true;
                    }).error(function(){
                        alert('cropping failed, check network and try again');
                    });
                    
                }
            }
        }
    }]);
});
