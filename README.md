# Install

```
$ go get -u github.com/zhangyoufu/unpbzx
```
The executable will be installed at `${GOPATH}/bin/unpbzx`.

# Usage

Make sure your `${PATH}` contains `${GOPATH}/bin`.

List macOS Installer payload
```
$ hdiutil attach -readonly /Applications/Install*.app/*/*/InstallESD.dmg
$ bsdtar -x -O -f /Volumes/InstallESD/Packages/Core.pkg Payload | unpbzx | cpio -it | head
.
./.DS_Store
./.file
./.vol
./Applications
./Applications/.DS_Store
./Applications/.localized
./Applications/App Store.app
./Applications/App Store.app/Contents
./Applications/App Store.app/Contents/Frameworks
```
