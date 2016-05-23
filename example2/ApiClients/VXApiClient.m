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
    [[VXAppDelegate sharedInstance].rootTabBarController dismissModalViewController];
}

- (void)doLoginAndPassRunBlock:(void (^) (void))runBlock {
    
    //SHOW THE LOGIN SCREEN AND WAIT FOR RESPONSE
    self.loginViewController.successBlock = runBlock;
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
