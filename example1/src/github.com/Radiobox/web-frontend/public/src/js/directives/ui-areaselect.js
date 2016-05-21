define(['./module'], function (directives) {
    'use strict';
    directives.directive('areaSelect', ['$interval', '$window', function ($interval, $window) {
        
        return {
            restrict: 'A',
            transclude: true,
            scope: {
                areaSelect: '=',
                aspect: '@'
            },
            templateUrl: '/partials/directives/ui-areaselect.tpl.html',
            link: function(scope, element) {
                scope.aspect = parseFloat(scope.aspect);
                var startCoords = {x:0,y:0};
                var currentCoords = {x:0,y:0};
                var lastCoords = {x:0,y:0};
                var started = false;
                var dragEl = element.find('div')[0];
                var initialEl = false;
                var actionType = false;
                var selfChange = false;
                var width = 0;
                var height = 0;
                
                function start(){
                    if(!scope.areaSelect || typeof(scope.areaSelect.left) == 'undefined' || typeof(scope.areaSelect.right) == 'undefined' || !typeof(scope.areaSelect.top) == 'undefined' || typeof(scope.areaSelect.bottom) == 'undefined'){
                        scope.areaSelect = { left: 0, right: 0, top: 0, bottom:0 };
                    }
                    scope.origAreaSelect = angular.copy(scope.areaSelect);
                    width = (100 - scope.areaSelect.left - scope.areaSelect.right) * dragEl.offsetWidth / 100;
                    height = (100 - scope.areaSelect.top - scope.areaSelect.bottom) * dragEl.offsetHeight / 100;
                    checkAspect();
                }
                start();
                
                function round(num){
                    return Math.floor( num * 100 ) / 100;
                }
                scope.$watch(function(){
                    return scope.areaSelect.left + ':' + scope.areaSelect.right + ':' + scope.areaSelect.top + ':' + scope.areaSelect.bottom;
                }, function(val){
                    if(!selfChange){
                        start();
                    }
                    selfChange = false;
                });
                scope.$watch(function(){
                    return element[0].offsetHeight;
                }, function(){
                    start();
                });
                function handleCoords(e){
                    
                    e.preventDefault();
                    
                    if(!initialEl)
                        initialEl = e.target;
                    
                    var target = dragEl;
                    
                    var offX  = e.pageX - angular.element(target).offset().left;
                    var offY  = e.pageY - angular.element(target).offset().top;
                    var x = (offX / target.offsetWidth * 100);
                    var y = (offY / target.offsetHeight * 100);
                    x = x < 0 ? 0 : (x > 100 ? 100 : x);
                    y = y < 0 ? 0 : (y > 100 ? 100 : y);
                    x = round(x);
                    y = round(y);
                    currentCoords = {x:x,y:y};
                    
                    width = (100 - scope.areaSelect.left - scope.areaSelect.right) * dragEl.offsetWidth / 100;
                    height = (100 - scope.areaSelect.top - scope.areaSelect.bottom) * dragEl.offsetHeight / 100;
                    
                    if(!started){
                        startCoords = lastCoords = {x:x,y:y};
                        started = true;
                        var $el = angular.element(e.target);
                        if($el.hasClass('handle'))
                            actionType = 'handle';
                        else if($el.parent().hasClass('handle'))
                            actionType = 'handle';
                        else if($el.hasClass('crop-overlay'))
                            actionType = 'move'
                        else
                            actionType = 'new';
                    } else {
                    
                        switch(actionType){
                            case 'move':
                                moveWindow(e);
                                break;
                            case 'handle':
                                handleGrab(e);
                                break;
                            default:
                                drawNew(e);
                        }
                        
                        scope.areaSelect.left = round(scope.areaSelect.left);
                        scope.areaSelect.right = round(scope.areaSelect.right);
                        scope.areaSelect.top = round(scope.areaSelect.top);
                        scope.areaSelect.bottom = round(scope.areaSelect.bottom);
                        
                    }
                    scope.$apply();
                    lastCoords = currentCoords;
                    
                }
                function setWidth(changeTo, MF, bMF){
                    selfChange = true;
                    if(MF != 'left'){
                        MF = 'right';
                        scope.areaSelect.right = ((changeTo * 100 / dragEl.offsetWidth) - 100 + scope.areaSelect.left) * -1;
                    } else {
                        scope.areaSelect.left  = ((changeTo * 100 / dragEl.offsetWidth) - 100 + scope.areaSelect.right) * -1;
                    }
                    if(scope.areaSelect[MF] < 0){
                        scope.areaSelect[MF] = 0;
                        getDimensions();
                        if(scope.aspect){
                            setHeight(width / scope.aspect, bMF);
                        }
                    }

                }
                function checkAspect(widthMF, heightMF){
                    getDimensions();
                    if(scope.aspect){
                        if(width/height > scope.aspect){
                            setWidth(height * scope.aspect, widthMF, heightMF);
                        } else {
                            setHeight(width / scope.aspect, heightMF, widthMF);
                        }
                    }
                }
                function setHeight(changeTo, MF, bMF){
                    if(dragEl.offsetHeight){
                        selfChange = true;
                        if(MF != 'top'){
                            MF = 'bottom';
                            scope.areaSelect.bottom = ((changeTo * 100 / dragEl.offsetHeight) - 100 + scope.areaSelect.top) * -1;
                        } else {
                            scope.areaSelect.top = ((changeTo * 100 / dragEl.offsetHeight) - 100 + scope.areaSelect.bottom) * -1;
                        }
                        if(scope.areaSelect[MF] < 0){
                            scope.areaSelect[MF] = 0;
                            getDimensions();
                            if(scope.aspect){
                                setWidth(height * scope.aspect, bMF);
                            }
                        }
                    }
                }
                function moveWindow(e){
                    selfChange = true;
                    var moveX = currentCoords.x - lastCoords.x;
                    var moveY = currentCoords.y - lastCoords.y;
                    var newCoords = {};
                    
                    newCoords.left = scope.areaSelect.left + moveX;
                    newCoords.right = scope.areaSelect.right - moveX;
                    newCoords.top = scope.areaSelect.top + moveY;
                    newCoords.bottom = scope.areaSelect.bottom - moveY;
                    if(newCoords.left >= 0 && newCoords.right >= 0){
                        scope.areaSelect.left += moveX;
                        scope.areaSelect.right -= moveX;
                    }
                    if(newCoords.top >= 0 && newCoords.bottom >= 0){
                        scope.areaSelect.top += moveY;
                        scope.areaSelect.bottom -= moveY;
                    }
                    
                }
                function getDimensions(){
                    width = (100 - scope.areaSelect.left - scope.areaSelect.right) * dragEl.offsetWidth / 100;
                    height = (100 - scope.areaSelect.top - scope.areaSelect.bottom) * dragEl.offsetHeight / 100;
                }
                function handleGrab(e){
                    var widthMF = 'right';
                    var heightMF = 'bottom';
                    var corner = -1;
                    if(initialEl.parentElement.classList.contains('handle'))
                        initialEl = initialEl.parentElement;
                    if(initialEl.className.indexOf('top') != -1){
                        if(100 - scope.areaSelect.bottom > currentCoords.y)
                            scope.areaSelect.top = currentCoords.y;
                        heightMF = 'top';
                        corner++;
                    }
                    if(initialEl.className.indexOf('bottom') != -1){
                        if(scope.areaSelect.top < currentCoords.y)
                            scope.areaSelect.bottom = 100 - currentCoords.y;
                        corner++;
                    }
                    if(initialEl.className.indexOf('left') != -1){
                        if(100 - scope.areaSelect.right > currentCoords.x)
                            scope.areaSelect.left = currentCoords.x;
                        widthMF = 'left';
                        corner++;
                    }
                    if(initialEl.className.indexOf('right') != -1){
                        if(scope.areaSelect.left < currentCoords.x)
                            scope.areaSelect.right = 100 - currentCoords.x;
                        corner++;
                    }
                    getDimensions();
                    if(scope.aspect){
                        if(corner){
                            if(width/height > scope.aspect){
                                setWidth(height * scope.aspect, widthMF, heightMF);
                            } else {
                                setHeight(width / scope.aspect, heightMF, widthMF);
                            }
                        } else {
                            if(initialEl.className.indexOf('left') == -1 && initialEl.className.indexOf('right') == -1){
                                setWidth(height * scope.aspect, widthMF, heightMF);
                            } else {
                                setHeight(width / scope.aspect, heightMF, widthMF);
                            }
                        }
                    }
                    
                }
                function drawNew(){
                    if(startCoords.x > currentCoords.x){
                        scope.areaSelect.left = currentCoords.x;
                        scope.areaSelect.right = 100 - startCoords.x;
                    } else {
                        scope.areaSelect.left = startCoords.x;
                        scope.areaSelect.right = 100 - currentCoords.x;
                    }
                    if(startCoords.y > currentCoords.y){
                        scope.areaSelect.top = currentCoords.y;
                        scope.areaSelect.bottom = 100 - startCoords.y;
                    } else {
                        scope.areaSelect.top = startCoords.y;
                        scope.areaSelect.bottom = 100 - currentCoords.y;
                    }
                    var widthMF = 'right';
                    var heightMF = 'bottom';
                    if(startCoords.x > currentCoords.x)
                        widthMF = 'left';
                    if(startCoords.y > currentCoords.y)
                        heightMF = 'top';
                    checkAspect(widthMF, heightMF);
                }
                function stopDrag(){
                    initialEl = false;
                    started = false;
                    actionType = false;
                    angular.element($window).unbind('mouseup touchend',stopDrag);
                    angular.element($window).unbind('mousemove touchmove',handleCoords);
                }
                scope.startDrag = function(e){
                    handleCoords(e);
                    angular.element($window).bind('mouseup touchend',stopDrag);
                    angular.element($window).bind('mousemove touchmove',handleCoords);
                }
            }
        }
    }]);
});
