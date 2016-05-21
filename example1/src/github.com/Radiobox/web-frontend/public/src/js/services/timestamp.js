define(['./module'], function (services) {
    'use strict';

    services.provider('$timestamp', function () {
        
        this.$get = function($interval) {
            
            var Timestamp = function(dt){
                
                Timestamp.asd = function(){
                    return this.getHours();
                };
                return new Date(dt);
            }
            var bum = new Timestamp();
            
            var wrappedService = function(func){
                if(!scope.ngModel){
                    scope.ngModel = new Date();
                } else if(typeof(scope.ngModel) == 'string'){
                    
                    scope.ngModel = new Date(scope.ngModel);
                }
            }
            return wrappedService;
        }

    });

});
