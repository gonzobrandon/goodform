/**
 * Defines the main routes in the application.
 * The routes you see here will be anchors '#/' unless specifically configured otherwise.
 */

define(['app'], function(app) {
    'use strict';
    return app.config(function($stateProvider, $urlRouterProvider, $locationProvider) {


        /* ***** TRAILING SLASHES *****   */

        $urlRouterProvider.rule(function($injector, $location) {
            var path = $location.path()
                // Note: misnomer. This returns a query object, not a search string
                , search = $location.search()
                , params
                ;
            
            // check to see if the path already ends in '/'
            if (path[path.length - 1] === '/') {
                return;
            }
            
            // If there was no search string / query params, return with a `/`
            if (Object.keys(search).length === 0) {
                return path + '/';
            }
            
            // Otherwise build the search string and return a `/?` prefix
            params = [];
            angular.forEach(search, function(v, k){
                params.push(k + '=' + v);
            });
            return path + '/?' + params.join('&');
        });
        
        /* ***** BEGIN ROUTES  *****   */

        var root = {
            name: 'root',
            abstract: true,
            url: '',
            views: {
                'header@': {
                    templateUrl: '/partials/header.tpl.html'
                },
                'footer@': {
                    templateUrl: '/partials/footer.tpl.html'
                }
            }
        }

        var hello = {
            name: 'hello',
            parent: root,
            url: '/',
            onEnter: function() {
            
            },
            onExit: function() {

            },
            views: {
                'content@': {
                    templateUrl: '/partials/hello.tpl.html',
                    controller: 'helloCtrl'
                }
            }
        }

        var artistCreate = {
            name: 'artistCreate',
            parent: root,
            url: '/artist/create/',
            onEnter: function() {
                
            },
            onExit: function() {
            },
            views: {
                'content@': {
                    templateUrl: '/partials/artist_create.tpl.html',
                    controller: 'artistCreateCtrl'
                }
            }
        }

        var artistRoot = {
            name: 'artistRoot',
            parent: root,
            onEnter: function() {
            },
            onExit: function() {
            },
            views: {
                'content@': {
                    templateUrl: '/partials/artist.tpl.html',
                    controller: 'artistCtrl'
                }
            }
        }

        var artistFollowers = {
            name: 'artist.followers',
            parent: artistRoot,
            onEnter: function() {
            },
            onExit: function() {
            },
            views: {
                'subBody': {
                    templateUrl: '/partials/slugSubviews/followers.tpl.html',
                    controller: 'artistFollowersCtrl'
                }
            }
        }

        var artistHome = {
            name: 'artist',
            url: '/artist/',
            parent: artistRoot,
            onEnter: function() {
            },
            onExit: function() {
            },
            views: {
                'subBody': {
                    templateUrl: '/partials/slugSubviews/artist_home.tpl.html'
/*                     controller: 'artistFollowersCtrl' */
                }
            }
        }
        
        /* DRUNKEN HEARTS ARTIST DEMO */
        var drunkenHearts = {
            name: 'drunkenhearts',
            url: '/drunkenhearts/',
            parent: root,
            views: {
                'content@': {
                    templateUrl: '/partials/artist_json.tpl.html',
                    controller: 'trickArtistCtrl'
                }
            }
        }



        /* TUMBLEWEEDWANDERERS ARTIST DEMO */
        var tumbleweedwanderers = {
            name: 'tumbleweedwanderers',
            url: '/tumbleweedwanderers/',
            parent: root,
            views: {
                'content@': {
                    templateUrl: '/partials/artist_json.tpl.html',
                    controller: 'trickArtistCtrl'
                }
            }
        }

        var userRoot = {
            name: 'userRoot',
            parent: root,
            views: {
                'content@': {
                    templateUrl: '/partials/user.tpl.html'
                }
            }
        }

        var userFollowers = {
            name: 'user.followers',
            parent: userRoot,
            onEnter: function() {
            },
            onExit: function() {
            },
            views: {
                'content@': {
                    templateUrl: '/partials/user.tpl.html',
                    controller: 'artistCtrl'
                }
            }
        }

        var userHome = {
            name: 'user',
            parent: userRoot,
            onEnter: function() {
            },
            onExit: function() {
            },
            views: {
                'content@': {
                    templateUrl: '/partials/user.tpl.html',
                    controller: 'artistCtrl'
                }
            }
        }

        var search = {
            name: 'search',
            parent: root,
            url: '/search/:query/',
            onEnter: function() {
            },
            onExit: function() {
            },
            views: {
                'content@': {
                    templateUrl: '/partials/search.tpl.html',
/*                     controller: 'artistCtrl' */
                }
            }
        }
        
        var passwordReset = {
            name: 'passwordReset',
            parent: root,
            url: '/password-reset/',
            views: {
                'content@': {
                    templateUrl: '/partials/manage/password-reset.tpl.html',
                    controller: 'passwordResetCtrl'
                }
            }
        }
        
        var emailVerify = {
            name: 'emailVerify',
            parent: root,
            url: '/email/verify/',
            views: {
                'content@': {
                    templateUrl: '/partials/manage/email-verify.tpl.html',
                    controller: 'emailVerifyCtrl'
                }
            }
        }
        
        var manage = {
            name: 'manage',
            parent: root,
            views: {
                'content@': {
                    templateUrl: '/partials/manage.tpl.html',
                    controller: 'manageCtrl'
                }
            }
        }

        var manageUser = {
            name: 'manage.user',
            parent: manage,
            url: '/manage/',
            views: {
                'subBody': {
                    templateUrl: '/partials/manage/user.tpl.html',
                    controller: 'manageUserCtrl'
                }
            }
        }

        var manageVenues = {
            name: 'manage.venues',
            parent: manage,
            url: '/manage/venues/',
            views: {
                'subBody': {
                    templateUrl: '/partials/manage/venues.tpl.html',
                    controller: 'manageVenuesCtrl'
                }
            }
        }

        var manageArtists = {
            name: 'manage.artists',
            parent: manage,
            url: '/manage/artists/',
            views: {
                'subBody': {
                    templateUrl: '/partials/manage/artists.tpl.html',
                    controller: 'manageArtistsCtrl'
                }
            }
        }

        var manageEvents = {
            name: 'manage.events',
            parent: manage,
            url: '/manage/artist/:artist/events/',
            views: {
                'subBody': {
                    templateUrl: '/partials/manage/events.tpl.html',
                    controller: 'manageArtistEventsCtrl'
                }
            }
        }

        var manageAlbums = {
            name: 'manage.albums',
            parent: manage,
            url: '/manage/artist/:artist/albums/',
            views: {
                'subBody': {
                    templateUrl: '/partials/manage/albums.tpl.html',
                    controller: 'manageArtistAlbumsCtrl'
                }
            }
        }

        var manageTracks = {
            name: 'manage.tracks',
            parent: manage,
            url: '/manage/album/:album/tracks/',
            views: {
                'subBody': {
                    templateUrl: '/partials/manage/tracks.tpl.html',
                    controller: 'manageAlbumTracksCtrl'
                }
            }
        }

        /* SLUG ROUTING */

        var routeSlug = function ($injector, $location) {
            var parts = $location.path().replace(/^\/|\/$/g, '').split('/');
            var rs = $injector.get('rootService');
            var $http = $injector.get('$http');
            var $state = $injector.get('$state');
            var url = $location.url();
            $http.get('/api/slugs/'+parts[0])
                .success(function(data){
                    // Here we look through states set in the stateProvider to find out if the subview in the url exists. If not, it goes to the default for the slug. If the slug type does not have a view in the allowed array it will go to search.
                    var target = data.response.target;
                    var allowed = [
                        'artist',
                        'user'
                    ];
                    if(allowed.indexOf(data.response.type) != -1){
                        var newState = data.response.type;
                        if($state.get(newState + '.' + parts[1]))
                            newState += '.' + parts[1];
                        rs.slugTarget = data.response.target;
                        $state.go(newState,{}, {location:false});
                    } else {
                        $location.url('/search/'+parts[0]).replace();
                    }
                })
                .error(function(){
                    $location.url('/search/'+parts[0]).replace();
                });
                        
        };

        $stateProvider
            .state(root)
            .state(hello)
            .state(emailVerify)
            .state(passwordReset)
            .state(drunkenHearts)
            .state(tumbleweedwanderers)
            .state(manage)
                .state(manageUser)
                .state(manageVenues)
                .state(manageArtists)
                .state(manageEvents)
                .state(manageAlbums)
                .state(manageTracks)
            .state(artistCreate)
            .state(artistRoot)
                .state(artistHome)
                .state(artistFollowers)
            .state(userRoot)
                .state(userHome)
                .state(userFollowers)
            .state(search)

        $urlRouterProvider.otherwise(routeSlug);
        $locationProvider.html5Mode(true);

        /* ***** END ROUTES  *****   */

    })
});