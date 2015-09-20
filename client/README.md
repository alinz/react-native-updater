## Spec

> This doc contains all materials require for client to show how it works. Currently it shows my random thoughts. I will try to organize it.

### feature

- prior to loading `react-native` app, and internal app must be executed.
- I'd like to have updater app to be written in react-native.
    - search how to instance of `react-native` can be executed in parallel or in sequence.
- [TASK] updater must download the `bundle.json` from configured location and compare with it's own version.
  - version must follow [Semantic Versioning 2.0.0](http://semver.org)
    - minor and major sections is a trigger to start the upgrade.
    - major requires updater to notify the user to download it from app store.
  - [TASK] we need a utility to parse and compare version.
  
