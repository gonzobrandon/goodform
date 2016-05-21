define(['./module'], function (directives) {
    'use strict';
    directives.directive('slider', ['$interval', '$window', function ($interval, $window) {
        
        return {
            restrict: 'E',
            scope: {
                ngModel: '='
            },
            template: '<div class="slide-control-pad pull-right col-xs-10" ng-clicktapstart="startVolumeChange($event)">'
                     +'    <div class="slide-control-bar pull-right col-xs-12">'
                     +'        <div class="slide-control-setter" style="right:{{(ngModel - 100) * -1}}%"></div>'
                     +'    </div>'
                     +'</div>',
            link: function(scope, element) {
                var volInt;
                var volEl;
                function setVolume(e){
                    
                    e.preventDefault();
                    var apply;
                    if(volEl){
                        target = volEl;
                        apply = true;
                    } else {
                        var target = angular.element(e.target);
                        if(!target.hasClass('slide-control-pad'))
                            target = target.parent();
                        if(!target.hasClass('slide-control-pad'))
                            target = target.parent();
                        target = target.find('div')[0];
                        volEl = target;
                    }
                    var offX  = (/* e.offsetX */ false || (e.clientX || e.pageX) - angular.element(target).offset().left)
                    var vol = (offX / target.offsetWidth * 100);
                    scope.ngModel = vol < 0 ? 0 : (vol > 100 ? 100 : vol);
                    scope.$apply();
                }
                function stopVolumeSet(){
                    volEl = false;
                    angular.element($window).unbind('mouseup touchend',stopVolumeSet);
                    angular.element($window).unbind('mousemove touchmove',setVolume);
                }
                scope.startVolumeChange = function(e){
                    setVolume(e);
                    angular.element($window).bind('mouseup touchend',stopVolumeSet);
                    angular.element($window).bind('mousemove touchmove',setVolume);
                }
            }
        }
    }]);
});
