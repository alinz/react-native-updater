@class RCTRootView;
@class UIView;

typedef void (^CompletionTaskBlock)(RCTRootView *);

@interface Updater : NSObject

@property (atomic, strong) CompletionTaskBlock beforeUpdaterLaunch;
@property (atomic, strong) CompletionTaskBlock beforeMainAppLaunch;

+ (id)instanceWithModuleName:(NSString *)moduleName;

- (id)initWithModuleName:(NSString *)moduleName;
- (void)launchUpdaterApp;
- (void)launchMainApp;

- (UIView *)view;

@end
