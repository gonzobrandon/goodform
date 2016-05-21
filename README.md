# Good Form

I'll walk through 3 examples of what I consider good form coding. Everything is code I have written - in a variety of languages and frameworks. Over the past 8 years, I have been working on proprietary and private software repos and I would love to open everything up to showcase - I'm bound by nondisclosure agreements.

## Example 1: golang, AngularJS, Node, Postgres

#### View this project at [http://radiobox.gonzobrandon.com/](http://radiobox.gonzobrandon.com/)

#### See the source for everything at [http://github.com/goodform/example1/](http://github.com/goodform/example1/)

This project was called Radiobox. I was a co-founder and the head developer of this project. The idea was a live-streaming backend delievered via AngularJS on desktop and a native iOS app on mobile.

Some basics of this project:

+ GruntJS build process (minifying, condensing and compbining all external and internal libs).

+ The demo link was used for development. The production site had a compressed minified JS file.

+ Golang and stretr web framework (extremely fast http response).

+ Single page AngularJS web application. The web page, CSS is served up statically, then the JS (angular router) calls the API to populate lists, boxes, etc.

+ Statically served archived audio files. Everything lives on AWS S3.

+ Live audio was delivered via Akamai CDN.

+ Responsive web design.

+ Religiously included all library dependencies and hard baked them into the repo. 3rd part repo updates have really bitten us in the pase.

#### Example Snippet 1: 

I chose angular's router to handle all pages. Here was is a snippet of the router declaration:

Notes: 

+ We use indentation for readability. Inline comments go a long way for a new developer, or when visiting this code months or years later.

+ Promises are used where possible to avoid nested callback hell.

+ Due to the early version of Angular Router, we had to code in exceptions for trailing slashes and searching when a path was not found.


```
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

      
        var userRoot = {
            name: 'userRoot',
            parent: root,
            views: {
                'content@': {
                    templateUrl: '/partials/user.tpl.html'
                }
            }
        }

// (OTHER ROUTES OMITTED FOR BREVITY OF THIS README)

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
```





