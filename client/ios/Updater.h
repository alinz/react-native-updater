/*
 * BSD License
 * Copyright (c) 2015, Ali Najafizadeh.
 * All rights reserved.
 */

@class RCTRootView;

typedef void (^UpdateRootViewBlock)(RCTRootView *);
typedef void (^SucceedBlock)();
typedef void (^FailedBlock)(NSError *);

@interface Updater : NSObject

@property (atomic, strong) UpdateRootViewBlock beforeUpdaterLaunch;
@property (atomic, strong) UpdateRootViewBlock beforeMainAppLaunch;

+ (id)instance;

- (id)initWithModuleName:(NSString *)moduleName;
- (void)launchUpdaterApp;
- (void)launchMainApp;
- (void)downloadMainAppFromURL:(NSURL *) url
              withSucceedBlock:(SucceedBlock)succeedBlock
                andFailedBlock:(FailedBlock)failedBlock ;

- (NSString *)loadCurrentVersion;
- (void)saveVersionAsCurrent:(NSString*)version;

- (id)view;

@end
