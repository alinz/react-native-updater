/*
 * BSD License
 * Copyright (c) 2015, Ali Najafizadeh.
 * All rights reserved.
 */

var React = require('react-native');
var {
  Component,
  NativeModules: {
    UpdaterManager
  }
} = React;

function versionToString(version, options) {
  const { major, minor, patch } = version;
  return `${major}.${minor}.${patch}`;
}

class UpdaterComponent extends Component {
  constructor(props) {
    super(props);
  }

  _getReleasesURL(version) {
    return `${this.props.domain}/releases/${version.major}`;
  }

  _getBundleURLFor(version) {
    return `${this.props.domain}/bundles/${versionToString(version)}`;
  }

  //simple versioning.
  _parseVersion(version) {
    let parsedVersion = null;
    const match = /^(\d+)\.(\d+)\.(\d+)$/.exec(version);

    if (match) {
      parsedVersion = {
        major: parseInt(match[0], 10),
        minor: parseInt(match[1], 10),
        patch: parseInt(match[2], 10)
      };
    }

    return parsedVersion;
  }

  async _getLocalVersion() {
    const version = await UpdaterManager.localVersion();
    return this._parseVersion(version);
  }

  async _process() {
    try {
      const localVersion = await this._getLocalVersion();
      console.log(localVersion);

    } catch(e) {
      console.log(e);
    }

    // this._getLocalVersion((version) => {
    //   fecth(this._getReleasesURL(version))
    // });
  }

  _compare(localVersion, remoteVersion) {
    const result = ['updateViaBundle', 'updateViaAppStore'];
    const pattern = /^([x\*])\.([x\*])\.([x\*])$/.exec(this.props.pattern);

    //localVersion's and remoteVersion's major must be the same
    if (pattern[0] === 'x') {
      if (localVersion.major !== remoteVersion.major) {

      }
    }

  }
};

UpdaterComponent.propTypes = {
  domain: React.PropTypes.string,
  pattern: React.PropTypes.string
};

UpdaterComponent.defaultTypes = {
  pattern: 'x.*.*'
}

module.exports = UpdaterComponent;
