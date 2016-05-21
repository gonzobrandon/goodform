define(['./module'], function (directives) {
    'use strict';
    directives.directive('slideDown', ['animationInterval', function (animationInterval) {
        
        return {
            restrict: 'A',
            scope: {
                slideDown: '='
            },
            link: function(scope, element, attrs) {
                var moving = false;
                var height = element[0].offsetHeight;
                var paddingTop = parseInt(element.computedStyle('padding-top'));
                var paddingBottom = parseInt(element.computedStyle('padding-bottom'));
                var target = scope.slideDown ? 1 : 0;
                var currentHeight = height*target;
                var currentPaddingTop = paddingTop*target;
                var currentPaddingBottom = paddingBottom*target;
                var nextHeight = false;
                var nextPaddingTop = false;
                var nextPaddingBottom = false;
                var int = false;
                var content = element.html();
                
                if(!target){
                    element.css({height: '0px', overflow: 'hidden', display: 'none'});
                }
                
                if(attrs.ngIf){
                    element.attr('ng-if', 'iffyif');
                }
                                
                var move = function(){
                    var ease = .2;
                    currentHeight = nextHeight;
                    var diff = target * height - currentHeight;
                    if(diff*ease > 50)
                        ease = 50 / diff;
                    nextHeight = currentHeight + diff*ease;
                    
                    
                    currentPaddingTop = nextPaddingTop;
                    var diffPaddingTop = target * paddingTop - currentPaddingTop;
                    nextPaddingTop = currentPaddingTop + diffPaddingTop*ease;
                    
                    
                    currentPaddingBottom = nextPaddingBottom;
                    var diffPaddingBottom = target * paddingBottom - currentPaddingBottom;
                    nextPaddingBottom = currentPaddingBottom + diffPaddingBottom*ease;
                    
                    if(diff > 1 || diff < -1){
                        element.css({height: nextHeight + 'px', paddingTop: nextPaddingTop + 'px', paddingBottom: nextPaddingBottom + 'px', overflow: 'hidden', display: 'block'});
                    } else {
                        if(target)
                            element.css({height: '', paddingBottom: '', paddingTop: '', overflow: '', display: ''});
                        else
                            element.css({height: '0px', paddingBottom: '0px', paddingTop: '0px', overflow: 'hidden', display: 'none'});
                        animationInterval.cancel(int);
                        int = false;
                    }
                }
                
                var getHeight = function(){
                    var disp = element.css('display');
                    var h = element.css('height');
                    element.css({display: '', height: ''});
                    var val = element[0].offsetHeight;
                    element.css({display: disp, height: h});
                    if(val)
                        height = val;
                    return val;
                }
                
                scope.$watch(function(){ return element.children().length }, function(val){
                    getHeight();
                });
                
                scope.$watch('slideDown', function(val){
                    getHeight();
                    target = val ? 1 : 0;
                    if(!int)
                        int = animationInterval(move);
                });
                scope.$on('$destroy', function() {
                    animationInterval.cancel(int);
                })

            }
        }
    }]);
});

