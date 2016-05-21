define(['./module'], function (services) {
    'use strict';

    services.provider('animationInterval', function () {

        var interval = false;
        var functions = Array();
        var promiseIndex = 0;
        
        this.$get = function($interval) {
        
            var run = function(){
                var length = 0;
                angular.forEach(functions, function(func){
                    func();
                    length++;
                });
                if(!length){
                    $interval.cancel(interval);
                    interval = false;
                }
            }
            
            var wrappedService = function(func){
                if(typeof(func) != 'function')
                    return false;
                var promise = promiseIndex++;
                functions[promise] = func;
                if(!interval)
                    interval = $interval(run, 25);
                return promise;
            }
            wrappedService.cancel = function(promise) {
                delete(functions[promise]);
            }
            return wrappedService;
        }

    });

});
