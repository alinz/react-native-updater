#import <Foundation/Foundation.h>
#import <UIKit/UIKit.h>

#import "Updater.h"
#import "RCTRootView.h"

dispatch_queue_t _serialQueue;

@implementation Updater {
  UINavigationController *_navigator;
  NSString *_moduleName;
}

@synthesize beforeUpdaterLaunch;
@synthesize beforeMainAppLaunch;

+ (id)instanceWithModuleName:(NSString *)moduleName {
  static Updater *updaterInstance = nil;
  static dispatch_once_t onceToken;

  dispatch_once(&onceToken, ^{
    updaterInstance = [[self alloc] initWithModuleName:moduleName];
  });

  return updaterInstance;
}

- (id) initWithModuleName:(NSString *)moduleName {
  self = [super init];

  if (self) {
    _serialQueue = dispatch_queue_create("com.example.name", DISPATCH_QUEUE_SERIAL);

    _navigator = [[UINavigationController alloc] init];
    [_navigator setNavigationBarHidden:YES animated:NO];

    _moduleName = moduleName;
  }

  return self;
}

- (void) launchUpdaterApp {
  if ([_navigator.viewControllers count] == 0) {
    NSURL *bundleURL = [[NSBundle mainBundle] URLForResource:@"updater" withExtension:@"jsbundle"];
    UIViewController *updaterViewController = [self rootViewWithModuleName:@"Updater"
                                                                 bundleURL:bundleURL];
    [_navigator pushViewController:updaterViewController animated:NO];
  } else {
    [_navigator popViewControllerAnimated:YES];
  }
}

- (void) launchMainApp {
  if ([_navigator.viewControllers count] == 1) {
    NSURL *bundleURL = [self savedMainAppPathAsURL];
    UIViewController *updaterViewController = [self rootViewWithModuleName:_moduleName
                                                                 bundleURL:bundleURL];
    [_navigator pushViewController:updaterViewController animated:NO];
  } else {
    NSLog(@"Error: either upadter is not launched or main app is already launched.");
  }
}

- (UIView *)view {
  return _navigator.view;
}

-(UIViewController *) rootViewWithModuleName:(NSString *)moduleName
                                   bundleURL:(NSURL *)bundleURL {

  RCTRootView *rootView = [[RCTRootView alloc] initWithBundleURL:bundleURL
                                                      moduleName:moduleName
                                               initialProperties:nil
                                                   launchOptions:nil];
  if ([moduleName isEqualToString:@"updater"]) {
    dispatch_sync(_serialQueue, ^{
      beforeUpdaterLaunch(rootView);
    });
  } else {
    dispatch_sync(_serialQueue, ^{
      beforeMainAppLaunch(rootView);
    });
  }

  UIViewController *viewController = [[UIViewController alloc] init];
  viewController.view = rootView;

  return viewController;
}

- (void) saveUpdateBundleWithContent:(NSString *)content {
  NSURL *urlPath = [self savedMainAppPathAsURL];

  [content writeToFile:[urlPath absoluteString]
            atomically:YES
              encoding:NSUTF8StringEncoding
                 error:nil];
}

- (NSURL *) savedMainAppPathAsURL {
  NSArray *paths = NSSearchPathForDirectoriesInDomains(NSDocumentDirectory,  NSUserDomainMask, YES);
  NSString *documentsDirectory = [paths objectAtIndex:0];
  NSString *appFile = [documentsDirectory stringByAppendingPathComponent:@"main.jsbundle"];

  appFile = [NSString stringWithFormat:@"file://%@", appFile];

  NSURL* bundleURL = [NSURL URLWithString:appFile];

  return bundleURL;
}

@end