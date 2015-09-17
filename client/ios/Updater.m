/*
 * BSD License
 * Copyright (c) 2015, Ali Najafizadeh.
 * All rights reserved.
 */
 
#import <Foundation/Foundation.h>
#import <UIKit/UIKit.h>

#import "Updater.h"
#import "RCTRootView.h"

dispatch_queue_t _serialQueue;
static Updater *updaterInstance = nil;

@implementation Updater {
  UINavigationController *_navigator;
  NSString *_moduleName;
}

@synthesize beforeUpdaterLaunch;
@synthesize beforeMainAppLaunch;

+ (id)instance {
  return updaterInstance;
}

- (id)initWithModuleName:(NSString *)moduleName {
  self = [super init];

  if (self) {
    _serialQueue = dispatch_queue_create("github.com/alinz/react-native-updater", DISPATCH_QUEUE_SERIAL);

    _navigator = [[UINavigationController alloc] init];
    [_navigator setNavigationBarHidden:YES animated:NO];

    _moduleName = moduleName;

    updaterInstance = self;
  }

  return self;
}

- (void)launchUpdaterApp {
  //make sure that the following code inside block executed by main thread. React-Native's requirement!
  dispatch_async(dispatch_get_main_queue(), ^{
    if ([_navigator.viewControllers count] == 0) {
      NSURL *bundleURL = [[NSBundle mainBundle] URLForResource:@"main" withExtension:@"jsbundle"];
      UIViewController *updaterViewController = [self rootViewWithModuleName:@"Updater"
                                                                   bundleURL:bundleURL];
      [_navigator pushViewController:updaterViewController animated:YES];
    } else {
      [_navigator popViewControllerAnimated:YES];
    }
  });
}

- (void)launchMainApp {
  //make sure that the following code inside block executed by main thread. React-Native's requirement!
  dispatch_async(dispatch_get_main_queue(), ^{
    if ([_navigator.viewControllers count] == 1) {
      NSURL *bundleURL = [self localURLForFilename:@"main.jsbundle"];
      UIViewController *updaterViewController = [self rootViewWithModuleName:_moduleName
                                                                   bundleURL:bundleURL];
      [_navigator pushViewController:updaterViewController animated:YES];
    } else {
      NSLog(@"Warning: either updater is not launched or main app is already launched.");
    }
  });
}

- (void)downloadMainAppFromURL:(NSURL *) url
              withSucceedBlock:(SucceedBlock)succeedBlock
                andFailedBlock:(FailedBlock)failedBlock {
  NSURLRequest *request = [NSURLRequest requestWithURL:url];
  [NSURLConnection sendAsynchronousRequest:request
                                     queue:[[NSOperationQueue alloc] init]
                         completionHandler:^(NSURLResponse *response, NSData *data, NSError *error) {
                           if (error) {
                             failedBlock(error);
                           }
                           if (data) {
                             [self saveUpdateBundleWithData:data];
                             succeedBlock();
                           }
                         }];
}

- (UIView *)view {
  return _navigator.view;
}

-(UIViewController *)rootViewWithModuleName:(NSString *)moduleName
                                   bundleURL:(NSURL *)bundleURL {

  RCTRootView *rootView = [[RCTRootView alloc] initWithBundleURL:bundleURL
                                                      moduleName:moduleName
                                               initialProperties:nil
                                                   launchOptions:nil];
  if ([moduleName isEqualToString:@"Updater"]) {
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

- (void)saveUpdateBundleWithData:(NSData *)data {
  NSString *url = [[self localURLForFilename:@"main.jsbundle"] path];
  [data writeToFile:url atomically:YES];
}

- (NSURL *)localURLForFilename:(NSString *)filename {
  NSArray *paths = NSSearchPathForDirectoriesInDomains(NSDocumentDirectory,  NSUserDomainMask, YES);
  NSString *documentsDirectory = [paths objectAtIndex:0];
  NSString *appFile = [documentsDirectory stringByAppendingPathComponent:filename];

  appFile = [NSString stringWithFormat:@"file://%@", appFile];

  NSURL* bundleURL = [NSURL URLWithString:appFile];

  return bundleURL;
}

- (NSString *)loadFileFromDocuments:(NSString *)filename {
  NSString *content = nil;
  NSURL *path = [self localURLForFilename:filename];

  if ([[NSFileManager defaultManager] fileExistsAtPath:[path path]]) {
    content = [NSString stringWithContentsOfFile:[path path]
                                        encoding:NSUTF8StringEncoding
                                           error:NULL];
  }

  return content;
}

- (NSString *)loadFileFromBundle:(NSString *)filename {
  NSString *content = nil;
  NSString *ext = [filename pathExtension];
  NSURL *path = [[NSBundle mainBundle] URLForResource:[filename stringByDeletingPathExtension]
                                        withExtension:ext];

  if ([[NSFileManager defaultManager] fileExistsAtPath:[path path]]) {
    content = [NSString stringWithContentsOfFile:[path path]
                                        encoding:NSUTF8StringEncoding
                                           error:NULL];
  }

  return content;
}

- (NSString *)loadCurrentVersion {
  NSString *versionFileName = @"bundle.version";
  NSString *version = nil;

  version = [self loadFileFromDocuments: versionFileName];

  if (version == nil) {
    version = [self loadFileFromBundle:versionFileName];
  }

  if (version == nil) {
    NSLog(@"Error, your app doesn't have bundle.version in neither bundle nor documents");
  }

  return version;
}

- (void)saveVersionAsCurrent:(NSString*)version {
  NSString *versionFileName = @"bundle.version";
  NSURL *path = [self localURLForFilename:versionFileName];
  [version writeToFile:[path path]
            atomically:YES
              encoding:NSUTF8StringEncoding
                 error:nil];
}

@end
