define(['./module', 'moment-timezone'], function (controllers, moment) {
    'use strict';

    controllers.controller('passwordResetCtrl', function (userService, $scope, $location) {
        $scope.password = '';
        $scope.save = function(){
            $scope.loading = true;
            userService.passwordReset($location.hash(), $scope.password).success(function(dat){
                userService.doLogin({
                    username: dat.response.username,
                    password: $scope.password
                }).success(function(){
                    $location.url('/');
                })
            }).error(function(dat){
                $scope.loading = false;
                $scope.errors = [dat.notifications.input.password, dat.notifications.input.reset_token]
            });
        }
    })

    controllers.controller('emailVerifyCtrl', function ($http, $location) {
        var qs = $location.search();
        if(qs.token){
            $http.patch('/api/email-verification/'+qs.token).success(function(dat){
                
            }).error(function(dat){
                
            });
        }
    })
    
    controllers.controller('manageCtrl', ['userService', '$scope',  function (userService, $scope) {
        $scope.userService = userService;
        $scope.loggedIn = false;
        var bindLogin = function(){
            $scope.loggedIn = true;
        };
        userService.onLogin(bindLogin, true, $scope);
    }]);

    controllers.controller('manageVenuesCtrl', ['$http', '$scope', '$timeout',  function ($http, $scope, $timeout) {
        $scope.loadingList = true;
        $scope.pushNew = function(data){
            $scope.venues.unshift(data.response);
            $scope.addNew = false;
        }
        $http.get('/api/venues/').success(function(data){
            $scope.venues = data.response.reverse();
            $scope.loadingList = false;
        });
        
        $scope.editorSlide = [];
        $scope.editorIf = [];
        var toggleTimeouts = [];
        $scope.toggleEditor = function(index){
            if($scope.editorSlide[index]){
                $scope.editorSlide[index] = false;
                // destroy form after a little time in case it is updated outside of this window
                if(toggleTimeouts[index]){
                    $timeout.cancel(toggleTimeouts[index])
                    toggleTimeouts[index] = false;
                }
                $timeout(function(){
                    if(!$scope.editorSlide[index])
                        $scope.editorIf[index] = false;
                }, 10000, true);
            } else {
                $scope.editorIf[index] = true;
                $scope.editorSlide[index] = true;
            }
            return false;
        }
        
    }]);

    controllers.controller('manageArtistsCtrl', ['$http', '$scope', '$timeout',  function ($http, $scope, $timeout) {
        $scope.loadingList = true;
        $scope.addNew = false;
        $scope.pushNew = function(data){
            $scope.artists.unshift(data.response);
            $scope.addNew = false;
        }
        $http.get('/api/artists?joins={"type":"full"}&type=admin').success(function(data){
            $scope.artists = data.response.reverse();
            $scope.loadingList = false;
        });
        
        $scope.editorSlide = [];
        $scope.editorIf = [];
        var toggleTimeouts = [];
        $scope.toggleEditor = function(index){
            if($scope.editorSlide[index]){
                $scope.editorSlide[index] = false;
                // destroy form after a little time in case it is updated outside of this window
                if(toggleTimeouts[index]){
                    $timeout.cancel(toggleTimeouts[index])
                    toggleTimeouts[index] = false;
                }
                $timeout(function(){
                    if(!$scope.editorSlide[index])
                        $scope.editorIf[index] = false;
                }, 10000, true);
            } else {
                $scope.editorIf[index] = true;
                $scope.editorSlide[index] = true;
            }
            return false;
        }
        
    }]);

    controllers.controller('manageArtistEventsCtrl', ['$http', '$scope', '$state', '$anchorScroll', '$timeout',  function ($http, $scope, $state, $anchorScroll, $timeout) {
        $anchorScroll();
        $scope.addNew = false;
        
        
        $scope.pushNew = function(data){
            $scope.events.unshift(data.response);
            $scope.addNew = false;
        }
        
        $scope.loadingArtist = true;
        $scope.loadingList = true;
        $http.get('/api/artists/'+$state.params.artist+'?joins={"type":"full"}&type=admin').success(function(data){
            
            $scope.artist = data.response;
            $scope.loadingArtist = false;
            $http.get('/api/events?joins={"type":"full"}&artist='+$state.params.artist).success(function(data){
                $scope.events = data.response;
                $scope.loadingList = false;
            });;
        });
        
        $scope.editorSlide = [];
        $scope.editorIf = [];
        var toggleTimeouts = [];
        $scope.toggleEditor = function(index){
            if($scope.editorSlide[index]){
                $scope.editorSlide[index] = false;
                // destroy form after a little time in case it is updated outside of this window
                if(toggleTimeouts[index]){
                    $timeout.cancel(toggleTimeouts[index])
                    toggleTimeouts[index] = false;
                }
                $timeout(function(){
                    if(!$scope.editorSlide[index])
                        $scope.editorIf[index] = false;
                }, 10000, true);
            } else {
                $scope.editorIf[index] = true;
                $scope.editorSlide[index] = true;
            }
            return false;
        }
        
    }]);

    controllers.controller('manageArtistAlbumsCtrl', ['$http', '$scope', '$state', '$anchorScroll', '$timeout',  function ($http, $scope, $state, $anchorScroll, $timeout) {
        $anchorScroll();
        $scope.addNew = false;
        
        
        $scope.pushNew = function(data){
            $scope.albums.unshift(data.response);
            $scope.addNew = false;
        }
        
        $scope.loadingArtist = true;
        $scope.loadingList = true;
        $http.get('/api/artists/'+$state.params.artist+'?joins={"type":"full"}&type=admin').success(function(data){
            
            $scope.artist = data.response;
            $scope.loadingArtist = false;
            $scope.albums = data.response.albums;
            $scope.loadingList = false;
        });
        
        $scope.editorSlide = [];
        $scope.editorIf = [];
        var toggleTimeouts = [];
        $scope.toggleEditor = function(index){
            if($scope.editorSlide[index]){
                $scope.editorSlide[index] = false;
                // destroy form after a little time in case it is updated outside of this window
                if(toggleTimeouts[index]){
                    $timeout.cancel(toggleTimeouts[index])
                    toggleTimeouts[index] = false;
                }
                $timeout(function(){
                    if(!$scope.editorSlide[index])
                        $scope.editorIf[index] = false;
                }, 10000, true);
            } else {
                $scope.editorIf[index] = true;
                $scope.editorSlide[index] = true;
            }
            return false;
        }
        
    }]);

    controllers.controller('manageAlbumTracksCtrl', ['$http', '$scope', '$state', '$anchorScroll', '$timeout',  function ($http, $scope, $state, $anchorScroll, $timeout) {
        $anchorScroll();
        $scope.addNew = false;
        
        $scope.pushNew = function(data){
            $scope.tracks.unshift(data.response);
            $scope.addNew = false;
        }
        
        $scope.loadingAlbum = true;
        $scope.loadingList = true;
        $http.get('/api/albums/'+$state.params.album+'?joins={"type":"full","tracks":{"type":"full"}}').success(function(data){
            
            $scope.album = data.response;
            $scope.loadingAlbum = false;
            $scope.tracks = data.response.tracks.response;
            $scope.loadingList = false;
        });
        
        var changing = false;
        var changeAgain = false;
        $scope.changeSort = function(){
            var input = [];
            angular.forEach($scope.tracks, function(val, key){
                input.push({
                    id: val.id,
                    track_number: key
                });
            });
            if(!changing){
                $http.patch('/api/tracks', input).success(function(){
                    changing = false;
                    if(changeAgain){
                        changeAgain = false;
                        $scope.changeSort();
                    }
                }).error(function(){
                    changing = false;
                });;
            } else {
                changeAgain = true;
            }
                
        }
        
        $scope.editorSlide = [];
        $scope.editorIf = [];
        var toggleTimeouts = [];
        $scope.toggleEditor = function(index){
            if($scope.editorSlide[index]){
                $scope.editorSlide[index] = false;
                // destroy form after a little time in case it is updated outside of this window
                if(toggleTimeouts[index]){
                    $timeout.cancel(toggleTimeouts[index])
                    toggleTimeouts[index] = false;
                }
                $timeout(function(){
                    if(!$scope.editorSlide[index])
                        $scope.editorIf[index] = false;
                }, 10000, true);
            } else {
                $scope.editorIf[index] = true;
                $scope.editorSlide[index] = true;
            }
            return false;
        }
        
    }]);

    controllers.controller('manageUserCtrl', ['$scope', '$modal', '$log', '$http', '$filter', '$cookieStore', 'rootService', 'userService',  function ($scope, $modal, $log, $http, $filter, $cookieStore, rootService, userService) {
        $scope.user = angular.copy(userService.userObj);
        
        if($scope.user.pic_square == 'http://static1.theradiobox.com/inhouse/web/anon-user.png')
            $scope.user.pic_square = '';

        $scope.loading = false;

        $scope.save = function () {

            $scope.loading = true;
            
            var input = {
                first_name: $scope.user.first_name || '',
                last_name: $scope.user.last_name || '',
                sex: $scope.user.sex || ''
            };
            if($scope.user.pic_square && $scope.user.pic_square != '')
                input.pic_square = $scope.user.pic_square;
            

            userService.profilePatch(input).success(function(data, status, headers, config) {
                angular.forEach(data.response, function(val, key){
                    userService.userObj[key] = val;
                });
                $scope.loading = false;
            }).error(function(data, status, headers, config) {
                $scope.errors = data.notifications.input;
                $scope.loading = false;
            });

        };


    }]);

});
