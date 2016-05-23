#import <Foundation/Foundation.h>
#import "AFHTTPSessionManager.h"

@interface VXUploadService : NSObject

- (void)obtainPreSignedURLRequestForPUTDataDaskWithBucket:(NSString *)bucket key:(NSString *)key contentType:(NSString *)contentType completion:(void(^)(NSURL *signedURL))completion;
- (void)beginUploadingJPEGImage:(UIImage*)image success:(void(^)(NSURL *uploadedImageURL))success;

@property (strong, nonatomic) AFURLSessionManager *backgroundSessionManager;
@property (strong, nonatomic) UIImage *selectedImage;
@property (strong, nonatomic) NSString *successKey;

@end
