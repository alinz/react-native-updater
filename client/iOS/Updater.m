
#import <Foundation/Foundation.h>
#import <UIKit/UIKit.h>

#import "Updater.h"
#import "RCTRootView.h"

@implementation Updater {
  UINavigationController *_navigator;
}

- (id) init {
  self = [super init];

  if (self) {
    _navigator = [[UINavigationController alloc] init];
    [_navigator setNavigationBarHidden:YES animated:NO];
  }

  return self;
}

- (void) launchUpdater {
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

}

- (void) beforeUpdaterLaunch:(RCTRootView *)launchRootView {

}

- (void) beforeMainAppLaucnh:(RCTRootView *)launchRootView {

}

-(UIViewController *) rootViewWithModuleName:(NSString *)moduleName
                                   bundleURL:(NSURL *)bundleURL{

  RCTRootView *rootView = [[RCTRootView alloc] initWithBundleURL:bundleURL
                                                      moduleName:moduleName
                                               initialProperties:nil
                                                   launchOptions:nil];

  UIViewController *viewController = [[UIViewController alloc] init];
  viewController.view = rootView;

  return viewController;
}

@end
