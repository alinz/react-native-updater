#import "Util.h"
#import <Security/Security.h>

@implementation Util

+ (NSData *)GenSecureRandomBytesWithSize:(int)size {
  void * bytes = malloc(size);
  SecRandomCopyBytes(kSecRandomDefault, size, bytes);
  NSData * data = [NSData dataWithBytes:bytes length:size];
  free(bytes);

  return data;
}

@end
