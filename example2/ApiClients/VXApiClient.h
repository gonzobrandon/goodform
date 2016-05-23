#import <Foundation/Foundation.h>
#import <AFNetworking/AFNetworking.h>
#import "VXLoginViewController.h"

typedef void(^SuccessBlock)(NSURLSessionDataTask *task, id responseObject);
typedef void(^FailBlock)(NSURLSessionDataTask *task, NSError *error);

@interface VXApiClient : AFHTTPSessionManager <VXLoginViewControllerDelegate>


@property (strong, nonatomic) VXLoginViewController *loginViewController;
@property (strong, nonatomic) UIViewController *rootViewController;
@property (strong, nonatomic) NSURLSessionDataTask *pauseMe;

- (void) dismissBasicLoginModal;
+ (id)sharedInstance;
- (NSURLSessionDataTask *)authGET:(NSString *)URLString parameters:(id)parameters success:(void ( ^ ) ( NSURLSessionDataTask *task , id responseObject ))success failure:(void ( ^ ) ( NSURLSessionDataTask *task , NSError *error ))failure;
- (NSURLSessionDataTask *)authPOST:(NSString *)URLString parameters:(id)parameters success:(void ( ^ ) ( NSURLSessionDataTask *task , id responseObject ))success failure:(void ( ^ ) ( NSURLSessionDataTask *task , NSError *error ))failure;
- (NSURLSessionDataTask *)authPUT:(NSString *)URLString parameters:(id)parameters success:(void ( ^ ) ( NSURLSessionDataTask *task , id responseObject ))success failure:(void ( ^ ) ( NSURLSessionDataTask *task , NSError *error ))failure;
- (NSURLSessionDataTask *)authDELETE:(NSString *)URLString parameters:(id)parameters success:(void ( ^ ) ( NSURLSessionDataTask *task , id responseObject ))success failure:(void ( ^ ) ( NSURLSessionDataTask *task , NSError *error ))failure;
- (void)authenticateWithEmail:(NSString *)email password:(NSString *)password success:(void ( ^ ) ( NSURLSessionDataTask *task , id responseObject ))success failure:(void ( ^ ) ( NSURLSessionDataTask *task , NSError *error ))failure;
- (void)loginWithEmail:(NSString*)email password:(NSString*)password withBlock:(void (^) (void))runBlock;

@end
