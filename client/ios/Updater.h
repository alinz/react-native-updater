@class RCTRootView;

@interface Updater : NSObject

+ (id)instanceWithModuleName:(NSString *)moduleName;

- (id) initWithModuleName:(NSString *)moduleName;
- (void) launchUpdaterApp;
- (void) launchMainApp;

- (void) beforeUpdaterLaunch:(RCTRootView *)launchRootView;
- (void) beforeMainAppLaunch:(RCTRootView *)launchRootView;

@end