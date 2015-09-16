#import "UpdaterManager.h"
#import "Updater.h"

@implementation UpdaterManager

RCT_EXPORT_MODULE();

RCT_EXPORT_METHOD(launchUpdaterApp:(NSString *)moduleName)
{
  Updater *updater = [Updater instanceWithModuleName:moduleName];
  [updater launchUpdaterApp];
}

RCT_EXPORT_METHOD(launchMainApp:(NSString *)moduleName)
{
  Updater *updater = [Updater instanceWithModuleName:moduleName];
  [updater launchMainApp];
}

@end
