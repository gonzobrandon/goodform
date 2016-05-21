define(['./module', './rootService', 'facebook'], function (services) {
    'use strict';

    services.provider('userService', [function () {

            var credsDefault = {
                refreshToken: null,
                userToken: null,
                expiresOn: null,
                accountUrl: null,
                profileUrl:null
            };
            var creds = angular.copy(credsDefault);

            this.$get = function($http, $cookieStore, $filter, rootService, $modal, $location, $timeout, $rootScope) {

                /**
                 * Private Methods
                 */
                
                var clientId = '00000001';
                var loggedInCallbacks = [];
                var forceLogin = false;
                var refreshTimeout = false;
                var fbToken, fbId;
                
                var setHeaders = function(token) {
                    if (!token) {
                        delete $http.defaults.headers.common['Authorization'];
                        return;
                    }
                    $http.defaults.headers.common['Authorization'] = 'Bearer ' + token;
                    
                };
                
                var refreshToken = function() {
                    return $http.post('/api/token/', 
                    {
                        client_id: clientId,
                        grant_type: 'refresh_token',
                        refresh_token: creds.refreshToken
                    }).success(function(data){
                        handleToken(data);
                    }).error(function(){
                        $cookieStore.put('uCreds', {});
                        getLoginData();
                        if(forceLogin){
                            wrappedService.openLogin();
                            forceLogin = false;
                        }
                    });    
                }
                
                var handleToken = function(response){
                    response = response.response;
                    creds = {
                        userToken : response.access_token,
                        refreshToken : response.refresh_token,
                        expiresOn: Math.floor(new Date().getTime()/1000) + response.expires_in,
                        accountUrl: response.user_account,
                        profileUrl: response.user_profile
                    }
                    $cookieStore.put('uCreds', creds);
                    wrappedService.userId = response.user_id;
                    
                    setHeaders(creds.userToken);
                    
                    var expiresIn = creds.expiresOn - Math.floor(new Date().getTime()/1000);
                    if(refreshTimeout)
                        $timeout.cancel(refreshTimeout);
                    refreshTimeout = $timeout(refreshToken, (expiresIn - 60)*1000);
                }

                var getLoginData = function () {
                    var cookieCreds = $cookieStore.get('uCreds');
                    if(typeof(cookieCreds) != 'undefined' && typeof(cookieCreds.userToken) != 'undefined')
                        creds = cookieCreds;
                    else
                        creds = credsDefault;
                    if (creds.userToken) {
                        wrappedService.getUser().success(function(){
                            wrappedService.isLoggedIn = true;
                            wrappedService.isInvalid = false;
                            wrappedService.loading = false;
                            
                        }).error(function(){
                            $cookieStore.put('uCreds', {});
                            getLoginData();
                            if(forceLogin){
                                wrappedService.openLogin();
                                forceLogin = false;
                            }
                        });
                        return true;
                    }
                    wrappedService.userRole = null;
                    wrappedService.userId = null;
                    wrappedService.isLoggedIn = false;
                    wrappedService.isInvalid = false;
                    wrappedService.userObj = {};
                    wrappedService.loading = false;
                    if(loggedInCallbacks.length > 0 && wrapper.openLogin)
                        wrapper.openLogin();
                    
                    setHeaders(false);
                    return false;
                };
                
                var fbHandler = function(response, succ, fail) {
/*                     GO TO SERVER, CHECK IF ACCOUNT EXISTS, IF NOT CREATE IT, SHARE FB TOKEN, LOG IN WITH CORRESPONDING RBOX TOKEN */
                    var account_exists = false;
                    if(response) {
                        fbToken = response.authResponse.accessToken
                        fbId = response.authResponse.userID
                    }
                    wrappedService.doLogin().success(function(data){
                        handleToken(data);
                        if(succ)
                            succ(data);
                        
                    }).error(function(data){
                        FB.api(
                            "/me",
                            function (response) {
                                $rootScope.$apply(function(){
                                    if (response && !response.error) {
                                        wrappedService.userObj.username = response.username;
                                        wrappedService.userObj.isFB = true;
                                        
                                        wrappedService.openSignup();
                                        
                                    }
                                    if(fail)
                                        fail(data);
                                });
                            }
                        );
                    })
                };
                
                FB.init({
/*                     FOR RADIOBOX.COM */
                    appId : '407288952710118'
/*                     FOR SPERADEV.COM */
/*                     appId : '491948537531912', */
                });

/*
                FB.getLoginStatus(function(response) {
                    fbHandler(response);
                });
*/
                
                var wrappedService = {
                    /**
                     * Public Methods
                     */
                    
                    isLoggedIn: false,
                    loading: true,
                    
                    facebookAuth: function(succ, fail) {
                        FB.login(function(response){
                            fbHandler(response, succ, fail);
                        });
                    },
                    currentDialog: false,
                    openSignup: function () {
                        if(wrappedService.currentDialog != 'signup'){
                            wrappedService.currentDialog = 'signup';
                            var modalInstance = $modal.open({
                                templateUrl: 'signupModalContent.html',
                                controller: function($scope,  $modalInstance) {
                                    $scope.loading = false;
                                    $scope.isFB = wrappedService.userObj.isFB;
                                    $scope.formData = {
                                        username: wrappedService.userObj.username || '',
                                        email: '',
                                        password: ''
                                    };
                
                                    $scope.signup = function () {
                
                
                                        $scope.loading = true;
                
                                        wrappedService.doSignup($scope.formData).success(function(data, status, headers, config) {
                
                                            $modalInstance.close();
                
                                        }).error(function(data, status, headers, config) {
                                             $scope.errors = data.notifications.input;
/*                                              $scope.hasErrors = Object.keys($scope.errors).length > 0; */
                                             $scope.loading = false;
                                        });
                
                                    };
                                    $scope.facebookAuth = function() {
                                        $scope.loading = true;
                                        wrappedService.facebookAuth($modalInstance.close,function(){
                                            $scope.isFB = wrappedService.userObj.isFB;
                                            $scope.formData = {
                                                username: wrappedService.userObj.username || '',
                                                email: '',
                                                password: ''
                                            };
                                            $scope.loading = false;
                                        });
                                    };
                
                                    $scope.cancel = function () {
                                        $modalInstance.dismiss('cancel');
                                    };
                                }
                            });
                            modalInstance.result.finally(function () {
                                wrappedService.currentDialog = false;
                            });
                        }
            
                    },
                    openLogin: function () {
                        if(wrappedService.currentDialog != 'login'){
                            wrappedService.currentDialog = 'login';
                            var modalInstance = $modal.open({
                                templateUrl: 'loginModalContent.html',
                                controller: function($scope,  $modalInstance) {
                
                                    $scope.loading = false;
                                    
                                    $scope.formData = {
                                        username:'',
                                        password: ''
                                    };
                                    $scope.openSignup = function(){
                                        $modalInstance.dismiss('signup');
                                        wrappedService.openSignup();
                                    }
                                    $scope.login = function () {
                
                                        $scope.loading = true;
                
                                        wrappedService.doLogin($scope.formData).success(function(data, status, headers, config) {
                                            $modalInstance.close();
                                        }).error(function(data, status){
                                            $scope.errors = ['Invalid Credentials'];
                                            $scope.loading = false;
                                        });
                
                                    };
                                    $scope.sendToken = function(){
                                        $scope.loading = true;
                                        $scope.errors = false;
                                        $scope.checkEmail = false;
                                        $http.post('/api/password-reset',{
                                            email_or_username: $scope.formData.username,
                                            client_id: clientId
                                        }).success(function(dat){
                                            $scope.checkEmail = true;
                                        }).error(function(dat){
                                            $scope.errors = dat.notifications.err;
                                            $scope.loading = false;
                                        });
                                    };
                                    $scope.facebookAuth = function() {
                                        var close = function(){
                                            $modalInstance.close();
                                        }
                                        wrappedService.facebookAuth(close, close);
                                    };
                
                                    $scope.cancel = function () {
                                        $modalInstance.dismiss('cancel');
                                    };
                                },
                                resolve: {
                
                                }
                            });

                            modalInstance.result.finally(function () {
                                wrappedService.currentDialog = false;
                            });
                        }
                    },
                    getUser: function() {
                        if(creds.accountUrl)
                            return $http.get(
                                creds.accountUrl
                            ).success(function(data, success) {
                                wrappedService.userObj = data.response;
                            $http.get(
                                creds.profileUrl
                            ).success(function(data, success) {
                                for(var key in data.response)
                                    wrappedService.userObj[key] = data.response[key];
                            
                                while(loggedInCallbacks.length > 0)
                                    loggedInCallbacks.shift()();
                            });
                        });
                    },
                    doLogin:  function(aLogin) {
                         var user = {};
                         if(!fbId){
                            user = {
                                username: aLogin.username,
                                password: aLogin.password,
                                grant_type: 'password',
                                client_id: clientId
                            }
                         } else {
                            user = {
                                client_id: clientId, 
                                grant_type: 'facebook',
                                facebook_id: fbId,
                                facebook_token: fbToken
                            }
                         }
                        return $http.post('/api/token',user).success(function(data, success) {
                            handleToken(data);
                            getLoginData();

                        }).error(function(data, status, headers, config) {
                            wrappedService.isInvalid = true;
                        });

                    },
                    doSignup:  function(user) {
                        var user = {
                            username: user.username,
                            password: user.password,
                            email: user.email,
                        }
                        if(fbToken){
                            user.facebook_user = fbId;
                            user.facebook_token = fbToken;
                            user.password = fbToken;
                        }
                        return $http.post( '/api/user', user).success(function(data, success) {
                            
                            wrappedService.doLogin(user);

                        }).error(function(data, status, headers, config) {
                            wrappedService.isInvalid = true;
                        });

                    },
                    doLogout: function(){
                        $cookieStore.put('uCreds', {});
                        fbToken = false;
                        fbId = false;
                        getLoginData();
                        if(refreshTimeout)
                            $timeout.cancel(refreshTimeout);
                        $location.url('/');
                    },
                    profilePatch: function(data){
                        return $http.patch(creds.profileUrl, data);
                    },
                    passwordReset: function(token, password){
                        return $http.patch('/api/user', {
                            reset_token: token,
                            password: password
                        })
                    },
                    
                    /* FORCE LOGIN + CALLBACKS */
                    
                    // Makes the login dialog pop up if nobody is logged in. Makes sure it fires after the cookies and tokens are sorted
                    forceLogin: function() {
                        if(!wrappedService.loading) {
                            wrappedService.openLogin();
                        } else {
                            forceLogin = true;
                        }
                    },
                    // Sets a callback for when the auth flow successfully logs in a user. Good to use for restricting pages, delaying rendering where login info is needed, or isolating user specific functionality like following, purchasing, etc.. 
                    onLogin: function(callback, force, unbindScope) {
                        if(wrappedService.isLoggedIn) {
                            if(typeof(callback) == 'function')
                                callback();
                        } else {
                            if(typeof( unbindScope ) == 'object' && typeof( unbindScope.$on ) == 'function'){
                                 unbindScope.$on('$destroy', function(){
                                    wrappedService.offLogin(callback);
                                });
                            }
                            if(typeof(callback) == 'function')
                                loggedInCallbacks.push(callback);

                            if(force)
                                wrappedService.forceLogin();
                            
                        }
                    },
                    // Unbind onLogin callbacks. Make sure to relese callbacks where necessary to prevent stacking and weird stuff
                    offLogin: function(callback) {
                        angular.forEach(loggedInCallbacks, function(obj, key){
                            if(obj === callback)
                                loggedInCallbacks.splice(key,1);
                        });
                    },

                    /**
                     * Public Properties
                     */
                    userObj: {},
                    
                    fbObj: {}

                }
                
                var cookieCreds = $cookieStore.get('uCreds');
                if(typeof(cookieCreds) != 'undefined' && typeof(cookieCreds.refreshToken) != 'undefined'){
                    creds = cookieCreds;
                    refreshToken().success(getLoginData);
                } else {
                    getLoginData();
                }



                return wrappedService
            }

    }]);

});
