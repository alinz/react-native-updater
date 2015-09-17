/*
 * BSD License
 * Copyright (c) 2015, Ali Najafizadeh.
 * All rights reserved.
 */

const React = require('react-native');
const {
  Component,
  NativeModules: {
    UpdaterManager
  }
} = React;

const platform = 'ios';

function versionToString(version, options) {
  const { major, minor, patch } = version;
  return `${major}.${minor}.${patch}`;
}

class UpdaterComponent extends Component {
  constructor(props) {
    super(props);
  }

  _getReleasesURL(version) {
    return version?
      `${this.props.domain}/${platform}/releases/${version.major}` :
      `${this.props.domain}/${platform}/releases`;
  }

  _getBundleURLFor(version) {
    return `${this.props.domain}/${platform}/bundles/${versionToString(version)}`;
  }

  _parseVersion(version) {
    let parsedVersion = null;
    const match = /^(\d+)\.(\d+)\.(\d+)$/.exec(version);

    if (match) {
      parsedVersion = {
        major: parseInt(match[1], 10),
        minor: parseInt(match[2], 10),
        patch: parseInt(match[3], 10)
      };
    }

    return parsedVersion;
  }

  async _getLocalVersion() {
    const version = await UpdaterManager.localVersion();
    return this._parseVersion(version);
  }

  _properAsyncFetch(url) {
    return new Promise((resolve, reject) => {
      fetch(url)
        .then((response) => response.text())
        .then((value) => JSON.parse(value))
        .then(resolve)
        .catch(reject);
    });
  }

  async _getReleases() {
    const releasesURL = this._getReleasesURL();
    const releases = await this._properAsyncFetch(this._getReleasesURL());

    return releases.map((release) => {
      release.version = this._parseVersion(release.version);
      return release;
    });
  }

  _compilePattern(versionPattern) {
    console.log("#####", versionPattern);
    const match = /^(x|\*)\.(x|\*)\.(x|\*)$/.exec(versionPattern);
    return {
      major: match[1] !== 'x',
      minor: match[2] !== 'x',
      patch: match[3] !== 'x'
    };
  }

  _findUpdateRelease(releases, localVersion) {
    const versionPattern = this._compilePattern(this.props.versionPattern);

    let foundRelease;
    releases.some((release) => {
      let result = true;

      if (!versionPattern.major) {
        result = release.version.major == localVersion.major;
      } else {
        result = release.version.major > localVersion.major;
      }

      if (result && !versionPattern.minor) {
        result = release.version.minor == localVersion.minor;
      } else if (result) {
        result = release.version.minor >= localVersion.minor;
      }

      if (result && !versionPattern.patch) {
        result = release.version.patch == localVersion.patch;
      } else if (result) {
        result = release.version.patch > localVersion.patch;
      }

      if (result) {
        foundRelease = release;
      }

      return result;
    });

    return foundRelease;
  }

  async _process() {
    try {
      const localVersion = await this._getLocalVersion();
      const releases = await this._getReleases();
      const foundRelease = this._findUpdateRelease(releases, localVersion);
      const latestRelease = releases[0];

      const softUpdate = !!foundRelease;
      const hardUpdate = foundRelease && foundRelease !== latestRelease

      this.props.onUpdate(
        softUpdate? foundRelease : false,
        hardUpdate? latestRelease : false
      );

    } catch(e) {
      console.log(e);
    }
  }

  render() {
    return null;
  }
};

UpdaterComponent.propTypes = {
  domain: React.PropTypes.string,
  versionPattern: React.PropTypes.string,
  onUpdate: React.PropTypes.func
};

UpdaterComponent.defaultProps = {
  versionPattern: 'x.*.*',
  onUpdate: (softUpdate, hardUpdate) => {}
};

module.exports = UpdaterComponent;
