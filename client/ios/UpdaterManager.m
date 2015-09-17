/*
 * BSD License
 * Copyright (c) 2015, Ali Najafizadeh.
 * All rights reserved.
 */
 
#import "UpdaterManager.h"
#import "Updater.h"

@implementation UpdaterManager

RCT_EXPORT_MODULE();

RCT_EXPORT_METHOD(launchUpdaterApp)
{
  Updater *updater = [Updater instance];
  [updater launchUpdaterApp];
}

RCT_EXPORT_METHOD(launchMainApp)
{
  Updater *updater = [Updater instance];
  [updater launchMainApp];
}

RCT_EXPORT_METHOD(downloadUpdate:(NSString *)path
             withSucceedCallback:(RCTPromiseResolveBlock)resolve
               andFailedCallback:(RCTPromiseRejectBlock)reject)
{
  NSURL *url = [NSURL URLWithString:path];
  Updater *updater = [Updater instance];
  [updater downloadMainAppFromURL:url
                 withSucceedBlock: ^{ resolve(@[]); }
                   andFailedBlock: ^(NSError *error) {
                     reject(error);
                   }];
}

RCT_EXPORT_METHOD(localVersion:(RCTPromiseResolveBlock)resolve
                      rejecter:(RCTPromiseRejectBlock)reject)
{
  Updater *updater = [Updater instance];
  NSString *currentVersion = [updater loadCurrentVersion];
  resolve(currentVersion);
}

RCT_EXPORT_METHOD(saveVersion:(NSString *)version)
{
  Updater *updater = [Updater instance];
  [updater saveVersionAsCurrent:version];
}
@end
