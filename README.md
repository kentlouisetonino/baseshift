## Description

![Screenshot from 2024-02-16 21-57-51](https://github.com/kentlouisetonino/baseshift/assets/69438999/62bffd01-0cc2-4a84-a41d-6a7fae23459b)

<br />

> - A CLI application for converting number systems.

> - This will support binary, decimal, octal, and hexadecimal systems.

> - The tools and technologies include Go, Bash, Snapcraft, and Linux.

> - Snap: https://snapcraft.io/baseshift

<br />
<br />



## Local Development
> - Note: `The root main.go is a symbolic link`.

> - Build and run the app.

```bash
# Change permission for the build script.
chmod +x build

# Run the app.
./app
```

<br />
<br />



## Installation
```bash
# If app is already installed.
sudo snap refresh baseshift

# If app is not yet installed.
sudo snap install baseshift

# Remove the app using snap.
sudo snap remove baseshift

# Running the app.
baseshift
```
