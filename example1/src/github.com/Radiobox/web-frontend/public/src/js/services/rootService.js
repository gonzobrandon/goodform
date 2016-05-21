define(['./module'], function (services) {
    'use strict';

    services.provider('rootService', function () {

        this.$get =  function($http, $rootScope) {
            
            
            /* EXTEND jQLITE HERE */
            angular.element.prototype = angular.extend(angular.element.prototype, {
                offset: function( elem ) {
                    
                    if(!elem){
                        elem = this[0];
                    }
                    var docElem, win,
                        box = { top: 0, left: 0 },
                        doc = elem && elem.ownerDocument;
            
                    if ( !doc ) {
                        return;
                    }
            
                    docElem = doc.documentElement;

                    if ( typeof elem.getBoundingClientRect !== 'undefined' ) {
                        box = elem.getBoundingClientRect();
                    }
                    win = window;
                    return {
                        top: box.top + win.pageYOffset - docElem.clientTop,
                        left: box.left + win.pageXOffset - docElem.clientLeft
                    }
                },
                computedStyle: function(strCssRule){
                    var oElm = this[0];
                    var strValue = '';
                    if(document.defaultView && document.defaultView.getComputedStyle){
                        strValue = document.defaultView.getComputedStyle(oElm, '').getPropertyValue(strCssRule);
                    }
                    else if(oElm.currentStyle){
                        strCssRule = strCssRule.replace(/\-(\w)/g, function (strMatch, p1){
                            return p1.toUpperCase();
                        });
                        strValue = oElm.currentStyle[strCssRule];
                    }
                    return strValue;
                },
                parentForm: function (el){
                    el = el || this;
                    var parent = el.parent();
                    
                    if(parent[0].tagName == 'BODY')
                        return false;
                    
                    if(parent[0].tagName == 'FORM')
                        return parent;
                    else
                        return this.parentForm(parent);
                }

            });
            
            
        
            $http.patch = function (url, data, config) {
                return $http(angular.extend(config || {}, {
                    method: 'PATCH',
                    url: url,
                    data: data
                }));
            };

            $http.defaults.headers.common['Accept'] = 'application/vnd.radiobox.encapsulated+json';
            $http.defaults.headers.put['Content-Type'] = 'application/json';
            $http.defaults.headers.patch['Content-Type'] = 'application/json';
            $http.defaults.headers.post['Content-Type'] = 'application/json';

            $rootScope.isEmptyObj = function(obj) {
                var name;
                for ( name in obj ) {
                    return false;
                }
                return true;
            };

            $rootScope.seconds2HMS = function(time) {
                var seconds = Math.floor(time);
                var minutes = Math.floor(seconds/60);
                var hours = 0;
                if (minutes > 0 && minutes < 59) {
                    seconds -= minutes*60;
                } else if (minutes > 59) {
                    hours = Math.floor(minutes/60);
                    minutes -= hours * 60;
                };

                seconds = seconds < 10 ? '0' + seconds : seconds;
                minutes = minutes < 10 ? '0' + minutes : minutes;

                if (hours > 0) {
                    return hours + ":" + minutes + ":" + seconds;
                } else {
                    return minutes + ":" + seconds;
                }
            };
            
            var wrappedService = {
                /**
                 * Public Methods
                 */
                alerts: [
/*                     { type: 'danger', msg: 'Oh snap! This provides messages to users from any module' }, */
/*                     { type: 'success', msg: 'Well done! You successfully read this important alert message.' } */
                ],

                addAlert: function(type, msg) {
                    this.alerts.push({type: type, msg: msg});
                },

                closeAlert: function(index) {
                    this.alerts.splice(index, 1);
                },

                slugTarget: {}
                
                /**
                 * Public Properties
                 */

            }

            return wrappedService
        }

    });

});
