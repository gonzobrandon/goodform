#import "VXUploadService.h"
#import "S3.h"


@implementation VXUploadService

- (void)obtainPreSignedURLRequestForPUTDataDaskWithBucket:(NSString *)bucket key:(NSString *)key contentType:(NSString *)contentType completion:(void(^)(NSURL *signedURL))completion
{

    NSParameterAssert(bucket);
    NSParameterAssert(key);
    NSParameterAssert(contentType);
    NSParameterAssert(completion);
    
    NSString *AWSSecretKey = kAWSUploadKeySecret;
    AWSStaticCredentialsProvider *credentialsProvider = [AWSStaticCredentialsProvider credentialsWithAccessKey:kAWSUploadKeyValue secretKey:AWSSecretKey];
    
    AWSServiceConfiguration *serviceConfiguration = [AWSServiceConfiguration configurationWithRegion:AWSRegionUSEast1 credentialsProvider:credentialsProvider];
    serviceConfiguration.maxRetryCount = 10;
    
    AWSS3PreSignedURLBuilder *builder = [[AWSS3PreSignedURLBuilder alloc] initWithConfiguration:serviceConfiguration];
    
    AWSS3GetPreSignedURLRequest *request = [[AWSS3GetPreSignedURLRequest alloc] init];
    request.bucket = bucket;
    request.key = key;
    request.HTTPMethod = AWSHTTPMethodPUT;
    request.expires = [NSDate dateWithTimeIntervalSinceNow:(60.0 * 60.0)];
    request.contentType = contentType;
    
    BFTask *task = [builder getPreSignedURL:request];
    [task continueWithBlock:^id(BFTask *task) {
        
        NSURL *signedURL = task.result;
        completion(signedURL);
       
        return nil;
    }];
}

- (void)uploadJPEGImageData:(NSData *)imageData withSignedURL:(NSURL *)signedURL success:(void(^)(NSURL *uploadedImageURL))success failure:(void(^)(NSError *error))failure
{
    NSParameterAssert(imageData);
    NSParameterAssert(signedURL);
    NSParameterAssert(success);
    NSParameterAssert(failure);
    
    NSMutableURLRequest *URLRequest = [NSMutableURLRequest requestWithURL:signedURL];
    
    NSString *contentType = @"image/jpeg";
    [URLRequest setValue:contentType forHTTPHeaderField:@"Content-Type"];
    
    URLRequest.HTTPMethod = @"PUT";
    URLRequest.HTTPBody = imageData;
    
    /*
     * self.backgroundSessionManager is an instance of AFURLSessionManager
     */
    NSURLSessionDataTask *dataTask = [self.backgroundSessionManager dataTaskWithRequest:URLRequest completionHandler:^(NSURLResponse *response, id responseObject, NSError *error) {

        if(error)
        {
            failure(error);
        }
        else
        {
            NSURL *uploadedImageURL = [NSURL URLWithString:URLRequest.URL.path];

            success(uploadedImageURL);
        }
    }];
    
    [dataTask resume];
}

- (void)beginUploadingJPEGImage:(UIImage *)image success:(void(^)(NSURL *uploadedImageURL))success {
 
    self.selectedImage = image;
    
    NSURLSessionConfiguration *sessionConfiguration = [NSURLSessionConfiguration defaultSessionConfiguration];
    
    self.backgroundSessionManager = [[AFURLSessionManager alloc] initWithSessionConfiguration:sessionConfiguration];

    NSString *bucket = kAWSUploadsBucketName;
    NSString *filename = [[VXUtils randomStringWithLength:16] stringByAppendingFormat:@"%@", @".jpg"];
    NSString *key    = [[VXUtils newUUID] stringByAppendingFormat:@"%@%@", @"/", filename];
    NSString *contentType = @"image/jpeg";
    
    __weak typeof(self) weakSelf = self;
    [self obtainPreSignedURLRequestForPUTDataDaskWithBucket:bucket key:key contentType:contentType completion:^(NSURL *signedURL) {
        
        __strong typeof(self) strongSelf = weakSelf;
        
        NSData *imageData = UIImageJPEGRepresentation(strongSelf.selectedImage, 1.0f);
        //[strongSelf startDataTaskForJPEGImageData:imageData withSignedURL:signedURL];
        
        [self uploadJPEGImageData:imageData withSignedURL:signedURL success:success failure:^(NSError *error) {
    
            //Fail Code
            
        }];
    
    }];
    
}

//- (void)startDataTaskForJPEGImageData:(NSData *)imageData withSignedURL:(NSURL *)signedURL
//{
//    [self uploadJPEGImageData:imageData withSignedURL:signedURL success:^(NSURL *uploadedImageURL) {
//        
//        
//    } failure:^(NSError *error) {
//        
//        //Fail Code
//        
//    }];
//}

@end
