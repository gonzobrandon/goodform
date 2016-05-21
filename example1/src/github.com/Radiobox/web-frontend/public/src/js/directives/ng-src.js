define(['./module'], function (directives) {
    'use strict';
    directives.directive('ngSrc', [ function () {
        
        /* SET IMAGE PREFIX HERE */
        var cdnUrl = '/img/';
        
        return {
            restrict: 'A',
            priority: 98,
            link: function(scope, element, attr) {
                attr.$observe('ngSrc', function(value) {
                    if(element[0].tagName == 'IMG'){
                        if(attr.ngSrc.indexOf('/cdn/') === 0){
                            attr.$set('ngSrc', cdnUrl + attr.ngSrc.substring(5));
                            attr.$set('src', cdnUrl + attr.ngSrc.substring(5));
                        }
                    }
                });
            }
        }
    }]);
});
