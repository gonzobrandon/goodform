define(['./module'], function (directives) {
    'use strict';
    directives.directive('dragSort', ['$interval', '$window', function ($interval, $window) {
        
        return {
            restrict: 'A',
            priority: 1001,
            scope: {
                dragSort: '=',
                onChange: '&'
            },
            link: function(scope, element) {
                var volInt;
                var dragEl;
                var $dragEl;
                var apply;
                var topStart;
                var startDragPos;
                var maxHeight;
                var slotSize;
                var $coveredEl;
                var bounds;
                scope.$watch('dragSort', function(val){
                    angular.forEach(element.find('button'), function(val,key){
                        angular.element(val).unbind('mousedown touchstart', scope.startMove);
                        angular.element(val).bind('mousedown touchstart', scope.startMove);
                    });
                });
                
                function moveIt(e){
                    
                    e.preventDefault();
                    
                    var offY  = (/* e.offsetY */ false || (e.clientY || e.pageY));
                    if(!startDragPos)
                        startDragPos = offY;
                    var move = topStart + offY - startDragPos;
                    move = move < 0 ? 0 : (move > maxHeight ? maxHeight : move);
                    $dragEl.css('top', move + 'px');
                    
                    var picked = false;
                    angular.forEach(bounds, function(val, key){
                        if(move + 20 > val.bound)
                            $coveredEl = val;
                    });
                    
                    scope.$apply();
                }
                scope.$watch(function(){ return $coveredEl; }, function(val, prev){
                    element.children().css('margin-top', '0px');
                    if(val && $dragEl.attr('index') > val.attr('index'))
                        val.css('margin-top', slotSize + 'px');
                    else if(val && $dragEl.attr('index') <= val.attr('index'))
                        val.next().css('margin-top', slotSize + 'px');
                    
                    if(val && prev)
                        element.children().css('transition', 'margin-top 250ms ease-out');
                });
                function stopMove(){
                    
                    element.children().css('transition', '');
                    var moveToRow = $coveredEl.attr('index');
                    var targetRow = $dragEl.attr('index');
                    
                    if(moveToRow != targetRow){
                        var moveObj = scope.dragSort.splice(targetRow,1);
                        scope.dragSort.splice(moveToRow,0,moveObj[0]);
                        if(scope.onChange)
                            scope.onChange();
                    }
                    
                    $dragEl.removeClass('floating');
                    $dragEl.css('top', '');
                    $coveredEl.css('margin-top', '');
                    element.css('border-bottom', '');
                    volInt = false;
                    dragEl = false;
                    $dragEl = false;
                    apply = false;
                    topStart = false;
                    startDragPos = false;
                    maxHeight = false;
                    slotSize = false;
                    $coveredEl = false;
                    bounds = false;
                    scope.$apply();
                    angular.element($window).unbind('mouseup touchend',stopMove);
                    angular.element($window).unbind('mousemove touchmove',moveIt);
                }
                scope.startMove = function(e){
                    
                    bounds = element.children();
                    angular.forEach(bounds, function(val, key){
                        bounds[key] = val = angular.element(val);
                        bounds[key].bound = val.offset().top - val.parent().offset().top;
                    });
                    
                    if(dragEl){
                        apply = true;
                    } else {
                        var target = angular.element(e.target);
                        if(!target.hasClass('drag-sort-row'))
                            target = target.parent();
                        if(!target.hasClass('drag-sort-row'))
                            target = target.parent();
                        dragEl = target[0];
                        $dragEl = target;
                    }
                    slotSize = dragEl.offsetHeight;
                    maxHeight = element[0].offsetHeight - dragEl.offsetHeight;
                    topStart = $dragEl.offset().top - $dragEl.parent().offset().top;
                    $dragEl.addClass('floating');
/*                     $coveredEl = $dragEl.next(); */
                    element.css('border-bottom', 'none');
                    
                    
                    moveIt(e);
                    angular.element($window).bind('mouseup touchend',stopMove);
                    angular.element($window).bind('mousemove touchmove',moveIt);
                }
            }
        }
    }]);
});
