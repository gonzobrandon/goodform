#import "VXUser.h"
#import "VXLoginViewController.h"
#import "VXApiClient.h"

@interface VXAuthenticateService : NSObject

@property (nonatomic, strong) UIStoryboard *mainStoryboard;

+ (id)sharedInstance;
- (NSURLSessionDataTask*)doLoginThenRunBlock:(void (^)(void))runThis;
- (void) userWithLogin;
- (void) tryBasicLoginWithUsername:(NSString*)username password:(NSString*)password successBlock:(NSString* (^)(void))successBlock;
- (void) presentBasicLoginModal;
- (void) dismissBasicLoginModal;

@end
