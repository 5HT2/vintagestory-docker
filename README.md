# vintagestory-docker

This Dockerfile and run script are all you need to run a Vintage Story server with no further configuration required.

You do need to ensure your server has port `42420` opened, as Vintage Story requires that.

## Usage

Your `~/.env` should look like so: 

```bash
# Edit the path to where you would like vintagestory config files to be saved
export VINTAGESTORY_PATH="$HOME/vintagestory"
```

Next, all you need to do in order to run the server is:

```bash
git clone https://github.com/l1ving/vintagestory-docker
cd vintagestory-docker
# Build the image, only required when updating the version
docker build -t vintagestory .
chmod +x run.sh

# Run the server. You can run this command again to re-create the server
# in case you edited the config file path, or to forcefully restart it.
./run.sh
```

### Updating

This Dockerfile will be kept up to date as Vintage Story updates (feel free to open a pull request if it is not).

Once you've followed the usage instructions and want to apply a new update, just do

```bash
cd vintagestory-docker
git pull
docker build -t vintagestory .
./run.sh
```

### Using an older version

You can use an older version of Vintage Story if needed, by building with the `VERSION` env variable set, eg

```bash
VERSION=1.16.1 docker build -t vintagestory .
```
