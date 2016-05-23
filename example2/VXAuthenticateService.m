#import <Foundation/Foundation.h>
#import "VXAuthenticateService.h"
#import "VXRootTabBarController.h"
#import "VXAppDelegate.h"
#import <SSKeychain/SSKeychain.h>
#import <TSMessage.h>

@implementation VXAuthenticateService {
    
    VXLoginViewController *_loginViewController;
}

+ (id)sharedInstance {
    static VXAuthenticateService *sharedManager = nil;
    static dispatch_once_t onceToken;
    dispatch_once(&onceToken, ^{
        sharedManager = [[self alloc] init];
    });
    return sharedManager;
}

- (id) init
{
    self.mainStoryboard = [UIStoryboard storyboardWithName:@"Main" bundle:nil];

    
    self = [super init];
    if (self) {
    
    }
    return self;
}


-(void)userWithLogin {
    
    [self presentBasicLoginModal];

}

- (void) tryBasicLoginWithUsername:(NSString*)username password:(NSString*)password successBlock:(NSString* (^)(void))successBlock{
    
    NSLog(@"Attempting Login via tryBasicLoginWithUsername");
    
    VXApiClient *api = [[VXApiClient alloc] init];
    
    [api authenticateWithEmail:username password:password success:^(NSURLSessionDataTask *task, id responseObject) {
        
        VXUser *user = [[VXUser alloc] init];
        user.authenticationToken = responseObject[@"authentication_token"];
        user.email = responseObject[@"email"];
        user.username = responseObject[@"username"];
        user.name = responseObject[@"name"];
        
        VXAppDelegate *appDelegate = [VXAppDelegate sharedInstance];
        appDelegate.currentUser = user;
        
        [SSKeychain setPassword:responseObject[@"authentication_token"] forService:kVXKeychainService account:responseObject[@"email"]];

        [self dismissBasicLoginModal];
        
    } failure:^(NSURLSessionDataTask *task, NSError *error) {
        [TSMessage showNotificationWithTitle:@"Error:" subtitle:error.localizedDescription type:TSMessageNotificationTypeError];
    }];
    
}


- (NSURLSessionDataTask*)doLoginThenRunBlock:(void (^) (void))runThis {
    
    VXApiClient *api = [[VXApiClient alloc] init];
    NSDictionary *parameters = @{@"email": @"email@domain.com", @"password": @"passphrase"};   
    
    return [api POST:kSessionFormat parameters:parameters success:^(NSURLSessionDataTask *task, id responseObject) {
        
        VXUser *user = [[VXUser alloc] init];
        user.authenticationToken = responseObject[@"authentication_token"];
        user.email = responseObject[@"email"];
        user.username = responseObject[@"username"];
        user.name = responseObject[@"name"];
        
        VXAppDelegate *appDelegate = [VXAppDelegate sharedInstance];
        appDelegate.currentUser = user;
        
        [SSKeychain setPassword:responseObject[@"authentication_token"] forService:kVXKeychainService account:responseObject[@"email"]];
        
        runThis();
        
        [self dismissBasicLoginModal];
        
    } failure:^(NSURLSessionDataTask *task, NSError *error) {
        [TSMessage showNotificationWithTitle:@"Error:" subtitle:error.localizedDescription type:TSMessageNotificationTypeError];
    }];
    
}


- (void)presentBasicLoginModal
{
    [[VXRootTabBarController sharedInstance] showModalViewController:_loginViewController];
};

- (void)dismissBasicLoginModal
{
    [[VXRootTabBarController sharedInstance] dismissModalViewController];
};


@end
