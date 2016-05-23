# Good Form

I'll walk through 3 examples of what I consider good form coding. Everything is code I have written - in a variety of languages and frameworks. Over the past 8 years, I have been working on proprietary and private software repos and I would love to open everything up to showcase - I'm bound by nondisclosure agreements.

## Example 1: golang, AngularJS, Node, Postgres

#### View this project:
[http://radiobox.gonzobrandon.com/](http://radiobox.gonzobrandon.com/)

#### See the complete project source:
[https://github.com/gonzobrandon/goodform/tree/master/example1](https://github.com/gonzobrandon/goodform/tree/master/example1). 

(see go/src/github.com/Radiobox path for most of my original work)

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


## Example 2: iOS Cocoa (Objective C)

Objective C is strange. Syntactically, accessing properites with a dot, and passing arguments between `[]` separated by a space is difficult to love. However, I do enjoy the delegate design pattern. And Apple's Cocoa/Touch API is fun to work with. 

Below is an example of a class I wrote to log into a RESTful JSON API and pass authenticated and un-authenticated `GET`, `POST` and `PUT` methods to the backend. It was designed to be a singleton (one instance only) and pass things off to an Authentication Service that actually makes the `NSURLSessionDataTask` call.

I regard this a good form code because it is very object oriented and implements encapsulation and adheres to separation of concern design.

```
#import "VXApiClient.h"
#import "VXAppDelegate.h"
#import <AFNetworking.h>
#import <SSKeychain/SSKeychain.h>
#import <TSMessage.h>

@implementation VXApiClient {

}

+ (id)sharedInstance {
    static VXApiClient *sharedInstance = nil;
    static dispatch_once_t onceToken;
    dispatch_once(&onceToken, ^{
        sharedInstance = [[self alloc] initWithBaseURL:[NSURL URLWithString:kVXApiURL]];
    });
    return sharedInstance;
}

- (id)initWithBaseURL:(NSURL *)url
{
    self = [super initWithBaseURL:url];
    if(!self)
        return nil;

    self.rootViewController = [VXAppDelegate sharedInstance].rootViewController;
    self.loginViewController = [[UIStoryboard storyboardWithName:@"Main" bundle:nil] instantiateViewControllerWithIdentifier:@"LoginViewController"];
    self.loginViewController.delegate = self;
    
    
    return self;
}


- (NSURLSessionDataTask *)dataTaskWithRequest:(NSURLRequest *)request
                            completionHandler:(void (^)(NSURLResponse *response, id responseObject, NSError *error))completionHandler
{

    //reset headers
    [[self requestSerializer] setValue:nil forHTTPHeaderField:@"X-AUTH-TOKEN"];
    [[self requestSerializer] setValue:nil forHTTPHeaderField:@"X-AUTH-EMAIL"];
    
    return [super dataTaskWithRequest:request completionHandler:completionHandler];
}


- (NSURLSessionDataTask *)authPOST:(NSString *)URLString parameters:(id)parameters success:(void ( ^ ) ( NSURLSessionDataTask *task , id responseObject ))success failure:(void ( ^ ) ( NSURLSessionDataTask *task , NSError *error ))failure {
    
    NSURLSessionDataTask* (^authPOSTBlock)(NSString*, id, SuccessBlock, FailBlock) = ^ NSURLSessionDataTask* (NSString *aURLString, id aParameters, SuccessBlock aSuccess, FailBlock aFailure ) {
        
        [[self requestSerializer] setValue:[VXAppDelegate sharedInstance].currentUser.authenticationToken forHTTPHeaderField:@"X-AUTH-TOKEN"];
        [[self requestSerializer] setValue:[VXAppDelegate sharedInstance].currentUser.email forHTTPHeaderField:@"X-AUTH-EMAIL"];
        
        return [self POST:aURLString parameters:aParameters success:aSuccess failure:aFailure];
        
    };
    
    if (![VXAppDelegate sharedInstance].currentUser) {
        
        [self doLoginAndPassRunBlock:^{
            authPOSTBlock(URLString, parameters, success, failure);
        }];
        
        NSURLSessionDataTask *dummyTask = [[NSURLSessionDataTask alloc] init];
        return dummyTask;
        
    } else {
        return authPOSTBlock(URLString, parameters, success, failure);
    }
    
}

- (NSURLSessionDataTask *)authPUT:(NSString *)URLString parameters:(id)parameters success:(void ( ^ ) ( NSURLSessionDataTask *task , id responseObject ))success failure:(void ( ^ ) ( NSURLSessionDataTask *task , NSError *error ))failure {
    
    NSURLSessionDataTask* (^authPUTBlock)(NSString*, id, SuccessBlock, FailBlock) = ^ NSURLSessionDataTask* (NSString *aURLString, id aParameters, SuccessBlock aSuccess, FailBlock aFailure ) {
        
        [[self requestSerializer] setValue:[VXAppDelegate sharedInstance].currentUser.authenticationToken forHTTPHeaderField:@"X-AUTH-TOKEN"];
        [[self requestSerializer] setValue:[VXAppDelegate sharedInstance].currentUser.email forHTTPHeaderField:@"X-AUTH-EMAIL"];
        
        return [self PUT:aURLString parameters:aParameters success:aSuccess failure:aFailure];
        
    };
    
    if (![VXAppDelegate sharedInstance].currentUser) {
        
        [self doLoginAndPassRunBlock:^{
            authPUTBlock(URLString, parameters, success, failure);
        }];
        
        NSURLSessionDataTask *dummyTask = [[NSURLSessionDataTask alloc] init];
        return dummyTask;
        
    } else {
        return authPUTBlock(URLString, parameters, success, failure);
    }
    
}

- (NSURLSessionDataTask *)authDELETE:(NSString *)URLString parameters:(id)parameters success:(void ( ^ ) ( NSURLSessionDataTask *task , id responseObject ))success failure:(void ( ^ ) ( NSURLSessionDataTask *task , NSError *error ))failure {
    
    NSURLSessionDataTask* (^authDELETEBlock)(NSString*, id, SuccessBlock, FailBlock) = ^ NSURLSessionDataTask* (NSString *aURLString, id aParameters, SuccessBlock aSuccess, FailBlock aFailure ) {
        
        [[self requestSerializer] setValue:[VXAppDelegate sharedInstance].currentUser.authenticationToken forHTTPHeaderField:@"X-AUTH-TOKEN"];
        [[self requestSerializer] setValue:[VXAppDelegate sharedInstance].currentUser.email forHTTPHeaderField:@"X-AUTH-EMAIL"];
        
        return [self DELETE:aURLString parameters:aParameters success:aSuccess failure:aFailure];
        
    };
    
    if (![VXAppDelegate sharedInstance].currentUser) {
        
        [self doLoginAndPassRunBlock:^{
            authDELETEBlock(URLString, parameters, success, failure);
        }];
        
        NSURLSessionDataTask *dummyTask = [[NSURLSessionDataTask alloc] init];
        return dummyTask;
        
    } else {
        return authDELETEBlock(URLString, parameters, success, failure);
    }
    
}

- (NSURLSessionDataTask *)authGET:(NSString *)URLString parameters:(id)parameters success:(void ( ^ ) ( NSURLSessionDataTask *task , id responseObject ))success failure:(void ( ^ ) ( NSURLSessionDataTask *task , NSError *error ))failure {

    NSURLSessionDataTask* (^authGETBlock)(NSString*, id, SuccessBlock, FailBlock) = ^ NSURLSessionDataTask* (NSString *aURLString, id aParameters, SuccessBlock aSuccess, FailBlock aFailure ) {
        
        [[self requestSerializer] setValue:[VXAppDelegate sharedInstance].currentUser.authenticationToken forHTTPHeaderField:@"X-AUTH-TOKEN"];
        [[self requestSerializer] setValue:[VXAppDelegate sharedInstance].currentUser.email forHTTPHeaderField:@"X-AUTH-EMAIL"];
        
        return [self GET:aURLString parameters:aParameters success:aSuccess failure:aFailure];
        
    };
   
    if (![VXAppDelegate sharedInstance].currentUser) {
       
        [self doLoginAndPassRunBlock:^{
            authGETBlock(URLString, parameters, success, failure);
        }];
        
        NSURLSessionDataTask *dummyTask = [[NSURLSessionDataTask alloc] init];
        return dummyTask;

    } else {
        return authGETBlock(URLString, parameters, success, failure);
    }
    
}


- (void) authenticateWithEmail:(NSString *)email password:(NSString *)password success:(void ( ^ ) ( NSURLSessionDataTask *task , id responseObject ))success failure:(void ( ^ ) ( NSURLSessionDataTask *task , NSError *error ))failure {
 
    NSDictionary *parameters = @{@"email": email, @"password": password};
    
    [self POST:kSessionFormat
    parameters:parameters constructingBodyWithBlock:nil
       success:success
       failure:failure];
    
}

- (void) dismissBasicLoginModal {
//    [self.loginViewController dismissViewControllerAnimated:YES completion:nil];
    [[VXAppDelegate sharedInstance].rootTabBarController dismissModalViewController];
}

- (void)doLoginAndPassRunBlock:(void (^) (void))runBlock {
    
    //SHOW THE LOGIN SCREEN AND WAIT FOR RESPONSE
    self.loginViewController.successBlock = runBlock;
//    [[VXAppDelegate sharedInstance].rootTabBarController presentViewController:self.loginViewController animated:YES completion:nil];
    [[VXAppDelegate sharedInstance].rootTabBarController showModalViewController:self.loginViewController];

}

- (void)loginWithEmail:(NSString*)email password:(NSString*)password withBlock:(void (^) (void))runBlock {
    
    NSDictionary *parameters = @{@"email": email, @"password": password};
  
    [self POST:kSessionFormat parameters:parameters success:^(NSURLSessionDataTask *task, id responseObject) {
        
        NSLog(@"%@", responseObject);
        
        VXUser *user = [[VXUser alloc] init];
        user.authenticationToken = responseObject[@"data"][@"authentication_token"];
        user.email = responseObject[@"data"][@"email"];
        user.username = responseObject[@"data"][@"username"];
        user.name = responseObject[@"data"][@"name"];

        VXAppDelegate *appDelegate = [VXAppDelegate sharedInstance];
        appDelegate.currentUser = user;

        [[NSUserDefaults standardUserDefaults] setObject:[NSKeyedArchiver archivedDataWithRootObject:user] forKey:@"currentUser"];
        

        NSLog(@"VXApiClient::loginWithEmail - Attempting to login");
        
        runBlock();
 
        [self dismissBasicLoginModal];

    } failure:^(NSURLSessionDataTask *task, NSError *error) {
        
        NSHTTPURLResponse *response = (NSHTTPURLResponse *)task.response;
        NSInteger statuscode = response.statusCode;
        
        if (statuscode == 401) {
            [TSMessage showNotificationWithTitle:@"Bad Login:" subtitle:@"Bad email or password. Try Again." type:TSMessageNotificationTypeError];
        } else {
            [TSMessage showNotificationWithTitle:@"Error:" subtitle:error.localizedDescription type:TSMessageNotificationTypeError];
        }


    }];
    
}

@end
```







