
define(['./module', 'jwplayer'], function (directives) {
    'use strict';
    directives.directive('playStatic', ['playerService', function (playerService) {

        return {
            restrict: 'AEC',
            scope: {
              mediaLink: '@playStatic'
            },
            link: function(scope, element, attrs) {

                var clickingCallback = function(e) {
/*                     alert('clicked smoehtin '); */
                }

                element.bind('click', clickingCallback);

/*                 scope.$on('$destroy', function () { alert('gone, bitch'); }); */

            }
        }
    }]);
});
