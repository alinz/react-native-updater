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
             withSucceedCallback:(RCTResponseSenderBlock)succeed
               andFailedCallback:(RCTResponseErrorBlock)failed)
{
  NSURL *url = [NSURL URLWithString:path];
  Updater *updater = [Updater instance];
  [updater downloadMainAppFromURL:url
                 withSucceedBlock: ^{ succeed(@[]); }
                   andFailedBlock: ^(NSError *error) {
                     failed(error);
                   }];
}
@end
